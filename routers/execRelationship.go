package routers

import (
	"net/http"

	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
)

func ExecRelationship(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The ID parameter is mandatory", http.StatusBadRequest)
		return
	}

	var t models.Relationship
	t.UserID = userID
	t.UserRelationshipID = ID

	status, err := bd.CreateRelationship(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the relationship "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "It has not been possible to insert "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
