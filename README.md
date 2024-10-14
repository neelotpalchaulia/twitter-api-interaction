# Twitter API Interaction Using Go

## Introduction

This project demonstrates how to interact with the Twitter API using the Go programming language. It provides functionalities to post and delete tweets using OAuth 1.0a authentication, leveraging Twitter’s API endpoints. This project serves as a hands-on learning experience for developers looking to understand how to build Go applications that interact with third-party APIs. 

## Table of Contents
- [Introduction](#introduction)
- [Setup Instructions](#setup-instructions)
  - [Prerequisites](#prerequisites)
  - [Setting Up a Twitter Developer Account](#setting-up-a-twitter-developer-account)
  - [Environment Variables](#environment-variables)
- [Running the Application](#running-the-application)
- [Functionality](#functionality)
  - [1. Posting a Tweet](#1-posting-a-tweet)
  - [2. Deleting a Tweet](#2-deleting-a-tweet)
- [Error Handling](#error-handling)
- [Project Structure](#project-structure)

## Setup Instructions

### Prerequisites
- **Go 1.18 or later**: Download and install Go from [Go's official website](https://go.dev/doc/install).
- **Twitter Developer Account**: Required to access Twitter’s API. Follow the setup instructions below.
- **Git**: For version control and managing the project repository.

### Setting Up a Twitter Developer Account
1. **Sign Up**: Visit the [Twitter Developer Platform](https://developer.twitter.com) and apply for a Developer account.
2. **Create an App**: Under the "Projects & Apps" section, create a new app to access the API.
3. **Generate API Credentials**:
   - Navigate to the "Keys and Tokens" tab.
   - Click on **Generate** for the following credentials:
     - `API Key`
     - `API Key Secret`
     - `Access Token`
     - `Access Token Secret`
   - **Note**: Store these securely; they are required for making authenticated API requests.

### Environment Variables
Create a `.env` file in the project root and add your Twitter API credentials:

```plaintext
API_KEY=your_api_key
API_SECRET_KEY=your_api_secret_key
ACCESS_TOKEN=your_access_token
ACCESS_TOKEN_SECRET=your_access_token_secret
```

This `.env` file helps keep sensitive information out of your codebase, making it easier to manage and share your project.

## Running the Application

1. **Clone the Repository**:
   
   ```bash
   git clone https://github.com/neelotpalchaulia/twitter-api-interaction.git
   cd twitter-api-interaction
   ```

2. **Install Dependencies**:

   Ensure the necessary Go packages are installed:
   ```bash
   go mod tidy
   ```

3. **Run the Program**:
   Start the application with:
   ```bash
   go run main.go
   ```

   Follow the on-screen instructions to post or delete tweets.

## Functionality

### 1. Posting a Tweet
The program allows users to post a new tweet by providing the tweet content through the console. It sends a JSON payload containing the tweet content to the Twitter API using a POST request. If successful, the tweet's ID is returned and displayed.

**Example of a Request:**

```json
{
    "text": "Exploring Twitter API with Go!"
}
```

**Example of a Successful Response:**

```json
{
    "data": {
        "id": "1456789123456789012",
        "text": "Exploring Twitter API with Go!"
    }
}
```

### 2. Deleting a Tweet
The program can delete tweets using either:
- The ID of the last posted tweet (stored in memory).
- A user-provided tweet ID.

It sends a DELETE request to the Twitter API to remove the specified tweet. Users receive feedback on whether the operation was successful.

**Example of a Request:**
```http
DELETE https://api.twitter.com/2/tweets/1456789123456789012
```

**Example of a Successful Response:**
```json
{
    "message": "Tweet deleted successfully!"
}
```

## Error Handling
The application includes robust error handling to ensure smooth user experience:
- **Missing `.env` File**: Logs an error if the `.env` file is not found, prompting the user to check their setup.
- **Invalid API Credentials**: If API keys or tokens are incorrect, Twitter returns a `401 Unauthorized` status. The program handles this by displaying a user-friendly message.
- **Empty Tweet Content**: Prevents posting tweets with empty content by validating input.
- **Network Issues**: Catches network errors (e.g., timeouts) and logs them for easier debugging.
- **Invalid Tweet ID**: Handles errors when attempting to delete tweets with non-existent IDs, providing feedback to the user.

## Project Structure
```
twitter-api-interaction/
│
├── main.go               # Entry point of the program
├── .env                  # Stores environment variables (not included in the repo)
├── README.md             # Documentation of the project
├── go.mod                # Module dependencies
└── go.sum                # Checksums for module dependencies
```

