package app

import (
	"fmt"
	"strings"
)

// ValidateStoryInput validates the story title and content
func ValidateStoryInput(title, content string) error {
	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("title is required")
	}
	if strings.TrimSpace(content) == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}
