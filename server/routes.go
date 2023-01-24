package server

import (
	"chi-webapi/server/healthcheck"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetRoutes(m *chi.Mux) {
	healthcheck.Routes(m)
	m.NotFound(http.NotFound)
}
