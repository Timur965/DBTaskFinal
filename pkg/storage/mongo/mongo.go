package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	db     *mongo.Client
	nameDb string
}

func New(namedb, port string) (*Mongo, error) {
	mongoOpts := options.Client().ApplyURI("mongodb://" + namedb + ":" + port + "/")
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	s := Mongo{db: client, nameDb: namedb}
	return &s, nil
}

func (c *Mongo) Posts() ([]storage.Post, error) {
	collection := c.db.Database(c.nameDb).Collection("posts")
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

func (c *Mongo) AddPost(post storage.Post) error {
	collection := c.db.Database(c.nameDb).Collection("posts")
	_, err := collection.InsertOne(context.Background(), post)
	return err
}
func (c *Mongo) UpdatePost(post storage.Post) error {
	collection := c.db.Database(c.nameDb).Collection("posts")
	filter := bson.D{{Key: "title", Value: post.Title}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "author_id", Value: post.AuthorID},
			{Key: "content", Value: post.Content},
			{Key: "created_at", Value: post.CreatedAt},
		}},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
func (c *Mongo) DeletePost(post storage.Post) error {
	collection := c.db.Database(c.nameDb).Collection("posts")
	filter := bson.D{{Key: "title", Value: post.Title}}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}
