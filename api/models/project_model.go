package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project representa a estrutura de um projeto no banco de dados.
type Project struct {
	// Identificação
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`    // ID do projeto (gerado pelo MongoDB)
	OwnerUID primitive.ObjectID `bson:"owner_uid" json:"owner_uid"` // ObjectID do dono do projeto
	Members  []Member           `bson:"members" json:"members"`     // Array de ObjectIDs dos membros do projeto

	// Informações Básicas
	Title          string `bson:"title" json:"title"`                     // Título do projeto
	Year           int    `bson:"year" json:"year"`                       // Ano de criação/desenvolvimento do projeto
	Location       string `bson:"location" json:"location"`               // Localização onde o projeto será realizado
	TargetAudience string `bson:"target_audience" json:"target_audience"` // Público-alvo do projeto

	// Detalhes do Projeto
	Details     string `bson:"details" json:"details"`         // Descrição detalhada do projeto
	Impact      string `bson:"impact" json:"impact"`           // Impacto esperado do projeto
	Methodology string `bson:"methodology" json:"methodology"` // Metodologia de realização do projeto
	Icon        string `bson:"icon" json:"icon"`               // Ícone ou emoji para o projeto
	Status      string `bson:"status" json:"status"`           // Status do projeto (ex: "active", "completed", "archived")

	// Classificação
	Course      string `bson:"course" json:"course"`             // Curso relacionado ao projeto
	OtherCourse string `bson:"other_course" json:"other_course"` // Especificação do curso se "Outro" for selecionado
	Area        string `bson:"area" json:"area"`                 // Área de atuação do projeto
	Complexity  string `bson:"complexity" json:"complexity"`     // Nível de complexidade (low, medium, high)

	// ODS (Objetivos de Desenvolvimento Sustentável)
	ODS []string `bson:"ods" json:"ods"` // Array de números dos ODS relacionados

	// Recursos Necessários
	Budget       float64 `bson:"budget" json:"budget"`             // Orçamento estimado em reais
	Materials    string  `bson:"materials" json:"materials"`       // Recursos materiais necessários
	Partnerships string  `bson:"partnerships" json:"partnerships"` // Parcerias existentes ou potenciais

	// Termos e Condições
	TermsAccepted bool `bson:"terms" json:"terms"`             // Se o usuário aceitou os termos de uso
	PublicData    bool `bson:"public_data" json:"public_data"` // Se o usuário concordou em tornar os dados públicos

	// Metadados
	CreatedAt time.Time `bson:"created_at" json:"created_at"` // Data e hora de criação do projeto
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"` // Data e hora da última atualização do projeto
}

// CreateProjectRequest define a estrutura da requisição para criar um novo projeto.
type CreateProjectRequest struct {
	// Informações Básicas
	Title          string `json:"title" bson:"title"`
	Year           int    `json:"year" bson:"year"`
	Location       string `json:"location" bson:"location"`
	TargetAudience string `json:"target_audience" bson:"target_audience"`

	// Detalhes do Projeto
	Details     string `json:"details" bson:"details"`
	Impact      string `json:"impact" bson:"impact"`
	Methodology string `json:"methodology" bson:"methodology"`
	Icon        string `json:"icon" bson:"icon"`
	Status      string `json:"status" bson:"status"`

	// Classificação
	Course     string `json:"course" bson:"course"`
	Area       string `json:"area" bson:"area"`
	Complexity string `json:"complexity" bson:"complexity"`

	// ODS (Objetivos de Desenvolvimento Sustentável)
	ODS []string `json:"ods" bson:"ods"`

	// Recursos
	Budget       float64 `json:"budget" bson:"budget"`
	Materials    string  `json:"materials" bson:"materials"`
	Partnerships string  `json:"partnerships" bson:"partnerships"`

	// Termos
	TermsAccepted bool `json:"terms_accepted" bson:"terms_accepted"`
	PublicData    bool `json:"public_data" bson:"public_data"`

	// Membros (emails que serão processados para adicionar como membros)
	MemberEmails []string `json:"member_emails,omitempty" bson:"member_emails,omitempty"`
}

// ProjectUpdate representa os campos que podem ser atualizados em um projeto.
type UpdateProjectRequest struct {
	// Identificação
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`              // ID do projeto (gerado pelo MongoDB)
	OwnerUID       primitive.ObjectID `json:"owner_uid,omitempty" bson:"owner_uid,omitempty"` // ObjectID do dono do projeto
	Members        []Member           `json:"members,omitempty" bson:"members,omitempty"`     // Array de ObjectIDs dos membros do projeto
	Title          *string            `json:"title,omitempty" bson:"title,omitempty"`
	Year           *int               `json:"year,omitempty" bson:"year,omitempty"`
	Location       *string            `json:"location,omitempty" bson:"location,omitempty"`
	TargetAudience *string            `json:"target_audience,omitempty" bson:"target_audience,omitempty"`

	// Detalhes do Projeto
	Details     *string `json:"details,omitempty" bson:"details,omitempty"`
	Impact      *string `json:"impact,omitempty" bson:"impact,omitempty"`
	Methodology *string `json:"methodology,omitempty" bson:"methodology,omitempty"`
	Icon        *string `json:"icon,omitempty" bson:"icon,omitempty"`
	Status      *string `json:"status,omitempty" bson:"status,omitempty"`

	// Classificação
	Course      *string `json:"course,omitempty" bson:"course,omitempty"`
	OtherCourse *string `json:"other_course,omitempty" bson:"other_course,omitempty"`
	Area        *string `json:"area,omitempty" bson:"area,omitempty"`
	Complexity  *string `json:"complexity,omitempty" bson:"complexity,omitempty"`

	// ODS (Objetivos de Desenvolvimento Sustentável)
	ODS []string `json:"ods,omitempty" bson:"ods,omitempty"`

	// Metadados
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`

	// Recursos
	Budget       *float64 `json:"budget,omitempty" bson:"budget,omitempty"`
	Materials    *string  `json:"materials,omitempty" bson:"materials,omitempty"`
	Partnerships *string  `json:"partnerships,omitempty" bson:"partnerships,omitempty"`

	// Termos
	TermsAccepted *bool `json:"terms_accepted,omitempty" bson:"terms_accepted,omitempty"`
	PublicData    *bool `json:"public_data,omitempty" bson:"public_data,omitempty"`

	// Membros (emails que serão processados para adicionar como membros)
	MemberEmails []string `json:"member_emails,omitempty" bson:"member_emails,omitempty"`
}
