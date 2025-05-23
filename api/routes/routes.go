package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/handlers"
	"github.com/codinomello/weebie-go/api/middleware"
)

type Router struct {
	mux *http.ServeMux
}

// Utilitário para mapear métodos
type MethodSwitch struct {
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
) http.Handler {
	// Inicializa handlers
	authHandler := handlers.NewAuthHandler(authCtrl)
	userHandler := handlers.NewUserHandler(userCtrl)
	//projectHandler := handlers.NewProjectHandler(projectCtrl)
	//memberHandler := handlers.NewMemberHandler(memberCtrl)

	// Cria router principal
	mainRouter := http.NewServeMux()

	// 1. Rotas estáticas (HTML e arquivos)
	staticRouter := http.NewServeMux()
	handlers.HandleTemplPublicRoutes(staticRouter)

	// Serve arquivos estáticos (imagens e JS)
	fileServer := http.FileServer(http.Dir("../images"))
	staticRouter.Handle("/images/", http.StripPrefix("/images/", fileServer))

	// 2. Rotas de API
	apiRouter := http.NewServeMux()

	// 2.1 Rotas de autenticação (/api/auth)
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("/api/auth/register", MethodSwitch{
		Post: authHandler.RegisterUser(),
	}.ServeHTTP)

	authRouter.HandleFunc("/api/auth/token", MethodSwitch{
		Post: authHandler.CreateToken(),
	}.ServeHTTP)

	authRouter.HandleFunc("/api/auth/verify", MethodSwitch{
		Post: authHandler.VerifyToken(),
	}.ServeHTTP)

	authRouter.HandleFunc("/api/auth/session", MethodSwitch{
		Delete: authHandler.RevokeSession(),
	}.ServeHTTP)

	authRouter.HandleFunc("/api/auth/refresh", MethodSwitch{
		Post: authHandler.RefreshToken(),
	}.ServeHTTP)

	// Aplica middlewares às rotas de autenticação
	authWithMiddlewares := middleware.CORS(authRouter)
	authWithMiddlewares = middleware.JSONContentType(authWithMiddlewares)
	apiRouter.Handle("/api/auth/", authWithMiddlewares)

	// 2.2 Rotas protegidas
	protectedRouter := http.NewServeMux()
	protectedRouter.HandleFunc("/api/users/{uid}", MethodSwitch{
		Get:    userHandler.GetUser(),
		Put:    userHandler.UpdateUser(),
		Delete: userHandler.DeleteUser(),
	}.ServeHTTP)

	protectedRouter.HandleFunc("/api/projects", MethodSwitch{
		//Get:  projectHandler.GetProjects(),
		//Post: projectHandler.CreateProject(),
	}.ServeHTTP)

	protectedRouter.HandleFunc("/api/members", MethodSwitch{
		//Get: memberHandler.GetMembers(),
	}.ServeHTTP)

	// Aplica middleware de autenticação nas rotas protegidas
	protectedWithAuth := middleware.AuthMiddleware(protectedRouter)
	apiRouter.Handle("/api/users/", protectedWithAuth)
	apiRouter.Handle("/api/projects/", protectedWithAuth)
	apiRouter.Handle("/api/members/", protectedWithAuth)

	// Aplica middlewares gerais da API
	apiWithMiddlewares := middleware.CORS(apiRouter)
	apiWithMiddlewares = middleware.JSONContentType(apiWithMiddlewares)

	// 3. Monta estrutura final de roteamento
	mainRouter.Handle("/", staticRouter)
	mainRouter.Handle("/api/", apiWithMiddlewares)

	// Aplica middleware de logging global
	return middleware.LoggingMiddleware(mainRouter)
}

func (m MethodSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
}
