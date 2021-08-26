package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Storage struct {
	db  *mongo.Database
}

func NewStorage(url string, dbname string) (*Storage, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return &Storage{
		db:  client.Database(dbname),
	}, nil
}

func (store *Storage) FindBoard(id string) (*Board, error) {
	filter := bson.D{{"id", id}}
	var res Board
	err := store.db.Collection("boards").FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Print("cannot find board by id", err)
	}
	return &res, nil
}

func (store *Storage) CreateBoard(board *Board) error {
	_, err := store.db.Collection("boards").InsertOne(context.Background(), board)
	if err != nil {
		log.Print("cannot find board by id", err)
	}
	return err
}

func (store *Storage) SaveBoard(board *Board) (*Board, error) {
	filter := bson.D{{"id", board.Id}}
	_, err := store.db.Collection("boards").ReplaceOne(context.Background(), filter, board)
	if err != nil {
		log.Print("cannot find board by id", err)
	}
	return board, err
}
