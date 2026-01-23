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
- **Application Layer (Use Cases):** - `CreateBookUseCase`: Business logic for creating books, decoupled from the database and HTTP layers using DTOs (Data Transfer Objects).

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
├── docker-compose.yml       # Infrastructure as Code
└── go.mod                   # Dependency Manager
```

### ⚡ Architecture Flow
`Main (Injection)` -> `UseCase (Logic)` -> `Repository (Interface)` -> `Database (MySQL)`

### 🛠️ How to Run (Development)
**1. Start the Database:**
```bash
docker compose up -d
```

**2. Run the Application:**
```bash
go run cmd/api/main.go
```

### 🚧 Status
In development.

✅ Entities & Database Modeling
✅ Infrastructure (Docker + GORM Connection)
✅ Repository Pattern Implemented
✅ Create Book UseCase (Business Logic)
⏳ API Handlers (Web Layer) - Next Step

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
- **Camada de Aplicação (Use Cases):** - `CreateBookUseCase`: Lógica de negócio para criação de livros, desacoplada do banco e da camada HTTP usando DTOs (Data Transfer Objects).

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
├── docker-compose.yml       # Infraestrutura como Código
└── go.mod                   # Gerenciador de Dependências
```

### ⚡ Fluxo da Arquitetura
`Main (Injeção) -> UseCase (Lógica) -> Repository (Interface) -> Database (MySQL)`

### 🛠️ Como Rodar (Desenvolvimento)
**1. Subir o Banco de Dados:**
```bash
docker compose up -d
```

**2. Rodar a Aplicação:**
```bash
go run cmd/api/main.go
```

### 🚧 Status
Em desenvolvimento.

✅ Entidades e Modelagem de Banco
✅ Infraestrutura (Docker + Conexão GORM)
✅ Padrão Repository Implementado
✅ UseCase de Criação de Livros (Lógica de Negócio)
⏳ API Handlers (Camada Web) - Próximo Passo