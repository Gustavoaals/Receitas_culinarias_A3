package handlers

CREATE TABLE receitas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    descricao TEXT,
    tempo_preparo INTEGER NOT NULL,
    rendimento INTEGER NOT NULL,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);