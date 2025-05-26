package models

// estrutura de dados de um ODS para o endpoint /api/ods
type ODS struct {
	Number      int      `json:"number"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	ImageURL    string   `json:"image_url"`
	Targets     []string `json:"targets"`
}
