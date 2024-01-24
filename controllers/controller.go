package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/vinolivae/personalitiesAPI/database"
	m "github.com/vinolivae/personalitiesAPI/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []m.Personality
	db.DB.Find(&personalities)

	// Encoda uma resposta em json, e a resposta a ser encodada é "personalities".
	json.NewEncoder(w).Encode(personalities)
}

func DetailPersonality(w http.ResponseWriter, r *http.Request) {
	var personality []m.Personality

	//Retorna as variáveis da rota, quero o id.
	vars := mux.Vars(r)
	id := vars["id"]

	db.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality m.Personality

	//vamos decodar o body que recebemos via interface pro gorm enviar pro banco
	json.NewDecoder(r.Body).Decode(&personality)
	db.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	var personality m.Personality

	vars := mux.Vars(r)
	id := vars["id"]

	db.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality m.Personality

	vars := mux.Vars(r)
	id := vars["id"]

	db.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	db.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
