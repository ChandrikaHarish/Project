package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Place string             `json:"place,omitempty" bson:"place,omitempty"`
}
