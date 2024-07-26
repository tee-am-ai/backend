package helper

import (
	"context"
	"fmt"

	"github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBInfo struct {
	DBString string
	DBName   string
}

func MongoConnect(mconn DBInfo) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		return nil, err
	}
	return client.Database(mconn.DBName), nil
}

func InsertOneDoc(db *mongo.Database, col string, doc any) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("users")
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("email tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}
