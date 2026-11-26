package models

type Ingrediente struct {
	ID            int64   `json:"id"`
	ReceitaID     int64   `json:"receita_id"`
	Nome          string  `json:"nome"`
	Quantidade    float64 `json:"quantidade"`
	UnidadeMedida string  `json:"unidade_medida"`
}

type IngredienteInput struct {
	Nome          string  `json:"nome"`
	Quantidade    float64 `json:"quantidade"`
	UnidadeMedida string  `json:"unidade_medida"`
}
