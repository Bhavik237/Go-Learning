package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// URL of the wrapped API
	url := "http://10.125.14.182:9000/v1/im/admin/token?secret=test"

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		log.Printf("Error creating request: %v", err)
		return
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		log.Printf("Error reading response: %v", err)
		return
	}

	// Set Content-Type header for the response
	w.Header().Set("Content-Type", "application/json")
	// Write the response body to the client
	w.Write(body)
}

func main() {
	// Handle "/api/hello" endpoint
	http.HandleFunc("/api/hello", helloHandler)

	// Start the HTTP server
	fmt.Println("Go server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
