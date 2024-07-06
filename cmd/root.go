/*
Copyright © 2024 cqani HERE cqanimohamed4@gmail.com
*/
package cmd

import (
	"aya/innit"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aya",
	Short: "A minimalist version control system with fewer than 6 commands, offering a straightforward alternative to Git.",
	Long: `This is an open-source version control system focused on simplicity. With fewer than 6 commands,
	 		it offers a minimalistic and easy-to-use alternative to Git. Perfect for those who prefer
	  		a straightforward tool without the complexity of numerous commands. Ideal for both beginners and 
	  		experienced developers who want a simpler way to manage their projects`,
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(innit.InitCmd)
}


