package routers

import (
	"encoding/json"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"net/http"
)

/*ShowProfile allows to extract the values of the  profile*/
func ShowProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying search the registry"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
