package main

import (
	"log"
	"net/http"

	"com.exersice.level-5/crud"
	postget "com.exersice.level-5/postGet"
	"github.com/gorilla/mux"
)

func handerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get", postget.Get).Methods("GET")
	router.HandleFunc("/post", postget.Post).Methods("POST")

	router.HandleFunc("/getproducts", crud.GetAll).Methods("GET")
	router.HandleFunc("/setproduct", crud.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handerRequest()
}
