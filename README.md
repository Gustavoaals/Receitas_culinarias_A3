<<<<<<< HEAD
# API de Gerenciamento de Receitas Culinárias A3
=======
# API de Receitas Culinárias
>>>>>>> 95ac416 (implementado o gerenciamento de ingredientes)

Back-end acadêmico em Go, com `net/http`, PostgreSQL e `pgx`.

## Requisitos

- Go 1.23 ou superior
- PostgreSQL
- Banco de dados `sistema_receitas`

## Configuração

Configure as variáveis de ambiente no seu terminal ou em um arquivo `.env` local (não versionado):

```bash
export HTTP_ADDR=:8080
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD='sua_senha_local'
export DB_NAME=sistema_receitas
export DB_SSLMODE=disable
```

O projeto lê um `.env` local se existir, mas esse arquivo está ignorado pelo Git. Não crie nem envie arquivos com senhas, tokens ou chaves para o GitHub.

## Banco de dados

Crie o banco:

```sql
CREATE DATABASE sistema_receitas;
```

Execute o esquema:

```bash
psql -U postgres -d sistema_receitas -f database/schema.sql
```

## Instalação e execução

Na raiz do projeto:

```bash
go mod tidy
go run .
```

Verifique o servidor:

```bash
curl http://localhost:8080/health
```

## Rotas de ingredientes

| Método | Rota | Descrição |
|---|---|---|
| GET | `/receitas/{id}/ingredientes` | Lista ingredientes da receita |
| POST | `/receitas/{id}/ingredientes` | Adiciona ingrediente à receita |
| PUT | `/ingredientes/{id}` | Atualiza ingrediente |
| DELETE | `/ingredientes/{id}` | Exclui ingrediente |

Exemplo de corpo para POST e PUT:

```json
{
  "nome": "Farinha de trigo",
  "quantidade": 2,
  "unidade_medida": "xícaras"
}
```

## Segurança antes de enviar ao GitHub

- Mantenha `.env` apenas na sua máquina.
- Nunca coloque senhas, tokens, API keys ou strings de conexão no código.
- Confira `git status` antes de `git add`.
- O `.gitignore` impede que novos arquivos ignorados sejam adicionados, mas não remove arquivos que já estejam rastreados pelo Git.
