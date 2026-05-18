package main

import (
	"log"
	"net/http"

	"projeto-teste-crud/config"
	"projeto-teste-crud/db"
	"projeto-teste-crud/handlers"
	"projeto-teste-crud/routes"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := db.Connect(cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("falha ao conectar no banco: %v", err)
	}
	defer conn.Close()

	h := handlers.NewConsultaHandler(conn)
	mux := http.NewServeMux()
	routes.Register(mux, h)
	routes.RegisterFrontend(mux)

	addr := ":" + cfg.ServerPort
	log.Printf("API rodando em %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("erro no servidor: %v", err)
	}
}
