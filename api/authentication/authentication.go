package authentication

import (
	"context"
	"fmt"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// Inicializa o Firebase App e o cliente de autenticação.
func InitializeFirebaseAuth() (*auth.Client, error) {
	firebaseConfigPath := os.Getenv("FIREBASE_CREDENTIALS")
	if firebaseConfigPath == "" {
		return nil, fmt.Errorf("variável de ambiente 'FIREBASE_CREDENTIALS' não encontrada")
	}

	if _, err := os.Stat(firebaseConfigPath); os.IsNotExist(err) {
		return nil, err
	}

	// Credenciais do Firebase
	opt := option.WithCredentialsFile(firebaseConfigPath)

	// Inicializa o Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	// Inicializa o cliente de autenticação
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Verifica um token de autenticação do Firebase
func VerifyIDToken(authClient *auth.Client, idToken string) (*auth.Token, error) {
	// Verifica o token de ID
	token, err := authClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Gera um UID do Firebase
func GenerateFirebaseUID(authClient *auth.Client) (string, error) {
	// Gerar um email único baseado em timestamp
	email := fmt.Sprintf("%d@weebie.com", time.Now().UnixNano())

	// Criar um novo usuário
	user := (&auth.UserToCreate{}).
		Email(email).
		Password("senha123") // Senha padrão

	createdUser, err := authClient.CreateUser(context.Background(), user)
	if err != nil {
		return "", err
	}

	return createdUser.UID, nil
}

// Busca um usuário do Firebase pelo UID
func GetFirebaseUser(authClient *auth.Client, uid string) (*auth.UserRecord, error) {
	user, err := authClient.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Busca um usuário do Firebase pelo email
func GetFirebaseUserByEmail(authClient *auth.Client, email string) (*auth.UserRecord, error) {
	user, err := authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Deleta um usuário do Firebase pelo UI
func DeleteFirebaseUser(authClient *auth.Client, uid string) error {
	err := authClient.DeleteUser(context.Background(), uid)
	if err != nil {
		return err
	}
	return nil
}

// Atualiza o email do usuário no Firebase
func UpdateFirebaseUserEmail(authClient *auth.Client, uid string, newEmail string) error {
	user := (&auth.UserToUpdate{}).
		Email(newEmail)

	_, err := authClient.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return err
	}
	return nil
}

// Atualiza a senha do usuário no Firebase
func UpdateFirebaseUserPassword(authClient *auth.Client, uid string, newPassword string) error {
	user := (&auth.UserToUpdate{}).
		Password(newPassword)

	_, err := authClient.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return err
	}
	return nil
}

// Atualiza o nome do usuário no Firebase
func UpdateFirebaseUserName(authClient *auth.Client, uid string, newName string) error {
	user := (&auth.UserToUpdate{}).
		DisplayName(newName)

	_, err := authClient.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return err
	}
	return nil
}

// Atualiza o telefone do usuário no Firebase
func UpdateFirebaseUserPhone(authClient *auth.Client, uid string, newPhone string) error {
	user := (&auth.UserToUpdate{}).
		PhoneNumber(newPhone)

	_, err := authClient.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return err
	}
	return nil
}

// Atualiza o foto do usuário no Firebase
func UpdateFirebaseUserPhoto(authClient *auth.Client, uid string, newPhoto string) error {
	user := (&auth.UserToUpdate{}).
		PhotoURL(newPhoto)

	_, err := authClient.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return err
	}
	return nil
}
