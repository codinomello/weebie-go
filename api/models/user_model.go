package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`                          // ID do usuário
	UID       string             `bson:"uid" json:"uid"`                                   // UID do usuário
	Name      string             `bson:"name" json:"name"`                                 // Nome do usuário
	Email     string             `bson:"email" json:"email"`                               // E-mail do usuário
	Password  string             `bson:"password" json:"password"`                         // Senha do usuário
	Phone     string             `bson:"phone" json:"phone"`                               // Telefone do usuário
	Age       int                `bson:"age" json:"age"`                                   // Idade do usuário
	Address   string             `bson:"address" json:"address"`                           // Endereço do usuário
	CPF       string             `bson:"cpf" json:"cpf"`                                   // CPF do usuário
	RG        string             `bson:"rg" json:"rg"`                                     // RG do usuário
	Sex       rune               `bson:"sex" json:"sex"`                                   // Sexo do usuário ('M' ou 'F')
	Role      string             `bson:"role" json:"role"`                                 // Papel do usuário (ex: admin, user)
	Status    string             `bson:"status" json:"status"`                             // Status do usuário (ex: active, inactive)
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`                     // Data de criação do usuário
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`                     // Data de atualização do usuário
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"` // Data de exclusão do usuário (opcional)
}
