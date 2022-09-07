package MongoDB

import (
	"BurninProject/aop/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoConn struct {
	clientOptions *options.ClientOptions
	client        *mongo.Client
	collections   *mongo.Collection
}

//var mongoConn *MongoConn

func InitMongoConn(url, user, password, dbname string) *MongoConn {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConn := &MongoConn{}
	//Set client options
	//construct url: mongodb://username:password@127.0.0.1:27017/dbname
	mongoUrl := "mongodb://" + user + ":" + password + "@" + url + "/" + dbname
	mongoUrl = "mongodb://127.0.0.1:27017/"
	mongoConn.clientOptions = options.Client().ApplyURI(mongoUrl)

	//Connect to MongoDB
	var err error
	mongoConn.client, err = mongo.Connect(ctx, mongoConn.clientOptions)
	if err != nil {
		logger.Logger.ErrorF("connect to mongodb error: ", err)
	}

	//check the connection
	err = mongoConn.client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Logger.ErrorF("connect to mongodb error: ", err)
	}
	mongoConn.collections = mongoConn.client.Database(dbname).Collection("Account")
	return mongoConn
}

func (mongoConn *MongoConn) CloseMongoConn() {
	err := mongoConn.client.Disconnect(context.TODO())
	if err != nil {
		logger.Logger.ErrorF("disconnect mongodb connect is error: ", err)
		return
	}
	logger.Logger.InfoF("connection to MongoDB closed.")
}

func (mongoConn *MongoConn) InsertOneData(data interface{}) {
	//insert dat into database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertResult, err := mongoConn.collections.InsertOne(ctx, data)
	if err != nil {
		logger.Logger.ErrorF("mongodb InsertOneData error: ", err)
	}
	fmt.Println("Inserted a single document = ", insertResult.InsertedID)
}

func (mongoConn *MongoConn) InsertManyData(datas []interface{}) {
	// insert datas into database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertManyResult, err := mongoConn.collections.InsertMany(ctx, datas)
	if err != nil {
		logger.Logger.ErrorF("mongodb InsertManyData err: %s", err)
	}
	fmt.Println("successfully inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func (mongoConn *MongoConn) FindOne(key string, value interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{key, value}}
	singleResult := mongoConn.collections.FindOne(ctx, filter)
	if singleResult != nil {
		fmt.Println(singleResult)
	}
	return singleResult
}

func (mongoConn *MongoConn) FindLimit(pagae int64, filter bson.D) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 分页查询选项设置
	// Pass these options to the Find method
	findOptions := options.Find()
	// use setlimit and setskip to implement pageable query
	findOptions.SetLimit(pagae)
	// specifies the number of documents to skip before returning.
	findOptions.SetSkip(pagae * (5 - 1))

	cur, err := mongoConn.collections.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, context.TODO())
}

func (mongoConn *MongoConn) Delect(filter bson.D) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	deleteResult, err := mongoConn.collections.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the tests collection\n", deleteResult.DeletedCount)
}

func (mongoConn *MongoConn) DelectAll() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := mongoConn.collections.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
