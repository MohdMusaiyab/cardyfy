package db
import (
    "context"
    "log"
    "os"
    "sync"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

var (
    clientInstance *mongo.Client
    clientOnce     sync.Once
    mongoURI       string
)

// Connect initializes the MongoDB client singleton, loads .env and Mongo URI
func Connect() *mongo.Client {
    clientOnce.Do(func() {
        // Load .env file
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file found, continuing with environment variables")
        }

        mongoURI = os.Getenv("MONGODB_URI")
        if mongoURI == "" {
            log.Fatal("MONGODB_URI environment variable is not set")
        }

        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        clientOptions := options.Client().ApplyURI(mongoURI)
        client, err := mongo.Connect(ctx, clientOptions)
        if err != nil {
            log.Fatal("Error connecting to MongoDB:", err)
        }

        // Ping DB to check connection is alive
        if err := client.Ping(ctx, nil); err != nil {
            log.Fatal("Could not ping to MongoDB:", err)
        }

        clientInstance = client
        log.Println("Successfully connected to MongoDB")
    })

    return clientInstance
}

// Disconnect properly disconnects the Mongo client (deferred in main)
func Disconnect() {
    if clientInstance != nil {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        if err := clientInstance.Disconnect(ctx); err != nil {
            log.Println("Error disconnecting MongoDB client:", err)
        } else {
            log.Println("Disconnected from MongoDB")
        }
    }
}