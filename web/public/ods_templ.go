// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package public

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import components "github.com/codinomello/weebie-go/web/components"

func Ods() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<html lang=\"pt-BR\"><head><!-- Metadados -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Meta().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!-- Título e descrição -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Title("Weebie - Objetivos de Desenvolvimento Sustentável").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- Scripts -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ScriptHead().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<!-- Estilos -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.StyleHead().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<!-- Links e Favicon -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Link("house_with_garden.png").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<!-- Firebase SDK --><script type=\"module\">\r\n            import { initializeApp } from 'https://www.gstatic.com/firebasejs/11.8.1/firebase-app.js';\r\n            import { getAuth, onAuthStateChanged } from 'https://www.gstatic.com/firebasejs/11.8.1/firebase-auth.js';\r\n\r\n            const firebaseConfig = {\r\n                apiKey: \"AIzaSyAI0Tc7GssKwWwtVdrz6OaK6KFACx58N5U\",\r\n                authDomain: \"weebie-go.firebaseapp.com\",\r\n                projectId: \"weebie-go\",\r\n                storageBucket: \"weebie-go.appspot.com\",\r\n                messagingSenderId: \"321509944065\",\r\n                appId: \"1:321509944065:web:199a546b7055f461ec4900\",\r\n                measurementId: \"G-S5CG0CLRVS\"\r\n            };\r\n\r\n            console.log('Inicializando Firebase...');\r\n            const app = initializeApp(firebaseConfig);\r\n            const auth = getAuth(app);\r\n            console.log('Firebase inicializado com sucesso');\r\n\r\n            onAuthStateChanged(auth, (user) => {\r\n                if (user) {\r\n                    console.log('✅ Usuário logado:', {\r\n                        uid: user.uid,\r\n                        email: user.email,\r\n                        displayName: user.displayName,\r\n                        photoURL: user.photoURL || 'Nenhuma foto de perfil disponível'\r\n                    });\r\n\r\n                    // Verifica se o usuário está autenticado via Google\r\n                    const isGoogleUser = user.providerData.some(profile => profile.providerId === 'google.com');\r\n                    if (isGoogleUser && user.photoURL) {\r\n                        console.log('📸 Foto de perfil do Google:', user.photoURL);\r\n                    } else if (isGoogleUser) {\r\n                        console.log('⚠️ Usuário Google sem foto de perfil definida');\r\n                    } else {\r\n                        console.log('⚠️ Usuário não autenticado via Google');\r\n                    }\r\n                } else {\r\n                    console.log('❌ Nenhum usuário logado');\r\n                }\r\n            });\r\n\r\n            window.checkAuth = () => {\r\n                const user = auth.currentUser;\r\n                console.log('Usuário atual:', user ? {\r\n                    uid: user.uid,\r\n                    email: user.email,\r\n                    photoURL: user.photoURL || 'Nenhuma foto'\r\n                } : 'Nenhum');\r\n            };\r\n            console.log('💡 Use window.checkAuth() no console para verificar o usuário');\r\n        </script></head><body class=\"bg-gray-100 dark:bg-gray-900 min-h-screen flex flex-col\"><!-- Barra de navegação -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Nav("Objetivos de Desenvolvimento Sustentável", "earth_globe.png").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<main class=\"flex-grow container mx-auto px-4 py-8\"><div class=\"text-center mb-12\"><h1 class=\"text-4xl font-bold text-blue dark:text-blue-300 mb-4\"><span class=\"text-yellow-600\">🌎</span> ODS</h1><p class=\"text-lg text-gray-600 dark:text-gray-300 max-w-3xl mx-auto\">Os 17 Objetivos de Desenvolvimento Sustentável (ODS) da ONU para transformar nosso mundo</p></div><!-- Estado de carregamento --><div id=\"loading\" class=\"flex justify-center py-12\"><div class=\"animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue\"></div></div><!-- Grid de ODS --><div id=\"ods-grid\" class=\"hidden grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6\"><!-- Os cards dos ODS serão inseridos aqui pelo JavaScript --></div><!-- Modal para detalhes do ODS --><div id=\"ods-modal\" class=\"hidden fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4\"><div class=\"bg-white dark:bg-gray-800 rounded-xl max-w-2xl w-full p-6 shadow-2xl\"><div class=\"flex justify-between items-start mb-4\"><div><h3 id=\"ods-modal-title\" class=\"text-2xl font-bold text-yellow-600 dark:text-yellow-300\"></h3><p id=\"ods-modal-number\" class=\"text-lg text-gray-500 dark:text-gray-400\"></p></div><button onclick=\"closeODSModal()\" class=\"text-gray-500 hover:text-gray-700 dark:hover:text-gray-300\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div><div class=\"flex flex-col md:flex-row gap-6\"><div class=\"flex-1 flex flex-col items-center\"><img id=\"ods-modal-image\" src=\"\" alt=\"\" class=\"w-full max-w-xs rounded-lg mb-4\"> <span id=\"ods-modal-emoji\" class=\"text-6xl\"></span></div><div class=\"flex-1\"><p id=\"ods-modal-description\" class=\"text-gray-700 dark:text-gray-300 mb-4\"></p><div class=\"bg-blue-50 dark:bg-blue-900 p-4 rounded-lg\"><h4 class=\"font-bold text-blue-800 dark:text-blue-200 mb-2\">Metas principais:</h4><ul id=\"ods-modal-targets\" class=\"list-disc list-inside text-gray-700 dark:text-gray-300 space-y-1\"></ul></div></div></div><div class=\"mt-6 flex justify-end\"><button onclick=\"closeODSModal()\" class=\"px-4 py-2 bg-yellow-500 text-white rounded-lg hover:bg-yellow-600 transition\">Fechar</button></div></div></div></main><!-- Rodapé -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.FooterMain().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<script>\r\n            document.addEventListener('DOMContentLoaded', () => {\r\n                // Carregar ODS do endpoint\r\n                fetch('/api/ods')\r\n                    .then(response => response.json())\r\n                    .then(data => {\r\n                        const loading = document.getElementById('loading');\r\n                        const odsGrid = document.getElementById('ods-grid');\r\n                        \r\n                        loading.classList.add('hidden');\r\n                        odsGrid.classList.remove('hidden');\r\n                        \r\n                        data.forEach(ods => {\r\n                            const odsCard = document.createElement('div');\r\n                            odsCard.className = 'bg-white dark:bg-gray-800 rounded-xl shadow-md overflow-hidden border border-gray-200 dark:border-gray-700 hover:shadow-lg transition cursor-pointer';\r\n                            odsCard.innerHTML = `\r\n                                <div onclick=\"showODSDetails(${ods.number})\">\r\n                                    <img src=\"${ods.image_url}\" alt=\"ODS ${ods.number}\" class=\"w-full h-48 object-cover\">\r\n                                    <div class=\"p-4\">\r\n                                        <div class=\"flex items-center gap-3 mb-2\">\r\n                                            <span class=\"text-3xl\">${ods.emoji}</span>\r\n                                            <h3 class=\"text-xl font-bold text-yellow-600 dark:text-yellow-300\">ODS ${ods.number}</h3>\r\n                                        </div>\r\n                                        <p class=\"text-gray-700 dark:text-gray-300\">${ods.title}</p>\r\n                                    </div>\r\n                                </div>\r\n                            `;\r\n                            odsGrid.appendChild(odsCard);\r\n                        });\r\n                    })\r\n                    .catch(error => {\r\n                        console.error('Erro ao carregar ODS:', error);\r\n                        document.getElementById('loading').innerHTML = `\r\n                            <div class=\"text-center py-12\">\r\n                                <p class=\"text-red-500\">Falha ao carregar os ODS. Tente recarregar a página.</p>\r\n                                <button onclick=\"window.location.reload()\" class=\"mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600\">\r\n                                    Recarregar\r\n                                </button>\r\n                            </div>\r\n                        `;\r\n                    });\r\n            });\r\n            \r\n            // Função para mostrar detalhes do ODS\r\n            window.showODSDetails = function(odsNumber) {\r\n                fetch(`/api/ods/${odsNumber}`)\r\n                    .then(response => response.json())\r\n                    .then(ods => {\r\n                        document.getElementById('ods-modal-title').textContent = ods.title;\r\n                        document.getElementById('ods-modal-number').textContent = `ODS ${ods.number}`;\r\n                        document.getElementById('ods-modal-emoji').textContent = ods.emoji;\r\n                        document.getElementById('ods-modal-image').src = ods.image_url;\r\n                        document.getElementById('ods-modal-image').alt = `ODS ${ods.number} - ${ods.title}`;\r\n                        document.getElementById('ods-modal-description').textContent = ods.description;\r\n                        \r\n                        const targetsList = document.getElementById('ods-modal-targets');\r\n                        targetsList.innerHTML = '';\r\n                        ods.targets.forEach(target => {\r\n                            const li = document.createElement('li');\r\n                            li.textContent = target;\r\n                            targetsList.appendChild(li);\r\n                        });\r\n                        \r\n                        document.getElementById('ods-modal').classList.remove('hidden');\r\n                    })\r\n                    .catch(error => {\r\n                        console.error('Erro ao carregar detalhes do ODS:', error);\r\n                        alert('Erro ao carregar detalhes do ODS');\r\n                    });\r\n            };\r\n            \r\n            // Função para fechar modal ODS\r\n            window.closeODSModal = function() {\r\n                document.getElementById('ods-modal').classList.add('hidden');\r\n            };\r\n        </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
