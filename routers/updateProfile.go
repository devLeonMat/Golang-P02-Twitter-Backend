package routers

import (
	"encoding/json"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"net/http"
)

/*UpdateProfile update the profile user */
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.UpdateUser(t, userID)
	if err != nil {
		http.Error(w, "An error occurred while trying to update the user profile "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Has not been update the user profile ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
