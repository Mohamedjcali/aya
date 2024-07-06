package add
/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com
*/
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var Version string
var Writer string
var Description string
var Pranch string

func AddVersion() error {
	// Check if the project is initialized by looking for the .aya directory
	ayaDir := ".aya"
	destDir := ".aya/versions/main"
	// Validate the version string to be numbers only
	matched, err := regexp.MatchString(`^\d+(\.\d+)*$`, Version)
	if err != nil {
		return fmt.Errorf("error validating version: %v", err)
	}
	if !matched {
		return errors.New("version name must be in the format x.y.z with numbers only")
	}
	if _, err := os.Stat(ayaDir); os.IsNotExist(err) {
		return errors.New("project is not initialized. Please run 'aya init' first")
	}

	if Pranch != "" {
		destDir = ".aya/versions/" + Pranch
		Last_pranch(Pranch)
	} else {
		destDir = Check_last_pranch(".aya-config")
	}
	err = CreateAndCopyFolder(".", destDir, Version)
	if err != nil {
		return fmt.Errorf("error saving the commit version: %v", err)
	}
	err = UpdateVersion((destDir + "/" + Version))
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	diffresult, err := CompareVersions()
	if err != nil {
		return fmt.Errorf("error for comparing the versions: %v", err)
	}
	err = version_info(diffresult, Version, Pranch, Description)
	if err != nil {
		return fmt.Errorf("error for saving version data: %v", err)
	}
	return nil
}

func Check_last_pranch(filePath string) string {
	defaultDir := ".aya/versions/main"
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultDir
		}
		return defaultDir
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return ".aya/versions/" + strings.TrimSpace(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// handle the error here if needed
		return defaultDir
	}

	return defaultDir
}
func Last_pranch(content string) error {
	fileName := ".aya-config"

	// Open the file in read-write mode
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file content
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Replace the first line
	if len(lines) > 0 {
		lines[0] = content
	} else {
		// If the file is empty, just add the content as the first line
		lines = append(lines, content)
	}

	// Move the file pointer to the beginning
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	// Truncate the file to zero length
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	// Write the updated content back to the file
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush to file: %w", err)
	}
	return nil
}
