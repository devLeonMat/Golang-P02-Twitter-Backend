package routers

import (
	"encoding/json"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"net/http"
	"time"
)

/*SaveTweet allows to save the tweet*/
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	registry := models.SaveTweet{
		UserID:  userID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.CreateTweet(registry)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the tweet, try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "the tweet insert has not been possible", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
