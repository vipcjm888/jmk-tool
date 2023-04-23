/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Contact me with email",
	Long:  `Contact me with email`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Email:vipcjm888@vip.qq.com")
	},
}

func init() {
	rootCmd.AddCommand(contactCmd)
}
