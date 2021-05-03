package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gin-moudle/internal/config"
)

var (
	mongoClient *mongo.Database
)

type mgo struct{}

func NewM() *mgo {
	return &mgo{}
}

func InitMongo(c *config.Mongo) {
	dst := "mongodb://" + c.Host + ":" + c.Port
	clientOptions := options.Client().ApplyURI(dst).SetMaxPoolSize(c.Maxpool)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	mongoClient = client.Database("test")

}

func (m *mgo) InsertOne(table string, data interface{}) error {

	_, err := mongoClient.Collection(table).InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return err
}
