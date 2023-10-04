package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/controllers"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/repositories"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/services"
)

const (
	servicePort   = "8080"
	listBooksPath = "/books"
)

type Controllers interface {
	ListBooks(http.ResponseWriter, *http.Request)
}

type app struct {
	c Controllers
}

func StartService() {
	app := &app{
		c: controllers.NewBookstoreController(services.NewBookstoreService(repositories.NewBookstoreRepository())),
	}
	app.startServer()
}

func (a *app) startServer() {
	gracefullTTL := time.Duration(5 * time.Second)
	r := mux.NewRouter()
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
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	log.Println("Server up and running!")

	// Block waiting for stopping signal
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), gracefullTTL)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	srv.Shutdown(ctx)
	log.Println("Server shutdown complete!")
	os.Exit(0)
}

func (a *app) addRoutes(r *mux.Router) {
	r.HandleFunc(listBooksPath, a.c.ListBooks).Methods(http.MethodGet)
}
