package middlew

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"net/http"
)
/*CheckingDB  midleware to valid database status*/
func CheckingDB(next http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		if bd.CheckConnection()==0 {
			http.Error(writer, "Lost connection", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}

}