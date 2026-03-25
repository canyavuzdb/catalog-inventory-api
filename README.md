# Catalog Inventory API

This project is a backend service for managing product catalog and inventory data.

It is built with Go and exposes a GraphQL API. The goal is to practice and demonstrate how to design a clean, scalable backend system with a layered architecture and real database integration.

---

## Why this project exists

I wanted to build something closer to a real backend service instead of a small demo application.

This project focuses on:

- designing a clean project structure
- building a GraphQL API with Go
- working with PostgreSQL in a realistic way
- separating responsibilities across resolver, service, and repository layers
- preparing a codebase that can grow into a larger service over time

---

## Tech stack

- Go (Golang)
- GraphQL (gqlgen)
- PostgreSQL
- GORM
- Docker
- GitHub Actions

Planned additions:

- filtering and sorting
- pagination improvements
- category support
- validation and better error handling
- migration files
- tests
- Redis caching
- JWT authentication
- OpenTelemetry

---

## Current features

- list products
- create products
- GraphQL Playground
- PostgreSQL integration
- layered architecture
- basic CI workflow

This is still the first working version, and the current focus is to evolve it into a more production-like service step by step.

---

## Project structure

The project is organized to keep responsibilities separate:

- `cmd/api`: application entry point
- `internal/domain`: core entities
- `internal/service`: business logic
- `internal/repository`: database access
- `internal/graph`: GraphQL schema and resolvers
- `internal/database`: database connection and setup

---

## How it works

```mermaid
flowchart TD
    A[Client] --> B[GraphQL API]
    B --> C[Resolver Layer]
    C --> D[Service Layer]
    D --> E[Repository Layer]
    E --> F[(PostgreSQL)]