package main

import (
	"fmt"

	"github.com/rommelbonfim/E-Commerce-GO/database"
	"github.com/rommelbonfim/E-Commerce-GO/models"
	"github.com/rommelbonfim/E-Commerce-GO/routes"
)

func main() {
	models.Products = []models.Product{
		{
			Id:         1,
			Nome:       "exemplo",
			Descricao:  "Descrição do Produto 1",
			Preco:      100.00,
			Quantidade: 10,
			Categoria:  "Eletrônicos",
			ImagemURL:  "http://example.com/produto1.jpg",
			Disponivel: true,
		},
		{
			Id:         2,
			Nome:       "exemplo 2",
			Descricao:  "Descrição do Produto 2",
			Preco:      250.50,
			Quantidade: 5,
			Categoria:  "Vestuário",
			ImagemURL:  "http://example.com/produto2.jpg",
			Disponivel: false,
		},
	}
	database.DbConnection()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
