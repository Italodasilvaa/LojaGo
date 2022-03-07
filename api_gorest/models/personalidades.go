package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
} //QUANDO FOR FAZER O RETORNO DA STRUCT EM JSON USAR ESSA SINTAXE

var Personalidades []Personalidade
