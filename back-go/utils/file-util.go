package utils

import (
	"os"
)

// read file and return content as string
func ReadFile(filename string) (string, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

// save to file 
func SaveFile(filename string, content string) error {
    return os.WriteFile(filename, []byte(content), 0644)
}