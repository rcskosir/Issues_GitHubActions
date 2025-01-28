package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to read the file and parse sections
func readFile(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	sections := make(map[string][]string)
	var currentSection string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Identify the section headers
		if strings.HasPrefix(line, "##") {
			// Set the current section as the version or date
			currentSection = strings.TrimSpace(line)
			sections[currentSection] = []string{}
		} else if line == "FEATURES:" || line == "ENHANCEMENTS:" || line == "BUG FIXES:" {
			// Set the current section as FEATURE, ENHANCEMENTS, or BUG FIXES
			currentSection = line
			if _, exists := sections[currentSection]; !exists {
				sections[currentSection] = []string{}
			}
		} else if currentSection != "" {
			// Append lines under the correct section
			sections[currentSection] = append(sections[currentSection], line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read the file: %v", err)
	}

	return sections, nil
}

// Function to append a new entry in alphabetical order
func appendNewEntry(filePath string, section string, entry string) error {
	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	// Write the section header if it's not already there
	_, err = file.WriteString(fmt.Sprintf("\n%s:\n", section))
	if err != nil {
		return fmt.Errorf("failed to write section header: %v", err)
	}

	// Write the entry under the correct section
	_, err = file.WriteString(fmt.Sprintf("* %s\n", entry))
	if err != nil {
		return fmt.Errorf("failed to write entry: %v", err)
	}

	return nil
}

// Function to categorize the entry based on keywords
func categorizeEntry(entry string) string {
	// Basic categorization based on keywords
	switch {
	case strings.Contains(entry, "[FEATURE]"):
		return "FEATURES"
	case strings.Contains(entry, "[ENHANCEMENT]"):
		return "ENHANCEMENTS"
	case strings.Contains(entry, "[BUG]"):
		return "BUG FIXES"
	}
}

func main() {
	// Ensure there are enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run update_changelog.go \"Your changelog entry here\"")
		return
	}

	// Get the input string from the command-line arguments (skip the first argument)
	input := os.Args[1]

	// Specify the changelog file path
	filePath := "CHANGELOG.md" // Replace this with the path to your changelog file

	// Read the current content of the file into sections
	sections, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Categorize the entry
	section := categorizeEntry(input)

	// Append the new entry to the appropriate section
	err = appendNewEntry(filePath, section, input)
	if err != nil {
		fmt.Println("Error appending new entry:", err)
		return
	}

	fmt.Println("Changelog updated successfully!")
}
