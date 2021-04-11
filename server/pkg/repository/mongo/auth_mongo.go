package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userTableName = "users"
	databaseName  = "test"
)

type AuthMongo struct {
	client *mongo.Client
}

func NewAuthMongo(mc *MongoClient) *AuthMongo {
	return &AuthMongo{client: mc.client}
}

func (r *AuthMongo) CreateUser(user models.User) (uuid.UUID, error) {
	// var id uuid.UUID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := r.client.Database(databaseName).Collection(userTableName)
	res, err := collection.InsertOne(ctx, bson.D{ //? key could be reflected like config from Struct
		{"id", user.Id}, //TODO!
		{"name", user.Name},
		{"username", user.Username},
		{"password", user.Password},
	})
	if err != nil {
		return uuid.Nil, err
	}
	id := res.InsertedID
	uid, ok := id.(uuid.UUID)
	if !ok {
		logrus.Error("NE OK")
		//TODO how?
	}
	return uid, nil
}

func (r *AuthMongo) GetUser(username, passwordHash string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{"username", username},
		{"password", passwordHash},
	}
	collection := r.client.Database(databaseName).Collection(userTableName)
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		logrus.Error("record does not exist")
	} else {
		logrus.Error(err)
	}

	return user, err
}
