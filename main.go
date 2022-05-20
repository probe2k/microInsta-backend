package main

import (
	"fmt"
	"context"
	"encoding/hex"
	"encoding/json"
	"../model"
	"log"
	"net/http"
	"strings"
	"strconv"
	"time"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"../setup"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}


type postFilter struct {
	LimitedPosts []model.Post `json:"Posts"`
	LowerId string `json:"lowerId"`
}

var listPost = setup.ConnectPostsDB()
var listUser = setup.ConnectUsersDB()

var key, _ = getKey()

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post model.Post
	PID := strings.TrimPrefix(r.URL.Path, "/api/posts/")
	id, _ := primitive.ObjectIDFromHex(PID)
	filter := bson.M{"_id": id}
	err := listPost.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		setup.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post model.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.Timestamp = time.Now()
	result, err := listPost.InsertOne(context.TODO(), post)
	if err != nil {
		setup.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	result, err := listUser.insertOne(context.TODO(), user)
	if err != nil {
		setup.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	ID := strings.TrimPrefix(r.URL.Path, "/api/users/")
	id, _ := primitive.ObjectIDFromHex(ID)

	err := listUser.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		setup.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func postsByUser(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/json")
	var posts []model.Post
	var user model.User
	ID := strings.TrimPrefix(r.URL.Path, "/api/posts/users/")
	ID = string.Split(ID, "?")[0]
	query := r.URL.Query()
	range, _ := strconv.ParseInt(query["limit"][0], 10, 64)
	lowerID := query["lowerid"]
	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": id}
	err := listUser.FindOne(context.TODO(), filter).Decode(&user)
	cur, err := listPost.Find(context.TODO(), bson.M{"author": user.Name})
	if err != nil {
		setup.GetError(err, w)
		return
	}
	counter := 0
	var retId string
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var post model.Post
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		if len(lowerID) > 0 {
			if strings.Compare(post.ID.Hex(), lowerID[0]) == 1 {
				posts = append(posts, post)
				counter += 1
			}
		} else {
			posts = append(posts, post)
			counter += 1
		}
		retId = post.ID.Hex()
		if int64(counter) == range {
			break
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	retObj := PostFilter{LimitedPosts: posts, LowerId: retId}
	ret, err := json.Marshal(retObj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.FPrintf(w, string(ret))
}

func main() {
	http.HandleFunc("/api/posts", createPost)
	http.HandleFunc("/api/posts", getPosts)
	http.HandleFunc("/api/posts/users", postsByUser)
	http.HandleFunc("/api/users/", createUser)
	http.HandleFunc("/api/users/", getUser)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
