package user

import (
	"context"
	"go-api/config"
	user_model "go-api/module/user/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func FindService(query bson.M) ([]user_model.User, error) {
	var users []user_model.User
	var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.Find(ctx, query)

	if err != nil { return nil, err }

	defer result.Close(ctx)
	for result.Next(ctx) {
			var singleUser user_model.User

			if err = result.Decode(&singleUser); err != nil {
				return nil, err;
			}

			users = append(users, singleUser)
	}

	return users, nil
}

func FindOneService(query bson.M) (user_model.User , error) {
	var user user_model.User
	var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, query).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
} 