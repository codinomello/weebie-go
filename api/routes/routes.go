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
	handlers.HandleTemplPrivateRoutes(staticRouter)

	// Serve arquivos estáticos (imagens e JS)
	imagesFileServer := http.FileServer(http.Dir("../images"))
	staticRouter.Handle("/images/", http.StripPrefix("/images/", imagesFileServer))

	scriptsFileServer := http.FileServer(http.Dir("../scripts"))
	staticRouter.Handle("/scripts/", http.StripPrefix("/scripts/", scriptsFileServer))

	// 2. Rotas de API
	apiRouter := http.NewServeMux()

	// 2.1 Rotas de autenticação (/api/auth) - SEM MIDDLEWARE DE AUTH
	authRouter := http.NewServeMux()

	// Registrar rotas de autenticação diretamente no authRouter
	authRouter.HandleFunc("/register", MethodSwitch{
		Post: authHandler.RegisterUser(),
	}.ServeHTTP)

	authRouter.HandleFunc("/login", MethodSwitch{
		Post: authHandler.LoginWithToken(),
	}.ServeHTTP)

	authRouter.HandleFunc("/token", MethodSwitch{
		Post: authHandler.CreateToken(),
	}.ServeHTTP)

	authRouter.HandleFunc("/verify", MethodSwitch{
		Post: authHandler.VerifyToken(),
	}.ServeHTTP)

	authRouter.HandleFunc("/session", MethodSwitch{
		Delete: authHandler.RevokeSession(),
	}.ServeHTTP)

	authRouter.HandleFunc("/refresh", MethodSwitch{
		Post: authHandler.RefreshToken(),
	}.ServeHTTP)

	// Aplica middlewares às rotas de autenticação (CORS e JSON apenas)
	authWithMiddlewares := middleware.CORS(authRouter)
	authWithMiddlewares = middleware.JSONContentType(authWithMiddlewares)

	// Monta as rotas de auth no apiRouter
	apiRouter.Handle("/auth/", http.StripPrefix("/auth", authWithMiddlewares))

	// 2.2 Rotas protegidas
	protectedRouter := http.NewServeMux()
	protectedRouter.HandleFunc("/user/{uid}", MethodSwitch{
		Get:    userHandler.GetUser(),
		Put:    userHandler.UpdateUser(),
		Delete: userHandler.DeleteUser(),
	}.ServeHTTP)

	protectedRouter.HandleFunc("/project/{uid}", MethodSwitch{
		//Get:  projectHandler.GetProjects(),
		//Post: projectHandler.CreateProject(),
	}.ServeHTTP)

	protectedRouter.HandleFunc("/member/{uid}", MethodSwitch{
		//Get: memberHandler.GetMembers(),
	}.ServeHTTP)

	// Aplica middleware de autenticação nas rotas protegidas
	protectedWithAuth := middleware.AuthMiddleware(protectedRouter)
	protectedWithCORS := middleware.CORS(protectedWithAuth)
	protectedWithJSON := middleware.JSONContentType(protectedWithCORS)

	apiRouter.Handle("/user/", http.StripPrefix("/user", protectedWithJSON))
	apiRouter.Handle("/project/", http.StripPrefix("/project", protectedWithJSON))
	apiRouter.Handle("/member/", http.StripPrefix("/member", protectedWithJSON))

	// 3. Monta estrutura final de roteamento
	mainRouter.Handle("/", staticRouter)
	mainRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))

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
