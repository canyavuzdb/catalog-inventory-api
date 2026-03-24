# Catalog Inventory API

This project is a backend service for managing product catalog and inventory data.

It is built with Go and exposes a GraphQL API. The main goal is to practice and demonstrate how to design a clean, scalable backend system using modern tools and patterns.

---

## Why this project exists

I wanted to build something that is closer to a real-world backend system instead of small demo apps.

This project focuses on:

- designing a proper project structure
- working with GraphQL in Go
- handling relational data with PostgreSQL
- thinking about performance and scalability early
- preparing a codebase that can grow into a microservice later

---

## Tech stack

- Go (Golang)
- GraphQL (gqlgen)
- PostgreSQL
- GORM
- Docker

Planned additions:

- Redis (for caching)
- JWT authentication
- OpenTelemetry (for tracing)

---

## What it does (current scope)

- list products
- get product details
- list categories
- basic inventory tracking
- filtering and pagination

This is the first version, so the focus is on getting a solid and clean foundation.

---

## Project structure

The project is organized to keep responsibilities separate:

- cmd/api: application entry point
- internal/domain: core entities
- internal/service: business logic
- internal/repository: database layer
- internal/graph: GraphQL schema and resolvers
- internal/database: database setup

---
How to run locally

Clone the repository:

```bash
git clone https://github.com/your-username/catalog-inventory-api.git
cd catalog-inventory-api
## How it works
```


```mermaid
flowchart TD
    A[Client] --> B[GraphQL API]
    B --> C[Resolver Layer]
    C --> D[Service Layer]
    D --> E[Repository Layer]
    E --> F[(PostgreSQL)]

    F --> E
    E --> D
    D --> C
    C --> B
    B --> A

    
