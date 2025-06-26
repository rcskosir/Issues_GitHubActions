package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ChangelogEntry represents a single entry from the input file.
type ChangelogEntry struct {
	Type    string // e.g., "BUG FIXES:", "ENHANCEMENTS:", "FEATURES:"
	Content string // The actual text of the entry
}

// readEntriesFromFile reads and parses entries from the specified text file.
func readEntriesFromFile(filePath string) ([]ChangelogEntry, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("error opening entries file: %w", err)
	}
	defer file.Close()

	var entries []ChangelogEntry
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
			fmt.Printf("Warning: Skipping line with invalid format in entries file: %s\n", line)
			continue
		}
		entries = append(entries, ChangelogEntry{Type: entryType, Content: strings.TrimSpace(trimmedContent)})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading entries file: %w", err)
	}
	return entries, nil
}

// updateChangelog processes the changelog file based on new entries.
func updateChangelog(changelogFilePath string, newEntries []ChangelogEntry) error {
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
		entriesToInsert[entry.Type] = append(entriesToInsert[entry.Type], entry.Content)
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
				break // Only one header match per line
			}
		}
	}

	// Handle headers that were not found in the original changelog, similar to Python's behavior (prepend)
	for header, entries := range entriesToInsert {
		if !insertedHeaders[header] {
			// Prepend the header and its entries to the beginning of the file content
			// This might not be ideal for all changelog formats, but matches the Python script's logic
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

	newEntries, err := readEntriesFromFile(entriesFilePath)
	if err != nil {
		fmt.Println("Error reading new entries:", err)
		os.Exit(1)
	}

	if len(newEntries) == 0 {
		fmt.Println("No valid entries found in the entries file. No updates made to CHANGELOG.md.")
		return
	}

	if err := updateChangelog(changelogFilePath, newEntries); err != nil {
		fmt.Println("Error updating changelog:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully updated CHANGELOG.md with entries from", entriesFilePath)
}