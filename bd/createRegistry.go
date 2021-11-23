package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*CreateNewRegistry create new registry in the database */
func CreateNewRegistry(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("User")
	u.Password, _ = PasswordEncrypt(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
