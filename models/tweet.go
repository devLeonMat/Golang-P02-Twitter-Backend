package models

/*Tweet catch de body */
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
