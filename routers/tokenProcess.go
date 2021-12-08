package routers

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"strings"
)

var Email string

/*userID is the returned ID of the model, which will be used in all EndPoints. */
var userID string

/*TokenProcess token process to extract its values */
func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("lmr-go-developer")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, found, _ := bd.CheckIfUserExist(claims.Email)
		if found == true {
			Email = claims.Email
			userID = claims.ID.Hex()
		}
		return claims, found, userID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalid")
	}
	return claims, false, string(""), err
}
