package handlers


-- Criação da tabela de receitas
CREATE TABLE receitas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    descricao TEXT,
    tempo_preparo INTEGER NOT NULL,
    rendimento INTEGER NOT NULL,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criação da tabela de ingredientes
CREATE TABLE ingredientes (
    id SERIAL PRIMARY KEY,
    receita_id INTEGER NOT NULL,
    nome VARCHAR(100) NOT NULL,
    quantidade DECIMAL(10,2) NOT NULL,
    unidade_medida VARCHAR(50) NOT NULL,

    CONSTRAINT fk_receita
        FOREIGN KEY (receita_id)
        REFERENCES receitas(id)
        ON DELETE CASCADE
);
