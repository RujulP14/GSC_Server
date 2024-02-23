// utils/gemini_utils.go

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// GeminiAPIURL is the base URL for the Gemini AI API
const GeminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key="

// GetGeminiAPIKey retrieves the API key from the environment variables
func GetGeminiAPIKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	return apiKey
}

// MakeGeminiAPIRequest makes a POST request to the Gemini AI API
func MakeGeminiAPIRequest(apiKey string, payload map[string]interface{}) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create JSON payload: %v", err)
	}

	url := fmt.Sprintf("%s%s", GeminiAPIURL, apiKey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return responseBody, nil
}

// ParseGeminiAPIResponse parses the response JSON from the Gemini AI API
func ParseGeminiAPIResponse(responseBody []byte) (string, error) {
	var chatbotResponse struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(responseBody, &chatbotResponse); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %v", err)
	}

	var generatedText string
	if len(chatbotResponse.Candidates) > 0 {
		for _, part := range chatbotResponse.Candidates[0].Content.Parts {
			generatedText += part.Text
		}
	}

	return generatedText, nil
}
