package bd

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	us, found, _ := CheckIfUserExist(email)
	if found == false {
		return us, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(us.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return us, false
	}
	return us, true
}
