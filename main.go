package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define the structure of the request to the Python backend
type ChatRequest struct {
	UserInput string `json:"user_input"`
}

// Define the structure of the response from the Python backend
type ChatResponse struct {
	Response string `json:"response"`
}

func main() {
	// Initialize Gin
	r := gin.Default()

	// Serve static files (like HTML, CSS, JS)
	r.Static("/static", "./static") // Serve files from the "static" folder

	// Serve the HTML file for the browser interface
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Define the POST endpoint for the frontend
	r.POST("/chat", func(c *gin.Context) {
		var chatReq ChatRequest

		// Parse the incoming request JSON
		if err := c.BindJSON(&chatReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Send the request to the Python backend
		pythonURL := "http://127.0.0.1:8000/chat" // Python backend URL
		jsonData, _ := json.Marshal(chatReq)
		resp, err := http.Post(pythonURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error calling Python backend:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
			return
		}
		defer resp.Body.Close()

		// Parse the Python backend response
		body, err := io.ReadAll(resp.Body) // Use io.ReadAll instead of ioutil.ReadAll
		if err != nil {
			log.Println("Error reading response body:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		var chatResp ChatResponse
		if err := json.Unmarshal(body, &chatResp); err != nil {
			log.Println("Error unmarshalling response:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process backend response"})
			return
		}

		// Return the backend response to the user
		c.JSON(http.StatusOK, chatResp)
	})

	// Start the Gin server on port 8000
	r.Run(":8080")
}
