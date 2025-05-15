package routes

import (
	"net/http"
	"strings"

	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/handlers"
	"github.com/codinomello/weebie-go/api/repository"
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

func SetupRoutes(userRepo *repository.MongoDBUserRepository, projectRepo *repository.MongoDBProjectRepository, memberRepo *repository.MongoDBMemberRepository) http.Handler {
	// Inicializa os controllers
	userController := controllers.NewUserController(userRepo)
	// projectController := controllers.NewProjectController(projectRepo, userRepo, memberRepo)
	// memberController := controllers.NewMemberController(memberRepo)

	// Inicializa handlers
	userHandler := handlers.NewUserHandler(userController)
	// projectHandler := handlers.NewProjectHandler(projectController)
	// memberHandler := handlers.NewMemberHandler(memberController)

	// Roteadores HTTP
	publicMux := http.NewServeMux()
	privateMux := http.NewServeMux()

	// Serve imagens
	fs := http.FileServer(http.Dir("../images"))
	publicMux.Handle("/images/", http.StripPrefix("/images/", fs))

	// Rotas públicas de usuário
	publicMux.HandleFunc("/users", MethodSwitch{
		Post: userHandler.SignUpUserHandler(),
	}.ServeHTTP)

	publicMux.HandleFunc("/users/signin", MethodSwitch{
		Post: userHandler.SignInUserHandler(),
	}.ServeHTTP)

	publicMux.HandleFunc("/users/signout", MethodSwitch{
		Post: userHandler.SignOutUserHandler(),
	}.ServeHTTP)

	// Rotas privadas de usuário
	privateMux.HandleFunc("/users/{id}", MethodSwitch{
		Get:    userHandler.GetUserHandler(),
		Put:    userHandler.UpdateUserHandler(),
		Delete: userHandler.DeleteUserHandler(),
	}.ServeHTTP)

	// // Rotas privadas de projeto
	// privateMux.HandleFunc("/projects", MethodSwitch{
	// 	Get:  projectHandler.GetUserProjects,
	// 	Post: projectHandler.CreateProject,
	// }.ServeHTTP)

	// privateMux.HandleFunc("/projects/{id}", MethodSwitch{
	// 	Get:    projectHandler.GetProject,
	// 	Put:    projectHandler.UpdateProject,
	// 	Delete: projectHandler.DeleteProject,
	// }.ServeHTTP)

	// // Rotas privadas de membro
	// privateMux.HandleFunc("/projects/{id}/members", MethodSwitch{
	// 	Get:  memberHandler.GetProjectMembers,
	// 	Post: memberHandler.AddMember,
	// }.ServeHTTP)

	// privateMux.HandleFunc("/projects/{id}/members/{memberId}", MethodSwitch{
	// 	Delete: memberHandler.DeleteMemberHandler,
	// }.ServeHTTP)

	// privateMux.HandleFunc("/projects/{id}/members/{memberId}/role", MethodSwitch{
	// 	Put: memberHandler.UpdateMemberRole,
	// }.ServeHTTP)

	// Rotas estáticas (HTML, CSS, JS)
	SetupPublicRoutes(publicMux)
	SetupPrivateRoutes(privateMux)

	// Aplicando middleware de autenticação
	authProtected := AuthMiddleware(privateMux)

	// Roteador principal
	mainMux := http.NewServeMux()
	mainMux.Handle("/", publicMux)
	mainMux.Handle("/projects", authProtected)
	mainMux.Handle("/projects/", authProtected)
	mainMux.Handle("/users/", authProtected)
	mainMux.Handle("/profile", authProtected)

	return mainMux
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Exemplo: validação de header (coloque autenticação real aqui)
		if strings.HasPrefix(r.URL.Path, "/users/") || strings.HasPrefix(r.URL.Path, "/projects") || r.URL.Path == "/profile" {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				http.Error(w, "não autorizado", http.StatusUnauthorized)
				return
			}
			// Aqui você deve validar o token e injetar o userID no contexto
			// Por exemplo:
			// ctx := context.WithValue(r.Context(), "userID", extrairUserIDDoToken(auth))
			// r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
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
