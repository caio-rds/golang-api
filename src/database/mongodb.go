package database

import (
	"context"
	"errors"
	"github.com/caio-rds/golang-api/src/model/user/requests"
	"github.com/caio-rds/golang-api/src/model/user/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var database *mongo.Database
var ctx = context.TODO()

type Db struct {
	clientDatabase *mongo.Database
}

func NewDatabase() *Db {
	clientOptions := options.Client().ApplyURI("mongodb+srv://tartarinha:481Ckt1cXTcbWuU8@tartarinhamaruga.tabxo.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var db Db
	db.clientDatabase = client.Database("go")
	return &db
}

func (d *Db) NewDbUser(source requests.Request) (response.Response, error) {

	var search = requests.FindByUsernameRequest{
		Username: source.Username,
	}

	result, err := d.GetUser(search)
	if err != nil {
		return response.Response{}, err
	}

	if result != nil {
		return response.Response{}, errors.New("username already exists")
	}

	resultDB, err := d.clientDatabase.Collection("users").InsertOne(ctx, source)
	if err != nil {
		return response.Response{}, err
	}

	var insertedId, ok = resultDB.InsertedID.(primitive.ObjectID)
	if !ok {
		return response.Response{}, errors.New("error on insert")
	}

	var resp = response.Response{
		ID:        insertedId,
		Username:  source.Username,
		Email:     source.Email,
		Name:      source.Name,
		BirthDate: source.BirthDate,
	}

	return resp, nil
}

func (d *Db) GetUser(source requests.FindByUsernameRequest) (*response.Response, error) {
	src, err := bson.Marshal(source)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	err = bson.Unmarshal(src, &filter)
	if err != nil {
		return nil, err
	}
	var result response.Response
	err = d.clientDatabase.Collection("users").FindOne(ctx, source).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

//func InsertOne(source any) error {
//	_, err := database.Collection(source.collection).InsertOne(ctx, source)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func FindOne(source any) any {
//	result := database.Collection(source.collection).FindOne(ctx, source)
//	if result.Err() != nil {
//		return result.Err()
//	}
//	return result
//}
