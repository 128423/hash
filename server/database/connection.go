package database

import (
	"context"
	"time"

	"github.com/128423/hash/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func GetUser(UserId int64) (*models.User, error) {
	var result *models.User
	collection := Client.Database("teste").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": UserId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetProduct(ProductId int64) (*models.Product, error) {
	var result *models.Product
	collection := Client.Database("teste").Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": ProductId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
