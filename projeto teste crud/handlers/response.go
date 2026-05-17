package handlers

import (
	"encoding/json"
	"net/http"
)

// JSON padroniza respostas JSON da API.
func JSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// Error devolve erros em um formato simples e consistente.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"erro": message})
}
