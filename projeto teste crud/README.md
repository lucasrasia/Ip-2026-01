# Projeto Teste CRUD — Sistema de Agendamento de Consultas Médicas

Este projeto implementa uma API REST simples em Go com PostgreSQL para gerenciar consultas médicas.

## Estrutura

```text
projeto teste crud/
├── cmd/main.go
├── config/config.go
├── db/postgres.go
├── handlers/
├── models/
├── routes/
└── sql/
    ├── schema.sql
    └── seed.sql
```

## Pré-requisitos

- Go 1.22+
- PostgreSQL 13+

## Como rodar

1. Entre na pasta do projeto:

```bash
cd "projeto teste crud"
```

2. (Opcional) copie o arquivo de exemplo de ambiente:

```bash
cp .env.example .env
```

> Observação: a aplicação lê variáveis de ambiente do sistema operacional.

3. Defina as variáveis de ambiente necessárias (exemplo Linux/macOS):

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=saude
export DB_SSLMODE=disable
export SERVER_PORT=8080
```

4. Baixe as dependências e execute:

```bash
go mod tidy
go run ./cmd
```

## Configurando PostgreSQL

1. Crie o banco:

```sql
CREATE DATABASE saude;
```

2. Execute o schema:

```bash
psql -U postgres -d saude -f sql/schema.sql
```

3. (Opcional) popule dados iniciais:

```bash
psql -U postgres -d saude -f sql/seed.sql
```

## Endpoints REST

- `GET /consultas` — lista consultas
- `GET /consultas/{id}` — busca por ID
- `POST /consultas` — cria consulta
- `PUT /consultas/{id}` — atualiza consulta
- `DELETE /consultas/{id}` — remove consulta

## Exemplos de requisição

### Criar consulta

```bash
curl -X POST http://localhost:8080/consultas \
  -H "Content-Type: application/json" \
  -d '{
    "nome_paciente": "Juliana Lima",
    "cpf_paciente": "321.654.987-00",
    "telefone": "(11) 97777-1111",
    "especialidade": "Ortopedia",
    "nome_medico": "Dr. Paulo Mendes",
    "data_consulta": "2026-06-01",
    "horario": "10:30:00",
    "status": "agendada",
    "observacoes": "Primeira consulta"
  }'
```

### Listar consultas

```bash
curl http://localhost:8080/consultas
```

### Buscar por ID

```bash
curl http://localhost:8080/consultas/1
```

### Atualizar consulta

```bash
curl -X PUT http://localhost:8080/consultas/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nome_paciente": "Juliana Lima",
    "cpf_paciente": "321.654.987-00",
    "telefone": "(11) 97777-1111",
    "especialidade": "Ortopedia",
    "nome_medico": "Dra. Renata Alves",
    "data_consulta": "2026-06-01",
    "horario": "11:00:00",
    "status": "confirmada",
    "observacoes": "Horário ajustado"
  }'
```

### Deletar consulta

```bash
curl -X DELETE http://localhost:8080/consultas/1
```
