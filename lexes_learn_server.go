package main

import (
	"context"
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"lexes_learn_server/common"
	"lexes_learn_server/controllers/healthcheck"
	"lexes_learn_server/controllers/register"
	"lexes_learn_server/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Initialize system variables
	common.StartUp()


	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Add your routes as needed
	r := mux.NewRouter()

	routes := router.Routes{
		router.Route{
			Name:    "Index",
			Path:    "/",
			Method:  http.MethodGet,
			Handler: healthcheck.IndexHandler,
		},
		router.Route{
			Name:    "Login",
			Path:    "/users/login",
			Method:  http.MethodGet,
			Handler: nil,
		},
		router.Route{
			Name:    "Register",
			Path:    "/users/register",
			Method:  http.MethodGet,
			Handler: register.Register,
		},
	}

	for _, route := range routes {
		r.Name(route.Name).
			Methods(route.Method).
			Path(route.Path).
			Handler(route.Handler)
	}

	// CORS (Cross Origin Resource sharing)
	// Makes this server accessible by Javascript in browser client side
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{
		"Content-Type",
		"Content-Length",
		"Content-Event-Type",
		"Accept",
		"X-CSRF-Token",
		"Accept-Encoding",
		"Authorization",
	})
	methods := handlers.AllowedMethods([]string{
		http.MethodPost,
		http.MethodGet,
		http.MethodPut,
		http.MethodOptions,
	})

	server := &http.Server{
		Addr:  common.AppConfig.Server,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: handlers.CORS(origins, headers, methods)(r), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server running on port 8000")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)
	// Block until we receive our signal.
	<-channel
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run server.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

