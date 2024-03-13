package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect() {
	godotenv.Load()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}
	fmt.Println("Connected to MongoDB!")

	filter := bson.D{{}}
	empColn := client.Database("Ecomm").Collection("employee")

	cursorEmp, err1 := empColn.Find(context.TODO(), filter)
	if err1 != nil {
		fmt.Println(err)
	}
	var results []bson.M
	if err = cursorEmp.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	// Now 'results' contains all documents that match the filter
	for _, result := range results {
		fmt.Println(result)
	}
}
