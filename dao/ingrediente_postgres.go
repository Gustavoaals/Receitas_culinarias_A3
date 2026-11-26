package dao

import (
	"context"
	"errors"

	"api-receitas/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ingredientePostgres struct {
	db *pgxpool.Pool
}

func NewIngredienteDAO(db *pgxpool.Pool) IngredienteDAO {
	return &ingredientePostgres{db: db}
}

func (d *ingredientePostgres) ListarPorReceita(ctx context.Context, receitaID int64) ([]models.Ingrediente, error) {
	var existe bool
	if err := d.db.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM receitas WHERE id = $1)`, receitaID).Scan(&existe); err != nil {
		return nil, err
	}
	if !existe {
		return nil, ErrNotFound
	}

	rows, err := d.db.Query(ctx, `
		SELECT id, receita_id, nome, quantidade, unidade_medida
		FROM ingredientes
		WHERE receita_id = $1
		ORDER BY id
	`, receitaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itens := make([]models.Ingrediente, 0)
	for rows.Next() {
		var item models.Ingrediente
		if err := rows.Scan(&item.ID, &item.ReceitaID, &item.Nome, &item.Quantidade, &item.UnidadeMedida); err != nil {
			return nil, err
		}
		itens = append(itens, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return itens, nil
}

func (d *ingredientePostgres) Criar(ctx context.Context, receitaID int64, input models.IngredienteInput) (models.Ingrediente, error) {
	var item models.Ingrediente
	err := d.db.QueryRow(ctx, `
		INSERT INTO ingredientes (receita_id, nome, quantidade, unidade_medida)
		SELECT id, $2, $3, $4
		FROM receitas
		WHERE id = $1
		RETURNING id, receita_id, nome, quantidade, unidade_medida
	`, receitaID, input.Nome, input.Quantidade, input.UnidadeMedida).Scan(
		&item.ID, &item.ReceitaID, &item.Nome, &item.Quantidade, &item.UnidadeMedida,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return models.Ingrediente{}, ErrNotFound
	}
	return item, err
}

func (d *ingredientePostgres) Atualizar(ctx context.Context, id int64, input models.IngredienteInput) (models.Ingrediente, error) {
	var item models.Ingrediente
	err := d.db.QueryRow(ctx, `
		UPDATE ingredientes
		SET nome = $2, quantidade = $3, unidade_medida = $4
		WHERE id = $1
		RETURNING id, receita_id, nome, quantidade, unidade_medida
	`, id, input.Nome, input.Quantidade, input.UnidadeMedida).Scan(
		&item.ID, &item.ReceitaID, &item.Nome, &item.Quantidade, &item.UnidadeMedida,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return models.Ingrediente{}, ErrNotFound
	}
	return item, err
}

func (d *ingredientePostgres) Excluir(ctx context.Context, id int64) error {
	result, err := d.db.Exec(ctx, `DELETE FROM ingredientes WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
