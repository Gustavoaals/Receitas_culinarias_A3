package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"api-receitas/dao"
	"api-receitas/models"
)

type IngredienteHandler struct {
	dao dao.IngredienteDAO
}

func NewIngredienteHandler(ingredienteDAO dao.IngredienteDAO) *IngredienteHandler {
	return &IngredienteHandler{dao: ingredienteDAO}
}

func pathID(r *http.Request, key string) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue(key), 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("identificador inválido")
	}
	return id, nil
}

func decodeIngrediente(r *http.Request) (models.IngredienteInput, error) {
	var input models.IngredienteInput
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, errors.New("corpo JSON inválido")
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return input, errors.New("envie apenas um objeto JSON")
	}

	input.Nome = strings.TrimSpace(input.Nome)
	input.UnidadeMedida = strings.TrimSpace(input.UnidadeMedida)

	if input.Nome == "" {
		return input, errors.New("o nome é obrigatório")
	}
	if input.Quantidade <= 0 {
		return input, errors.New("a quantidade deve ser maior que zero")
	}
	if input.UnidadeMedida == "" {
		return input, errors.New("a unidade de medida é obrigatória")
	}
	return input, nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeDAOError(w http.ResponseWriter, err error) {
	if errors.Is(err, dao.ErrNotFound) {
		http.Error(w, "Receita ou ingrediente não encontrado", http.StatusNotFound)
		return
	}
	http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
}

func (h *IngredienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	receitaID, err := pathID(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itens, err := h.dao.ListarPorReceita(r.Context(), receitaID)
	if err != nil {
		writeDAOError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, itens)
}

func (h *IngredienteHandler) Criar(w http.ResponseWriter, r *http.Request) {
	receitaID, err := pathID(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input, err := decodeIngrediente(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item, err := h.dao.Criar(r.Context(), receitaID, input)
	if err != nil {
		writeDAOError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (h *IngredienteHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input, err := decodeIngrediente(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item, err := h.dao.Atualizar(r.Context(), id, input)
	if err != nil {
		writeDAOError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (h *IngredienteHandler) Excluir(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.dao.Excluir(r.Context(), id); err != nil {
		writeDAOError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
