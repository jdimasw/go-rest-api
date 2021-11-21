package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/service"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	client := app.NewDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, client)
	productController := controller.NewProductController(productService)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/api/products", productController.Create)
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)

	log.Fatal(http.ListenAndServe(":8080", router))

	// server := http.Server{
	// 	Addr: "localhost:8080",
	// 	Handler: router
	// }
}
