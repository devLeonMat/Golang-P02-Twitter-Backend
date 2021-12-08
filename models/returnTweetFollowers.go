package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*ReturnTweetsFollowers is the structure of the tweets from de followers */
type ReturnTweetsFollowers struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID             string             `bson:"userID" json:"userId,omitempty"`
	UserRelationshipID string             `bson:"userRelationshipID" json:"userRelationshipID,omitempty"`
	Tweet              struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
