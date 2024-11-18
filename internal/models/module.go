package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Module struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	CourseID primitive.ObjectID   `bson:"course_id" json:"courseId"`
	Title    string               `bson:"title" json:"title"`
	Content  string               `bson:"content" json:"content"`
	QuizIDs  []primitive.ObjectID `bson:"quizzes" json:"quizzes"`
}
