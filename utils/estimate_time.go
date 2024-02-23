package utils

import (
	"math"
	"strings"
)

// EstimateReadingTime calculates the estimated reading time for the given content
func EstimateReadingTime(content string) int {
	// Average reading speed in words per minute
	averageReadingSpeed := 200

	// Count the number of words in the content
	words := strings.Fields(content)
	numWords := len(words)

	// Calculate the estimated reading time
	readingTime := float64(numWords) / float64(averageReadingSpeed)

	// Round the reading time to the nearest whole number
	roundedReadingTime := int(math.Round(readingTime))

	return roundedReadingTime
}
