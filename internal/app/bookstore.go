package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/controllers"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/repositories"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/services"

	"github.com/rs/zerolog"
)

const (
	gracefullTTL = time.Duration(5 * time.Second)
	servicePort  = "8080"

	listBooksPath = "/books"

	contentTypeHeaderKey                  = "Content-Type"
	contentTypeHeaderValueApplicationJSON = "application/json"
)

type Controllers interface {
	ListBooks(http.ResponseWriter, *http.Request)
}

type app struct {
	c   Controllers
	log *zerolog.Logger
}

func StartService() {
	logger := zerolog.New(os.Stdout)
	repositories := repositories.NewBookstoreRepository(&logger)
	services := services.NewBookstoreService(repositories, &logger)
	controllers := controllers.NewBookstoreController(services, &logger)
	app := &app{
		c:   controllers,
		log: &logger,
	}
	app.startServer()
}

func (a *app) startServer() {
	r := mux.NewRouter()
	a.addMiddlewares(r)
	a.addRoutes(r)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", servicePort),
		// Good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Running server in a goroutine so that it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			a.log.Fatal().Err(err).Msg("error starting server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	a.log.Info().Msg("Server up and running!")

	// Block waiting for stopping signal
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), gracefullTTL)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	srv.Shutdown(ctx)
	a.log.Info().Msg("Server shutdown complete!")
	os.Exit(0)
}

func (a *app) addRoutes(r *mux.Router) {
	r.HandleFunc(listBooksPath, a.c.ListBooks).Methods(http.MethodGet)
}

func (a *app) addMiddlewares(r *mux.Router) {
	r.Use(loggingMiddleware)
	r.Use(setContentHeaderMiddleware)
}
