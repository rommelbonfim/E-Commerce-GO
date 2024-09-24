package models

type Product struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
	Categoria  string  `json:"categoria"`
	ImagemURL  string  `json:"imagem_url"`
	Disponivel bool    `json:"disponivel"`
}

var Products []Product
