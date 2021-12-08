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

	router.HandleFunc("/showProfile", middlew.CheckingDB(middlew.ValidJWT(routers.ShowProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckingDB(middlew.ValidJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/saveTweet", middlew.CheckingDB(middlew.ValidJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckingDB(middlew.ValidJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckingDB(middlew.ValidJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckingDB(middlew.ValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckingDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckingDB(middlew.ValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckingDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/execRelationship", middlew.CheckingDB(middlew.ValidJWT(routers.ExecRelationship))).Methods("POST")
	router.HandleFunc("/removeRelationship", middlew.CheckingDB(middlew.ValidJWT(routers.RemoveRelationship))).Methods("DELETE")
	router.HandleFunc("/searchRelationship", middlew.CheckingDB(middlew.ValidJWT(routers.SearchRelationship))).Methods("GET")

	router.HandleFunc("/listUsers", middlew.CheckingDB(middlew.ValidJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/readTweetsFollowers", middlew.CheckingDB(middlew.ValidJWT(routers.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
