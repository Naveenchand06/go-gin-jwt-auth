package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Naveenchand06/go-gin-jwt-auth/constants"
	"github.com/Naveenchand06/go-gin-jwt-auth/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID                primitive.ObjectID  `json:"id" bson:"id"`
	Name              string              `json:"name"`
	Email 			  string    		  `json:"email" validate:"email,required"`
	Phone 			  string    		  `json:"phone" validate:"required"`
	Password          string    		  `json:"password,omitempty" validate:"required,min=6"` 
	// FirstName		  string    		  `json:"first_name" validate:"required, min=2, max=100"`
	// LastName		      string    		  `json:"last_name" validate:"required, min=2, max=100"`
	// UserID             string    		  `json:"user_id"`
	// UserType           string 			  `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	// Token 			  string 			  `json:"token"`
	// RefreshToken 	  string  			  `json:"refresh_token"`
	CreatedAt 		  time.Time 		  `json:"created_at,omitempty"`
	UpdatedAt 		  time.Time 		  `json:"updated_at,omitempty"`
}

func (u *User) IsUserExists(db *mongo.Client) (*User, bool) {
	users := database.OpenCollection(db, constants.USERCOLL)
	var user User
	err := users.FindOne(context.Background(), bson.M{"email": u.Email}).Decode(&user)
	fmt.Println("Got user si --> ", user)
	fmt.Println("Got user err is --> ", err)

	// * If exitsts then user already exists
	return &user, !(err == mongo.ErrNoDocuments)
}

func (u *User) RegisterUser(db *mongo.Client) (*User, error) {
	// * Users Collection
	users := database.OpenCollection(db, constants.USERCOLL)
	// * Checking whether the user exits or not
	var user User
	err := users.FindOne(context.Background(), bson.M{"email": u.Email}).Decode(&user)
	fmt.Println("Got user si --> ", user)
	fmt.Println("Got user err is --> ", err)

	// * If exitsts then user already exists
	if  _, ok := u.IsUserExists(db); ok {
		return nil, errors.New("user already exists")
	}
	// * If new user, Store it in DB and return the created User
	res, err := users.InsertOne(context.Background(), u)
	if err != nil {
		return nil, err
	}
	u.ID = res.InsertedID.(primitive.ObjectID)
	u.Password = ""
	return u, nil
}


func (u *User) LoginUser(db *mongo.Client) (*User, error) {
	if user, ok := u.IsUserExists(db); ok {
		return user, nil
	} else {
		return nil, errors.New("user does not exists")
	}
}