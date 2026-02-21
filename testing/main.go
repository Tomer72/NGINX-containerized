package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func checkServer(url string, expectedStatus int) {
	fmt.Printf("[INFO] Testing %s (Expecting %d)...", url, expectedStatus)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[FAILURE] Connection error : %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != expectedStatus {
		fmt.Printf("[FAILURE] Got status %d instead of %d\n", resp.Status, expectedStatus)
		fmt.Printf("Response: %s\n", string(bodyBytes))
		os.Exit(1)
	}

	fmt.Println("[SUCCESS] Passed all tests")
}

func main() {
	fmt.Println("--- Starting Tests ---")

	checkServer("http://localhost:8081", 200)
	checkServer("http://localhost:8082", 503)

	fmt.Println("-----------------------")
	fmt.Println("All tests passed successfully")

}



