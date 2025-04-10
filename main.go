package main

import (
	_ "KaduHod/muscles_api/docs"
	"KaduHod/muscles_api/src/auth"
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/swaggo/http-swagger" // http-swagger middleware
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/joho/godotenv"
)

type Middleware func(http.Handler) http.Handler
// MethodMiddleware cria um middleware que verifica se o método HTTP da requisição
// corresponde ao método especificado.
func MethodMiddleware(method string) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Verifica se o método da requisição corresponde ao método esperado
            if r.Method != method {
                http.NotFound(w, r)
                return
            }
            // Se o método estiver correto, continua para o próximo handler
            next.ServeHTTP(w, r)
        })
    }
}

func Use(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func Logger() Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            log.Printf("[ %s ] %s",r.Method, r.URL.Path)
            next.ServeHTTP(w, r)
        })
    }
}
// @title Musculo Eskeletal Api
// @version 1.0
// @description API for Muscles System
// @host gymapi.kadu.tec.br
// @BasePath /api/v1
func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal(err)
    }
    db := database.ConnetionMysql()
    defer db.Close()
    redis := database.NewRedis()
    defer redis.Conn.Close()
    musclesRepository := repository.MuscleRepository{Db: db}
    movementRepository := repository.MovementRepository{Db: db}
    jointRepository := repository.JointRepository{Db: db}
    ammRepository := repository.AmmRepository{Db: db}
    userRepository := repository.UserRepository{Db: db}
    tokenRepository := repository.TokenRepository{Db: db}
    githubService := services.GitHubService{}
    sessionService := services.SessionService{Redis: redis}
    tokenService := services.NewTokenService(&userRepository, &tokenRepository)
    csrfService := services.NewCsrfService(&sessionService)
    authService := auth.ApiAuthService{
        TokenRepository: &tokenRepository,
        TokenService: &tokenService,
    }
    controller := controllers.Controller{
        UserRepository: &userRepository,
        SessionService: &sessionService,
        GitHubService: &githubService,
        TokenService: &tokenService,
        TokenRepository: &tokenRepository,
    }
    musculoSkeletalController := controllers.MusculoSkeletalController{
        Controller: controller,
        MuscleRepository: &musclesRepository,
        MovementRepository: &movementRepository,
        JointRepository: &jointRepository,
        AmmRepository: &ammRepository,
    }
    loginController := controllers.LoginController{
        Controller: controller,
        GitHubService: &githubService,
        SessionService: &sessionService,
        UserRepository: &userRepository,
        CsrfService: csrfService,
    }
    userController := controllers.UserController{
        Controller: controller,
        TokenService: &tokenService,
        UserRepository: &userRepository,
        SessionService: &sessionService,
    }
    server := chi.NewRouter()
    server.Use(middleware.GetHead)
    server.Use(middleware.Logger)
    server.Use(middleware.Recoverer)
    server.Use(middleware.RealIP)
    server.Group(func(r chi.Router) {
        r.Use(authService.Middleware)
        r.Get("/api/v1/muscles/groups", musculoSkeletalController.ListMuscleGroups)
        r.Get("/api/v1/muscles/portions", musculoSkeletalController.ListMusclePortions)
        r.Get("/api/v1/muscles/movement-map", musculoSkeletalController.ListAmm)
        r.Get("/api/v1/muscles", musculoSkeletalController.ListMuscles)
        r.Get("/api/v1/joints", musculoSkeletalController.ListJoints)
        r.Get("/api/v1/movements", musculoSkeletalController.ListMoviments)
    })

    server.Get("/docs/", httpSwagger.WrapHandler)
    server.Get("/", controller.Index)
    server.Get("/auth/github", loginController.Auth)
    server.Get("/tokens", userController.ListTokens)
    server.Group(func(r chi.Router) {
        r.Use(csrfService.Middleware)
        server.Post("/token", userController.CreateToken)
        server.Delete("/token/{id}", userController.DeleteToken)
    })
    server.Get("/info", controller.Info)
    server.Get("/dashboard", controller.Dashboard)
    http.ListenAndServe(":3005", server)
}
