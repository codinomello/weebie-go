package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// Dados básicos
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"` // ID do usuário
	Name  string             `bson:"name" json:"name"`        // Nome do usuário
	Email string             `bson:"email" json:"email"`      // E-mail do usuário
	Phone string             `bson:"phone" json:"phone"`      // Telefone do usuário
	Age   int                `bson:"age" json:"age"`          // Idade do usuário
	CPF   string             `bson:"cpf" json:"cpf"`          // CPF do usuário
	RG    string             `bson:"rg" json:"rg"`            // RG do usuário
	Sex   string             `bson:"sex" json:"sex"`          // Sexo do usuário ('M' ou 'F')

	// Endereço
	CEP          string `bson:"cep" json:"cep"`                   // CEP do usuário
	Address      string `bson:"address" json:"address"`           // Endereço do usuário
	Number       string `bson:"number" json:"number"`             // Número do endereço
	Complement   string `bson:"complement" json:"complement"`     // Complemento do endereço
	Neighborhood string `bson:"neighborhood" json:"neighborhood"` // Bairro do usuário
	City         string `bson:"city" json:"city"`                 // Cidade do usuário
	State        string `bson:"state" json:"state"`               // Estado do usuário

	// Autenticação
	Provider string `json:"provider" binding:"required"`
	Password string `bson:"password" json:"password"`    // Senha do usuário
	IDToken  string `bson:"-" json:"id_token,omitempty"` // ID do token do usuário
	UID      string `bson:"uid" json:"uid"`              // UID do usuário
	Role     string `bson:"role" json:"role"`            // Papel do usuário (ex: admin, user)
	Status   string `bson:"status" json:"status"`        // Status do usuário (ex: active, inactive)

	// Informações de tempo
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`                     // Data de criação do usuário
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`                     // Data de atualização do usuário
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"` // Data de exclusão do usuário (nil se não excluído)

	// Aceito os termos de uso
	Terms bool `bson:"terms" json:"terms"` // Aceito os termos de uso
}

type CreateUserRequest struct {
	// Dados básicos
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
	CPF   string `json:"cpf" binding:"required"`
	RG    string `json:"rg" binding:"required"`
	Sex   string `json:"sex"`

	// Endereço
	CEP          string `json:"cep"`
	Address      string `json:"address"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`

	// Autenticação
	Provider string `json:"provider" binding:"required"`
	Password string `json:"password" binding:"required"`
	IDToken  string `json:"id_token" binding:"required"`
	UID      string `json:"uid"`
	Role     string `json:"role"`
	Status   string `json:"status"`

	// Informações de tempo
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`

	// Aceito os termos de uso
	Terms bool `json:"terms" binding:"required"`
}

// UserUpdateRequest representa os dados que podem ser atualizados
type UserUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Age      int    `json:"age,omitempty"`
	Address  string `json:"address,omitempty"`
	CPF      string `json:"cpf,omitempty"`
	RG       string `json:"rg,omitempty"`
	Sex      rune   `json:"sex,omitempty"`
	Role     string `json:"role,omitempty"`
	Status   string `json:"status,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id"`
	UID       string             `json:"uid"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone,omitempty"`
	Age       int                `json:"age,omitempty"`
	Address   string             `json:"address,omitempty"`
	CPF       string             `json:"cpf,omitempty"`
	RG        string             `json:"rg,omitempty"`
	Sex       string             `json:"sex,omitempty"`
	Role      string             `json:"role"`
	Status    string             `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// Converte User para UserResponse removendo dados sensíveis
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		UID:       u.UID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Age:       u.Age,
		Address:   u.Address,
		CPF:       u.CPF,
		RG:        u.RG,
		Sex:       u.Sex,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// Verifica se os dados do usuário são válidos
func (u *User) IsValid() bool {
	return u.UID != "" && u.Email != "" && u.Name != ""
}

// Define valores padrão para o usuário
func (u *User) SetDefaults() {
	if u.Role == "" {
		u.Role = "user"
	}
	if u.Status == "" {
		u.Status = "active"
	}
}
