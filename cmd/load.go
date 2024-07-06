/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"aya/add"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "he loads already saed versions",
	Long: `this comand he is going to load already saved versions from .aya/versions folder to current directory
	he coping everything that is there into here`,
	Run: func(cmd *cobra.Command, args []string) {
		err := load()
		if err != nil {
			return
		} else {
			fmt.Println("we loaded the version succesfully")
		}
	},
}
var LoadVersion string
var LoadPranch string

func init() {
	rootCmd.AddCommand(loadCmd)

	loadCmd.Flags().StringVarP(&LoadVersion, "version", "v", "", "enter the version to load")
	loadCmd.Flags().StringVarP(&LoadPranch, "pranch", "p", "", "enter the pranch to load from")
}

func load() error {
	ayaDir := ".aya"
	destDir := ".aya/versions/main"
	if _, err := os.Stat(ayaDir); os.IsNotExist(err) {
		return errors.New("project is not initialized. Please run 'aya init' then save versions to load")
	}
	filePath := ".aya/refs/aya.json"
	// Read the JSON file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Unmarshal the JSON into a struct
	var vc add.ProjectInfo
	err = json.Unmarshal(file, &vc)
	if err != nil {
		return err
	}
	// Validate the version string to be numbers only
	if LoadVersion != "" {
		matched, err := regexp.MatchString(`^\d+(\.\d+)*$`, LoadVersion)
		if err != nil {
			return fmt.Errorf("error validating version: %v", err)
		}
		if !matched {
			return errors.New("version name must be in the format x.y.z with numbers only")
		}
	}
	if LoadPranch != "" {
		destDir = ".aya/versions/" + LoadPranch + "/" + LoadVersion
		add.Last_pranch(LoadPranch)
	} else {
		destDir = add.Check_last_pranch(".aya-config") + "/" + LoadVersion
	}
	if LoadVersion == "" && LoadPranch == "" {
		destDir = vc.LastVersion
	}

	if LoadVersion == "" {
		destDir = vc.LastVersion
	}
	fmt.Println(destDir)
	var answer string
	fmt.Println("are you sure to load(y,n):",LoadVersion)
	fmt.Scan(&answer)
	if answer == "y" || answer == "yes"{
		err = cleanAndCopy(destDir, ".")
		if err != nil {
			return err
		}
	}else {
		fmt.Println("you canceled loading")
		return nil
	}
	
	return nil
}
func cleanAndCopy(destDir, targetDir string) error {
	// Delete contents of targetDir
	err := os.RemoveAll(targetDir)
	if err != nil {
		return err
	}

	// Recreate the targetDir
	err = os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Copy contents of destDir to targetDir
	return copyDir(destDir, targetDir)
}

func copyDir(srcDir, destDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return copyFile(path, destPath)
	})
}

func copyFile(srcFile, destFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	srcFileInfo, err := src.Stat()
	if err != nil {
		return err
	}

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return os.Chmod(destFile, srcFileInfo.Mode())
}
