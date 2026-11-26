# API de Gerenciamento de Receitas Culinárias A3

Estrutura inicial do projeto acadêmico.

## Tecnologias
- Go
- net/http
- PostgreSQL
- database/sql
- JSON
- Autenticação por API KEY

## Organização
- `config/`: configuração da conexão com o banco.
- `models/`: entidades Receita e Ingrediente.
- `handlers/`: tratamento das requisições HTTP.
- `services/`: regras de negócio.
- `dao/`: interfaces e acesso aos dados.
- `middleware/`: autenticação por API KEY.
- `routes/`: registro das rotas.
- `database/`: script SQL do banco.
