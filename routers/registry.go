package routers

import (
	"encoding/json"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"net/http"
)

/*Registry  is the function to create registry in the database */
func Registry(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the request"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "the minimum len in password is 6", 400)
		return
	}
	_, found, _ := bd.CheckIfUserExist(t.Email)
	if found == true {
		http.Error(w, "A user with this email already exist", 400)
		return
	}

	_, status, err := bd.CreateNewRegistry(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to save the record"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "failed to inset the record"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
