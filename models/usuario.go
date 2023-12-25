package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Nome string `bson:"nome,omitempty"`
	Senha string `bson:"senha,omitempty"`
}