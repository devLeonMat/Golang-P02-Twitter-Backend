package routers

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/*UploadBanner upload banner to server */
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var file2 string = "uploads/banners/" + userID + "." + extension

	f, err := os.OpenFile(file2, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying image ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = userID + "." + extension
	status, err = bd.UpdateUser(user, userID)
	if err != nil || status == false {
		http.Error(w, "Error saving the banner in the database ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
