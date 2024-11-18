package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Assignment struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ModuleID     primitive.ObjectID `bson:"module_id" json:"moduleId"`
    Description  string              `bson:"description" json:"description"`
    DueDate      time.Time           `bson:"due_date" json:"dueDate"`
}
