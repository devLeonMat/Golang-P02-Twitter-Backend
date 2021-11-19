package middlew

import (
	"net/http"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
)

func  checkingDB(next http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		if bd.CheckConnection()==0 {
			http.Error(writer, "Lost connection", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}

}