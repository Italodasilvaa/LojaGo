package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/italodasilvaa/go-rest-api/database"
	"github.com/italodasilvaa/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	var p []models.Personalidade
	database.DB.Find(&p) //TODAS PERSONALIDADES PASSANDO O ENDEREÇO DE MEMORARIA DA ESTRUTURA
	json.NewEncoder(w).Encode(p)
}

func RetornarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //AO INSTALAR O GORILLA/MUX TEM A OPCAO DO MUX ONDE CONSIGO PERGAR OS VALORES DA REQUEST ATRAVEZ DE VARS
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id) //PARA UTILIZAR APENAS UMA PASSA O ENDEREÇO DA PERSONALIDADE E ID
	json.NewEncoder(w).Encode(personalidade)

	// for _, personalidade := range models.Personalidades {
	// 	if strconv.Itoa(personalidade.Id) == id { //STRCON PARA CONVERTER O FORMATO
	// 		json.NewEncoder(w).Encode(personalidade) //PARA EXIBIR O RESULTADO USAR O ENCODER PASSANDO A RESPONSEWRITE E DEPOIS ENCODE COM A INFORMACAO RETORNADA
	// 	}
	// }
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade) //QUANDO O DADO E JSON E PRECISA CONVERTER PARA O GO DECODIFICANDO
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade) //DADO JSON E QUER EXIBIR FAZER UM NEWENCODER

}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //AO INSTALAR O GORILLA/MUX TEM A OPCAO DO MUX ONDE CONSIGO PERGAR OS VALORES DA REQUEST ATRAVEZ DE VARS
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //AO INSTALAR O GORILLA/MUX TEM A OPCAO DO MUX ONDE CONSIGO PERGAR OS VALORES DA REQUEST ATRAVEZ DE VARS
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade) //QUANDO O DADO E JSON E PRECISA CONVERTER PARA O GO DECODIFICANDO
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade) //DADO JSON E QUER EXIBIR FAZER UM NEWENCODER

}
