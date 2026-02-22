package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func checkServer(url string, expectedStatus int) {
	/*
	Function that recieves a url and expected HTTP response from the NGINX server, 
	and returns valid (0 exit code) or invalid (1 exit code)
	*/
	fmt.Printf("[INFO] Testing %s (Expecting %d)...", url, expectedStatus)

	// Sending HTTP GET request to the server
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("[FAILURE] Connection error : %v\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	if response.StatusCode != expectedStatus {
		fmt.Printf("[FAILURE] Got status %d instead of %d\n", response.StatusCode, expectedStatus)
		fmt.Printf("Response: %s\n", string(bodyBytes))
		os.Exit(1)
	}

	fmt.Println("[SUCCESS] Passed")
}

func main() {

	// Fetch URLs from environment variables (.env file).
	url1 := os.Getenv("SERVER1_URL")
	url2 := os.Getenv("SERVER2_URL")

	if url1 == ""  {
		fmt.Println("[WARNING] SERVER1_URL not found, using default")
		url1 = "http://nginx-service:8081"
	}
	if url2 == ""  {
		fmt.Println("[WARNING] SERVER2_URL not found, using default")
		url2 = "http://nginx-service:8082"
	}

	fmt.Println("--- Starting Tests ---")

	checkServer(url1, 200)
	checkServer(url2, 503)

	fmt.Println("-----------------------")
	fmt.Println("All tests passed successfully")

}



