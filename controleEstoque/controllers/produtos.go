package controllers

import (
	"controleEstoque/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidadeConvertido, err := strconv.Atoi(quantidade) //Converter para INT
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		produto := models.Produto{}

		produto.Nome = nome
		produto.Descricao = descricao
		produto.Quantidade = quantidadeConvertido
		produto.Preco = precoConvertido

		models.CriarNovoProduto(produto)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeleteProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.ProdutoPorId(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}
		precoConvert, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float:", err)
		}
		quantidadeConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade para int:", err)
		}

		produto := models.Produto{}
		produto.Id = idConvert
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = precoConvert
		produto.Quantidade = quantidadeConvert

		models.AtualizaProduto(produto)
	}
	http.Redirect(w, r, "/", 301)
}

// type Produtos struct {
// 	Nome, Descricao string
// 	Preco           float64
// 	Id, Quantidade  int
// }
