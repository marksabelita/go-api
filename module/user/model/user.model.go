package user_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Name     string             `json:"name,omitempty" validate:"required"`
}