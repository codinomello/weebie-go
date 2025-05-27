package authentication

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/codinomello/weebie-go/api/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
)

// AuthService define a interface para operações de autenticação
type AuthService interface {
	Initialize() (*auth.Client, error)
	VerifyToken(idToken string) (*auth.Token, error)
	GenerateFirebaseUID() (string, error)
	GetUser(uid string) (*auth.UserRecord, error)
	GetUserByEmail(email string) (*auth.UserRecord, error)
	DeleteUser(uid string) error
	UpdateUserEmail(uid, newEmail string) error
	UpdateUserPassword(uid, newPassword string) error
	UpdateUserName(uid, newName string) error
	UpdateUserPhone(uid, newPhone string) error
	UpdateUserPhoto(uid, newPhoto string) error
	CreateDefaultAdmin(db *mongo.Database) error
}

// Implementa a interface AuthService para o Firebase
type FirebaseAuthentication struct {
	Client *auth.Client
}

// Cria uma nova instância do serviço de autenticação
func NewFirebaseAuthentication() *FirebaseAuthentication {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar Firebase: %v", err)
		return nil
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Erro ao criar client de autenticação: %v", err)
		return nil
	}

	return &FirebaseAuthentication{
		Client: client,
	}
}

// Inicializa o cliente Firebase
func (fa *FirebaseAuthentication) Initialize() (*auth.Client, error) {
	firebaseConfigPath := os.Getenv("FIREBASE_CREDENTIALS")
	if firebaseConfigPath == "" {
		return nil, errors.New("variável de ambiente 'FIREBASE_CREDENTIALS' não encontrada")
	}

	if _, err := os.Stat(firebaseConfigPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "arquivo de credenciais não encontrado")
	}

	opt := option.WithCredentialsFile(firebaseConfigPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "falha ao inicializar Firebase App")
	}

	fa.Client, err = app.Auth(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "falha ao obter cliente de autenticação")
	}

	return fa.Client, nil
}

// Verifica um token de ID do Firebase
func (fa *FirebaseAuthentication) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fa.Client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, errors.Wrap(err, "falha ao verificar token")
	}
	return token, nil
}

// Gera um novo UID no Firebase
func (fa *FirebaseAuthentication) GenerateFirebaseUID() (string, error) {
	email := fmt.Sprintf("%d@weebie.com", time.Now().UnixNano())
	user := (&auth.UserToCreate{}).
		Email(email).
		Password("senha123")

	createdUser, err := fa.Client.CreateUser(context.Background(), user)
	if err != nil {
		return "", errors.Wrap(err, "falha ao criar usuário no Firebase")
	}
	return createdUser.UID, nil
}

// Obtém um usuário pelo UID
func (fa *FirebaseAuthentication) GetUser(uid string) (*auth.UserRecord, error) {
	user, err := fa.Client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, errors.Wrap(err, "falha ao obter usuário")
	}
	return user, nil
}

// Obtém um usuário pelo email
func (fa *FirebaseAuthentication) GetUserByEmail(email string) (*auth.UserRecord, error) {
	user, err := fa.Client.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, errors.Wrap(err, "falha ao obter usuário por email")
	}
	return user, nil
}

// Remove um usuário do Firebase
func (fa *FirebaseAuthentication) DeleteUser(uid string) error {
	if err := fa.Client.DeleteUser(context.Background(), uid); err != nil {
		return errors.Wrap(err, "falha ao deletar usuário")
	}
	return nil
}

// Atualiza o email de um usuário
func (fa *FirebaseAuthentication) UpdateUserEmail(uid, newEmail string) error {
	user := (&auth.UserToUpdate{}).Email(newEmail)
	_, err := fa.Client.UpdateUser(context.Background(), uid, user)
	return errors.Wrap(err, "falha ao atualizar email")
}

// Atualiza a senha de um usuário
func (fa *FirebaseAuthentication) UpdateUserPassword(uid, newPassword string) error {
	user := (&auth.UserToUpdate{}).Password(newPassword)
	_, err := fa.Client.UpdateUser(context.Background(), uid, user)
	return errors.Wrap(err, "falha ao atualizar senha")
}

// Atualiza o nome de um usuário
func (fa *FirebaseAuthentication) UpdateUserName(uid, newName string) error {
	user := (&auth.UserToUpdate{}).DisplayName(newName)
	_, err := fa.Client.UpdateUser(context.Background(), uid, user)
	return errors.Wrap(err, "falha ao atualizar nome")
}

// Atualiza o telefone de um usuário
func (fa *FirebaseAuthentication) UpdateUserPhone(uid, newPhone string) error {
	user := (&auth.UserToUpdate{}).PhoneNumber(newPhone)
	_, err := fa.Client.UpdateUser(context.Background(), uid, user)
	return errors.Wrap(err, "falha ao atualizar telefone")
}

// Atualiza a foto de um usuário
func (fa *FirebaseAuthentication) UpdateUserPhoto(uid, newPhoto string) error {
	user := (&auth.UserToUpdate{}).PhotoURL(newPhoto)
	_, err := fa.Client.UpdateUser(context.Background(), uid, user)
	return errors.Wrap(err, "falha ao atualizar foto")
}

// Cria um usuário admin padrão
func (fa *FirebaseAuthentication) CreateDefaultAdmin(db *mongo.Database) error {
	ctx := context.Background()
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_INITIAL_PASSWORD")

	// Verificar se o admin já existe no MongoDB
	filter := bson.M{"email": adminEmail, "role": "admin"}
	if count, err := db.Collection("users").CountDocuments(ctx, filter); err != nil {
		return errors.Wrap(err, "falha ao verificar admin existente")
	} else if count > 0 {
		log.Println("usuário admin já existe no MongoDB")
		return nil
	}

	// Verificar se o admin já existe no Firebase
	if _, err := fa.GetUserByEmail(adminEmail); err == nil {
		log.Println("usuário admin já existe no Firebase")
		return nil
	} else if !auth.IsUserNotFound(err) {
		return errors.Wrap(err, "falha ao verificar usuário no Firebase")
	}

	// Criar usuário no Firebase
	params := (&auth.UserToCreate{}).
		Email(adminEmail).
		EmailVerified(true).
		Password(adminPassword).
		DisplayName("Administrator").
		Disabled(false)

	firebaseUser, err := fa.Client.CreateUser(ctx, params)
	if err != nil {
		return errors.Wrap(err, "falha ao criar usuário admin no Firebase")
	}

	// Hash da senha para MongoDB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		_ = fa.DeleteUser(firebaseUser.UID)
		return errors.Wrap(err, "falha ao gerar hash da senha")
	}

	// Criar documento no MongoDB
	adminUser := models.User{
		UID:       firebaseUser.UID,
		Name:      "Administrador Padrão",
		Email:     adminEmail,
		Password:  string(hashedPassword),
		Phone:     "+5511999999999",
		Age:       30,
		Address:   "Endereço Administrativo",
		CPF:       "000.000.000-00",
		RG:        "00.000.000-0",
		Sex:       "male",
		Role:      "admin",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := db.Collection("users").InsertOne(ctx, adminUser); err != nil {
		_ = fa.DeleteUser(firebaseUser.UID)
		return errors.Wrap(err, "falha ao criar usuário admin no MongoDB")
	}

	// Definir claims customizados
	claims := map[string]interface{}{"role": "admin", "status": "active"}
	if err := fa.Client.SetCustomUserClaims(ctx, firebaseUser.UID, claims); err != nil {
		log.Printf("aviso: falha ao definir claims: %v\n", err)
	}

	log.Println("🎫 usuário admin criado com sucesso")
	return nil
}
