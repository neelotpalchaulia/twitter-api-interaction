package main // Defines that the following code is a part of the main package

import (
	"bufio"   // provides buffered I/O ops, allows reading inputs from users
	"fmt"     // provides funcs for formatting strings
	"os"      // Provides funcs to interact with the OS, allows access to env variables and handles i/o
	"strings" // provides funcs to manipulate string values

	// bytes and encoding/json for data-->JSON
	"bytes"
	"encoding/json"
	"io"       // basic i/o funcs
	"log"      //provides logging func for errors
	"net/http" // Allows us to make http requests

	"github.com/dghubble/oauth1" // Manages OAuth 1.0a authentication for API requests
	"github.com/joho/godotenv"
)

func main() {

	var tweetIDValue string // Variable to store the tweet ID

	// creates a new buffered reader - "reader" that reads input from the standard input (os.Stdin). It will be used to read user input from the console.
	reader := bufio.NewReader(os.Stdin)

	choice := ""           // Variable to store the user's choice
	for choice != "quit" { // Loop to keep the program running until the user chooses to exit

		// Display options for the user
		fmt.Println("Choose an option:")
		fmt.Println("1. Post a new tweet")
		fmt.Println("2. Delete the previously created tweet")
		fmt.Println("3. Delete a tweet by providing an ID")
		fmt.Println("Type 'quit' to exit")
		fmt.Print("Enter your choice: ")

		// Read user input
		choice, _ := reader.ReadString('\n') // this method reads until a newline char and saved in var choice
		choice = strings.TrimSpace(choice)   // the input which is saved in choice is trimmed

		// Handle user's choice using a switch statement
		switch choice {
		case "1":

			fmt.Print("Enter the content of your tweet: ")
			// Defining the tweet content yto be posted
			tweetContent, _ := reader.ReadString('\n')
			tweetContent = strings.TrimSpace(tweetContent)

			// Post the tweet and store its ID
			fmt.Println("Posting a new tweet...")
			tweetID := createTweet(tweetContent)
			fmt.Printf("Tweet posted with ID: %s\n", tweetID)
			tweetIDValue = tweetID // Store the tweet ID for deletion beacause the tweetID is to be used in multiple cases of the switch statement

		case "2":

			// Delete the previously created tweet using the stored tweetIDValue
			if tweetIDValue == "" {
				fmt.Println("No tweet has been created yet. Please create a tweet first.")
			} else {
				fmt.Println("Deleting the previously created tweet...")
				deleteTweet(tweetIDValue)
				fmt.Printf("Tweet with ID %s has been deleted.\n", tweetIDValue)
				tweetIDValue = "" // Reset tweetID after deletion
			}

		case "3":

			// Prompt the user for the tweet ID to delete
			fmt.Print("Enter the tweet ID to delete: ")
			inputTweetID, _ := reader.ReadString('\n')
			inputTweetID = strings.TrimSpace(inputTweetID)

			// Delete the tweet based on user-provided ID
			fmt.Println("Deleting the tweet with the provided ID...")
			deleteTweet(inputTweetID)
			fmt.Printf("Tweet with ID %s has been deleted.\n", inputTweetID)

		case "quit":

			fmt.Println("Exiting the program. Goodbye!")
			return

		default:

			fmt.Println("Invalid choice. Please try again")
		}
	}
}

// Function to create a new tweet
func createTweet(tweetContent string) string {

	// godotenv.Load(): reads env vars from a .env file [ensuring secure storage and access to our API keys & tokens]
	err := godotenv.Load() // thus, makes the env vars accessible within the GO program
	// Logs an error message and stops the program if the .env file has errors
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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
	var responseData map[string]interface{}                     // declares a map to store the parsed JSON response. The map's keys are strings, and the values can be of any type (interface{}).
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

// Function to delete a tweet by its ID
func deleteTweet(tweetID string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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
