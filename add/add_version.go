package add
/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com

*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateAndCopyFolder(srcDir, destDir, newFolderName string) error {
	// Clean up paths
	srcDir = filepath.Clean(srcDir)
	destDir = filepath.Clean(destDir)

	// Ensure source directory exists
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return fmt.Errorf("source directory %s does not exist", srcDir)
	}

	// Create target directory path
	targetDir := filepath.Join(destDir, newFolderName)

	// Check if the target directory already exists
	if _, err := os.Stat(targetDir); !os.IsNotExist(err) {
		return fmt.Errorf("folder %s already exists", targetDir)
	}

	// Create the target directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	// Get list of folders to ignore
	ignoreList, err := getIgnoreList(filepath.Join(srcDir, ".ayaignore"))
	if err != nil {
		return err
	}

	// Copy directory contents
	var copyDir func(string, string) error
	copyDir = func(src, dst string) error {
		entries, err := os.ReadDir(src)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if contains(ignoreList, entry.Name()) {
				continue // Ignore listed folders
			}

			srcPath := filepath.Join(src, entry.Name())
			dstPath := filepath.Join(dst, entry.Name())

			if entry.IsDir() {
				if err := os.MkdirAll(dstPath, entry.Type()); err != nil {
					return err
				}
				if err := copyDir(srcPath, dstPath); err != nil {
					return err
				}
			} else {
				srcFile, err := os.Open(srcPath)
				if err != nil {
					return err
				}
				defer srcFile.Close()

				dstFile, err := os.Create(dstPath)
				if err != nil {
					return err
				}
				defer dstFile.Close()

				if _, err := io.Copy(dstFile, srcFile); err != nil {
					return err
				}

				// Copy file permissions
				srcInfo, err := os.Stat(srcPath)
				if err != nil {
					return err
				}
				if err := os.Chmod(dstPath, srcInfo.Mode()); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := copyDir(srcDir, targetDir); err != nil {
		return err
	}

	return nil
}

func getIgnoreList(filePath string) ([]string, error) {
	var ignoreList []string
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return ignoreList, nil // No .ayaignore file, no folders to ignore
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ignoreList = append(ignoreList, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ignoreList, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
