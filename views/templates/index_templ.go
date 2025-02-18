// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"pt-BR\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Weebie - Projetos Comunitários</title><link rel=\"icon\" type=\"image/x-icon\" href=\"../images/house.png\"><!-- Tailwind CSS --><script src=\"https://cdn.tailwindcss.com\"></script><!-- Flowbite CSS --><link href=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/1.6.5/flowbite.min.css\" rel=\"stylesheet\"><!-- HTMX --><script src=\"https://unpkg.com/htmx.org@1.9.5\"></script><style>\r\n        @import url('https://fonts.cdnfonts.com/css/euclid-circular-a');\r\n        * {\r\n            font-family: \"Euclid Circular A\";\r\n        }\r\n    </style><script>\r\n        tailwind.config = {\r\n          darkMode: 'class',\r\n          theme: {\r\n            extend: {\r\n              colors: {\r\n                red: '#FF0000',    // Vermelho\r\n                green: '#32CD32',  // Verde\r\n                blue: '#00BFFF',   // Azul\r\n                lilac: '#C8A2C8',  // Lilás\r\n                pink: '#FF1493',   // Rosa\r\n                orange: '#FF6F00', // Laranja\r\n                golden: '#FFD700'  // Ouro\r\n              }\r\n            }\r\n          }\r\n        }\r\n    </script></head><body class=\"bg-white dark:bg-gray-900\"><!-- Header Moderno --><nav class=\"bg-gradient-to-r from-orange to-golden dark:from-gray-800 dark:to-gray-700 px-4 py-3 shadow-lg\"><div class=\"container mx-auto flex items-center justify-between\"><!-- Logo e Nome --><div class=\"flex items-center space-x-3\"><span class=\"text-3xl\">🏘️</span><h1 class=\"text-2xl font-bold text-white dark:text-gray-100\">Weebie</h1></div><!-- Links de Navegação --><div class=\"hidden md:flex items-center space-x-8\"><a href=\"#\" class=\"text-white hover:text-golden dark:hover:text-blue transition-colors font-medium\">Início</a> <a href=\"#\" class=\"text-white hover:text-golden dark:hover:text-blue transition-colors font-medium\">Projetos</a> <a href=\"#\" class=\"text-white hover:text-golden dark:hover:text-blue transition-colors font-medium\">Sobre</a></div><!-- Botão de Tema e Menu Mobile --><div class=\"flex items-center space-x-4\"><!-- Botão de Tema Moderno --><button id=\"theme-toggle\" class=\"theme-toggle\"><span class=\"sr-only\">Alternar Modo Escuro</span><style>\r\n                        .theme-toggle {\r\n                            position: relative;\r\n                            width: 48px;\r\n                            height: 24px;\r\n                            border-radius: 9999px;\r\n                            background-color: rgba(255, 255, 255, 0.1);\r\n                            transition: background-color 0.3s;\r\n                        }\r\n                        .theme-toggle:hover {\r\n                            background-color: rgba(255, 255, 255, 0.2);\r\n                        }\r\n                        .theme-toggle::before {\r\n                            content: '';\r\n                            position: absolute;\r\n                            top: 2px;\r\n                            left: 2px;\r\n                            width: 20px;\r\n                            height: 20px;\r\n                            border-radius: 50%;\r\n                            background-color: white;\r\n                            transition: transform 0.3s, background-color 0.3s;\r\n                        }\r\n                        .dark .theme-toggle::before {\r\n                            transform: translateX(24px);\r\n                            background-color: #FFD700;\r\n                        }\r\n                    </style></button><!-- Menu Mobile --><button class=\"md:hidden text-white p-2 rounded-lg hover:bg-white/10 transition-colors\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button></div></div></nav><!-- Hero Section --><section class=\"container mx-auto px-4 py-16 text-center\"><div class=\"max-w-2xl mx-auto\"><h2 class=\"text-4xl md:text-5xl font-bold text-orange dark:text-golden mb-6\"><span class=\"text-blue\">👋</span> Construindo Comunidades Melhores Juntos</h2><p class=\"text-lg text-gray-600 dark:text-gray-300 mb-8\">Conectando boas ideias a pessoas que fazem acontecer</p><div class=\"flex flex-col sm:flex-row justify-center gap-4\"><button class=\"bg-green text-white px-6 py-3 rounded-lg hover:bg-[#2AA32A] transition flex items-center gap-2\"><span>🚀</span> Iniciar Projeto</button> <button class=\"bg-blue text-white px-6 py-3 rounded-lg hover:bg-[#009FCC] transition flex items-center gap-2\"><span>🔍</span> Explorar</button></div></div></section><!-- Projetos em Destaque --><section class=\"bg-gray-50 dark:bg-gray-800 py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\"><span class=\"text-pink\">🌟</span> Projetos Destacados</h3><div class=\"grid md:grid-cols-3 gap-6\"><!-- Card Projeto 1 --><div class=\"bg-white dark:bg-gray-700 rounded-xl shadow-sm hover:shadow-md transition p-6 border-l-4 border-green\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">🌻 Horta Urbana</h4><span class=\"bg-green/10 dark:bg-green/20 text-green dark:text-green/80 px-3 py-1 rounded-full text-sm\">Ativo</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Cultivo coletivo de alimentos orgânicos no centro da cidade</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-golden rounded-full w-3/4\"></div></div><span class=\"text-blue\">75%</span></div><button class=\"w-full text-pink hover:bg-pink/10 py-2 rounded-lg transition\">Apoiar Projeto →</button></div></div><!-- Card Projeto 2 --><div class=\"bg-white dark:bg-gray-700 rounded-xl shadow-sm hover:shadow-md transition p-6 border-l-4 border-blue\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">📚 Biblioteca Livre</h4><span class=\"bg-blue/10 dark:bg-blue/20 text-blue dark:text-blue/80 px-3 py-1 rounded-full text-sm\">Novo</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Pontos de compartilhamento de livros em espaços públicos</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-golden rounded-full w-2/5\"></div></div><span class=\"text-blue\">40%</span></div><button class=\"w-full text-pink hover:bg-pink/10 py-2 rounded-lg transition\">Apoiar Projeto →</button></div></div><!-- Card Projeto 3 --><div class=\"bg-white dark:bg-gray-700 rounded-xl shadow-sm hover:shadow-md transition p-6 border-l-4 border-pink\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">🎨 Arte na Rua</h4><span class=\"bg-pink/10 dark:bg-pink/20 text-pink dark:text-pink/80 px-3 py-1 rounded-full text-sm\">Concluído</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Oficinas de arte urbana para jovens da comunidade</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-golden rounded-full w-full\"></div></div><span class=\"text-blue\">100%</span></div><button class=\"w-full text-pink hover:bg-pink/10 py-2 rounded-lg transition\">Ver Resultados →</button></div></div></div></div></section><!-- Nossa Proposta --><section class=\"py-16\"><div class=\"container mx-auto px-4\"><div class=\"max-w-4xl mx-auto text-center\"><h3 class=\"text-3xl font-bold mb-8 text-orange dark:text-golden\"><span class=\"text-green\">❤️</span> Como Funcionamos</h3><div class=\"grid md:grid-cols-2 gap-8 text-left\"><div class=\"p-6 bg-golden/5 dark:bg-golden/10 rounded-xl\"><div class=\"text-4xl mb-4\">🤝</div><h4 class=\"text-xl font-semibold mb-3 dark:text-gray-200\">Conectamos</h4><p class=\"text-gray-600 dark:text-gray-400\">Unimos pessoas com ideias inovadoras a voluntários e apoiadores dispostos a colaborar</p></div><div class=\"p-6 bg-blue/5 dark:bg-blue/10 rounded-xl\"><div class=\"text-4xl mb-4\">🛠️</div><h4 class=\"text-xl font-semibold mb-3 dark:text-gray-200\">Implementamos</h4><p class=\"text-gray-600 dark:text-gray-400\">Oferecemos ferramentas e suporte para tirar projetos do papel</p></div></div><div class=\"mt-12 bg-white dark:bg-gray-700 shadow-lg rounded-xl p-8 border-t-4 border-pink\"><h4 class=\"text-2xl font-bold text-orange dark:text-golden mb-4\">Nossos Pilares</h4><ul class=\"space-y-4 text-gray-600 dark:text-gray-400 text-left\"><li class=\"flex items-center gap-3\"><span class=\"text-green\">✓</span> Transparência total nos processos</li><li class=\"flex items-center gap-3\"><span class=\"text-blue\">✓</span> Apoio comunitário sustentável</li><li class=\"flex items-center gap-3\"><span class=\"text-pink\">✓</span> Impacto social mensurável</li></ul></div></div></div></section><!-- Estatísticas --><section class=\"bg-gray-50 dark:bg-gray-800 py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">📈 Nosso Impacto</h3><div class=\"grid md:grid-cols-3 gap-8\"><div class=\"bg-white dark:bg-gray-700 p-6 rounded-xl text-center\"><div class=\"text-4xl text-green dark:text-golden\">120+</div><p class=\"text-gray-600 dark:text-gray-300\">Projetos Ativos</p><span class=\"text-sm text-orange dark:text-blue\">🏆 Em 15 cidades</span></div><div class=\"bg-white dark:bg-gray-700 p-6 rounded-xl text-center\"><div class=\"text-4xl text-blue dark:text-golden\">5.8K</div><p class=\"text-gray-600 dark:text-gray-300\">Voluntários</p><span class=\"text-sm text-orange dark:text-blue\">🤝 Comunidade ativa</span></div><div class=\"bg-white dark:bg-gray-700 p-6 rounded-xl text-center\"><div class=\"text-4xl text-pink dark:text-golden\">92%</div><p class=\"text-gray-600 dark:text-gray-300\">Sucesso</p><span class=\"text-sm text-orange dark:text-blue\">🎯 Metas alcançadas</span></div></div></div></section><!-- Depoimentos --><section class=\"py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">🗣️ Quem Apoia</h3><div class=\"grid md:grid-cols-2 gap-8\"><div class=\"bg-golden/5 dark:bg-gray-700 p-6 rounded-xl\"><div class=\"flex items-center gap-4 mb-4\"><span class=\"text-3xl\">👩🌾</span><div><h4 class=\"font-semibold dark:text-gray-200\">Maria Silva</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">Voluntária na Horta Comunitária</p></div></div><p class=\"text-gray-600 dark:text-gray-300\">\"A Weebie nos deu estrutura para transformar um terreno abandonado em fonte de alimentos para toda comunidade!\"</p></div><div class=\"bg-blue/5 dark:bg-gray-700 p-6 rounded-xl\"><div class=\"flex items-center gap-4 mb-4\"><span class=\"text-3xl\">👨🏫</span><div><h4 class=\"font-semibold dark:text-gray-200\">Carlos Mendes</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">Coordenador de Projetos</p></div></div><p class=\"text-gray-600 dark:text-gray-300\">\"A plataforma simplificou nossa captação de recursos e organização de voluntários de maneira incrível.\"</p></div></div></div></section><!-- FAQ --><section class=\"bg-gray-50 dark:bg-gray-800 py-16\"><div class=\"container mx-auto px-4 max-w-3xl\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">❓ Perguntas Frequentes</h3><div class=\"space-y-4\"><div class=\"bg-white dark:bg-gray-700 rounded-xl p-4\"><button class=\"flex justify-between items-center w-full\"><span class=\"font-semibold text-left dark:text-gray-200\">Como participar de um projeto?</span> <span class=\"text-orange dark:text-golden\">+</span></button><div class=\"mt-2 text-gray-600 dark:text-gray-400 hidden\">Escolha um projeto em nossa plataforma, clique em \"Participar\" e siga as instruções. É gratuito e aberto a todos!</div></div><!-- Adicione mais perguntas seguindo o mesmo padrão --></div></div></section><!-- Eventos --><section class=\"py-16\"><div class=\"container mx-auto px-4\"><h3 class=\"text-3xl font-bold text-center mb-12 text-orange dark:text-golden\">📅 Próximos Eventos</h3><div class=\"grid md:grid-cols-2 gap-6\"><div class=\"bg-white dark:bg-gray-700 p-6 rounded-xl border-l-4 border-green\"><div class=\"flex items-center gap-4\"><div class=\"text-center\"><div class=\"text-2xl font-bold text-orange dark:text-golden\">25</div><div class=\"text-sm\">Out</div></div><div><h4 class=\"font-semibold dark:text-gray-200\">Workshop de Compostagem</h4><p class=\"text-sm text-gray-600 dark:text-gray-400\">📌 Centro Comunitário - 14h</p></div></div></div><!-- Adicione mais eventos seguindo o mesmo padrão --></div></div></section><!-- Footer --><footer class=\"bg-gradient-to-r from-orange to-golden dark:from-gray-800 dark:to-gray-700 text-white\"><div class=\"container mx-auto px-4 py-8\"><div class=\"grid md:grid-cols-3 gap-8\"><div><div class=\"flex items-center gap-2 mb-4\"><span class=\"text-2xl\">🌍</span><h5 class=\"text-xl font-semibold\">Weebie</h5></div><p class=\"text-sm opacity-90\">Transformando ideias em ações comunitárias desde 2024</p></div><div><h6 class=\"text-lg font-medium mb-3\">Links Rápidos</h6><ul class=\"space-y-2 text-sm\"><li><a href=\"#\" class=\"hover:underline\">📘 Sobre Nós</a></li><li><a href=\"#\" class=\"hover:underline\">📩 Contato</a></li><li><a href=\"#\" class=\"hover:underline\">📜 Termos</a></li></ul></div><div><h6 class=\"text-lg font-medium mb-3\">Conectar</h6><div class=\"flex gap-4\"><a href=\"#\" class=\"p-2 bg-white/10 rounded-full hover:bg-white/20 transition\">📱</a> <a href=\"#\" class=\"p-2 bg-white/10 rounded-full hover:bg-white/20 transition\">💬</a> <a href=\"#\" class=\"p-2 bg-white/10 rounded-full hover:bg-white/20 transition\">📸</a></div></div></div><div class=\"border-t border-white/20 mt-8 pt-4 text-center text-sm\"><p>© 2025 Weebie - Feito com ❤️ para comunidades</p></div></div></footer><script>\r\n        document.addEventListener('DOMContentLoaded', () => {\r\n          const themeToggle = document.getElementById('theme-toggle');\r\n          const html = document.documentElement;\r\n          \r\n          // Verificar preferência salva\r\n          const savedTheme = localStorage.getItem('theme') || 'light';\r\n          html.classList.toggle('dark', savedTheme === 'dark');\r\n        \r\n          themeToggle.addEventListener('click', () => {\r\n            html.classList.toggle('dark');\r\n            localStorage.setItem('theme', html.classList.contains('dark') ? 'dark' : 'light');\r\n            updateThemeIcon();\r\n          });\r\n        \r\n          function updateThemeIcon() {\r\n            const isDark = html.classList.contains('dark');\r\n            themeToggle.querySelectorAll('svg').forEach(icon => {\r\n              icon.classList.toggle('hidden', isDark ? icon.getAttribute('fill') === '#fff' : icon.getAttribute('fill') !== '#fff');\r\n            });\r\n          }\r\n          updateThemeIcon();\r\n        });\r\n    </script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/2.2.0/flowbite.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
