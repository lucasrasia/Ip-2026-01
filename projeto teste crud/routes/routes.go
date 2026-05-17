package routes

import (
	"net/http"

	"projeto-teste-crud/handlers"
)

// Register configura rotas REST de consultas.
func Register(mux *http.ServeMux, h *handlers.ConsultaHandler) {
	mux.HandleFunc("/consultas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.ListarConsultas(w, r)
		case http.MethodPost:
			h.CriarConsulta(w, r)
		default:
			handlers.Error(w, http.StatusMethodNotAllowed, "método não permitido")
		}
	})

	mux.HandleFunc("/consultas/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.BuscarConsulta(w, r)
		case http.MethodPut:
			h.AtualizarConsulta(w, r)
		case http.MethodDelete:
			h.DeletarConsulta(w, r)
		default:
			handlers.Error(w, http.StatusMethodNotAllowed, "método não permitido")
		}
	})
}
