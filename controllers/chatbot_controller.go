// controllers/chatbot_controller.go

package controllers

import (
	"Server/utils"
	"net/http"
	"log"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// ChatbotRequest represents the request body for the chatbot API
type ChatbotRequest struct {
	Contents []struct {
		Role  string `json:"role"`
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
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
	// Marshal the request contents to JSON
    requestJSON, err := json.Marshal(request.Contents)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var result []map[string]interface{}

    // Unmarshal the JSON array into the slice
    if err := json.Unmarshal(requestJSON, &result); err != nil {
        log.Println("Error:", err)
        return
    }

    // log.Println(result)
	payload := map[string]interface{}{
		"contents": result,
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
