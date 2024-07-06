package cmd
/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com

*/
import (
	"archive/zip"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [url] [destination]",
	Short: "Copies a GitHub repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		destination := ""

		if len(args) > 1 {
			destination = args[1]
		} else {
			destination = extractRepoName(url)
		}

		if err := cloneRepo(url, destination); err != nil {
			fmt.Printf("Failed to copy repository: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}

func extractRepoName(url string) string {
	parts := strings.Split(url, "/")
	repoName := parts[len(parts)-1]
	repoName = strings.TrimSuffix(repoName, ".git")
	return repoName
}

func cloneRepo(url, destination string) error {
	// Convert GitHub repo URL to zip download URL
	zipURL := convertToZipURL(url)

	// Download the zip file
	zipFilePath, err := downloadZip(zipURL)
	if err != nil {
		return err
	}
	defer os.Remove(zipFilePath)

	// Unzip the file
	if err := unzip(zipFilePath, destination); err != nil {
		return err
	}

	return nil
}

func convertToZipURL(url string) string {
	url = strings.TrimSuffix(url, ".git")
	return url + "/archive/refs/heads/main.zip"
}

func downloadZip(zipURL string) (string, error) {
	// Create the file
	out, err := os.CreateTemp("", "repo-*.zip")
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(zipURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create progress bar
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	// Write the body to file
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, 0755)

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), f.Mode()); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
