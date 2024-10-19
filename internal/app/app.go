package app

import (
	"context"
	"database/sql"
	"github.com/nosliwmichael/go-rest-api/internal/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		db     *sql.DB
		stop   chan os.Signal

		router      *mux.Router
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
	app.getDatabaseInstance()
	app.createUserService()
	app.createRouter()

	return app
}

func (app *App) getDatabaseInstance() {
	app.db = database.GetDbInstance(app.config)
}

func (app *App) createUserService() {
	app.userRepo = repo.NewUserRepo(app.db)
	app.userService = service.NewUserService(app.userRepo)
	app.userHandler = handler.NewUserHandler(app.userService)
}

func (app *App) createRouter() {
	app.router = mux.NewRouter()
	app.router.Use(middleware.DefaultHeaderMiddleware, middleware.LogMiddleware)

	apiRouter := app.router.PathPrefix(app.config.Api.ContextRoot).Subrouter()
	apiRouter.HandleFunc(app.config.Api.Endpoints.User, app.userHandler.AddUser).Methods("POST")
	apiRouter.HandleFunc(app.config.Api.Endpoints.UserById, app.userHandler.GetUser).Methods("GET")
}

func (app *App) createServer() {
	app.server = http.Server{
		Addr:         app.config.HttpServer.Address,
		ReadTimeout:  app.config.HttpServer.ReadTimeout,
		WriteTimeout: app.config.HttpServer.WriteTimeout,
		Handler:      app.router,
	}
}

func (app *App) Start() {
	go app.startServer()
	signal.Notify(app.stop, syscall.SIGINT, syscall.SIGTERM)
	<-app.stop
	app.cleanup()
}

func (app *App) startServer() {
	log.Printf("Starting %s on %s", app.config.AppName, app.config.HttpServer.Address)
	if err := app.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (app *App) cleanup() {
	err := app.db.Close()
	if err != nil {
		log.Printf("Unable to close database connections: %v", err)
	} else {
		log.Println("Successfully shutdown database")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = app.server.Shutdown(ctx); err != nil {
		log.Printf("Failed to shutdown server gracefully: %v", err)
	} else {
		log.Println("Server gracefully shutdown")
	}
}
