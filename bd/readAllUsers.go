package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/*ReadAllUsers search the profile in the database*/
func ReadAllUsers(ID string, page int64, search string, typ string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("User")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Relationship
		r.UserID = ID
		r.UserRelationshipID = s.ID.Hex()

		include = false

		found, err = SearchRelationship(r)
		if typ == "new" && found == false {
			include = true
		}
		if typ == "follow" && found == true {
			include = true
		}

		if r.UserRelationshipID == ID {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
