package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	//Os handlers terão os seus headers configurados com base nesse middleware
	//A implementação está no router

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application-json")
		next.ServeHTTP(w, r)
	})
}
