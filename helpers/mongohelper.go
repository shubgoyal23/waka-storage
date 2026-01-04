package helpers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var MongoConn *mongo.Client
var MongoDb string

func MongoInit(uri string, dbName string) (f bool) {
	f = false
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		Logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Error("Failed to ping MongoDB", zap.Error(err))
		return
	}
	MongoConn = client
	MongoDb = dbName
	return true
}

// add doc to mongo db
func MongoAddManyDoc(collection string, doc []interface{}) (f bool) {
	f = false
	client := MongoConn.Database(MongoDb).Collection(collection)
	opts := options.InsertMany().SetOrdered(false)
	ints, err := client.InsertMany(context.TODO(), doc, opts)
	if err != nil {
		Logger.Error("Failed to insert many documents", zap.Error(err))
		return
	}
	if ints.InsertedIDs == nil {
		return
	}
	return true
}

// get many doc from mongo db
func MongoGetLastOneDoc(collection string, docInp interface{}) (f bool) {
	f = false
	client := MongoConn.Database(MongoDb).Collection(collection)
	// Find the last inserted document
	findOptions := options.FindOne().
		SetSort(bson.D{{Key: "time", Value: -1}})
	if err := client.FindOne(context.TODO(), bson.D{}, findOptions).Decode(docInp); err != nil {
		Logger.Error("Failed to find one document", zap.Error(err))
		return
	}
	f = true
	return
}

// // add doc to mongo db
// func MongoAddOncDoc(collection string, doc interface{}) (f bool) {
// 	f = false
// 	client := MongoConn.Database(MongoDb).Collection(collection)
// 	ints, err := client.InsertOne(context.TODO(), doc)
// 	if err != nil {
// 		Logger.Error("Failed to insert one document", zap.Error(err))
// 		return
// 	}
// 	if ints.InsertedID == nil {
// 		return
// 	}
// 	return true
// }

// // get many doc from mongo db
// func MongoGetManyDoc(collection string, filter any) (doc []bson.M, f bool) {
// 	doc = []bson.M{}
// 	f = false
// 	client := MongoConn.Database(MongoDb).Collection(collection)
// 	cursor, err := client.Find(context.TODO(), filter)
// 	if err != nil {
// 		Logger.Error("Failed to find documents", zap.Error(err))
// 		return
// 	}
// 	err = cursor.All(context.TODO(), &doc)
// 	if err != nil {
// 		Logger.Error("Failed to decode documents", zap.Error(err))
// 		return
// 	}
// 	f = true
// 	return
// }

// // get many doc from mongo db
// func MongoGetOneDoc(collection string, filter interface{}, docInp interface{}) (f bool) {
// 	f = false
// 	client := MongoConn.Database(MongoDb).Collection(collection)
// 	if err := client.FindOne(context.TODO(), filter).Decode(docInp); err != nil {
// 		Logger.Error("Failed to find one document", zap.Error(err))
// 		return
// 	}
// 	f = true
// 	return
// }

// // delete many doc from mongo db
// func MongoDeleteManyDoc(collection string, filter interface{}) (f bool) {
// 	f = false
// 	client := MongoConn.Database(MongoDb).Collection(collection)
// 	_, err := client.DeleteMany(context.TODO(), filter)
// 	if err != nil {
// 		Logger.Error("Failed to delete many documents", zap.Error(err))
// 		return
// 	}
// 	return true
// }

// // update many docs
// func MongoUpdateManyDoc(collection string, filter, update bson.M) error {
// 	client := MongoConn.Database(MongoDb).Collection(collection)
// 	res, err := client.UpdateMany(context.TODO(), filter, update)
// 	if err != nil {
// 		Logger.Error("Failed to update many documents", zap.Error(err))
// 		return err
// 	}
// 	total := res.ModifiedCount + res.UpsertedCount
// 	if res.MatchedCount != total {
// 		return fmt.Errorf("matched count and updated count mismatch")
// 	}
// 	return nil
// }

// // update one docs
// func MongoUpdateOneDoc(collection string, filter, update bson.M) error {
// 	client := MongoConn.Database(MongoDb).Collection(collection)
// 	res, err := client.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		Logger.Error("Failed to update one document", zap.Error(err))
// 		return err
// 	}
// 	if res.MatchedCount != res.UpsertedCount {
// 		return fmt.Errorf("matched count and updated count mismatch")
// 	}
// 	return nil
// }
