// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package public

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import components "github.com/codinomello/weebie-go/pages/components"

func About() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- HTMX e Tailwind CSS -->")
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<link rel=\"icon\" type=\"image/x-icon\" href=\"/images/icons/house_with_garden.png\"></head><body class=\"bg-gray-100 dark:bg-gray-900\"><!-- Header -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Nav("Sobre Nós").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<!-- Main Content --><main class=\"container mx-auto px-4 py-16 mt-20 flex-1\"><!-- Hero Section --><section class=\"text-center mb-16\"><h1 class=\"text-4xl md:text-5xl font-bold text-orange dark:text-golden mb-6\"><span class=\"text-blue\">🏡</span> Weebie: Transformando Comunidades</h1><p class=\"text-lg text-gray-600 dark:text-gray-300 max-w-2xl mx-auto\">Weebie é uma plataforma que conecta pessoas apaixonadas por criar impacto positivo, promovendo iniciativas sustentáveis e colaborativas para fortalecer comunidades locais.</p><div class=\"mt-8 flex flex-col sm:flex-row justify-center gap-4\"><a href=\"/project\" class=\"bg-green text-white text-bold px-6 py-3 rounded-lg hover:bg-green/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105\"><span>🌱</span> Explorar Projetos</a> <a href=\"/\" class=\"bg-blue text-white text-bold px-6 py-3 rounded-lg hover:bg-blue/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105\"><span>🚀</span> Começar Agora</a></div></section><!-- Mission Section --><section class=\"mb-16\"><h2 class=\"text-3xl font-bold text-center mb-8 text-orange dark:text-golden\"><span class=\"text-pink\">🌍</span> Nossa Missão</h2><div class=\"bg-white dark:bg-gray-800 rounded-2xl p-8 shadow-lg border border-gray-200 dark:border-gray-700\"><p class=\"text-gray-600 dark:text-gray-300 text-center max-w-3xl mx-auto\">No Weebie, acreditamos que comunidades fortes são construídas por pessoas unidas por um propósito comum. Nossa missão é fornecer ferramentas para criar, gerenciar e compartilhar projetos que promovam sustentabilidade, colaboração e impacto social mensurável.</p></div></section><!-- Values Section --><section class=\"mb-16\"><h2 class=\"text-3xl font-bold text-center mb-8 text-orange dark:text-golden\"><span class=\"text-green\">❤️</span> Nossos Valores</h2><div class=\"grid md:grid-cols-3 gap-6\"><div class=\"dark:bg-gray-700 p-6 rounded-xl border-l-4 border-blue shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl mb-4\">🤝</div><h3 class=\"text-xl font-semibold dark:text-gray-200 mb-2\">Colaboração</h3><p class=\"text-gray-600 dark:text-gray-400\">Unimos pessoas e organizações para transformar ideias em realidade.</p></div><div class=\"dark:bg-gray-700 p-6 rounded-xl border-l-4 border-green shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl mb-4\">🌱</div><h3 class=\"text-xl font-semibold dark:text-gray-200 mb-2\">Sustentabilidade</h3><p class=\"text-gray-600 dark:text-gray-400\">Promovemos práticas que respeitam o meio ambiente e as gerações futuras.</p></div><div class=\"dark:bg-gray-700 p-6 rounded-xl border-l-4 border-orange shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl mb-4\">📊</div><h3 class=\"text-xl font-semibold dark:text-gray-200 mb-2\">Transparência</h3><p class=\"text-gray-600 dark:text-gray-400\">Garantimos clareza em todos os processos e impactos dos projetos.</p></div></div></section><!-- Impact Section --><section class=\"mb-16\"><h2 class=\"text-3xl font-bold text-center mb-8 text-orange dark:text-golden\"><span class=\"text-blue\">📈</span> Nosso Impacto</h2><div class=\"grid md:grid-cols-3 gap-8\"><div class=\"dark:bg-gray-700 p-6 rounded-xl text-center border-b-4 border-red shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl text-green dark:text-golden\">120+</div><p class=\"text-gray-600 dark:text-gray-300\">Projetos Ativos</p><span class=\"text-sm text-orange dark:text-red\">🏆 Em 15 cidades</span></div><div class=\"dark:bg-gray-700 p-6 rounded-xl text-center border-b-4 border-green shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl text-blue dark:text-golden\">5.8K</div><p class=\"text-gray-600 dark:text-gray-300\">Voluntários</p><span class=\"text-sm text-orange dark:text-green\">🤝 Comunidade ativa</span></div><div class=\"dark:bg-gray-700 p-6 rounded-xl text-center border-b-4 border-blue shadow-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:scale-105\"><div class=\"text-4xl text-pink dark:text-golden\">92%</div><p class=\"text-gray-600 dark:text-gray-300\">Sucesso</p><span class=\"text-sm text-orange dark:text-blue\">🎯 Metas alcançadas</span></div></div></section><!-- Call to Action --><section class=\"mb-16 text-center\"><h2 class=\"text-3xl font-bold mb-6 text-orange dark:text-golden\"><span class=\"text-pink\">🚀</span> Junte-se ao Weebie!</h2><p class=\"text-lg text-gray-600 dark:text-gray-300 max-w-2xl mx-auto mb-8\">Faça parte de uma comunidade que transforma ideias em ações concretas. Crie, colabore e impacte o mundo ao seu redor.</p><a href=\"/form\" class=\"bg-orange text-white text-bold px-6 py-3 rounded-lg hover:bg-orange/80 transition flex items-center gap-2 hover:shadow-lg hover:scale-105 mx-auto\"><span>👉</span> Comece Agora</a></section></main><!-- Footer -->")
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
