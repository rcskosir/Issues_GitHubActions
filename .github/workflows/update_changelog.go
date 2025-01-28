package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Step 4: Find the first occurrence of the header and insert the string under it
	headerFound := false
	var newLines []string

	for _, line := range lines {
		newLines = append(newLines, line)

		// If the header is found, insert the cleaned string after it
		if !headerFound && strings.Contains(line, header) {
			newLines = append(newLines, fmt.Sprintf("%s\n", cleanedString))
			headerFound = true
		}
	}

	// If the header is not found, append the entry at the end of the file
	if !headerFound {
		fmt.Printf("Warning: '%s' not found in the changelog. Appending the entry at the top.\n", header)
		newLines = append([]string{
			"\n### " + header + "\n",
			cleanedString + "\n",
		}, newLines...)
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
		_, err := writer.WriteString(line)
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
