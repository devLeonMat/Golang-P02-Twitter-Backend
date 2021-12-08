package models

/*Relationship to save the relationship between two users */
type Relationship struct {
	UserID             string `bson:"userID" json:"userID"`
	UserRelationshipID string `bson:"userRelationshipID" json:"userRelationshipID"`
}
