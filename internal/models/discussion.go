package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Discussion struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ModuleID primitive.ObjectID `bson:"module_id" json:"moduleId"`
	UserID   primitive.ObjectID `bson:"user_id" json:"userId"`
	Content  string             `bson:"content" json:"content"`
	Replies  []string           `bson:"replies" json:"replies"`
}
