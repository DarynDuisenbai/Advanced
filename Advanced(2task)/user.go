package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/options"
)

type User struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
	Age   int    `bson:"age"`
}

func createUser(client *mongo.Client, user User) error {
	collection := client.Database("chatdb").Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	user.ID = result.InsertedID.(string)
	return nil
}

func getUserByID(client *mongo.Client, userID string) (User, error) {
	collection := client.Database("chatdb").Collection("users")
	var user User
	err := collection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	return user, err
}

func updateUser(client *mongo.Client, userID, newName string) (User, error) {
	collection := client.Database("chatdb").Collection("users")
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newName}}}}
	var updatedUser User
	err := collection.FindOneAndUpdate(context.Background(), bson.M{"_id": userID}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedUser)
	return updatedUser, err
}

func getAllUsers(client *mongo.Client) ([]User, error) {
	collection := client.Database("chatdb").Collection("users")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []User
	for cur.Next(context.Background()) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func deleteUser(client *mongo.Client, userID string) error {
	collection := client.Database("chatdb").Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": userID})
	return err
}
