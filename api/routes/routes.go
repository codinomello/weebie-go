package routes

import (
	"log"
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/handlers"
	"github.com/codinomello/weebie-go/api/middleware"
)

type Router struct {
	mux *http.ServeMux
}

// Utilitário para mapear métodos
type HTTPMethod struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Put    http.HandlerFunc
	Delete http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

func SetupRoutes(
	authCtrl *controllers.AuthController,
	userCtrl *controllers.UserController,
	projectCtrl *controllers.ProjectController,
	memberCtrl *controllers.MemberController,
	odsCtrl *controllers.ODSController,
	profileCtrl *controllers.ProfileController,
) http.Handler {
	// Inicializa handlers
	authHandler := handlers.NewAuthHandler(authCtrl)
	userHandler := handlers.NewUserHandler(userCtrl)
	projectHandler := handlers.NewProjectHandler(projectCtrl)
	templHandler := handlers.NewTemplHandler(profileCtrl)
	// memberHandler := handlers.NewMemberHandler(memberCtrl)
	// odsHandler := handlers.NewODSHandler(odsCtrl)

	// Cria router principal
	mainRouter := http.NewServeMux()

	// 1. Rotas estáticas (HTML e arquivos)
	staticRouter := http.NewServeMux()
	handlers.HandleTemplPublicRoutes(staticRouter)
	templHandler.HandleTemplPrivateRoutes(staticRouter) // Corrija para usar o método do handler

	// Serve arquivos estáticos (imagens e JS)
	imagesFileServer := http.FileServer(http.Dir("../images"))
	staticRouter.Handle("/images/", http.StripPrefix("/images/", imagesFileServer))

	scriptsFileServer := http.FileServer(http.Dir("../scripts"))
	staticRouter.Handle("/scripts/", http.StripPrefix("/scripts/", scriptsFileServer))

	// 2. Rotas de API (/api)
	apiRouter := http.NewServeMux()

	// 2.1 Rotas de autenticação (/api/auth) - SEM autenticação
	authRouter := http.NewServeMux()

	authRouter.HandleFunc("/register", HTTPMethod{
		Post: authHandler.RegisterUser(),
	}.ServeHTTP)
	authRouter.HandleFunc("/login", HTTPMethod{
		Post: authHandler.LoginWithToken(),
	}.ServeHTTP)
	authRouter.HandleFunc("/social", HTTPMethod{
		Post: authHandler.LoginWithSocial(),
	}.ServeHTTP)
	authRouter.HandleFunc("/token", HTTPMethod{
		Post: authHandler.CreateToken(),
	}.ServeHTTP)
	authRouter.HandleFunc("/verify", HTTPMethod{
		Post: authHandler.VerifyToken(),
	}.ServeHTTP)
	authRouter.HandleFunc("/session", HTTPMethod{
		Delete: authHandler.RevokeSession(),
	}.ServeHTTP)
	authRouter.HandleFunc("/refresh", HTTPMethod{
		Post: authHandler.RefreshToken(),
	}.ServeHTTP)

	// Aplica middlewares às rotas de autenticação
	authWithMiddlewares := middleware.CORS(authRouter)
	authWithMiddlewares = middleware.JSONContentType(authWithMiddlewares)

	// 2.2 Rotas protegidas - COM autenticação
	protectedRouter := http.NewServeMux()

	// Rota para criar projeto (POST /api/project)
	protectedRouter.HandleFunc("/project", HTTPMethod{
		Post: projectHandler.CreateProject(),
		Get:  projectHandler.GetProject(), // Para listar projetos
	}.ServeHTTP)

	// Rota para operações com ID específico (GET/PUT/DELETE /api/project/{id})
	protectedRouter.HandleFunc("/project/{id}", HTTPMethod{
		Get:    projectHandler.GetProject(),
		Put:    projectHandler.UpdateProject(),
		Delete: projectHandler.DeleteProject(),
	}.ServeHTTP)

	// Rotas de usuários
	protectedRouter.HandleFunc("/user", HTTPMethod{
		Get:  userHandler.GetUser(),
		Post: userHandler.UpdateUser(), // Para criar/atualizar perfil
	}.ServeHTTP)

	protectedRouter.HandleFunc("/user/{uid}", HTTPMethod{
		Get:    userHandler.GetUser(),
		Put:    userHandler.UpdateUser(),
		Delete: userHandler.DeleteUser(),
	}.ServeHTTP)

	// Rotas de membros
	// protectedRouter.HandleFunc("/member/{uid}", HTTPMethod{
	// 	Get: memberHandler.GetMember(),
	// }.ServeHTTP)

	// Rotas de ODS
	// protectedRouter.HandleFunc("/ods", HTTPMethod{
	// 	Get: odsHandler.GetAllODS(),
	// }.ServeHTTP)

	// CORREÇÃO: Aplicar middlewares na ordem correta
	protectedWithAuth := middleware.AuthMiddleware(protectedRouter)
	protectedWithCORS := middleware.CORS(protectedWithAuth)
	protectedWithJSON := middleware.JSONContentType(protectedWithCORS)

	// 3. Monta as rotas no apiRouter
	apiRouter.Handle("/auth/", http.StripPrefix("/auth", authWithMiddlewares))

	// CORREÇÃO: Registrar rotas protegidas sem StripPrefix adicional
	apiRouter.Handle("/project", protectedWithJSON)
	apiRouter.Handle("/project/", protectedWithJSON)
	apiRouter.Handle("/user", protectedWithJSON)
	apiRouter.Handle("/user/", protectedWithJSON)
	// apiRouter.Handle("/member/", protectedWithJSON)
	// apiRouter.Handle("/ods", protectedWithJSON)

	// 4. Monta estrutura final de roteamento
	mainRouter.Handle("/", staticRouter)
	mainRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))
	mainRouter.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Server is running", "path": "` + r.URL.Path + `"}`))
	})

	// Aplica middleware de logging global
	return middleware.LoggingMiddleware(mainRouter)
}

