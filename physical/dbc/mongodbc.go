package dbc

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func New(config *viper.Viper) *mongo.Client {
	connectionString := config.GetString("database.connection_string")

	if connectionString == "" {
		log.Fatalf("Database connection string is missing")
	}

	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}*/

	opts := options.Client()
	opts.Monitor = otelmongo.NewMonitor()
	opts.ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB")

	return client
}
