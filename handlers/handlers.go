package handlers

import (
	"github.com/gorilla/mux"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/middlew"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// Managers set port, handler and put to hear*/
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/registry", middlew.CheckingDB(routers.Registry)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckingDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
