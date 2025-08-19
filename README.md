# Go Currency API

API em **Go (Golang)** para consultar **cotações de moedas fiat** com **fallback automático de provedor** e documentação em **Swagger**.

> ✅ **Status:** MVP funcional para portfólio  
> Rotas ativas: `GET /health`, `GET /v1/rates` e docs em `GET /docs`.

---

## ✨ Destaques do projeto

- **Gin** como framework HTTP.
- **Fallback de provedor**: tenta `exchangerate.host`; se falhar/ exigir chave, usa `frankfurter.app` (sem API key).
- **Resposta indica o provedor** utilizado (`provider`).
- **Swagger UI** publicado em `/docs` (OpenAPI autogerado via `swag`).
- Arquitetura organizada em **camadas** (`cmd/`, `internal/`).

---

## 🚀 Como executar

### Pré‑requisitos
- Go **1.25+**
- Git
- (Opcional) `swag` CLI para regerar a documentação

### Clonar e rodar
```bash
git clone https://github.com/lenonmerlo/go-currency-api.git
cd go-currency-api

# executar
go run ./cmd/api
```

A API sobe em `http://localhost:8080`

### Variáveis de ambiente
- `PORT` (opcional): porta do servidor (padrão `8080`).

---

## 📚 Documentação (Swagger)

- UI: `http://localhost:8080/docs`  
- JSON: `http://localhost:8080/docs/doc.json`

Se você **precisar regerar** os arquivos do Swagger (pasta `docs/`):

```bash
# instalar o CLI (uma vez)
go install github.com/swaggo/swag/cmd/swag@latest

# garantir que o binário do Go está no PATH (Windows, sessão atual)
# $env:Path += ";$((go env GOPATH))\bin"

# gerar os arquivos
swag init -g cmd/api/main.go -o docs --parseInternal
```

As anotações principais estão no `cmd/api/main.go` (metadados) e
no handler `internal/http/controllers/rate_controller.go`.

---

## 🧪 Endpoints

### Healthcheck
```
GET /health
```
**200** – `{"status":"ok"}`

---

### Cotações (fiat)
```
GET /v1/rates?base=BRL&symbols=USD,EUR
```
**Parâmetros**
- `base` *(string, 3 letras)* – moeda base (ex.: `BRL`, `USD`)
- `symbols` *(string)* – lista separada por vírgula (ex.: `USD,EUR`)

**Exemplo de resposta – 200**
```json
{
  "base": "BRL",
  "provider": "frankfurter.app",
  "rates": {
    "USD": 0.18515,
    "EUR": 0.15841
  }
}
```

**Erros comuns**
- **400** – parâmetros inválidos (ex.: `base` diferente de 3 letras)
- **502** – nenhum provedor retornou taxa (raro; ocorre com entrada inválida ou indisponibilidade)

---

## 🧱 Arquitetura & pastas

```
go-currency-api/
├─ cmd/
│  └─ api/
│     └─ main.go                # bootstrap do servidor + Swagger UI
├─ internal/
│  ├─ router/
│  │  └─ routes.go              # registro das rotas públicas
│  ├─ http/
│  │  └─ controllers/
│  │     └─ rate_controller.go  # handler do /v1/rates
│  ├─ services/
│  │  └─ rate_service.go        # orquestra providers e validações
│  ├─ clients/
│  │  ├─ exchangerate/
│  │  │  └─ client.go           # tenta exchangerate.host
│  │  └─ frankfurter/
│  │     └─ client.go           # fallback sem API key
│  └─ domain/
│     └─ rate.go                # DTOs/structs de resposta
├─ docs/                        # artefatos do Swagger (gerados)
├─ go.mod / go.sum
└─ README.md
```

**Decisões técnicas**
- **Erro como valor** (padrão Go): retornos múltiplos `<T, error>` em clients/services.
- **Dependências externas** isoladas em `internal/clients/*`.
- **Camada de serviço** controla o **fallback** e normaliza entrada (`BRL`, `USD`, ...).

---

## 🔧 Comandos úteis

```bash
# baixar deps e organizar
go mod tidy

# build binário
go build -o bin/api ./cmd/api

# executar binário compilado
./bin/api
```

---

## 📄 Licença
MIT – sinta-se à vontade para reutilizar e evoluir.
