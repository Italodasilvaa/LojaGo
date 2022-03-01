package models

import (
	"controleEstoque/db"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectarDB()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos
}
func ProdutoPorId(id string) Produto {
	db := db.ConectarDB()
	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

	}

	defer db.Close()
	return produto
}

func AtualizaProduto(produto Produto) {
	db := db.ConectarDB()
	atualizarDadosNoBanco, err := db.Prepare("update  produtos set nome=$1,descricao=$2,preco=$3,quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizarDadosNoBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)
	defer db.Close()
}
func CriarNovoProduto(produto Produto) {

	db := db.ConectarDB()

	insertDadosNoBanco, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertDadosNoBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	defer db.Close()
}

func DeleteProduto(id string) {
	db := db.ConectarDB()
	deleteProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduto.Exec(id)
	defer db.Close()
}
