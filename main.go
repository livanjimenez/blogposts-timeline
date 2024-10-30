package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/livanjimenez/blogposts-timeline/cmd"
)

func startServer() {
	// Start the HTTP server and listen for requests
	http.HandleFunc("/postblog", cmd.PostBlog)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendPostRequest() {
	// Create a sample blog post
	post := cmd.BlogPost{
		Title:   "My first blog post",
		Content: "This is my first blog post. I hope you enjoy it!",
	}

	// Convert the post to JSON
	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// Wait a moment to ensure the server is up
	time.Sleep(5 * time.Second) // Give server some time to start

	// Send the POST request
	resp, err := http.Post("http://localhost:8080/postblog", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Failed to send POST request:", err)
	}
	defer resp.Body.Close()

	// Print response status
	fmt.Println("Response Status:", resp.Status)
}

// ISSUE: Error 500 when sending a POST request to the server
// need to debug issue^

func main() {
	// Start the server in a goroutine to prevent blocking
	go startServer()

	// Send a POST request after the server is up
	sendPostRequest()

	// Block forever so the server keeps running
	select {}
}
