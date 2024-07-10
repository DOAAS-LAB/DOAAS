/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package check

import (
	
	
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the versions of tools installed",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	
}
