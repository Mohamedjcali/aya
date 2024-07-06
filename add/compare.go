package add

/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com

*/
import (
	"bufio"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DiffResult struct {
	NewFiles     []string
	RemovedFiles []string
	NewLines     int
	RemovedLines int
	ChangedLines int
}

func CompareVersions() (DiffResult, error) {
	var err error
	filePath := ".aya/refs/aya.json"
	// Read the JSON file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return DiffResult{}, err
	}

	// Unmarshal the JSON into a struct
	var vc ProjectInfo
	err = json.Unmarshal(file, &vc)
	if err != nil {
		return DiffResult{}, err
	}
	oldDir := vc.LastVersion
	newDir := vc.NewestVersion
	var oldFiles, newFiles []string

	if oldDir != "" {
		oldFiles, err = listFiles(oldDir)
		if err != nil {
			return DiffResult{}, err
		}
	}

	if newDir != "" {
		newFiles, err = listFiles(newDir)
		if err != nil {
			return DiffResult{}, err
		}
	}

	newFilesSet := make(map[string]struct{})
	for _, file := range newFiles {
		newFilesSet[file] = struct{}{}
	}

	oldFilesSet := make(map[string]struct{})
	for _, file := range oldFiles {
		oldFilesSet[file] = struct{}{}
	}

	var diff DiffResult

	// Identify new and removed files
	for _, file := range newFiles {
		if _, exists := oldFilesSet[file]; !exists {
			diff.NewFiles = append(diff.NewFiles, file)
			newLineCount, err := countLines(filepath.Join(newDir, file))
			if err != nil {
				return DiffResult{}, err
			}
			diff.NewLines += newLineCount
		}
	}

	for _, file := range oldFiles {
		if _, exists := newFilesSet[file]; !exists {
			diff.RemovedFiles = append(diff.RemovedFiles, file)
			removedLineCount, err := countLines(filepath.Join(oldDir, file))
			if err != nil {
				return DiffResult{}, err
			}
			diff.RemovedLines += removedLineCount
		}
	}

	// Compare file contents for modified files
	for _, file := range newFiles {
		if _, exists := oldFilesSet[file]; exists {
			newLines, removedLines, changedLines, err := compareFiles(filepath.Join(oldDir, file), filepath.Join(newDir, file))
			if err != nil {
				return DiffResult{}, err
			}
			diff.NewLines += newLines
			diff.RemovedLines += removedLines
			diff.ChangedLines += changedLines
		}
	}

	return diff, nil
}

func listFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			files = append(files, relativePath)
		}
		return nil
	})
	return files, err
}

func compareFiles(oldFile, newFile string) (int, int, int, error) {
	oldLines, err := readLines(oldFile)
	if err != nil {
		return 0, 0, 0, err
	}

	newLines, err := readLines(newFile)
	if err != nil {
		return 0, 0, 0, err
	}

	oldLinesSet := make(map[string]int)
	for idx, line := range oldLines {
		oldLinesSet[line] = idx
	}

	newLinesSet := make(map[string]int)
	for idx, line := range newLines {
		newLinesSet[line] = idx
	}

	var newLineCount, removedLineCount, changedLineCount int

	// Count new and removed lines
	for idx, line := range newLines {
		if oldIdx, exists := oldLinesSet[line]; !exists {
			newLineCount++
		} else if oldIdx != idx {
			changedLineCount++
		}
	}

	for line := range oldLinesSet {
		if _, exists := newLinesSet[line]; !exists {
			removedLineCount++
		}
	}

	return newLineCount, removedLineCount, changedLineCount, nil
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func countLines(filePath string) (int, error) {
	lines, err := readLines(filePath)
	if err != nil {
		return 0, err
	}
	return len(lines), nil
}
