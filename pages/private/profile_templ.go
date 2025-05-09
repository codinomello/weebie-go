// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package private

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import components "github.com/codinomello/weebie-go/pages/components"

func Profile() templ.Component {
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
		templ_7745c5c3_Err = components.Title("Weebie - Perfil do Usuário").Render(ctx, templ_7745c5c3_Buffer)
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<!-- Fonte -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.StyleHead().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<!-- Flowbite CSS e Ícone -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Link("house_with_garden.png").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</head><body class=\"bg-gray-50 dark:bg-gray-900 flex flex-col min-h-screen\"><!-- Header --><nav class=\"bg-gradient-to-r from-orange to-golden dark:from-gray-800 dark:to-gray-700 px-4 py-4 shadow-xl fixed top-0 w-full\"><div class=\"container mx-auto flex items-center justify-between\"><div class=\"flex items-center space-x-3\"><span class=\"text-3xl text-white\">🏡</span><h1 class=\"text-2xl font-bold text-white\">Weebie</h1></div><!-- Botões de Controle --><div class=\"flex items-center space-x-4\"><!-- Botão de Tema -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ButtonToggleTheme().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<!-- Ícone de Conta --><button class=\"p-2 rounded-full text-white hover:bg-white/10 transition-colors duration-300\"><svg class=\"w-7 h-7 text-white dark:golden\" viewBox=\"0 0 24 24\" fill=\"currentColor\"><path d=\"M12 4a4 4 0 014 4 4 4 0 01-4 4 4 4 0 01-4-4 4 4 0 014-4m0 10c4.42 0 8 1.79 8 4v2H4v-2c0-2.21 3.58-4 8-4z\"></path></svg></button><!-- Menu Mobile --><button class=\"md:hidden text-white p-2 rounded-lg hover:bg-white/10 transition-colors\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button></div></div></nav><!-- Main Content --><main class=\"container mx-auto px-4 py-16 mt-20 flex-1\"><div class=\"grid md:grid-cols-3 gap-8\"><!-- Perfil do Usuário --><div class=\"md:col-span-1\"><div class=\"bg-white dark:bg-gray-800 rounded-2xl p-8 shadow-lg border border-gray-200 dark:border-gray-700\"><div class=\"flex flex-col items-center\"><div class=\"w-24 h-24 rounded-full bg-gradient-to-r from-orange to-golden flex items-center justify-center text-4xl text-white mb-4\"><span>👤</span></div><h2 class=\"text-2xl font-bold text-orange dark:text-golden mb-2\">Nome do Usuário</h2><p class=\"text-gray-600 dark:text-gray-400 mb-4\">email@exemplo.com</p><p class=\"text-gray-600 dark:text-gray-400 mb-6 text-center\">Membro desde: Janeiro de 2025</p><button class=\"w-full px-6 py-2 bg-gradient-to-r from-orange to-golden text-white rounded-lg hover:opacity-90 transition\">Editar Perfil</button></div></div></div><!-- Projetos do Usuário --><div class=\"md:col-span-2\"><h3 class=\"text-3xl font-bold mb-8 text-orange dark:text-golden\"><span class=\"text-pink\">🌟</span> Meus Projetos</h3><div class=\"grid md:grid-cols-2 gap-6\"><!-- Projeto 1 --><div class=\"dark:bg-gray-700 rounded-xl shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105 p-6 border-l-4 border-green\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">🌻 Horta Urbana</h4><span class=\"bg-green/10 dark:bg-green/20 text-green dark:text-green/80 px-3 py-1 rounded-full text-sm\">Ativo</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Cultivo coletivo de alimentos orgânicos no centro da cidade</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-blue rounded-full w-3/4\"></div></div><span class=\"text-blue\">75%</span></div><button class=\"w-full text-green hover:bg-green/10 py-2 rounded-lg transition\">Gerenciar Projeto →</button></div></div><!-- Projeto 2 --><div class=\"dark:bg-gray-700 rounded-xl shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105 p-6 border-l-4 border-blue\"><div class=\"flex items-start justify-between mb-4\"><h4 class=\"text-xl font-semibold dark:text-gray-200\">📚 Biblioteca Livre</h4><span class=\"bg-blue/10 dark:bg-blue/20 text-blue dark:text-blue/80 px-3 py-1 rounded-full text-sm\">Novo</span></div><p class=\"text-gray-600 dark:text-gray-400 mb-4\">Pontos de compartilhamento de livros em espaços públicos</p><div class=\"space-y-3\"><div class=\"flex items-center gap-2 text-sm\"><span class=\"text-orange dark:text-golden font-medium\">Progresso:</span><div class=\"flex-1 h-2 bg-gray-200 rounded-full\"><div class=\"h-2 bg-golden rounded-full w-2/5\"></div></div><span class=\"text-orange\">40%</span></div><button class=\"w-full text-blue hover:bg-blue/10 py-2 rounded-lg transition\">Gerenciar Projeto →</button></div></div></div><!-- Botão para Criar Novo Projeto --><div class=\"mt-8 text-center\"><button class=\"bg-green text-white text-bold px-6 py-3 rounded-lg hover:bg-green/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105 mx-auto\"><span>🚀</span> Criar Novo Projeto</button></div></div></div></main><!-- Footer -->")
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
