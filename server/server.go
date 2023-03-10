package server

import (
	"chi-webapi/server/configuration"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Run function to start server instance with config & router.
func Run(c *configuration.Config, r *chi.Mux) error {
	// Prepare server with CloudFlare recommendation timeouts config.
	// See: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Handler:      r,
		Addr:         c.Server.Addr,
		ReadTimeout:  c.Server.ReadTimeout,
		WriteTimeout: c.Server.WriteTimeout,
		IdleTimeout:  c.Server.IdleTimeout,
	}

	// Start server.
	return server.ListenAndServe()
}
