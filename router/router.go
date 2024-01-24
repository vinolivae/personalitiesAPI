package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/vinolivae/personalitiesAPI/controllers"
	m "github.com/vinolivae/personalitiesAPI/middleware"
)

func LoadRoutes() {
	r := mux.NewRouter()
	r.Use(m.ContentTypeMiddleware)

	r.HandleFunc("/", c.Home)
	r.HandleFunc("/api/personalities", c.ListPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", c.DetailPersonality).Methods("GET")
	r.HandleFunc("/api/personalities", c.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", c.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", c.UpdatePersonality).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
