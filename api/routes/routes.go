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
) http.Handler {
	// Inicializa handlers
	authHandler := handlers.NewAuthHandler(authCtrl)
	userHandler := handlers.NewUserHandler(userCtrl)
	projectHandler := handlers.NewProjectHandler(projectCtrl)
	// memberHandler := handlers.NewMemberHandler(memberCtrl)
	// odsHandler := handlers.NewODSHandler(odsCtrl)

	// Cria router principal
	mainRouter := http.NewServeMux()

	// 1. Rotas estáticas (HTML e arquivos)
	staticRouter := http.NewServeMux()
	handlers.HandleTemplPublicRoutes(staticRouter)
	handlers.HandleTemplPrivateRoutes(staticRouter)

	// Serve arquivos estáticos (imagens e JS)
	imagesFileServer := http.FileServer(http.Dir("../images"))
	staticRouter.Handle("/images/", http.StripPrefix("/images/", imagesFileServer))

	scriptsFileServer := http.FileServer(http.Dir("../scripts"))
	staticRouter.Handle("/scripts/", http.StripPrefix("/scripts/", scriptsFileServer))

	// 2. Rotas de autenticação (/api/auth)
	authRouter := http.NewServeMux()

	// Registrar rotas de autenticação diretamente no authRouter
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

	// Aplica middlewares às rotas de autenticação (CORS e JSON apenas)
	authWithMiddlewares := middleware.CORS(authRouter)
	authWithMiddlewares = middleware.JSONContentType(authWithMiddlewares)

	// 2.1 Rotas de API
	apiRouter := http.NewServeMux()

	// Monta as rotas de auth no apiRouter
	apiRouter.Handle("/auth/", http.StripPrefix("/auth", authWithMiddlewares))

	// 2.2 Rotas protegidas
	protectedRouter := http.NewServeMux()

	// Rotas de usuários
	protectedRouter.HandleFunc("/user/{uid}", HTTPMethod{
		Get:    userHandler.GetUser(),
		Put:    userHandler.UpdateUser(),
		Delete: userHandler.DeleteUser(),
	}.ServeHTTP)

	// Rotas de projetos
	protectedRouter.HandleFunc("/project", HTTPMethod{
		Get:    projectHandler.GetProject(),
		Post:   projectHandler.CreateProject(),
		Put:    projectHandler.UpdateProject(),
		Delete: projectHandler.DeleteProject(),
	}.ServeHTTP)

	// Rotas de membros
	protectedRouter.HandleFunc("/member/{uid}", HTTPMethod{
		// Get: memberHandler.GetMember(),
	}.ServeHTTP)

	// Rotas de ODS
	protectedRouter.HandleFunc("/ods/", HTTPMethod{
		// Get: odsHandler.GetAllODS(),
	}.ServeHTTP)

	// Aplica middleware de autenticação nas rotas protegidas
	protectedWithAuth := middleware.AuthMiddleware(protectedRouter)
	protectedWithCORS := middleware.CORS(protectedWithAuth)
	protectedWithJSON := middleware.JSONContentType(protectedWithCORS)

	apiRouter.Handle("/user/", http.StripPrefix("/user", protectedWithJSON))
	apiRouter.Handle("/project/", http.StripPrefix("/project", protectedWithJSON))
	apiRouter.Handle("/member/", http.StripPrefix("/member", protectedWithJSON))
	apiRouter.Handle("/ods/", http.StripPrefix("/ods", protectedWithJSON))

	// 3. Monta estrutura final de roteamento
	mainRouter.Handle("/", staticRouter)
	mainRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// Aplica middleware de logging global
	return middleware.LoggingMiddleware(mainRouter)
}

func (m HTTPMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}

func LogAvailableRoutes() {
	// Rotas de autenticação (/api/auth)
	log.Println("💫 rotas disponíveis:")
	log.Println(" ➕ POST   /api/auth/register") // Registrar novo usuário
	log.Println(" ➕ POST   /api/auth/login")    // Login com token Firebase
	log.Println(" ➕ POST   /api/auth/social")   // Login social (Google/GitHub)
	log.Println(" ➕ POST   /api/auth/token")    // Criar token JWT
	log.Println(" 🔍 POST   /api/auth/verify")   // Verificar token
	log.Println(" ❌ DELETE /api/auth/session")  // Revogar sessão (logout)
	log.Println(" 🔄 POST   /api/auth/refresh")  // Refresh token

	// Rotas protegidas de usuários (/api/user)
	log.Println(" 🔍 GET    /api/user/{uid}")   // Obter usuário
	log.Println(" ✏️  PUT    /api/user/{uid}") // Atualizar usuário
	log.Println(" ❌ DELETE /api/user/{uid}")   // Deletar usuário

	// Rotas protegidas de projetos (/api/project)
	log.Println(" 🔍 GET    /api/project/{uid}")   // Obter projeto
	log.Println(" ➕ POST   /api/project/{uid}")   // Criar projeto
	log.Println(" ✏️  PUT    /api/project/{uid}") // Atualizar projeto
	log.Println(" ❌ DELETE /api/project/{uid}")   // Deletar projeto

	// Rotas protegidas de membros (/api/member)
	log.Println(" 🔍 GET    /api/member/{uid}") // Obter membro (comentado no código)
}
