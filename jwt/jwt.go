package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"time"
)

func GenerateJwt(t models.User) (string, error) {
	myKey := []byte("lmr-go-developer")
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"bornDate":  t.BornDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
