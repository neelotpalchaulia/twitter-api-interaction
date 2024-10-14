package main

import (
	"bufio" // provides buffered I/O ops, allows reading inputs from users
	"fmt"
	"os"
	"strings" // provides funcs to manipulate string values	
)

func main() {

	var tweetID string
	// creates a new buffered reader - "reader" that reads input from the standard input (os.Stdin). It will be used to read user input from the console.
	reader := bufio.NewReader(os.Stdin)

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

	case "2":

		// Delete the previously created tweet using the stored tweetID
		if tweetID == "" {
			fmt.Println("No tweet has been created yet. Please create a tweet first.")
		} else {
			fmt.Println("Deleting the previously created tweet...")
			deleteTweet(tweetID)
			fmt.Printf("Tweet with ID %s has been deleted.\n", tweetID)
			tweetID = "" // Reset tweetID after deletion
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