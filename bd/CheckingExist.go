package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)
/*CheckIfUserExist valid if user exist */
func CheckIfUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("User")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
