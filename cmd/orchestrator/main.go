package main

import (
	"calculator_go/internal/http/middlewares"
	"calculator_go/internal/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	authHandler "calculator_go/internal/http/handlers/auth"
	calcHandler "calculator_go/internal/http/handlers/expression"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	db, err := storage.New("./database/storage.db")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/api/v1/register", authHandler.RegisterUserHandler(ctx, db))
	r.Post("/api/v1/login", authHandler.LoginUserHandler(ctx, db))
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthorizeJWTToken)

		r.Post("/api/v1/calculate", calcHandler.CreateExpressionHandler(ctx, db))
		r.Get("/api/v1/expressions", calcHandler.GetExpressionsHandler(ctx, db))
		r.Delete("/api/v1/expression/{id}", calcHandler.DeleteExpressionHandler(ctx, db))
	})

	host, ok := os.LookupEnv("ORCHESTRATOR_HOST")
	if !ok {
		log.Print("ORCHESTRATOR_HOST not set, using 0.0.0.0")
		host = "0.0.0.0"
	}

	port, ok := os.LookupEnv("ORCHESTRATOR_PORT")
	if !ok {
		log.Print("ORCHESTRATOR_PORT not set, using 8080")
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Printf("running Orchestrator server at %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
