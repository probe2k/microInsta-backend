package setup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"github.com/joho/godotenv"
)

func ConnectPostsDB() *mongo.Collection {
	clientData := options.Client().ApplyURI("mongodb+srv://Admin:" + os.Getenv("DBKEY") + "@cluster0.rre0y.mongodb.net/Database?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientData)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	collection := client.Database("go_rest_api").Collection("posts")

	return collection
}

func ConnectUsersDB() *mongo.Collection {
	clientData := options.Client().ApplyURI("mongodb+srv://Admin:" + os.Getenv("DBKEY") + "@cluster0.rre0y.mongodb.net/Database?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	collection := client.Database("go_rest_api").Collection("users")

	return collection
}

type ErrorResponse struct {
	StatusCode int `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var res = ErrorResponse {
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	msg, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(msg)
}