package healthcheck

import "github.com/go-chi/chi/v5"

const (
	groupURL  = "/hc"
	statusURL = "/status"
)

func Routes(m *chi.Mux) {
	m.Route(groupURL, func(r chi.Router) {
		r.Get(statusURL, getStatus)
	})
}
