package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/canyavuzdb/catalog-inventory-api/internal/database"
	"github.com/canyavuzdb/catalog-inventory-api/internal/graph"
	"github.com/canyavuzdb/catalog-inventory-api/internal/repository"
	"github.com/canyavuzdb/catalog-inventory-api/internal/service"
)

func main() {

	runSeed := false
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		runSeed = true
	}

	database.Connect(runSeed)

	productRepo := repository.NewProductRepository(database.DB)
	categoryRepo := repository.NewCategoryRepository(database.DB)

	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	resolver := &graph.Resolver{
		ProductService:  productService,
		CategoryService: categoryService,
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
