// controllers/chatbot_controller.go

package controllers

import (
	"Server/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// ChatbotRequest represents the request body for the chatbot API
type ChatbotRequest struct {
	InputText string `json:"inputText"`
}

// ChatbotResponse represents the response body for the chatbot API
type ChatbotResponse struct {
	Text string `json:"text"`
}

// GetChatbotResponse is the handler for the chatbot API call
func GetChatbotResponse(c *gin.Context) {
	apiKey := utils.GetGeminiAPIKey()

	var request ChatbotRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Construct the path to the prompt.txt file
	promptPath := filepath.Join(dir, "utils", "prompt.txt")

	// Read the text from prompt.txt
	promptBytes, err := os.ReadFile(promptPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Convert the content to string
	promptText := string(promptBytes)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"role": "user",
				"parts": []map[string]interface{}{
					{
						"text": promptText,
					},
				},
			},
			{
				"role": "model",
				"parts": []map[string]interface{}{
					{
						"text": "Hi",
					},
				},
			},
			{
				"role": "user",
				"parts": []map[string]interface{}{
					{
						"text": request.InputText,
					},
				},
			},
		},
		"safetySettings": []map[string]interface{}{
			{
				"category":  "HARM_CATEGORY_SEXUALLY_EXPLICIT",
				"threshold": "BLOCK_NONE",
			},
			{
				"category":  "HARM_CATEGORY_HATE_SPEECH",
				"threshold": "BLOCK_NONE",
			},
			{
				"category":  "HARM_CATEGORY_HARASSMENT",
				"threshold": "BLOCK_NONE",
			},
			{
				"category":  "HARM_CATEGORY_DANGEROUS_CONTENT",
				"threshold": "BLOCK_NONE",
			},
		},
	}

	responseBody, err := utils.MakeGeminiAPIRequest(apiKey, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	generatedText, err := utils.ParseGeminiAPIResponse(responseBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"text": generatedText})
}
