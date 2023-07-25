package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Students struct {
	RollNumber primitive.ObjectID `json:"_rollnumber,omitempty" bson:"_rollnumber,omitempty"`
	Name       string             `json:"name,omitempty"`
	Age        int                `json:"age,omitempty"`
	Stream     string             `json:"stream,omitempty"`
}
