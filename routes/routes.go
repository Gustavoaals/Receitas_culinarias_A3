package routes

import (
	"net/http"

	"api-receitas/dao"
	"api-receitas/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()
	ingredienteDAO := dao.NewIngredienteDAO(db)
	ingredienteHandler := handlers.NewIngredienteHandler(ingredienteDAO)

	mux.HandleFunc("GET /receitas/{id}/ingredientes", ingredienteHandler.Listar)
	mux.HandleFunc("POST /receitas/{id}/ingredientes", ingredienteHandler.Criar)
	mux.HandleFunc("PUT /ingredientes/{id}", ingredienteHandler.Atualizar)
	mux.HandleFunc("DELETE /ingredientes/{id}", ingredienteHandler.Excluir)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(r.Context()); err != nil {
			http.Error(w, "Banco de dados indisponível", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Servidor e banco de dados funcionando!"))
	})

	return mux
}
