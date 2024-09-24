package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rommelbonfim/E-Commerce-GO/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/products", controllers.AllProducts).Methods("Get")
	r.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("Get")
	r.HandleFunc("/api/products", controllers.NewProduct).Methods("Post")
	r.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("Delete")
	r.HandleFunc("/api/products/{id}", controllers.PutProduct).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
