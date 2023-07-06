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

func (d *Db) NewDbUser(source requests.Request) (response.User, error) {

	var searchUsername = requests.FindByUsername{
		Username: source.Username,
	}

	resultUsername, err := d.GetUserByName(searchUsername)
	if err != nil {
		return response.User{}, err
	}

	if resultUsername != nil {
		return response.User{}, errors.New("username already exists")
	}

	var searchEmail = requests.FindUserByEmail{
		Email: source.Email,
	}

	resultEmail, err := d.GetUserByEmail(searchEmail)
	if err != nil {
		return response.User{}, err
	}

	if resultEmail != nil {
		return response.User{}, errors.New("email already exists")
	}

	resultDB, err := d.clientDatabase.Collection("users").InsertOne(ctx, source)
	if err != nil {
		return response.User{}, err
	}

	var insertedId, ok = resultDB.InsertedID.(primitive.ObjectID)
	if !ok {
		return response.User{}, errors.New("error on insert")
	}

	var resp = response.User{
		ID:        insertedId,
		Username:  source.Username,
		Email:     source.Email,
		Name:      source.Name,
		BirthDate: source.BirthDate,
	}

	return resp, nil
}

func (d *Db) GetUserByName(source requests.FindByUsername) (*response.User, error) {
	src, err := bson.Marshal(source)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	err = bson.Unmarshal(src, &filter)
	if err != nil {
		return nil, err
	}
	var result response.User
	err = d.clientDatabase.Collection("users").FindOne(ctx, source).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (d *Db) GetUserByEmail(source requests.FindUserByEmail) (*response.User, error) {
	src, err := bson.Marshal(source)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	err = bson.Unmarshal(src, &filter)
	if err != nil {
		return nil, err
	}
	var result response.User
	err = d.clientDatabase.Collection("users").FindOne(ctx, source).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (d *Db) DeleteUser(source requests.FindByUsername) (*response.User, error) {
	var result response.User
	err := d.clientDatabase.Collection("users").FindOneAndDelete(ctx, source).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &result, err
}

func (d *Db) EditUser(source requests.EditUser, filter requests.FindByUsername) (*response.User, error) {
	update := bson.D{{"$set", source}}
	_, err := d.clientDatabase.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return d.GetUserByName(filter)
}
