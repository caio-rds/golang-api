package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username" binding:"required,min=3,max=30"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	BirthDate string             `json:"birthdate" bson:"birthdate" binding:"required"`
}
