package innit

/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com

*/
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ProjectInfo struct {
	ProjectName   string    `json:"Name"`
	DateCreated   time.Time `json:"dateCreated"`
	LastVersion   string    `json:"lastVersion"`
	NewestVersion string    `json:"newestVersion"`
	LastUpdated   time.Time `json:"lastUpdated"`
	BasicInfo     string    `json:"description"`
}

func setup(name string, discription string, folderPath string) error {
	project := ProjectInfo{
		ProjectName:   name,
		DateCreated:   time.Now(),
		NewestVersion: "1.0",
		LastVersion:   "0.0",
		LastUpdated:   time.Now(),
		BasicInfo:     discription,
	}
	projectJSON, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %w", err)
	}

	// Create the file path
	filePath := filepath.Join(folderPath, "aya.json")

	// Create the JSON file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Write JSON data to the file
	_, err = file.Write(projectJSON)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
