package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/canyavuzdb/catalog-inventory-api/internal/database"
	"github.com/canyavuzdb/catalog-inventory-api/internal/graph"
	"github.com/canyavuzdb/catalog-inventory-api/internal/repository"
	"github.com/canyavuzdb/catalog-inventory-api/internal/service"
)

func main() {
	database.Connect()

	productRepo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(productRepo)

	resolver := &graph.Resolver{
		ProductService: productService,
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: resolver},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
