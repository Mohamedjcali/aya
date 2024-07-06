package add

/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com

*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Version_info struct {
	VersionName  string    `json:"Name"`
	DateCreated  time.Time `json:"dateCreated"`
	Pranch       string    `json:"Pranch"`
	VersionInfo  string    `json:"description"`
	NewFiles     []string  `json:"NewFiles"`
	RemovedFiles []string  `json:"RemovedFiles"`
	NewLines     int       `json:"NewLines"`
	RemovedLines int       `json:"RemovedLines"`
	ChangedLines int       `json:"ChangedLines"`
}
type ProjectInfo struct {
	ProjectName   string    `json:"Name"`
	DateCreated   time.Time `json:"dateCreated"`
	LastVersion   string    `json:"lastVersion"`
	NewestVersion string    `json:"newestVersion"`
	LastUpdated   time.Time `json:"lastUpdated"`
	BasicInfo     string    `json:"description"`
}
func version_info(version DiffResult, version_name string, pranch string, info string) error {
	version_data := Version_info{
		VersionName:  version_name,
		DateCreated:  time.Now(),
		Pranch:       pranch,
		VersionInfo:  info,
		NewFiles:     version.NewFiles,
		RemovedFiles: version.RemovedFiles,
		NewLines:     version.NewLines,
		RemovedLines: version.RemovedLines,
		ChangedLines: version.ChangedLines,
	}
	folderPath := (".aya/refs/" + pranch)
	projectJSON, err := json.MarshalIndent(version_data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %w", err)
	}

	// Create the file path
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	// Join the folder path and version name to create the file path
	filePath := filepath.Join(folderPath, version_name+".json")

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
func UpdateVersion(newestVersion string) error {
	filePath := ".aya/refs/aya.json"
	// Read the JSON file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the JSON into a struct
	var vc ProjectInfo
	err = json.Unmarshal(file, &vc)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Update the fields
	vc.LastVersion = vc.NewestVersion
	vc.NewestVersion = newestVersion

	// Marshal the struct back to JSON
	updatedFile, err := json.MarshalIndent(vc, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write the updated JSON back to the file
	err = ioutil.WriteFile(filePath, updatedFile, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
