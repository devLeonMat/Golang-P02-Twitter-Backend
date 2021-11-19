package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/** user model   */
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	name      string             `bson:"name" json:"name,omitempty"`
	lastName  string             `bson:"lastName" json:"lastName,omitempty"`
	bornDate  time.Time          `bson:"bornDate" json:"bornDate,omitempty"`
	email     string             `bson:"email" json:"email"`
	password  string             `bson:"password" json:"password,omitempty"`
	avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	banner    string             `bson:"banner" json:"banner,omitempty"`
	biography string             `bson:"biography" json:"biography,omitempty"`
	location  string             `bson:"location" json:"location,omitempty"`
	webSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
