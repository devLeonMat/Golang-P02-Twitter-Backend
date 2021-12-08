package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*UpdateUser update the profile in the database*/
func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("User")

	registry := make(map[string]interface{})
	if len(u.Name) > 0 {
		registry["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		registry["lastname"] = u.LastName
	}
	registry["bornDate"] = u.BornDate
	if len(u.Avatar) > 0 {
		registry["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registry["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		registry["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		registry["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		registry["webSite"] = u.WebSite
	}

	updtString := bson.M{
		"$set": registry,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
