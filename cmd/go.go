/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "go browser",
	Long:  `go browser`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		if str[0:4] != "http" {
			str = "http://" + str
		}
		exec.Command(`xdg-open`, str).Start()
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
	goCmd.Flags().StringP("str", "S", "", "Search Str")
}
