package bd

import (
	"context"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/*ReadTweet delete a tweet */
func ReadTweet(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("Tweet")
	var results []*models.ReturnTweets
	condition := bson.M{
		"userId": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var registry models.ReturnTweets
		err := cursor.Decode(&registry)
		if err != nil {
			return results, false
		}
		results = append(results, &registry)
	}
	return results, true
}
