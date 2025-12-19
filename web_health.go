package main

import (
	"fmt"
	"net/http"
	"time"
)

// SiteResult holds the data we want to send back from the worker
type SiteResult struct {
	URL    string
	Status int
	Err    error
}

// checkUrl is our "worker" function.
// It takes a URL and a channel to send the result back to.
func checkUrl(url string, c chan SiteResult) {
	// 1. Record start time (optional, but fun to see speed)
	start := time.Now()

	// 2. Make the HTTP request
	// We set a short timeout so we don't hang forever on bad sites
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)

	// 3. Calculate how long it took
	duration := time.Since(start)

	// 4. Prepare the result
	result := SiteResult{
		URL: url,
	}

	if err != nil {
		result.Err = err
	} else {
		result.Status = resp.StatusCode
		resp.Body.Close() // Always close the body!
	}

	// 5. Send the result into the channel
	fmt.Printf("[%s] Checked %s\n", duration, url) // Log activity
	c <- result
}

func main() {
	// A list of safe, public URLs to test
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://golang.org",
		"https://www.doesntexist12345.com", // Intentionally bad URL to test errors
	}

	// 1. Create a channel to communicate between routines
	// We make it "buffered" (len(urls)) so workers don't get stuck waiting
	resultsChannel := make(chan SiteResult, len(urls))

	fmt.Println("Starting concurrent checks...")
	start := time.Now()

	// 2. Launch a goroutine for each URL
	for _, url := range urls {
		// The 'go' keyword starts a new thread
		go checkUrl(url, resultsChannel)
	}

	// 3. Collect the results
	// Since we launched 'len(urls)' tasks, we expect 'len(urls)' answers.
	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel // Wait here for a message

		if result.Err != nil {
			fmt.Printf("❌ %s is DOWN (Error: %v)\n", result.URL, result.Err)
		} else {
			fmt.Printf("✅ %s is UP (Status: %d)\n", result.URL, result.Status)
		}
	}

	fmt.Printf("\nTotal time taken: %s\n", time.Since(start))
}
