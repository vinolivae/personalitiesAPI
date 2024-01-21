package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/vinolivae/personalitiesAPI/controllers"
)

func LoadRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/", c.Home)
	r.HandleFunc("/api/personalities", c.ListPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", c.DetailPersonality).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
