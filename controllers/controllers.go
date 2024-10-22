package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	
	"net/http"
	
	"github.com/gorilla/mux"
)



func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Home")
}

func AllPeson(w http.ResponseWriter,r *http.Request){
	var list []models.Personalidades
	database.DB.Find(&list)
	json.NewEncoder(w).Encode(list)
}

func AllPesonID(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	id:= vars["id"]
    var person models.Personalidades
	database.DB.First(&person,id)
	json.NewEncoder(w).Encode(person)
}
func NewPerson(w http.ResponseWriter, r *http.Request){
	var newPerson models.Personalidades
	json.NewDecoder(r.Body).Decode(&newPerson)
	database.DB.Create(&newPerson)
	json.NewEncoder(w).Encode(newPerson)
	

}
func DeletePerson(w http.ResponseWriter, r *http.Request){
     var deletedPerson models.Personalidades
	 vars := mux.Vars(r)
	 id:= vars["id"]
	 database.DB.Delete(&deletedPerson,id)
	 json.NewEncoder(w).Encode(deletedPerson)

}
func EditPerson(w http.ResponseWriter, r *http.Request){
	 var person models.Personalidades
	 vars := mux.Vars(r)
	 id:= vars["id"]
	 database.DB.First(&person,id)
	 json.NewDecoder(r.Body).Decode(&person)
	 database.DB.Save(&person)
	 json.NewEncoder(w).Encode(&person)
}