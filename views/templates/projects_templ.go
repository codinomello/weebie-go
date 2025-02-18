// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Projects() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"pt-BR\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Cadastro de Projetos - Weebie</title><link rel=\"icon\" type=\"image/x-icon\" href=\"../images/cabinet.png\"><!-- Tailwind CSS --><script src=\"https://cdn.tailwindcss.com\"></script><!-- Flowbite CSS --><link href=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/1.6.5/flowbite.min.css\" rel=\"stylesheet\"><!-- HTMX --><script src=\"https://unpkg.com/htmx.org@1.9.5\"></script><!-- Alpine.js--><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script><style>\r\n        @import url('https://fonts.cdnfonts.com/css/euclid-circular-a');\r\n        * {\r\n            font-family: \"Euclid Circular A\";\r\n            transition: all 0.3s ease;\r\n        }\r\n    </style><script>\r\n        tailwind.config = {\r\n          darkMode: 'class',\r\n          theme: {\r\n            extend: {\r\n              colors: {\r\n                red: '#FF0000',    // Vermelho\r\n                green: '#32CD32',  // Verde\r\n                blue: '#00BFFF',   // Azul\r\n                lilac: '#C8A2C8',  // Lilás\r\n                pink: '#FF1493',   // Rosa\r\n                orange: '#FF6F00', // Laranja\r\n                golden: '#FFD700'  // Ouro\r\n              }\r\n            }\r\n          }\r\n        }\r\n    </script></head><body class=\"bg-gray-50 dark:bg-gray-900\" x-data=\"{ authors: [], newAuthor: &#39;&#39; }\"><!-- Header --><nav class=\"bg-gradient-to-r from-orange to-golden dark:from-gray-800 dark:to-gray-700 px-4 py-4 shadow-xl\"><div class=\"container mx-auto flex items-center justify-between\"><div class=\"flex items-center space-x-3\"><span class=\"text-3xl text-white\">🏘️</span><h1 class=\"text-2xl font-bold text-white\">Cadastro de Projetos</h1></div><!-- Botão de Tema e Menu Mobile --><div class=\"flex items-center space-x-4\"><!-- Botão de Tema Moderno --><button id=\"theme-toggle\" class=\"theme-toggle\"><span class=\"sr-only\">Alternar Modo Escuro</span><style>\r\n                        .theme-toggle {\r\n                            position: relative;\r\n                            width: 48px;\r\n                            height: 24px;\r\n                            border-radius: 9999px;\r\n                            background-color: rgba(255, 255, 255, 0.1);\r\n                            transition: background-color 0.3s;\r\n                        }\r\n                        .theme-toggle:hover {\r\n                            background-color: rgba(255, 255, 255, 0.2);\r\n                        }\r\n                        .theme-toggle::before {\r\n                            content: '';\r\n                            position: absolute;\r\n                            top: 2px;\r\n                            left: 2px;\r\n                            width: 20px;\r\n                            height: 20px;\r\n                            border-radius: 50%;\r\n                            background-color: white;\r\n                            transition: transform 0.3s, background-color 0.3s;\r\n                        }\r\n                        .dark .theme-toggle::before {\r\n                            transform: translateX(24px);\r\n                            background-color: #FFD700;\r\n                        }\r\n                    </style></button><!-- Menu Mobile --><button class=\"md:hidden text-white p-2 rounded-lg hover:bg-white/10 transition-colors\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button></div></div></nav><!-- Formulário --><main class=\"container mx-auto px-4 py-8 max-w-3xl\"><form class=\"bg-white dark:bg-gray-800 rounded-2xl p-8 shadow-lg border border-gray-200 dark:border-gray-700\"><!-- Seção Azul --><div class=\"mb-8 p-4 rounded-xl bg-blue/10 dark:bg-blue/20\"><h2 class=\"text-2xl font-bold text-blue dark:text-blue/80 mb-4\">Informações Básicas</h2><div class=\"grid md:grid-cols-2 gap-6\"><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">ID do Projeto</label> <input type=\"text\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-orange\"></div><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Ano</label> <input type=\"number\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-orange\"></div></div></div><!-- Seção Rosa --><div class=\"mb-8 p-4 rounded-xl bg-pink/10 dark:bg-pink/20\"><h2 class=\"text-2xl font-bold text-pink dark:text-pink/80 mb-4\">Detalhes do Projeto</h2><div class=\"space-y-6\"><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Título</label> <input type=\"text\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-orange\"></div><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Autores</label><div class=\"flex flex-wrap gap-2 p-2 border border-gray-300 dark:border-gray-600 rounded-lg\"><template x-for=\"(author, index) in authors\" :key=\"index\"><div class=\"flex items-center bg-green/10 dark:bg-green/20 px-3 py-1 rounded-full\"><span class=\"text-sm text-green dark:text-green/80\" x-text=\"author\"></span> <button @click=\"authors.splice(index, 1)\" class=\"ml-2 text-orange hover:text-[#FF8A33]\">×</button></div></template><input x-model=\"newAuthor\" @keydown.enter.prevent=\"if(newAuthor.trim()) { authors.push(newAuthor.trim()); newAuthor = &#39;&#39; }\" class=\"flex-1 bg-transparent outline-none px-2 min-w-[120px]\" placeholder=\"Adicionar autor...\"></div></div></div></div><!-- Seção Verde --><div class=\"mb-8 p-4 rounded-xl bg-green/10 dark:bg-green/20\"><h2 class=\"text-2xl font-bold text-green dark:text-green/80 mb-4\">Classificação</h2><div class=\"grid md:grid-cols-2 gap-6\"><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Grupo</label> <input type=\"text\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-orange\"></div><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Curso</label> <select class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-orange\"><option>Engenharia de Software</option> <option>Ciência da Computação</option> <option>Sistemas de Informação</option></select></div></div></div><!-- Seção ODS --><div class=\"mb-8 p-4 rounded-xl bg-golden/10 dark:bg-golden/20\"><h2 class=\"text-2xl font-bold text-golden dark:text-golden/80 mb-4\">ODS Relacionados</h2><div class=\"grid grid-cols-2 md:grid-cols-4 gap-3\"><template x-for=\"i in 17\" :key=\"i\"><label class=\"flex items-center space-x-2 p-2 rounded-lg bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600\"><input type=\"checkbox\" class=\"ods-checkbox border-orange checked:bg-orange dark:border-gray-500\"> <span class=\"text-sm font-medium text-gray-700 dark:text-gray-300\" x-text=\"i\"></span></label></template></div></div><!-- Botões --><div class=\"flex flex-col md:flex-row justify-end gap-4 mt-8\"><button type=\"reset\" class=\"px-6 py-2 bg-pink text-white text-gray-700 dark:text-gray-300 hover:bg-red rounded-lg transition\">Limpar Campos</button> <button type=\"submit\" class=\"px-6 py-2 bg-golden text-white rounded-lg hover:bg-green transition\">Salvar Projeto</button></div></form></main><script>\r\n        document.addEventListener('DOMContentLoaded', () => {\r\n            const themeToggle = document.getElementById('theme-toggle');\r\n            const html = document.documentElement;\r\n\r\n            themeToggle.addEventListener('click', () => {\r\n                html.classList.toggle('dark');\r\n                localStorage.setItem('theme', html.classList.contains('dark') ? 'dark' : 'light');\r\n            });\r\n\r\n            if (localStorage.getItem('theme') === 'dark') {\r\n                html.classList.add('dark');\r\n            }\r\n        });\r\n    </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
