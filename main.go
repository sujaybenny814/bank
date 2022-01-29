package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro1/database"
	hd "github.com/micro1/handlers"
)

func main() {
	DB := database.Init()
	h := hd.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/createBank", h.AddBank).Methods("POST")
	r.HandleFunc("/getAllBank", h.GetAllBank).Methods("GET")
	r.HandleFunc("/updateBank", h.UpdateBank).Methods("PATCH")
	r.HandleFunc("/deleteBank/{id}", h.DeleteBank).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
