package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()                     //pegou uma estancia do GIN
	r.GET("/:nome", controllers.Saudacao)           // registrou uma nova rota
	req, _ := http.NewRequest("GET", "/Italo", nil) //passou que a requisicao vai ser get com parametro e nenhum corpo para ela
	resposta := httptest.NewRecorder()              //armazenando a resposta todo status
	r.ServeHTTP(resposta, req)                      //realiza de fato a requisicao
	// if resposta.Code != http.StatusOK {           //verifica se os status sao iguais
	// 	t.Fatalf("Status error: valor recebido foi %d, e o esperado era %d", resposta.Code, http.StatusOK)
	// }
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") //verifica se os status sao iguais COM O ASSERT DO P TESTIFY
	mockDaResposta := `{"API diz:":"E ai Italo, tudo beleza?"}`
	respostaDaBody, _ := ioutil.ReadAll(resposta.Body)      //LER TUDO QUE TEM NO RESPOSTA BODY
	assert.Equal(t, mockDaResposta, string(respostaDaBody)) //FAZER COMPARACAO DE RESPOSTAS
}

func TestListaTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder() //armazenando a resposta todo status
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoProCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678912", nil)
	resposta := httptest.NewRecorder() //armazenando a resposta todo status
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder() //armazenando a resposta todo status
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock) //PEGOU O JS E ARMAZENOU DENTRO DA VARIAVEL CONVERTENDO ELE PARA BYTES E DEPOIS JSON
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	resposta := httptest.NewRecorder() //armazenando a resposta todo status
	r.ServeHTTP(resposta, req)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "98765432198", RG: "473456789"}
	valorJson, _ := json.Marshal(aluno) //TEM Q CONVERTER ALUNO PARA JSON
	pathDaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathDaEditar, bytes.NewBuffer(valorJson)) //TEM Q CONVERTER DE JSON PARA BYTES PARA SER PASSADO
	resposta := httptest.NewRecorder()                                           //armazenando a resposta todo status
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, aluno.CPF, alunoMockAtualizado.CPF)
	assert.Equal(t, aluno.RG, alunoMockAtualizado.RG)
	assert.Equal(t, aluno.Nome, alunoMockAtualizado.Nome)
}
