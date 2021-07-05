package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var articles []Article

type Article struct {
	Id      string
	Title   string
	Desc    string
	Content string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	rs := []byte("Welcome to the HomePage!")
	w.Write(rs)
}

func findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func findById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["id"]

	var article *Article
	for _, v := range articles {
		if v.Id == param {
			article = &v
			break
		}
	}

	if article != nil {
		json.NewEncoder(w).Encode(article)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Article Not Found!"))
	}

}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article

	d := json.NewDecoder(r.Body)
	d.Decode(&article)

	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article

	d := json.NewDecoder(r.Body)
	d.Decode(&article)

	var articleToUpdate *Article
	for i := 0; i < len(articles); i++ {
		if articles[i].Id == article.Id {
			articles[i] = article

			articleToUpdate = &articles[i]
			break
		}
	}

	if articleToUpdate != nil {
		json.NewEncoder(w).Encode(articleToUpdate)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Failed to update article!"))
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["id"]

	isDeleted := false
	for index, v := range articles {
		if v.Id == param {
			isDeleted = true
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}

	if isDeleted {
		w.Write([]byte("Successfully deleted article"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Failed to delete article Not Found"))
	}
}

func handleRequest() {
	// http.HandleFunc("/", homePage)

	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", findAll).Methods("GET")
	router.HandleFunc("/articles/{id}", findById).Methods("GET")
	router.HandleFunc("/articles", create).Methods("POST")
	router.HandleFunc("/articles", update).Methods("PUT")
	router.HandleFunc("/articles/{id}", delete).Methods("DELETE")

	log.Println("The server started listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	articles = []Article{
		{"1", "Hello", "Description 1", "Article Content"},
		{"2", "Hello 2", "Description 2", "Article Content"},
	}

	log.Println("Initializing requests handlers")
	handleRequest()

}
