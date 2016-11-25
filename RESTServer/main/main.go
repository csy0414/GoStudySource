package main

import (
	"log"
	"net/http"

	"github.com/csy0414/GoStudySource/RESTServer"
	"github.com/gorilla/mux"
)

func main() {
	RESTServer.NewMemoryDataAccess()
	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", RESTServer.GetHandler).Methods("GET")
	router.HandleFunc("/account/{id}", RESTServer.PutHandler).Methods("PUT")
	router.HandleFunc("/account/{id}", RESTServer.PostHandler).Methods("POST")
	router.HandleFunc("/account/{id}", RESTServer.DeleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8904", router))

}
