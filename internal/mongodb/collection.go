package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Collection struct {
	orig *mongo.Collection
	name string
}

func (c Collection) Get(filter interface{}) *Result {
	res := c.orig.FindOne(context.TODO(), filter)
	if err := res.Err(); err != nil {
		log.Printf("Error getting from Collection(%s): %s", c.name, err)
		return nil
	}
	return &Result{orig: res}
}

func (c Collection) Add(v interface{}) error {
	_, err := c.orig.InsertOne(context.TODO(), v)
	if err != nil {
		log.Printf("Error adding document to Collection(%s): %s", c.name, err)
	}
	return err
}

type Result struct {
	orig *mongo.SingleResult
}

func (r *Result) Decode(v interface{}) interface{} {
	if r == nil {
		return nil
	}
	err := r.orig.Decode(v)
	if err != nil {
		log.Printf("Error decoding Result into type %T: %s", v, err)
		return nil
	}
	return v
}
