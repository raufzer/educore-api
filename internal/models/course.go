package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	InstructorID primitive.ObjectID `bson:"instructor_id" json:"instructorId"`
	Modules      []Module           `bson:"modules" json:"modules"`
}
