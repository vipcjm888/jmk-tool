/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var bdCmd = &cobra.Command{
	Use:   "bd",
	Short: "search in baidu",
	Long:  `search in baidu`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		exec.Command(`xdg-open`, `http://www.baidu.com/s?wd=`+str).Start()
	},
}

func init() {
	rootCmd.AddCommand(bdCmd)
	bdCmd.Flags().StringP("str", "S", "", "Search Str")
}
