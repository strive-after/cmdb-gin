package mongo

import (
	"context"
	"gin-moudle/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Database
	ctx         = context.TODO()
)

type mgo struct{}

type CURD interface {
	InsertOne(table string, data interface{}) error
	InsertMany(table string, data []interface{}) error
	FindOne(table string, filter, result interface{}) error
	FindMany(table string, limit, index int64, filter, result interface{}) error
	UpdateMany(table string, filter, datas interface{}) error
	UpdateById(table string, id, data interface{}) error
	UpdateOne(table string, filter, data interface{}) error
	DeleteOne(table string, filter interface{}) error
	DeleteMany(table string, filter interface{}) error
}

func Newm() CURD {
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

	_, err := mongoClient.Collection(table).InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return err
}

func (m *mgo) InsertMany(table string, data []interface{}) error {
	_, err := mongoClient.Collection(table).InsertMany(ctx, data)
	if err != nil {
		return err
	}

	return err
}

func (m *mgo) FindOne(table string, filter, result interface{}) error {

	err := mongoClient.Collection(table).FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return err
}

func (m *mgo) FindMany(table string, limit, index int64, filter, result interface{}) error {
	//每次查询多少个
	findOptions := options.Find()
	if limit != 0 {
		findOptions.SetLimit(limit)
		findOptions.SetSkip(limit * index)
	}

	cur, err := mongoClient.Collection(table).Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, result)
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo) UpdateOne(table string, filter, data interface{}) error {

	_, err := mongoClient.Collection(table).UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo) UpdateMany(table string, filter, datas interface{}) error {

	_, err := mongoClient.Collection(table).UpdateMany(ctx, filter, datas)
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo) UpdateById(table string, id, data interface{}) error {

	_, err := mongoClient.Collection(table).UpdateByID(ctx, id, data)
	if err != nil {
		return err
	}
	return nil

}

func (m *mgo) DeleteOne(table string, filter interface{}) error {

	_, err := mongoClient.Collection(table).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *mgo) DeleteMany(table string, filter interface{}) error {

	_, err := mongoClient.Collection(table).DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil

}

type Tes struct {
	Name string
	Age  int
	Addr string
}