func (m HTTPMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORREÇÃO: Definir Content-Type como JSON por padrão para APIs
	w.Header().Set("Content-Type", "application/json")

	// Adicionar headers CORS aqui também como backup
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if m.Get != nil {
			m.Get(w, r)
			return
		}
	case http.MethodPost:
		if m.Post != nil {
			m.Post(w, r)
			return
		}
	case http.MethodPut:
		if m.Put != nil {
			m.Put(w, r)
			return
		}
	case http.MethodDelete:
		if m.Delete != nil {
			m.Delete(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"error": "Método não permitido", "method": "` + r.Method + `", "path": "` + r.URL.Path + `"}`))
}

func LogAvailableRoutes() {
	log.Println("💫 Rotas disponíveis:")

	// Rotas de autenticação (/api/auth)
	log.Println("🔓 Rotas de autenticação:")
	log.Println(" ➕ POST   /api/auth/register") // Registrar novo usuário
	log.Println(" ➕ POST   /api/auth/login")    // Login com token Firebase
	log.Println(" ➕ POST   /api/auth/social")   // Login social (Google/GitHub)
	log.Println(" ➕ POST   /api/auth/token")    // Criar token JWT
	log.Println(" 🔍 POST   /api/auth/verify")   // Verificar token
	log.Println(" ❌ DELETE /api/auth/session")  // Revogar sessão (logout)
	log.Println(" 🔄 POST   /api/auth/refresh")  // Refresh token

	// Rotas protegidas
	log.Println("🔒 Rotas protegidas:")

	// Usuários
	log.Println(" 🔍 GET    /api/user")         // Obter usuário atual
	log.Println(" 🔍 GET    /api/user/{uid}")   // Obter usuário por UID
	log.Println(" ✏️  PUT    /api/user/{uid}") // Atualizar usuário
	log.Println(" ❌ DELETE /api/user/{uid}")   // Deletar usuário

	// Projetos - CORRIGIDO
	log.Println(" 🔍 GET    /api/project")        // Listar projetos do usuário
	log.Println(" ➕ POST   /api/project")        // Criar projeto
	log.Println(" 🔍 GET    /api/project/{id}")   // Obter projeto por ID
	log.Println(" ✏️  PUT    /api/project/{id}") // Atualizar projeto
	log.Println(" ❌ DELETE /api/project/{id}")   // Deletar projeto

	// Outros
	log.Println(" 🔍 GET    /api/member/{uid}") // Obter membro
	log.Println(" 🔍 GET    /api/ods")          // Obter todos os ODS
}
