package database

import (
	"context"
	"fmt"

	"github.com/nem0z/cda/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	posts  *mongo.Collection
	ctx    context.Context
}

func Init() (*Mongo, error) {
	// init client
	options := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	// check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// load collection
	posts := client.Database("cda").Collection("posts")

	return &Mongo{
		client,
		posts,
		ctx,
	}, nil
}

func (m *Mongo) InsertOne(post *types.Post) (interface{}, error) {
	fmt.Println("Inserted", post.Index)
	result, err := m.posts.InsertOne(m.ctx, post)
	return result.InsertedID, err
}

func (m *Mongo) FindOne(id interface{}) (*types.Post, error) {
	post := &types.Post{}
	err := m.posts.FindOne(m.ctx, bson.M{"_id": id}).Decode(&post)
	return post, err
}

func (m *Mongo) FindAll() (*types.Post, error) {
	post := &types.Post{}
	err := m.posts.FindOne(m.ctx, bson.M{}).Decode(&post)
	return post, err
}

func (m *Mongo) Clear() error {
	_, err := m.posts.DeleteMany(m.ctx, bson.M{})
	return err
}
