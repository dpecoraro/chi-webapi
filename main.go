package main

import (
	"chi-webapi/server"
	"chi-webapi/server/configuration"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Configuration file not found!")
	}
	// Create router.
	r := chi.NewRouter()

	// Create config.
	c := configuration.Bootstrap()

	// Set a logger middleware.
	r.Use(middleware.Logger)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(c.Server.ReadTimeout))

	// Get router with all routes.
	server.GetRoutes(r)

	// Run server instance.
	if err := server.Run(c, r); err != nil {
		log.Fatal(err)
		return
	}
}
