package middlew

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/routers"
	"net/http"
)

/*ValidJWT valid the JWT in the response  */
func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.TokenProcess(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
