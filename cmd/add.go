package cmd

/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com

*/
import (
	"github.com/spf13/cobra"
	"aya/add"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new version to the version control system",
	Long:  `This command adds a new version to the version control system with a version name, author, and description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return add.AddVersion()
	},
}



func init() {
	rootCmd.AddCommand(addCmd)

	// Define flags
	addCmd.Flags().StringVarP(&add.Version, "version", "v", "", "Version name (numbers only)")
	addCmd.Flags().StringVarP(&add.Writer, "writer", "w", "", "Name of the writer")
	addCmd.Flags().StringVarP(&add.Description, "description", "d", "", "Description of the version")
	addCmd.Flags().StringVarP(&add.Pranch, "pranch", "p", "", "which pranch to save ")

	// Mark flags as required
	addCmd.MarkFlagRequired("version")
	addCmd.MarkFlagRequired("writer")
	addCmd.MarkFlagRequired("description")
}
