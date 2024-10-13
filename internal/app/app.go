package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/nosliwmichael/go-rest-api/internal/configuration"
	"github.com/nosliwmichael/go-rest-api/internal/handler"
	"github.com/nosliwmichael/go-rest-api/internal/middleware"
	"github.com/nosliwmichael/go-rest-api/internal/repo"
	"github.com/nosliwmichael/go-rest-api/internal/service"
)

type (
	App struct {
		config *configuration.Config
		server http.Server
		stop   chan os.Signal

		userRepo    repo.UserRepo
		userService service.UserService
		userHandler handler.UserHandler
	}
)

func NewApp() *App {
	app := &App{
		stop:   make(chan os.Signal, 1),
		config: configuration.LoadConfigs(),
	}
	app.createUserService()
	app.createRouter()

	return app
}

func (app *App) createUserService() {
	app.userRepo = repo.NewUserRepo()
	app.userService = service.NewUserService(app.userRepo)
	app.userHandler = handler.NewUserHandler(app.userService)
}

func (app *App) createRouter() {
	r := mux.NewRouter()
	r.Use(middleware.DefaultHeaderMiddleware, middleware.LogMiddleware)
	r.HandleFunc(app.config.Endpoints.User, app.userHandler.AddUser).Methods("POST")
	r.HandleFunc(app.config.Endpoints.UserByName, app.userHandler.GetUser).Methods("GET")
	app.server = http.Server{Addr: app.config.Address, Handler: r}
}

func (app *App) Start() {
	go app.startServer()
	signal.Notify(app.stop, syscall.SIGINT, syscall.SIGTERM)
	<-app.stop
	app.cleanup()
}

func (app *App) startServer() {
	log.Printf("Starting %s on %s", app.config.AppName, app.config.Address)
	if err := app.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (app *App) cleanup() {
	log.Println("Server gracefully shutdown")
}
