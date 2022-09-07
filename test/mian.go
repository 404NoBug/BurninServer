package main

import (
	"BurninProject/engine/MongoDB"
	"fmt"
)

type PlayerAccount struct {
	ACCOUNT  string
	PASSWORD string
}

func main() {
	mongoDB := MongoDB.InitMongoConn("127.0.0.1", "", "", "Burnin")
	playerAccount := PlayerAccount{ACCOUNT: "gy001", PASSWORD: "123456"}
	mongoDB.InsertOneData(playerAccount)
	N := mongoDB.FindOne("account", "gy001")
	if N != nil {
		var result interface{}
		err := N.Decode(&result)
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("test = ", result)
	}

	//if err := godotenv.Load(); err != nil {
	//	log.Println("No .env file found")
	//}
	//uri := "mongodb://127.0.0.1:27017/" //os.Getenv("MONGODB_URI")
	//if uri == "" {
	//	log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	//}
	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	//coll := client.Database("sample_mflix").Collection("movies")
	//title := ""
	//var result bson.M
	//err = coll.FindOne(context.TODO(), bson.D{{"sese", title}}).Decode(&result)
	//if err == mongo.ErrNoDocuments {
	//	fmt.Printf("No document was found with the title %s\n", title)
	//	return
	//}
	//if err != nil {
	//	panic(err)
	//}
	//jsonData, err := json.MarshalIndent(result, "", "    ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", jsonData)

}
