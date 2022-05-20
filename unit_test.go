package main

import (
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestGetPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/posts/6rapw0kk1bjazxrqh894qllx5qqy7scd", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getPosts)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Error("wrong status code")
	}

	expected := `{"PID": "s3u89f2", "_id":"6rapw0kk1bjazxrqh894qllx5qqy7scd", "title": "Test 1", "img": "{"/api/posts/img1.png", "/api/posts/img2.png"}", "desc": "This is a new post"}`

	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Error("Unexpected Body!")
	}
}

func TestCreatePost(t *testing.T) {
	var jsonData = []byte(`{"title": "test2", "img": "{"/api/posts/image/file1.png"}", "desc": "create post test"}`)
	req, err := http.NewRequest("POST", "/api/users/", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(createPost)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Error("Unexpected status code!")
	}

	expected := `{"PID": "238797as8d9f723asdfhsdf2121"}`
	if len(strings.TrimSpace(rr.Body.String())) != len(expected) {
		t.Error("Unexpected Body!")
	}
}

func TestCreateUser(t *testing.T) {
	var jsonData = []byte(`{"Name": "", "Mobile": "", "Address": ""}`)
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Error("Unexpected error!")
	}
	expected := `{"ID": "6rapw0kk1bjazxrqh894qllx5qqy7scd"}`
	if len(strings.TrimSpace(rec.Body.String())) != len(expected) {
		t.Error("Unexpected Body!")
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "api/users/6rapw0kk1bjazxrqh894qllx5qqy7scd", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getUser)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Error("Unexpected Error")
	}

	expected := `{"_id": "", "name": "", "mobile": "", "addres": "", "postCount": ""}`
	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Error("Unexpected Body!")
	}
}

func TestGetPostsByUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/6rapw0kk1bjazxrqh894qllx5qqy7scd", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(postsByUser)
	handler.ServeHTTP(rec, req)

	expcted := `{"Posts": [{"PID": "", "_id": "", "title": "", "image": {"", ""}, "desc": ""}], "lowerId": "6rapw0kk1bjazxrqh894qllx5qqy7scd"}`
	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Error("Unexpected Body!")
	}
}