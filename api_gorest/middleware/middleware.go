package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json") //TODA VEZ QUE FAZER A CHAMADA ELE VAI FORMATAR Q E DO TIPO JSON A RESPOSTA
		next.ServeHTTP(w, r)
	})
}
