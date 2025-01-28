package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func updateChangelog(inputString string, changelogFile string) {
	// Step 1: Identify the type of entry and clean the input string
	re := regexp.MustCompile(`^\[(BUG|ENHANCEMENT|FEATURE)\](.*)`)
	match := re.FindStringSubmatch(strings.TrimSpace(inputString))
	if match == nil {
		fmt.Println("Error: Input string must start with '[BUG]' or '[ENHANCEMENT]' or '[FEATURE]'.")
		return
	}

	changeType := match[1]
	cleanedString := strings.TrimSpace(match[2])

	// Step 2: Determine the appropriate header
	var header string
	switch changeType {
	case "BUG":
		header = "BUG FIXES:"
	case "ENHANCEMENT":
		header = "ENHANCEMENTS:"
	case "FEATURE":
		header = "FEATURES:"
	}

	// Step 3: Read the existing changelog file and look for the header
	file, err := os.Open(changelogFile)
	if err != nil {
		fmt.Println("Error: Unable to open the changelog file.")
		return
	}
	defer file.Close()

	var lines []string
	var sectionStartIdx, sectionEndIdx int
	var inSection bool
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lines = append(lines, line)

		// Check if we have found the header and begin collecting the section
		if !inSection && strings.Contains(line, header) {
			sectionStartIdx = i
			inSection = true
		}
		// If we're in the section, keep track of where the section ends
		if inSection && (strings.TrimSpace(line) == "") && sectionEndIdx == 0 {
			sectionEndIdx = i
			break
		}
	}

	// Step 4: Insert the entry under the header in alphabetical order
	var newLines []string
	headerFound := false
	if sectionStartIdx != 0 && sectionEndIdx != 0 {
		// Found the section, insert the new entry in sorted order
		var sectionEntries []string
		for _, line := range lines[sectionStartIdx+1 : sectionEndIdx] {
			sectionEntries = append(sectionEntries, line)
		}

		// Insert the new entry in the correct position (alphabetically)
		sectionEntries = append(sectionEntries, cleanedString)
		sort.Strings(sectionEntries) // Sort the entries alphabetically

		// Rebuild the lines with the sorted entries
		for i := 0; i < sectionStartIdx; i++ {
			newLines = append(newLines, lines[i])
		}
		newLines = append(newLines, fmt.Sprintf("### %s", header))
		for _, entry := range sectionEntries {
			newLines = append(newLines, entry)
		}
		for i := sectionEndIdx; i < len(lines); i++ {
			newLines = append(newLines, lines[i])
		}
		headerFound = true
	}

	// If the header is not found, add it to the top with the new entry
	if !headerFound {
		fmt.Printf("Warning: '%s' not found in the changelog. Appending the entry at the top.\n", header)
		newLines = append(newLines, fmt.Sprintf("\n### %s\n", header), cleanedString+"\n")
	} else {
		fmt.Println("Changes were added in alphabetical order.")
	}

	// Step 5: Write the updated content back to the file
	fileOut, err := os.Create(changelogFile)
	if err != nil {
		fmt.Println("Error: Unable to create the changelog file.")
		return
	}
	defer fileOut.Close()

	writer := bufio.NewWriter(fileOut)
	for _, line := range newLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to the file.")
			return
		}
	}

	writer.Flush()

	fmt.Printf("The change has been added to the changelog under %s.\n", header)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the input string as a command-line argument.")
		return
	}

	inputString := os.Args[1] // Get the entry from the command line
	updateChangelog(inputString, "CHANGELOG.md")
}
