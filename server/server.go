package server

import (
	"github.com/seantcanavan/tidder/user"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/user",      user.GetUserByEmailAddressHandler).Methods("GET")
	router.HandleFunc("/user/{id}", user.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/user/{id}", user.AddUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", user.DeleteUserHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
