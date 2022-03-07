package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos())
	r.Run()
}

func ExibeTodosAlunos() {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Rodrigo Ferreira",
	})
}
