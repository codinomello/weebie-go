// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package static

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import components "github.com/codinomello/weebie-go/templates/components"

func Project() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- Ícone --><link rel=\"icon\" type=\"image/x-icon\" href=\"../../icons/house.png\"><!-- HTMX e Tailwind CSS -->")
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<!-- Lista de ODS -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ScriptHeadODSList().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</head><body class=\"bg-gray-50 dark:bg-gray-900\"><!-- Header --><nav class=\"bg-gradient-to-r from-orange to-golden dark:from-gray-800 dark:to-gray-700 px-4 py-4 shadow-xl\"><div class=\"container mx-auto flex items-center justify-between\"><div class=\"flex items-center space-x-3\"><span class=\"text-3xl text-white\">🏘️</span><h1 class=\"text-2xl font-bold text-white\">Cadastro de Projetos</h1></div><!-- Botão de Tema e Menu Mobile --><div class=\"flex items-center space-x-4\"><!-- Botão de Tema Moderno --><button id=\"theme-toggle\" class=\"theme-toggle\"><span class=\"sr-only\">Alternar Modo Escuro</span><style>\n                        .theme-toggle {\n                            position: relative;\n                            width: 48px;\n                            height: 24px;\n                            border-radius: 9999px;\n                            background-color: rgba(255, 255, 255, 0.1);\n                            transition: background-color 0.3s;\n                        }\n                        .theme-toggle:hover {\n                            background-color: rgba(255, 255, 255, 0.2);\n                        }\n                        .theme-toggle::before {\n                            content: '';\n                            position: absolute;\n                            top: 2px;\n                            left: 2px;\n                            width: 20px;\n                            height: 20px;\n                            border-radius: 50%;\n                            background-color: white;\n                            transition: transform 0.3s, background-color 0.3s;\n                        }\n                        .dark .theme-toggle::before {\n                            transform: translateX(24px);\n                            background-color: #FFD700;\n                        }\n                    </style></button><!-- Menu Mobile --><button class=\"md:hidden text-white p-2 rounded-lg hover:bg-white/10 transition-colors\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button></div></div></nav><!-- Formulário --><main class=\"container mx-auto px-4 py-8 max-w-3xl\"><form class=\"bg-white dark:bg-gray-800 rounded-2xl p-8 shadow-lg border border-gray-200 dark:border-gray-700\"><!-- Seção Azul --><div class=\"mb-8 p-4 rounded-xl bg-blue/10 dark:bg-blue/20\"><h2 class=\"text-2xl font-bold text-blue dark:text-blue/80 mb-4\">Informações Básicas</h2><div class=\"grid md:grid-cols-2 gap-6\"><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">ID do Projeto</label> <input type=\"text\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 dark:text-gray-300 focus:ring-2 focus:ring-blue\"></div><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Data</label> <input type=\"date\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 dark:text-gray-300 focus:ring-2 focus:ring-blue\"></div></div></div><!-- Seção Rosa --><div class=\"mb-8 p-4 rounded-xl bg-pink/10 dark:bg-pink/20\"><h2 class=\"text-2xl font-bold text-pink dark:text-pink/80 mb-4\">Detalhes do Projeto</h2><div class=\"space-y-6\"><!-- Campo de Título --><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Título</label> <input type=\"text\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 dark:text-gray-300 focus:ring-2 focus:ring-pink\"></div><!-- Campo de Autores (e-mails) --><div class=\"space-y-4\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Autores</label> <input id=\"new-author-input\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 dark:text-gray-300 focus:ring-2 focus:ring-pink\"><!-- Div dos e-mails (inicialmente oculta) --><div id=\"authors-container\" class=\"hidden flex flex-wrap gap-2 px-4 py-2 border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 dark:text-gray-300 rounded-lg mb-4\"><!-- E-mails serão adicionados aqui dinamicamente --></div><!-- Contêiner para os botões --><div class=\"flex items-center gap-2\"><!-- Botão Adicionar --><button type=\"button\" id=\"add-author-btn\" class=\"mt-2 px-4 py-2 bg-pink text-gray-300 rounded-lg hover:bg-pink/80 transition\">Adicionar</button><!-- Botão de Informação --><button data-tooltip-target=\"email-tooltip\" data-tooltip-placement=\"right\" type=\"button\" class=\"text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300\"><svg class=\"w-4 h-4\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg></button></div><!-- Tooltip --><div id=\"email-tooltip\" role=\"tooltip\" class=\"inline-block absolute invisible z-10 py-2 px-3 text-sm font-medium dark:text-gray-300 bg-pink/10 rounded-lg shadow-sm opacity-0 transition-opacity duration-300\">Insira e-mails separados por vírgula ou espaço.</div></div></div></div><script>\n                document.addEventListener('DOMContentLoaded', () => {\n                    const authorsContainer = document.getElementById('authors-container');\n                    const newAuthorInput = document.getElementById('new-author-input');\n                    const addAuthorBtn = document.getElementById('add-author-btn');\n            \n                    // Função para validar e-mail\n                    const isValidEmail = (email) => {\n                        const regex = /^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$/;\n                        return regex.test(email);\n                    };\n            \n                    // Função para adicionar e-mails\n                    const addEmails = () => {\n                        const input = newAuthorInput.value.trim();\n                        if (input) {\n                            // Separa os e-mails por vírgula ou espaço\n                            const emails = input.split(/[,\\s]+/).filter(email => email.trim() !== '');\n            \n                            emails.forEach(email => {\n                                if (isValidEmail(email)) {\n                                    const emailTag = document.createElement('div');\n                                    emailTag.className = 'flex items-center bg-pink/10 dark:bg-pink/20 px-3 py-1 rounded-full';\n                                    emailTag.innerHTML = `\n                                        <span class=\"text-sm text-pink dark:text-pink/80\">${email}</span>\n                                        <button type=\"button\" class=\"ml-2 text-pink hover:text-pink\" onclick=\"this.parentElement.classList.add('hidden');\">×</button>\n                                    `;\n                                    authorsContainer.appendChild(emailTag);\n            \n                                    // Mostra a div de autores se estiver oculta\n                                    const checkIfEmpty = () => { \n                                        if (authorsContainer.children.length === 0) { // Verifica se não há elementos filhos\n                                            authorsContainer.classList.add('hidden');\n                                        } else {\n                                            authorsContainer.classList.remove('hidden'); // Garante que a classe seja removida se não estiver vazio\n                                        }\n                                    };\n                                    checkIfEmpty()\n\n                                } else {\n                                    alert(`E-mail inválido: ${email}`);\n                                }\n                            });\n            \n                            newAuthorInput.value = ''; // Limpa o campo de entrada\n                        }\n                    };\n            \n                    // Adicionar e-mails ao clicar no botão\n                    addAuthorBtn.addEventListener('click', addEmails);\n            \n                    // Adicionar e-mails ao pressionar Enter\n                    newAuthorInput.addEventListener('keydown', (e) => {\n                        if (e.key === 'Enter') {\n                            e.preventDefault(); // Evita o comportamento padrão do Enter\n                            addEmails();\n                        }\n                    });\n                });\n            </script><!-- Seção Verde --><div class=\"mb-8 p-4 rounded-xl bg-green/10 dark:bg-green/20\"><h2 class=\"text-2xl font-bold text-green dark:text-green/80 mb-4\">Classificação</h2><div class=\"grid md:grid-cols-2 gap-6\"><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Grupo</label> <input type=\"text\" class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:text-gray-300 dark:bg-gray-700 focus:ring-2 focus:ring-green\"></div><div class=\"space-y-2\"><label class=\"block text-sm font-medium text-gray-700 dark:text-gray-300\">Curso</label> <select class=\"w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:text-gray-300 dark:bg-gray-700 focus:ring-2 focus:ring-green\"><option></option> <option>Administração</option> <option>Edificações</option> <option>Enfermagem</option> <option>Informática</option> <option>Mecatrônica</option></select></div></div></div><!-- Seção ODS --><div class=\"mb-8 p-4 rounded-xl bg-golden/10 dark:bg-golden/20\"><h2 class=\"text-2xl font-bold text-golden dark:text-golden/80 mb-4\">ODS Relacionados</h2><div class=\"grid grid-cols-2 md:grid-cols-4 gap-3\" id=\"ods-container\"><!-- ODS são inseridos aqui via JavaScript --></div><div id=\"ods-popover\" class=\"hidden absolute z-50 bg-white dark:bg-gray-800 p-4 rounded-lg shadow-lg ods-popover\"><h3 class=\"font-bold text-lg text-gray-900 dark:text-white mb-2\" id=\"ods-title\"></h3><p class=\"text-sm text-gray-600 dark:text-gray-300\" id=\"ods-description\"></p></div></div><!-- Botões --><div class=\"flex flex-col md:flex-row justify-end gap-4 mt-8\"><!-- Botão limpar (vermelho) --><button type=\"button\" id=\"clear-btn\" class=\"px-6 py-2 bg-red text-white rounded-lg hover:bg-red-600 active:bg-red-700 transition\">Limpar Campos</button><!-- Botão salvar (verde) --><button type=\"submit\" class=\"px-6 py-2 bg-green text-white rounded-lg hover:bg-green-600 active:bg-green-700 transition\">Salvar Projeto</button></div><script>\n                document.addEventListener('DOMContentLoaded', () => {\n                    const clearBtn = document.getElementById('clear-btn');\n                    const form = document.querySelector('form');\n\n                    // Função para limpar todos os campos do formulário\n                    clearBtn.addEventListener('click', () => {\n                        form.reset(); // Limpa todos os campos do formulário\n                    });\n                });\n            </script></form></main><script src=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/1.6.5/flowbite.min.js\"></script><script>\n        document.addEventListener('DOMContentLoaded', () => {\n            const themeToggle = document.getElementById('theme-toggle');\n            const html = document.documentElement;\n            const odsContainer = document.getElementById('ods-container');\n            const odsPopover = document.getElementById('ods-popover');\n            const odsTitle = document.getElementById('ods-title');\n            const odsDescription = document.getElementById('ods-description');\n\n            // Tema\n            themeToggle.addEventListener('click', () => {\n                html.classList.toggle('dark');\n                localStorage.setItem('theme', html.classList.contains('dark') ? 'dark' : 'light');\n            });\n\n            if (localStorage.getItem('theme') === 'dark') html.classList.add('dark');\n\n            // Gerar ODS\n            odsList.forEach(ods => {\n                const odsItem = document.createElement('div');\n                odsItem.className = 'flex items-center justify-between p-2 rounded-lg bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600';\n                odsItem.innerHTML = `\n                    <label class=\"flex items-center gap-2 flex-1\">\n                        <input type=\"checkbox\" name=\"ods\" value=\"${ods.number}\" \n                               class=\"w-4 h-4 border-golden rounded focus:ring-golden dark:bg-gray-900\">\n                        <span class=\"text-sm font-medium text-gray-700 dark:text-gray-300\">\n                            ${ods.emoji} ODS ${ods.number}\n                        </span>\n                    </label>\n                    <button type=\"button\" class=\"text-golden hover:text-golden/80 px-2 ods-info\" \n                            data-number=\"${ods.number}\">+</button>\n                `;\n                odsContainer.appendChild(odsItem);\n            });\n\n            // Popover ODS\n            document.querySelectorAll('.ods-info').forEach(button => {\n                button.addEventListener('click', (e) => {\n                    const odsNumber = parseInt(e.target.dataset.number);\n                    const ods = odsList.find(o => o.number === odsNumber);\n                    \n                    // Posicionar popover\n                    const rect = e.target.getBoundingClientRect();\n                    odsPopover.style.top = `${rect.top + window.scrollY}px`;\n                    odsPopover.style.left = `${rect.left + rect.width + 10}px`;\n                    \n                    // Atualizar conteúdo\n                    odsTitle.textContent = ods.name;\n                    odsDescription.textContent = ods.description;\n                    odsPopover.classList.remove('hidden');\n                });\n            });\n\n            // Fechar popover\n            document.addEventListener('click', (e) => {\n                if (!e.target.closest('.ods-info') && !e.target.closest('#ods-popover')) {\n                    odsPopover.classList.add('hidden');\n                }\n            });\n        });\n    </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
