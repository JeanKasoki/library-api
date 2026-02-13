# Library API 📚

[🇺🇸 English](#-project) | [🇧🇷 Português](#-projeto)

---

## 🇺🇸 Project
API for library management developed in Go (Golang), focusing on Clean Architecture, Scalability, and High Performance.

### 🚀 Technologies
- **Language:** Go 1.22+
- **Database:** MySQL 9.0
- **Containerization:** Docker & Docker Compose
- **ORM:** GORM (Object Relational Mapper)
- **Architecture:** Clean Architecture + Standard Go Project Layout

### 🏗️ Architecture & Features (Current Progress)
- **Domain Layer:** Entities defined (`Book`, `User`, `Loan`) with JSON mapping.
- **Infrastructure Layer:** - Dockerized MySQL 9.0 instance.
  - Database connection using GORM driver.
  - **Auto Migrations:** Tables are automatically created/updated based on Go structs.
  - **Repository Pattern:** Isolated database operations (`BookRepository`).
**Web Handler:** HTTP Handlers managing Requests/Responses (`BookHandler`).
- **Application Layer (Use Cases):** - `CreateBookUseCase`: Business logic for creating books, decoupled using DTOs.
  - `ListBooksUseCase`: Logic for retrieving all books and mapping Entities to Output DTOs.
  - `GetBookUseCase`: Logic for retrieving a single book by ID.
  - `UpdateBookUseCase`: Logic for updating book details by ID.

### 📂 Project Structure
The project follows the **Standard Go Project Layout**:
```bash
.
├── cmd/api/main.go          # Application Entrypoint (Dependency Injection)
├── internal/
│   ├── entity/              # Domain Entities (Core Business Objects)
│   ├── usecase/             # Business Logic (Managers)
│   └── infra/
│       ├── database/        # DB Connection
│       └── repository/      # Data Access Layer (Implements GORM)
│       └── web/
│           └── handler/     # HTTP Handlers (Controllers)
├── docker-compose.yml       # Infrastructure as Code
└── go.mod                   # Dependency Manager
```

### ⚡ Architecture Flow
`Main (Injection) -> Handler (HTTP) -> UseCase (Logic) -> Repository (Interface) -> Database (MySQL)`

### 🛠️ How to Run (Development)
**1. Start the Database:**
```bash
docker compose up -d
```

**2. Run the Application:**
```bash
go run cmd/api/main.go
```

**3. Test the Endpoint (POST):**
```bash
curl -X POST http://localhost:8080/books -d '{"titulo": "The Go Programming Language", "autor": "Alan A. A. Donovan", "isbn": "978-0134190440", "ano_publicacao": 2015}'
```
**4. Test: List Books (GET): Open in browser or run:**
```bash
curl http://localhost:8080/books
```

**5. Test: Find Book By ID (GET): Open in browser or run:**
```bash
curl "http://localhost:8080/book?id=1"
```

**6. Test: Update Book (PUT):**
```bash
curl -X PUT "http://localhost:8080/book?id=1" -d '{"titulo": "The Go Programming Language (Updated)", "autor": "Alan Donovan", "isbn": "978-0134190440", "ano_publicacao": 2024}'
```

### 🚧 Status
In development.

✅ Entities & Database Modeling
✅ Infrastructure (Docker + GORM Connection)
✅ Repository Pattern Implemented
✅ Feature: Create Book (POST /books) - Done
✅ Feature: List Books (GET /books) - Done
✅ Feature: Find Book By ID (GET /book?id=x) - Done 
✅ Feature: Update Book (PUT /book?id=x) - Done
⏳ Feature: Delete Book (DELETE /book?id=x) - Next Step

---

## 🇧🇷 Projeto
API para gerenciamento de biblioteca desenvolvida em Go (Golang), focando em Clean Architecture, Escalabilidade e Alta Performance.

### 🚀 Tecnologias
- **Linguagem:** Go 1.22+
- **Banco de Dados:** MySQL 9.0
- **Containerização:** Docker & Docker Compose
- **ORM:** GORM (Object Relational Mapper)
- **Arquitetura:** Clean Architecture + Standard Go Project Layout

### 🏗️ Arquitetura & Funcionalidades (Progresso Atual)
- **Camada de Domínio:** Entidades definidas (`Book`, `User`, `Loan`) com mapeamento JSON.
- **Camada de Infraestrutura:** - Instância MySQL 9.0 rodando em Docker.
  - Conexão com banco de dados usando driver GORM.
  - **Auto Migrations:** Tabelas são criadas/atualizadas automaticamente baseadas nas structs do Go.
  - **Repository Pattern:** Operações de banco de dados isoladas (`BookRepository`).
  - **Web Handler:** Handlers HTTP gerenciando Requisições/Respostas (`BookHandler`).
- **Camada de Aplicação (Use Cases):** - `CreateBookUseCase`: Lógica de negócio para criação de livros, desacoplada do banco e da camada HTTP usando DTOs (Data Transfer Objects).
- `ListBooksUseCase`: Lógica para recuperar todos os livros e mapear Entidades para DTOs de saída.
- `GetBookUseCase`: Lógica para recuperar um único livro por ID.
- `UpdateBookUseCase`: Lógica para atualizar os dados de um livro pelo ID.

### 📂 Estrutura do Projeto
O projeto segue o **Standard Go Project Layout**:

```bash
.
├── cmd/api/main.go          # Ponto de Entrada (Injeção de Dependência)
├── internal/
│   ├── entity/              # Entidades de Domínio (O coração do projeto)
│   ├── usecase/             # Regras de Negócio (Gerentes)
│   └── infra/
│       ├── database/        # Conexão com Banco
│       └── repository/      # Acesso a Dados (Implementação GORM)
│       └── web/
│           └── handler/     # Handlers HTTP (Controladores)
├── docker-compose.yml       # Infraestrutura como Código
└── go.mod                   # Gerenciador de Dependências
```

### ⚡ Fluxo da Arquitetura
`Main (Injeção) -> Handler (HTTP) -> UseCase (Lógica) -> Repository (Interface) -> Database (MySQL)`

### 🛠️ Como Rodar (Desenvolvimento)
**1. Subir o Banco de Dados:**
```bash
docker compose up -d
```

**2. Rodar a Aplicação:**
```bash
go run cmd/api/main.go
```

**3. Testar a Rota (POST):**
```bash
curl -X POST http://localhost:8080/books -d '{"titulo": "The Go Programming Language", "autor": "Alan A. A. Donovan", "isbn": "978-0134190440", "ano_publicacao": 2015}'
```

**4. Testar: Listar Livros (GET): Abra no navegador ou rode:**
```bash
curl http://localhost:8080/books
```

**5. Testar: Buscar Livro por ID (GET): Abra no navegador ou rode:**
```bash
curl "http://localhost:8080/book?id=1"
```

**6. Testar: Atualizar Livro (PUT):**
```bash
curl -X PUT "http://localhost:8080/book?id=1" -d '{"titulo": "The Go Programming Language (Atualizado)", "autor": "Alan Donovan", "isbn": "978-0134190440", "ano_publicacao": 2024}'
```

### 🚧 Status
Em desenvolvimento.

✅ Entidades e Modelagem de Banco
✅ Infraestrutura (Docker + Conexão GORM)
✅ Padrão Repository Implementado
✅ Feature: Criar Livro (POST /books) - Feito
✅ Feature: Listar Livros (GET /books) - Feito
✅ Feature: Buscar Livro por ID (GET /book?id=x) - Feito 
✅ Feature: Atualizar Livro (PUT /book?id=x) - Feito
⏳ Feature: Deletar Livro (DELETE /book?id=x) - Próximo Passo