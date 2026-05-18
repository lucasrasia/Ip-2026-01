package routes

import (
	"net/http"

	"projeto-teste-crud/handlers"
)

func Register(mux *http.ServeMux, h *handlers.ConsultaHandler) {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))

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

func RegisterFrontend(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "web/index.html")
	})
}
