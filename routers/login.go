package routers

import (
	"encoding/json"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/jwt"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or password is invalid :"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "email user is required:"+err.Error(), 400)
		return
	}

	document, exist := bd.TryLogin(t.Email, t.Password)
	if !exist {
		http.Error(w, "User or password is invalid", 400)
		return
	}
	
	jwtKey, err := jwt.GenerateJwt(document)
	if err != nil {
		http.Error(w, "ocurred an error during the jwt generator:"+err.Error(), 400)
		return
	}

	resp:=models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	expirationTime := time.Now().Add(24*time.Hour)
	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
	
}
