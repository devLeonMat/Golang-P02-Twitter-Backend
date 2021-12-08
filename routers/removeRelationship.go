package routers

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"net/http"
)

/*RemoveRelationship makes the deletion of the relationship */
func RemoveRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relationship
	t.UserID = userID
	t.UserRelationshipID = ID

	status, err := bd.DeleteRelationship(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete relationship "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "The relationship has not been"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
