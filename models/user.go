package models

type User struct {
	// Nome do usuário
	Name string `json:"name"`
	// E-mail do usuário
	Email string `json:"email"`
	// Senha do usuário
	Password string `json:"password"`
	// Telefone do usuário
	Phone string `json:"phone"`
	// Idade do usuário
	Age int `json:"age"`
	// Endereço do usuário
	Address string `json:"address"`
	// CPF do usuário
	CPF string `json:"cpf"`
	// RG do usuário
	RG string `json:"rg"`
	// Sexo do usuário (M ou F)
	Sex rune `json:"sex"`
	// Data de criação do usuário
	CreatedAt string `json:"created_at"`
	// Data de atualização do usuário
	UpdatedAt string `json:"updated_at"`
}
