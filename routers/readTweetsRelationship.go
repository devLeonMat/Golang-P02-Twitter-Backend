package routers

import (
	"encoding/json"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"net/http"
	"strconv"
)

/*ReadTweetsFollowers read the tweets of the followers */
func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "you must send the page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter with a value greater than 0", http.StatusBadRequest)
		return
	}

	response, correct := bd.ReadTweetFollowers(userID, page)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
