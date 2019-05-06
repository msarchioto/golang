/****************************
* Author: Michele Sarchioto	*
* Project: SAP-Trenitalia  	*
* Description: MongoDB Test	*
* Date: 14/02/2019			*
****************************/

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Items You will be using this Trainer type later in the program
type Items struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	ItemID   string             `json:"itemID" bson:"itemID"`
	Priority int                `json:"priority" bson:"priority"`
	Desc     string             `json:"desc" bson:"desc"`
}

func main() {

	fmt.Println("MongoDB Start & connect")

	// *** Connect *** //
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	errorHandler(err)

	// Check the connection, simple ping, no context
	err = client.Ping(context.TODO(), nil)
	errorHandler(err)

	// check the connection, method with context and 2 seconds timeout
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	fmt.Println("Connected to MongoDB!")

	// collection
	collection := client.Database("test").Collection("items")

	// *** CRUD OPERATIONS *** //

	// create new items to insert
	item1 := Items{primitive.NewObjectID(), "ID_A_001", 15, "Description 1"}
	item2 := Items{primitive.NewObjectID(), "ID_A_002", 1, "Description 2"}
	item3 := Items{primitive.NewObjectID(), "ID_A_003", 30, "Description 3"}

	// *** insert single document *** //
	bsonDocument := convertItemsToBson(&item1)
	insertResult, err := collection.InsertOne(context.TODO(), bsonDocument)
	errorHandler(err)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// *** insert multiple documents *** //
	doc2 := convertItemsToBson(&item2)
	doc3 := convertItemsToBson(&item3)
	bsonDocumentMult := []interface{}{doc2, doc3}
	insertManyResult, err := collection.InsertMany(context.TODO(), bsonDocumentMult)
	errorHandler(err)
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// *** Update documents *** //
	filter := bson.M{"itemID": "ID_A_001"}

	// define update: increase priority by 1
	update := bson.D{
		{"$inc", bson.D{
			{"priority", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	errorHandler(err)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// *** find single item and write it into a struct instance *** //
	var resultItem Items
	err = collection.FindOne(context.TODO(), filter).Decode(&resultItem)
	errorHandler(err)
	fmt.Printf("Found a single document: %+v\n", resultItem)

	// *** Find multiple elements with cursor *** //
	emptyFilter := bson.M{} // empty filter

	// Options, limit 2
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Items

	cur, err := collection.Find(context.TODO(), emptyFilter, findOptions)
	errorHandler(err)

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Items
		err := cur.Decode(&elem)
		errorHandler(err)
		results = append(results, &elem)

	}
	// check error on cursor
	err = cur.Err()
	errorHandler(err)

	// close cursor
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// *** Delete the collection *** //
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.M{})
	errorHandler(err)
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// *** Close connection *** //
	err = client.Disconnect(context.TODO())
	errorHandler(err)
	fmt.Println("Connection to MongoDB closed.")

}

// errorHandler to avoid verbose error handling
func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// convertItemsToBson converts Items struct to bson.M
func convertItemsToBson(item *Items) primitive.M {

	documentToReturn := bson.M{"itemID": item.ItemID, "priority": item.Priority, "desc": item.Desc}
	return documentToReturn

}
