package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers" //
	"github.com/gorilla/mux"
	"github.com/italodasilvaa/go-rest-api/controllers"
	"github.com/italodasilvaa/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware) //UTILIZADO PARA PADRONIZAR O RETORNO DO JSON ALINHADO
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornarPersonalidade).Methods("Get") //SEMPRE QUE TIVER UM PARAMETRO A SER PASSADO USAR{} E A INFOR DENTRO
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) //HANDLERS SERVE PARA LIBERAR O ACESSO A API DE PORTAS DIFERENTES
}
