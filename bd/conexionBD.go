package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*MongoCN es el objeto de conexión a la BD */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://rleon2020:1PppUaDYqYUEHh4Q@cluster0.e0fnv.mongodb.net/twitter-clone?retryWrites=true&w=majority")

/*ConnectDB for connect to db */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection successful")
	return client
}

/*CheckConnection for check to db */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
