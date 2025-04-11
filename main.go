package main

import (
	"KaduHod/muscles_api/src/auth"
	"KaduHod/muscles_api/src/cache"
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "KaduHod/muscles_api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"

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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your Token.
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
    cacheService := cache.CacheService{
        Redis: redis,
    }
    controller := controllers.Controller{
        UserRepository: &userRepository,
        SessionService: &sessionService,
        GitHubService: &githubService,
        TokenService: &tokenService,
        TokenRepository: &tokenRepository,
        CacheService: &cacheService,
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
    server.Get("/auth/github", loginController.Auth)
    server.Get("/docs/*", httpSwagger.Handler(
        httpSwagger.URL("http://localhost:3005/docs/doc.json"), //The url pointing to API definition
    ))
    server.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
    server.Get("/", controller.Index)
    server.Group(func(r chi.Router) {
        r.Use(csrfService.Middleware)
        r.Use(cors.Handler(cors.Options{
            // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
            AllowedOrigins:   []string{"https://*", "http://*"},
            // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
            AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
            AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
            ExposedHeaders:   []string{"Link"},
            AllowCredentials: false,
            MaxAge:           300, // Maximum value not ignored by any of major browsers
        }))
        r.Get("/tokens", userController.ListTokens)
        r.Post("/token", userController.CreateToken)
        r.Delete("/token/{id}", userController.DeleteToken)
    })
    server.Get("/info", controller.Info)
    server.Get("/dashboard", controller.Dashboard)
    server.Group(func(r chi.Router) {
        r.Use(authService.Middleware)
        r.Use(cacheService.Middleware)
        r.Get("/api/v1/muscles/groups", musculoSkeletalController.ListMuscleGroups)
        r.Get("/api/v1/muscles/portions", musculoSkeletalController.ListMusclePortions)
        r.Get("/api/v1/muscles/movement-map", musculoSkeletalController.ListAmm)
        r.Get("/api/v1/muscles", musculoSkeletalController.ListMuscles)
        r.Get("/api/v1/joints", musculoSkeletalController.ListJoints)
        r.Get("/api/v1/movements", musculoSkeletalController.ListMoviments)
    })
    http.ListenAndServe(":3005", server)
}
