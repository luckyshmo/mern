package mongo

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	client *mongo.Client
}

func NewUserMongo(uc *MongoClient) *UserMongo {
	return &UserMongo{client: uc.client}
}

func (r *UserMongo) GetAll() ([]models.User, error) {
	var userList []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := r.client.Database(databaseName).Collection(userTableName)
	cur, err := collection.Find(ctx, bson.D{})
	for cur.Next(ctx) {
		var result models.User
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		userList = append(userList, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return userList, err
}

func (r *UserMongo) GetById(userId uuid.UUID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{"id", user.Id}}
	collection := r.client.Database(databaseName).Collection(userTableName)
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		logrus.Error("record does not exist")
	} else {
		logrus.Error(err)
	}

	return user, err
}
