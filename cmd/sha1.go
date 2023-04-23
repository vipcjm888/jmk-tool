/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "sha1 encode",
	Long:  `sha1 encode`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		hash := sha1.Sum([]byte(str))
		hashStr := hex.EncodeToString(hash[:])
		fmt.Println(hashStr)
	},
}

func init() {
	rootCmd.AddCommand(sha1Cmd)
	sha1Cmd.Flags().StringP("str", "S", "", "Operated Str")
	sha1Cmd.MarkFlagRequired("str")
}
