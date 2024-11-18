package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quiz struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ModuleID  primitive.ObjectID `bson:"module_id" json:"moduleId"`
	Questions []string         `bson:"questions" json:"questions"`
	MaxScore  int                `bson:"max_score" json:"maxScore"`
}


