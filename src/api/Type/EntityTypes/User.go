package EntityTypes

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" validate:"required"`
	Age   int                `json:"age,omitempty" validate:"required"`
	Document
}
