package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*ReturnTweets is the structure of the tweets from de followers */
type ReturnTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userID" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
