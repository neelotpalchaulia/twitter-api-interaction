package main // indicates that this is the main package of the program

import ( //imports packages needed for the program
	// bytes and encoding/json for data-->JSON
	"bytes"
	"encoding/json"
	"io" // basic i/o funcs
	"fmt"      // provides funcs for formatting strings
	"log"      //provides logging func for errors
	"net/http" // Allows us to make http requests
	"os"       // Provides funcs to interact with the OS, allows access to env variables and handles i/o

	"github.com/dghubble/oauth1" // Manages OAuth 1.0a authentication for API requests
	"github.com/joho/godotenv"   // Reads env variables from a .env file
)

// function to load the .env file
// init() - special func in Go that runs before the main(), used here to load env variables
func init() {
	// godotenv.Load(): reads env vars from a .env file [ensuring secure storage and access to our API keys & tokens]
	err := godotenv.Load() // thus, makes the env vars accessible within the GO program
	// Logs an error message and stops the program if the .env file has errors
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Function to create a new tweet
func createTweet(tweetContent string) string {

	// Load Twitter API keys from environment variables
	// OAuth1 setup with credentials
	config := oauth1.NewConfig(os.Getenv("API_KEY"), os.Getenv("API_SECRET_KEY"))         // sets up oauth1 config using api key and secret
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET")) // creates an oauth token using access token and secret

	// Create an HTTP client using OAuth1 config for making authenticated requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API endpoint to create a tweet
	tweetURL := "https://api.twitter.com/2/tweets"

	// Create the JSON payload containing the tweet content
	requestBody, err := json.Marshal(map[string]string{
		"text": tweetContent,
	})
	if err != nil {
		// Logs an error message and stops the program if the .env file has errors
		log.Fatalf("Error creating JSON request body: %v", err)
	}

	// Create a new HTTP POST request with the tweet content
	req, err := http.NewRequest("POST", tweetURL, bytes.NewBuffer(requestBody)) //creates a new http req, here a POST req
	if err != nil {
		log.Fatalf("Error creating the POST request: %v", err)
	}

	// This is to be used when we want to achieve Bearer Token (OAuth 2.0) authentication - common for app-only authentication
	// Settings Authorization header with the Bearer Token
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("BEARER_TOKEN")))
	// In the above, "Authorization" is the key for the header to auth the request
	// In scenarios adopting OAuth2 authentication, we can omit configuring OAuth1 @ line 32, 33 and 36

	// Set the Content-Type header to indicate that the req body contains JSON data
	req.Header.Set("Content-Type", "application/json")

	// Finally making/sending the request using the HTTP client
	resp, err := httpClient.Do(req) // Executes the HTTP req and returns a response which is saved in resp
	if err != nil {
		log.Fatalf("Error sending the POST request: %v", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }

	defer resp.Body.Close() // ensures the the resp body is closed after reading and freeing the resources

	// Check if the tweet was successfully posted
	if resp.StatusCode == 201 {
		fmt.Println("Tweet posted successfully!")
	} else {
		fmt.Printf("Failed to post tweet. Status code: %d\n", resp.StatusCode)
	}

	// Parse the response to extract the tweet ID
    var responseData map[string]interface{} // declares a map to store the parsed JSON response. The map's keys are strings, and the values can be of any type (interface{}).
    if err := json.Unmarshal(body, &responseData); err != nil { // json.unmarshal(): converts response in JSON --> map in responseData
        log.Fatalf("Error parsing JSON response: %v", err)
    }

	// Extract the tweet ID from the response
    tweetID, ok := responseData["data"].(map[string]interface{})["id"].(string)
    if !ok {
        log.Fatalf("Error extracting tweet ID from response")
    }

    return tweetID
}
