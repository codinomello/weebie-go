// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package public

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import components "github.com/codinomello/weebie-go/templates/components"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"pt-BR\"><head><!-- Meta Tags -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Meta().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!-- Título -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Title("Weebie - Projetos Comunitários").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- Ícone --><link rel=\"icon\" type=\"image/x-icon\" href=\"/hou\"><!-- HTMX e Tailwind CSS -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ScriptHead().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<!-- Flowbite CSS e Ícone -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Link().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<!-- Fonte -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.StyleFont().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</head><body class=\"dark:bg-gray-900\"><!-- Header Moderno --><nav class=\"bg-gradient-to-r from-golden to-orange dark:from-gray-800 dark:to-gray-700 px-4 py-3 shadow-lg\"><div class=\"container mx-auto flex items-center justify-between\"><!-- Logo e Nome --><div class=\"flex items-center space-x-3\"><span class=\"text-3xl\">🏡</span><h1 class=\"text-2xl font-bold text-white dark:text-gray-100\">Weebie</h1></div><!-- Links de Navegação --><div class=\"hidden md:flex items-center space-x-8\"><a href=\"index.html\" class=\"text-white hover:text-red hover:underline hover:underline-red font-medium transition-all duration-300\">Início</a> <a href=\"project.html\" class=\"text-white hover:text-green hover:underline hover:underline-green font-medium transition-all duration-300\">Projetos</a> <a href=\"about.html\" class=\"text-white hover:text-blue hover:underline hover:underline-blue font-medium transition-all duration-300\">Sobre</a></div><!-- Botões de Controle --><div class=\"flex items-center space-x-6\"><!-- Botão de Tema --><button id=\"theme-toggle\" class=\"theme-toggle\"><span class=\"sr-only\">Alternar Modo Escuro</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.StyleToggleTheme().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</button><!-- Ícone de Conta --><button class=\"p-2 rounded-full text-white hover:35 transition-colors duration-300\"><svg class=\"w-7 h-7 text-white dark:golden\" viewBox=\"0 0 24 24\" fill=\"currentColor\"><path d=\"M12 4a4 4 0 014 4 4 4 0 01-4 4 4 4 0 01-4-4 4 4 0 014-4m0 10c4.42 0 8 1.79 8 4v2H4v-2c0-2.21 3.58-4 8-4z\"></path></svg></button><!-- Menu Mobile --><button class=\"md:hidden text-white p-2 rounded-lg hover:10 transition-colors\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button></div></div></nav><!-- Hero Section --><section class=\"container mx-auto px-4 py-16 text-center\"><div class=\"max-w-2xl mx-auto\"><h2 class=\"text-4xl md:text-5xl font-bold text-orange dark:text-golden mb-6\"><span class=\"text-blue\">👋</span> Comunidades que Inspiram e Transformam!</h2><p class=\"text-lg text-gray-600 dark:text-gray-300 mb-8\">Conectando boas ideias a pessoas que fazem acontecer</p><div class=\"flex flex-col sm:flex-row justify-center gap-4\"><button class=\"bg-red text-white text-bold px-6 py-3 rounded-lg hover:bg-red/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105\"><span>💼</span> Seus Projetos</button> <button class=\"bg-green text-white text-bold px-6 py-3 rounded-lg hover:bg-green/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105\"><span>🚀</span> Iniciar Projeto</button> <button class=\"bg-blue text-white text-bold px-6 py-3 rounded-lg hover:bg-blue/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105\"><span>🔍</span> Explorar</button></div></div></section><!-- Projetos em Destaque --><section class=\"bg-gray-50 dark:bg-gray-800 py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\"><span class=\"text-pink\">🌟</span> Projetos Destacados</h3><div class=\"grid md:grid-cols-3 gap-6\"><!-- Card Projeto 1 --><div class=\"dark:bg-gray-700 rounded-xl shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105 p-6 border-l-4 border-green\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">🌻 Horta Urbana</h4><span class=\"bg-green/10 dark:bg-green/20 text-green dark:text-green/80 px-3 py-1 rounded-full text-sm\">Ativo</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Cultivo coletivo de alimentos orgânicos no centro da cidade</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-blue rounded-full w-3/4\"></div></div><span class=\"text-blue\">75%</span></div><button class=\"w-full text-green hover:bg-green/10 py-2 rounded-lg transition\">Apoiar Projeto →</button></div></div><!-- Card Projeto 2 --><div class=\"dark:bg-gray-700 rounded-xl shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105 p-6 border-l-4 border-blue\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">📚 Biblioteca Livre</h4><span class=\"bg-blue/10 dark:bg-blue/20 text-blue dark:text-blue/80 px-3 py-1 rounded-full text-sm\">Novo</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Pontos de compartilhamento de livros em espaços públicos</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-golden rounded-full w-2/5\"></div></div><span class=\"text-orange\">40%</span></div><button class=\"w-full text-blue hover:bg-blue/10 py-2 rounded-lg transition\">Apoiar Projeto →</button></div></div><!-- Card Projeto 3 --><div class=\"dark:bg-gray-700 rounded-xl shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105 p-6 border-l-4 border-pink\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">🎨 Arte na Rua</h4><span class=\"bg-pink/10 dark:bg-pink/20 text-pink dark:text-pink/80 px-3 py-1 rounded-full text-sm\">Concluído</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Oficinas de arte urbana para jovens da comunidade de Guarulhos</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-golden rounded-full w-full\"></div></div><span class=\"text-green\">100%</span></div><button class=\"w-full text-pink hover:bg-pink/10 py-2 rounded-lg transition\">Ver Resultados →</button></div></div></div></div></section><!-- Nossa Proposta --><section class=\"py-16\"><div class=\"container mx-auto px-4\"><div class=\"max-w-4xl mx-auto text-center\"><h3 class=\"text-3xl font-bold mb-8 text-orange dark:text-golden\"><span class=\"text-green\">❤️</span> Como Funcionamos</h3><div class=\"grid md:grid-cols-2 gap-8 text-left\"><div class=\"dark:bg-gray-700 p-6 rounded-xl border-t-4 border-b-4 border-copper shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl mb-4\">🤝</div><h4 class=\"text-xl font-semibold mb-3 dark:text-gray-200\">Conectamos</h4><p class=\"text-gray-600 dark:text-gray-400\">Unimos pessoas com ideias inovadoras a voluntários e apoiadores dispostos a colaborar</p></div><div class=\"dark:bg-gray-700 p-6 rounded-xl border-t-4 border-b-4 border-silver shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl mb-4\">🛠️</div><h4 class=\"text-xl font-semibold mb-3 dark:text-gray-200\">Implementamos</h4><p class=\"text-gray-600 dark:text-gray-400\">Oferecemos ferramentas e suporte para tirar projetos do papel</p></div></div><div class=\"mt-12 dark:bg-gray-700 shadow-lg rounded-xl p-8 border-t-4 border-teal shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><h4 class=\"text-2xl font-bold text-orange dark:text-golden mb-4\">Nossos Pilares</h4><ul class=\"space-y-4 text-gray-600 dark:text-gray-400 text-left\"><li class=\"flex items-center gap-3\"><span class=\"text-green\">✓</span> Transparência total nos processos</li><li class=\"flex items-center gap-3\"><span class=\"text-blue\">✓</span> Apoio comunitário sustentável</li><li class=\"flex items-center gap-3\"><span class=\"text-pink\">✓</span> Impacto social mensurável</li></ul></div></div></div></section><!-- Estatísticas --><section class=\"bg-gray-50 dark:bg-gray-800 py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">📈 Nosso Impacto</h3><div class=\"grid md:grid-cols-3 gap-8\"><div class=\"dark:bg-gray-700 p-6 rounded-xl text-center border-b-4 border-red shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl text-green dark:text-golden\">120+</div><p class=\"text-gray-600 dark:text-gray-300\">Projetos Ativos</p><span class=\"text-sm text-orange dark:text-red\">🏆 Em 15 cidades</span></div><div class=\"dark:bg-gray-700 p-6 rounded-xl text-center border-b-4 border-green shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl text-blue dark:text-golden\">5.8K</div><p class=\"text-gray-600 dark:text-gray-300\">Voluntários</p><span class=\"text-sm text-orange dark:text-green\">🤝 Comunidade ativa</span></div><div class=\"dark:bg-gray-700 p-6 rounded-xl text-center border-b-4 border-blue shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl text-pink dark:text-golden\">92%</div><p class=\"text-gray-600 dark:text-gray-300\">Sucesso</p><span class=\"text-sm text-orange dark:text-blue\">🎯 Metas alcançadas</span></div></div></div></section><!-- Depoimentos --><section class=\"py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">🗣️ Quem Apoia</h3><div class=\"grid md:grid-cols-2 gap-8\"><div class=\"dark:bg-gray-700 p-6 rounded-xl border-r-4 border-lilac shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"flex items-center gap-4 mb-4\"><span class=\"text-3xl\">👩🌾</span><div><h4 class=\"font-semibold dark:text-gray-200\">Maria Silva</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">Voluntária na Horta Comunitária</p></div></div><p class=\"text-gray-600 dark:text-gray-300\">\"A Weebie nos deu estrutura para transformar um terreno abandonado em fonte de alimentos para toda comunidade!\"</p></div><div class=\"dark:bg-gray-700 p-6 rounded-xl border-r-4 border-orange shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"flex items-center gap-4 mb-4\"><span class=\"text-3xl\">👨🏫</span><div><h4 class=\"font-semibold dark:text-gray-200\">Carlos Mendes</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">Coordenador de Projetos</p></div></div><p class=\"text-gray-600 dark:text-gray-300\">\"A plataforma simplificou nossa captação de recursos e organização de voluntários de maneira incrível.\"</p></div></div></div></section><!-- FAQ --><section class=\"bg-gray-50 dark:bg-gray-800 py-16\"><div class=\"container mx-auto px-4 max-w-3xl\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">❓ Perguntas Frequentes</h3><div class=\"space-y-4\"><!-- Dropdown 1 --><div class=\"dark:bg-gray-700 rounded-xl p-4 border-l-4 border-silver shadow-lg\"><button id=\"dropdown-button-1\" data-collapse-toggle=\"dropdown-content-1\" class=\"flex justify-between items-center w-full\"><span class=\"font-semibold text-left dark:text-gray-200\">Como participar de um projeto?</span> <span class=\"text-silver dark:text-silver\">+</span></button><div id=\"dropdown-content-1\" class=\"mt-2 text-gray-600 dark:text-gray-400 hidden\">Escolha um projeto em nossa plataforma, clique em \"Participar\" e siga as instruções. É gratuito e aberto a todos!</div></div><!-- Adicione mais perguntas seguindo o mesmo padrão --></div></div></section><!-- Eventos --><section class=\"py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">📅 Próximos Eventos</h3><div class=\"grid md:grid-cols-2 gap-6\"><div class=\"dark:bg-gray-700 p-6 rounded-xl border-l-4 border-teal shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"flex items-center gap-4\"><div class=\"text-center\"><div class=\"text-2xl font-bold text-orange dark:text-golden\">25</div><div class=\"text-sm dark:text-gray-300\">Maio</div></div><div><h4 class=\"font-semibold dark:text-gray-200\">Motratech</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">📌 Colégio Eniac - 16:50h</p></div></div></div><div class=\"dark:bg-gray-700 p-6 rounded-xl border-l-4 border-brown shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"flex items-center gap-4\"><div class=\"text-center\"><div class=\"text-2xl font-bold text-orange dark:text-golden\">12</div><div class=\"text-sm dark:text-gray-300\">Abril</div></div><div><h4 class=\"font-semibold dark:text-gray-200\">Workshop de Compostagem</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">📌 Centro Comunitário - 14:20h</p></div></div></div><!-- Adicione mais eventos seguindo o mesmo padrão --></div></div></section><!-- Footer -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.FooterIndex().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<!-- Scripts -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ScriptBody().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
