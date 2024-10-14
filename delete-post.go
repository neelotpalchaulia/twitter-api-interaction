package main // Defines that the following code is a part of the main package

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// function to load the .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Function to delete a tweet by its ID
func deleteTweet(tweetID string) {

	// Load Twitter API keys from environment variables
	config := oauth1.NewConfig(os.Getenv("API_KEY"), os.Getenv("API_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	// Create an HTTP client using OAuth1 config for making authenticated requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API endpoint to delete a tweet
	tweetURL := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)

	// Create a new HTTP DELETE request
	req, err := http.NewRequest("DELETE", tweetURL, nil)
	if err != nil {
		log.Fatalf("Error creating DELETE request: %v", err)
	}

	// This is to be used when we want to achieve Bearer Token (OAuth 2.0) authentication - common for app-only authentication
	// Settings Authorization header with the Bearer Token
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("BEARER_TOKEN")))
	// In the above, "Authorization" is the key for the header to auth the request

	// Send the request using the HTTP client
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending DELETE request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the tweet was successfully deleted
	if resp.StatusCode == 200 {
		fmt.Println("Tweet deleted successfully!")
	} else {
		fmt.Printf("Failed to delete tweet. Status code: %d\n", resp.StatusCode)
	}
}
