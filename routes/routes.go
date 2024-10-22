package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest(){
	r := mux.NewRouter()
	r.Use(middleware.ConfigSetJson)
	r.HandleFunc("/",controllers.Home).Methods("Get")
	r.HandleFunc("/api",controllers.AllPeson).Methods("Get")
	r.HandleFunc("/api/new",controllers.NewPerson).Methods("Post")
	r.HandleFunc("/api/{id}",controllers.AllPesonID).Methods("Get")
	r.HandleFunc("/api/delete/{id}",controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/api/edit/{id}",controllers.EditPerson).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000",handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}