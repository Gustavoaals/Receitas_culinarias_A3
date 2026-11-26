package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-receitas/config"
	"api-receitas/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	db, err := config.ConnectPostgres(ctx, cfg)
	cancel()
	if err != nil {
		log.Fatalf("não foi possível conectar ao PostgreSQL: %v", err)
	}
	defer db.Close()

	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      routes.New(db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Servidor iniciado em http://localhost%s", cfg.HTTPAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("erro ao iniciar servidor: %v", err)
		}
	}()

	<-shutdownCtx.Done()

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("erro ao encerrar servidor: %v", err)
	}
}
