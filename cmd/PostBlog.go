package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ConnectToDB() *sql.DB {
	connStr := "host=localhost port=5432 user=root dbname=dev password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Check if the connection is actually working
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to the database")
	return db
}


type BlogPost struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func PostBlog(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return 
    }

    var post BlogPost
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return 
    }

    db := ConnectToDB()
    defer db.Close()

    _, err = db.Exec("INSERT INTO blogposts (title, content) VALUES ($1, $2)", post.Title, post.Content)
    if err != nil {
        http.Error(w, "Error inserting blog post", http.StatusInternalServerError)
        return 
    }
    fmt.Fprintln(w, "Inserted a blog post")
}
