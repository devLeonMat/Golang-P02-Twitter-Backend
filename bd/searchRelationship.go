package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*SearchRelationship search the profile in the database*/
func SearchRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("Relationship")

	condition := bson.M{
		"userID":             t.UserID,
		"userRelationshipID": t.UserRelationshipID,
	}

	var result models.Relationship
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
