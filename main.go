package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Check for URL argument
	if (len(os.Args) < 2 || os.Args[1] == "help") {
		fmt.Println("Swirl 1.0")
		fmt.Println("------------------")
		fmt.Println("Usage: swirl <url>")
		return
	}

	// Get URL from argument
	url := os.Args[1]

	// Create an HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching content:", err)
		return
	}

	defer resp.Body.Close() // Ensure body is closed

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the page content
	fmt.Println(string(body))
}
