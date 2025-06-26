package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ChangelogEntry represents a single entry from the input file.
type ChangelogEntry struct {
	Type    string // e.g., "BUG FIXES:", "ENHANCEMENTS:", "FEATURES:"
	Content string // The actual text of the entry
}

// readEntryAndPRFromFile reads and parses a single entry and the PR number from the specified text file.
// It expects the file to contain a single line with the PR tag [GH-XXXXX] at the beginning,
// followed by the actual changelog entry.
func readEntryAndPRFromFile(filePath string) ([]ChangelogEntry, string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0o644)
	if err != nil {
		return nil, "", fmt.Errorf("error opening entries file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var entryType, trimmedContent string

		if strings.HasPrefix(line, "[BUG]") {
			entryType = "BUG FIXES:"
			trimmedContent = strings.TrimPrefix(line, "[BUG]")
		} else if strings.HasPrefix(line, "[ENHANCEMENT]") {
			entryType = "ENHANCEMENTS:"
			trimmedContent = strings.TrimPrefix(line, "[ENHANCEMENT]")
		} else if strings.HasPrefix(line, "[FEATURE]") {
			entryType = "FEATURES:"
			trimmedContent = strings.TrimPrefix(line, "[FEATURE]")
		} else {
			// Skip lines that don't match the expected format
			fmt.Printf("Skipping line with invalid format in entries file: %s\n", line)
			continue
		}
		entries = append(entries, ChangelogEntry{Type: entryType, Content: strings.TrimSpace(trimmedContent)})
	}

	if err := scanner.Err(); err != nil {
		return nil, "", fmt.Errorf("error reading entries file: %w", err)
	}

	// Regex to find [GH-XXXXX] at the beginning of the line
	re := regexp.MustCompile(`^\[GH-\d+\]`)
	prMatch := re.FindString(fullEntryLine)

	var prNumber string
	var entryWithoutPR string

	if prMatch != "" {
		prNumber = prMatch
		// Remove the PR part from the beginning of the entry line
		entryWithoutPR = strings.TrimSpace(strings.TrimPrefix(fullEntryLine, prMatch))
	} else {
		// If no PR number found at the beginning, the whole line is the entry, and PR number is empty
		entryWithoutPR = fullEntryLine
		prNumber = ""
	}

	var entryType, trimmedContent string
	if strings.HasPrefix(entryWithoutPR, "[BUG]") {
		entryType = "BUG FIXES:"
		trimmedContent = strings.TrimPrefix(entryWithoutPR, "[BUG]")
	} else if strings.HasPrefix(entryWithoutPR, "[ENHANCEMENT]") {
		entryType = "ENHANCEMENTS:"
		trimmedContent = strings.TrimPrefix(entryWithoutPR, "[ENHANCEMENT]")
	} else if strings.HasPrefix(entryWithoutPR, "[FEATURE]") {
		entryType = "FEATURES:"
		trimmedContent = strings.TrimPrefix(entryWithoutPR, "[FEATURE]")
	} else {
		// If the entry format (e.g., [BUG]) is also missing, treat the whole line as content, and return error for invalid format.
		return nil, "", fmt.Errorf("invalid entry format: entry must start with '[BUG]', '[ENHANCEMENT]', or '[FEATURE]' after optional PR tag")
	}

	// Return a slice with just one entry, as the changelog_entry.txt workflow currently generates a single line.
	// The content here is just the trimmed message, the PR will be appended later.
	return []ChangelogEntry{{Type: entryType, Content: strings.TrimSpace(trimmedContent)}}, prNumber, nil
}

// updateChangelog processes the changelog file based on new entries and appends the PR number.
func updateChangelog(changelogFilePath string, newEntries []ChangelogEntry, prNumber string) error {
	// Read current changelog content
	changelogFile, err := os.OpenFile(changelogFilePath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("error opening changelog file: %w", err)
	}
	defer changelogFile.Close()

	var currentChangelogLines []string
	scanner := bufio.NewScanner(changelogFile)
	for scanner.Scan() {
		currentChangelogLines = append(currentChangelogLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading changelog file: %w", err)
	}

	// Create a map to group new entries by their type
	entriesToInsert := make(map[string][]string)
	for _, entry := range newEntries {
		// Append the PR number to the content if available
		contentWithPR := entry.Content
		if prNumber != "" {
			contentWithPR = fmt.Sprintf("%s %s", entry.Content, prNumber)
		}
		entriesToInsert[entry.Type] = append(entriesToInsert[entry.Type], contentWithPR)
	}

	var updatedLines []string
	insertedHeaders := make(map[string]bool) // To track if a header's entries have been inserted

	for i := 0; i < len(currentChangelogLines); i++ {
		line := currentChangelogLines[i]
		updatedLines = append(updatedLines, line) // Add the current line

		// Check if the current line is a header we care about
		for header := range entriesToInsert {
			if strings.HasPrefix(line, header) && !insertedHeaders[header] {
				// Insert all entries for this header immediately after the header line
				for _, entryContent := range entriesToInsert[header] {
					updatedLines = append(updatedLines, entryContent)
				}
				insertedHeaders[header] = true // Mark as inserted
				break                          // Only one header match per line
			}
		}
	}

	// Handle headers that were not found in the original changelog (prepend them)
	for header, entries := range entriesToInsert {
		if !insertedHeaders[header] {
			// Prepend the header and its entries to the beginning of the file content
			prependLines := []string{"", fmt.Sprintf("### %s", header)} // Add empty line and then formatted header
			for _, entryContent := range entries {
				prependLines = append(prependLines, entryContent)
			}
			updatedLines = append(prependLines, updatedLines...)
			insertedHeaders[header] = true
		}
	}

	// Overwrite the changelog file with updated content
	changelogFile.Seek(0, 0)    // Reset file pointer to beginning
	changelogFile.Truncate(0) // Clear existing content

	writer := bufio.NewWriter(changelogFile)
	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to changelog file: %w", err)
		}
	}
	return writer.Flush()
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run update_changelog.go CHANGELOG.md <new_entries_file.txt>")
		os.Exit(1)
	}

	changelogFilePath := os.Args[1]
	entriesFilePath := os.Args[2]

	newEntries, prNumber, err := readEntryAndPRFromFile(entriesFilePath)
	if err != nil {
		fmt.Println("Error reading new entries:", err)
		os.Exit(1)
	}

	if len(newEntries) == 0 {
		fmt.Println("No valid entries found in the entries file. No updates made to CHANGELOG.md.")
		return
	}

	if err := updateChangelog(changelogFilePath, newEntries, prNumber); err != nil {
		fmt.Println("Error updating changelog:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully updated CHANGELOG.md with entries from", entriesFilePath)
}