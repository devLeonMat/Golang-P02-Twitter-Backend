package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim Structure to process jwt  */
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"id" json:"_id,omitempty"`
	jwt.StandardClaims
}
