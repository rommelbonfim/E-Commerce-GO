package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rommelbonfim/E-Commerce-GO/database"
	"github.com/rommelbonfim/E-Commerce-GO/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	var p []models.Product
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.First(&product, id)
	json.NewEncoder(w).Encode(product)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	json.NewDecoder(r.Body).Decode(&newProduct)
	database.DB.Create(&newProduct)
	json.NewEncoder(w).Encode(newProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.Delete(&product, id)
	json.NewEncoder(w).Encode(product)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.First(&product, id)
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}
