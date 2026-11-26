package dao

import (
	"context"
	"errors"

	"api-receitas/models"
)

var ErrNotFound = errors.New("receita ou ingrediente não encontrado")

type IngredienteDAO interface {
	ListarPorReceita(ctx context.Context, receitaID int64) ([]models.Ingrediente, error)
	Criar(ctx context.Context, receitaID int64, input models.IngredienteInput) (models.Ingrediente, error)
	Atualizar(ctx context.Context, id int64, input models.IngredienteInput) (models.Ingrediente, error)
	Excluir(ctx context.Context, id int64) error
}
