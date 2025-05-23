package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`                          // ID do usuário
	IDToken      string             `bson:"-" json:"id_token,omitempty"`                      // ID do token do usuário
	UID          string             `bson:"uid" json:"uid"`                                   // UID do usuário
	Name         string             `bson:"name" json:"name"`                                 // Nome do usuário
	Email        string             `bson:"email" json:"email"`                               // E-mail do usuário
	Password     string             `bson:"password" json:"password"`                         // Senha do usuário
	Phone        string             `bson:"phone" json:"phone"`                               // Telefone do usuário
	CPF          string             `bson:"cpf" json:"cpf"`                                   // CPF do usuário
	RG           string             `bson:"rg" json:"rg"`                                     // RG do usuário
	Sex          rune               `bson:"sex" json:"sex"`                                   // Sexo do usuário ('M' ou 'F')
	Age          int                `bson:"age" json:"age"`                                   // Idade do usuário
	Address      string             `bson:"address" json:"address"`                           // Endereço do usuário
	Complement   string             `bson:"complement" json:"complement"`                     // Complemento do endereço
	Number       string             `bson:"number" json:"number"`                             // Número do endereço
	Neighborhood string             `bson:"neighborhood" json:"neighborhood"`                 // Bairro do usuário
	City         string             `bson:"city" json:"city"`                                 // Cidade do usuário
	State        string             `bson:"state" json:"state"`                               // Estado do usuário
	CEP          string             `bson:"cep" json:"cep"`                                   // CEP do usuário
	Terms        bool               `bson:"terms" json:"terms"`                               // Aceito os termos de uso
	Role         string             `bson:"role" json:"role"`                                 // Papel do usuário (ex: admin, user)
	Status       string             `bson:"status" json:"status"`                             // Status do usuário (ex: active, inactive)
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`                     // Data de criação do usuário
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`                     // Data de atualização do usuário
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"` // Data de exclusão do usuário (nil se não excluído)
}

type UserCreateRequest struct {
	IDToken string `json:"id_token" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone,omitempty"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	CPF     string `json:"cpf,omitempty"`
	RG      string `json:"rg,omitempty"`
	Sex     rune   `json:"sex,omitempty"`
	Role    string `json:"role,omitempty"`
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
	Sex       rune               `json:"sex,omitempty"`
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
