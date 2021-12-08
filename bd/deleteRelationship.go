package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"time"
)

/*DeleteRelationship delete the relationship in the data base */
func DeleteRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("Relationship")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
