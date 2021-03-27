# Descrição

O sistema é uma API REST que tem como principal responsabilidade a gerencia de dados de favorecidos e sua respectiva conta bancária.
É possível salvar, editar, deletar, e buscar os favorecidos, inclusive de forma paginada para uma melhor performance na busca pelo cliente.

# Pré-requisitos

- Go 1.16
- Sqlite3

# Dependências

- github.com/go-playground/universal-translator
- github.com/google/uuid v1.2.0
- github.com/gorilla/mux v1.8.0
- github.com/leodido/go-urn v1.2.1
- github.com/mattn/go-sqlite3 v1.14.6
- github.com/stretchr/testify v1.7.0
- gopkg.in/go-playground/validator.v9

# Como executar

Para executar o projeto é necessário rodar o comando abaixo na raix do projeto

```sh
go run main.go
```

Após executado o comando a API estará disponível na porta local 8080

# Exemplo de requisições

## Buscar dados do favorecido

```sh
curl -X GET http://localhost:8080/api/useraccount/58dbd91e-a89f-4520-8f6b-7d60b5ae7e1c
```

## Buscar dados de favorecidos com paginação

```sh
curl -X GET http://localhost:8080/api/useraccount?limit=10&offset=0
```

## Salvar novo favorecido

```sh
curl -X POST 'http://localhost:8080/api/useraccount' -H 'Content-Type: application/json' --data-raw '{
    "name": "felipe teste",
    "email": "email@teste.com",
    "document": "06271429904",
    "document_type": "NATURAL",
    "bank": "Banco do Brasil",
    "agency": 459763,
    "agency_digit": 2,
    "account_number": 9475634,
    "account_digit": 1,
    "account_type": "corrente",
    "status": "rascunho"
}'
```

## Atualizar favorecido

```sh
curl -X PUT 'http://localhost:8080/api/useraccount/5810e83a-88fe-45e5-bf41-aa36c39298be' -H 'Content-Type: application/json' --data-raw '{
    "name": "felipe teste",
    "email": "teste@teste.com",
    "document": "06271429904",
    "bank": "Banco do Brasil",
    "agency": 459763,
    "agency_digit": 2,
    "account_number": 9475634,
    "account_digit": 1,
    "account_type": "corrente",
    "status": "validado"
}'
```

## Deletar favorecido

```sh
curl -X DELETE 'http://localhost:8080/api/useraccount/5810e83a-88fe-45e5-bf41-aa36c39298be'
```

# Testes

Para rodar os testes unitários e de integração executar o comando abaixo na raiz do projeto

```sh
go test ./...
```