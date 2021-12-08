package routers

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"net/http"
)

/*DeleteTweet allows you to delete a specific Tweet */
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, userID)
	if err != nil {
		http.Error(w, "An error occurred while try to delete the tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
