package innit

/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com
*/
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new version control system",
	Long:  `This command initializes a new version control system by creating a .aya directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		initProject()
	},
}
var name string
var BasicInfo string

func initProject() {

	// Define the .aya directory
	ayaDir := ".aya"
	fmt.Println("please enter the project name: ")
	fmt.Scan(&name)

	// Clear the buffer
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Println("please enter the project basic info:")
	BasicInfo, _ = reader.ReadString('\n')
	BasicInfo = BasicInfo[:len(BasicInfo)-1] // Remove the newline character

	fmt.Println("Project Name:", name)
	fmt.Println("Project Basic Info:", BasicInfo)
	// Check if the .aya directory already exists
	if _, err := os.Stat(ayaDir); !os.IsNotExist(err) {
		fmt.Println("Project is already initialized.")
		return
	}
	// Create the .aya directory
	err := os.Mkdir(ayaDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create directory %s: %v\n", ayaDir, err)
		return
	}

	// Make the .aya directory hidden on Windows
	if runtime.GOOS == "windows" {
		err := exec.Command("attrib", "+H", ayaDir).Run()
		if err != nil {
			fmt.Printf("Failed to hide directory %s: %v\n", ayaDir, err)
			return
		}
	}

	// Create necessary files and directories inside the .aya directory
	files := []string{
		"HEAD",
		"config",
		"objects",
		"refs",
		"versions",
	}

	for _, file := range files {
		path := filepath.Join(ayaDir, file)
		var fileMode os.FileMode
		if file == "objects" || file == "refs" || file == "versions" {
			err = os.Mkdir(path, 0755)
		} else {
			fileMode = 0644
			_, err = os.Create(path)
		}
		if err != nil {
			fmt.Printf("Failed to create %s: %v\n", path, err)
			return
		}
		if fileMode != 0 {
			err = os.Chmod(path, fileMode)
			if err != nil {
				fmt.Printf("Failed to set permissions for %s: %v\n", path, err)
			}
		}
	}
	err = setup(name, BasicInfo, ".aya/refs")
	if err != nil {
		fmt.Println("Error creating project JSON:", err)
	}
	// Create .ayaignore file
	create_ignore_file()
	// Create .aya-config file
	create_config_file()
	fmt.Println("Initialized empty Aya repository in", ayaDir)
}
func create_ignore_file() {
	ayaIgnoreFile := filepath.Join(".", ".ayaignore")

	// Create the .ayaignore file
	file, err := os.Create(ayaIgnoreFile)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n", ayaIgnoreFile, err)
		return
	}
	defer file.Close()

	// Content to be written into the .ayaignore file
	content := `
	.aya
	go.mod
	go.sum
	.ayaignore
	.ayaconfig
	node_modules`

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Failed to write to %s: %v\n", ayaIgnoreFile, err)
		return
	}
}
func create_config_file() {
	ayaconfigFile := filepath.Join(".", ".aya-config")

	// Create the .aya-config file
	file, err := os.Create(ayaconfigFile)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n", ayaconfigFile, err)
		return
	}
	defer file.Close()

	// Content to be written into the .aya-config file
	content := `main`

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Failed to write to %s: %v\n", ayaconfigFile, err)
		return
	}
}
