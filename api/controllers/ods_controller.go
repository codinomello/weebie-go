package controllers

import (
	"errors"

	"github.com/codinomello/weebie-go/api/models"
)

type ODSController struct {
	odsList []models.ODS
}

func NewODSController() *ODSController {
	return &ODSController{
		odsList: loadODSData(), // Carrega dados na inicialização
	}
}

// Get all ODS (lógica pura - sem HTTP)
func (c *ODSController) GetAllODS() []models.ODS {
	return c.odsList
}

// Get ODS by number (lógica pura - sem HTTP)
func (c *ODSController) GetODSByNumber(number int) (*models.ODS, error) {
	for _, ods := range c.odsList {
		if ods.Number == number {
			return &ods, nil
		}
	}
	return nil, ErrODSNotFound // Erro customizado
}

// --- Helpers ---
var ErrODSNotFound = errors.New("ODS não encontrada")

func loadODSData() []models.ODS {
	return []models.ODS{
		{
			Number:      1,
			Title:       "Erradicação da Pobreza",
			Description: "Acabar com a pobreza em todas as suas formas, em todos os lugares",
			Emoji:       "❌",
			ImageURL:    "/images/assets/ods/1.jpg",
			Targets: []string{
				"Até 2030, erradicar a pobreza extrema para todas as pessoas em todos os lugares",
				"Implementar sistemas de proteção social adequados",
			},
		},
		{
			Number:      2,
			Title:       "Fome Zero e Agricultura Sustentável",
			Description: "Acabar com a fome, alcançar a segurança alimentar e melhoria da nutrição e promover a agricultura sustentável",
			Emoji:       "🍎",
			ImageURL:    "/images/assets/ods/2.jpg",
			Targets: []string{
				"Até 2030, acabar com a fome e garantir o acesso de todas as pessoas a alimentos seguros, nutritivos e suficientes",
				"Duplicar a produtividade agrícola e a renda dos pequenos produtores de alimentos",
			},
		},
		{
			Number:      3,
			Title:       "Saúde e Bem-Estar",
			Description: "Assegurar uma vida saudável e promover o bem-estar para todos, em todas as idades",
			Emoji:       "🏥",
			ImageURL:    "/images/assets/ods/3.jpg",
			Targets: []string{
				"Reduzir a taxa de mortalidade materna global para menos de 70 mortes por 100.000 nascidos vivos",
				"Acabar com as epidemias de AIDS, tuberculose, malária e doenças tropicais negligenciadas",
			},
		},
		{
			Number:      4,
			Title:       "Educação de Qualidade",
			Description: "Assegurar a educação inclusiva e equitativa e de qualidade, e promover oportunidades de aprendizagem ao longo da vida para todos",
			Emoji:       "📚",
			ImageURL:    "/images/assets/ods/4.jpg",
			Targets: []string{
				"Garantir que todas as meninas e meninos completem o ensino primário e secundário gratuito, equitativo e de qualidade",
				"Aumentar o número de jovens e adultos com habilidades relevantes para emprego e empreendedorismo",
			},
		},
		{
			Number:      5,
			Title:       "Igualdade de Gênero",
			Description: "Alcançar a igualdade de gênero e empoderar todas as mulheres e meninas",
			Emoji:       "♀️",
			ImageURL:    "/images/assets/ods/5.jpg",
			Targets: []string{
				"Acabar com todas as formas de discriminação contra todas as mulheres e meninas em toda parte",
				"Eliminar todas as formas de violência contra todas as mulheres e meninas nas esferas públicas e privadas",
			},
		},
		{
			Number:      6,
			Title:       "Água Potável e Saneamento",
			Description: "Garantir disponibilidade e manejo sustentável da água e saneamento para todos",
			Emoji:       "💧",
			ImageURL:    "/images/assets/ods/6.jpg",
			Targets: []string{
				"Até 2030, alcançar o acesso universal e equitativo a água potável e segura para todos",
				"Melhorar a qualidade da água, reduzindo a poluição e eliminando despejo de produtos químicos perigosos",
			},
		},
		{
			Number:      7,
			Title:       "Energia Limpa e Acessível",
			Description: "Garantir acesso à energia barata, confiável, sustentável e renovável para todos",
			Emoji:       "⚡",
			ImageURL:    "/images/assets/ods/7.jpg",
			Targets: []string{
				"Até 2030, assegurar o acesso universal, confiável, moderno e a preços acessíveis a serviços de energia",
				"Aumentar substancialmente a participação de energias renováveis na matriz energética global",
			},
		},
		{
			Number:      8,
			Title:       "Trabalho Decente e Crescimento Econômico",
			Description: "Promover o crescimento econômico sustentado, inclusivo e sustentável, emprego pleno e produtivo e trabalho decente para todos",
			Emoji:       "💼",
			ImageURL:    "/images/assets/ods/8.jpg",
			Targets: []string{
				"Sustentar o crescimento econômico per capita de acordo com as circunstâncias nacionais",
				"Promover políticas orientadas para o desenvolvimento que apoiem as atividades produtivas",
			},
		},
		{
			Number:      9,
			Title:       "Indústria, Inovação e Infraestrutura",
			Description: "Construir infraestruturas resilientes, promover a industrialização inclusiva e sustentável e fomentar a inovação",
			Emoji:       "🏗️",
			ImageURL:    "/images/assets/ods/9.jpg",
			Targets: []string{
				"Desenvolver infraestrutura de qualidade, confiável, sustentável e resiliente",
				"Promover a industrialização inclusiva e sustentável e aumentar a participação da indústria no PIB",
			},
		},
		{
			Number:      10,
			Title:       "Redução das Desigualdades",
			Description: "Reduzir a desigualdade dentro dos países e entre eles",
			Emoji:       "⚖️",
			ImageURL:    "/images/assets/ods/10.jpg",
			Targets: []string{
				"Até 2030, progressivamente alcançar e sustentar o crescimento da renda dos 40% da população mais pobre",
				"Garantir a igualdade de oportunidades e reduzir desigualdades de resultados",
			},
		},
		{
			Number:      11,
			Title:       "Cidades e Comunidades Sustentáveis",
			Description: "Tornar as cidades e os assentamentos humanos inclusivos, seguros, resilientes e sustentáveis",
			Emoji:       "🏙️",
			ImageURL:    "/images/assets/ods/11.jpg",
			Targets: []string{
				"Até 2030, garantir o acesso de todos à habitação segura, adequada e a preço acessível",
				"Proporcionar acesso a sistemas de transporte seguros, acessíveis e sustentáveis para todos",
			},
		},
		{
			Number:      12,
			Title:       "Consumo e Produção Responsáveis",
			Description: "Assegurar padrões de produção e de consumo sustentáveis",
			Emoji:       "🔄",
			ImageURL:    "/images/assets/ods/12.jpg",
			Targets: []string{
				"Implementar o Plano Decenal de Programas sobre Produção e Consumo Sustentáveis",
				"Até 2030, reduzir pela metade o desperdício de alimentos per capita mundial",
			},
		},
		{
			Number:      13,
			Title:       "Ação Contra a Mudança Global do Clima",
			Description: "Tomar medidas urgentes para combater a mudança climática e seus impactos",
			Emoji:       "🌍",
			ImageURL:    "/images/assets/ods/13.jpg",
			Targets: []string{
				"Reforçar a resiliência e a capacidade de adaptação a riscos relacionados ao clima",
				"Integrar medidas da mudança do clima nas políticas, estratégias e planejamentos nacionais",
			},
		},
		{
			Number:      14,
			Title:       "Vida na Água",
			Description: "Conservação e uso sustentável dos oceanos, dos mares e dos recursos marinhos",
			Emoji:       "🐠",
			ImageURL:    "/images/assets/ods/14.jpg",
			Targets: []string{
				"Até 2025, prevenir e reduzir significativamente a poluição marinha de todos os tipos",
				"Gerir de forma sustentável e proteger os ecossistemas marinhos e costeiros",
			},
		},
		{
			Number:      15,
			Title:       "Vida Terrestre",
			Description: "Proteger, recuperar e promover o uso sustentável dos ecossistemas terrestres",
			Emoji:       "🌳",
			ImageURL:    "/images/assets/ods/15.jpg",
			Targets: []string{
				"Até 2020, assegurar a conservação, recuperação e uso sustentável de ecossistemas terrestres",
				"Promover a implementação da gestão sustentável de todos os tipos de florestas",
			},
		},
		{
			Number:      16,
			Title:       "Paz, Justiça e Instituições Eficazes",
			Description: "Promover sociedades pacíficas e inclusivas para o desenvolvimento sustentável",
			Emoji:       "🕊️",
			ImageURL:    "/images/assets/ods/16.jpg",
			Targets: []string{
				"Reduzir significativamente todas as formas de violência e as taxas de mortalidade relacionada",
				"Promover o Estado de Direito em nível nacional e internacional",
			},
		},
		{
			Number:      17,
			Title:       "Parcerias e Meios de Implementação",
			Description: "Fortalecer os meios de implementação e revitalizar a parceria global para o desenvolvimento sustentável",
			Emoji:       "🤝",
			ImageURL:    "/images/assets/ods/17.jpg",
			Targets: []string{
				"Fortalecer a mobilização de recursos internos para melhorar a capacidade nacional",
				"Aumentar o apoio internacional para a capacitação eficaz e direcionada nos países em desenvolvimento",
			},
		},
	}
}
