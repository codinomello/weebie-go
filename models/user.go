package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	// ID do usuário
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// Nome do usuário
	Name string `bson:"name" json:"name"`
	// E-mail do usuário
	Email string `bson:"email" json:"email"`
	// Senha do usuário
	Password string `bson:"password" json:"password"`
	// Telefone do usuário
	Phone string `bson:"phone" json:"phone"`
	// Idade do usuário
	Age int `bson:"age" json:"age"`
	// Endereço do usuário
	Address string `bson:"address" json:"address"`
	// CPF do usuário
	CPF string `bson:"cpf" json:"cpf"`
	// RG do usuário
	RG string `bson:"rg" json:"rg"`
	// Sexo do usuário (M ou F)
	Sex rune `bson:"sex" json:"sex"`
	// Data de criação do usuário
	CreatedAt string `bson:"created_at" json:"created_at"`
	// Data de atualização do usuário
	UpdatedAt string `bson:"updated_at" json:"updated_at"`
}
