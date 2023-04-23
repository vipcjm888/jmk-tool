/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "sha256 encode",
	Long:  `sha256 encode`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		hash := sha256.Sum256([]byte(str))
		hashStr := hex.EncodeToString(hash[:])
		fmt.Println(hashStr)
	},
}

func init() {
	rootCmd.AddCommand(sha256Cmd)
	sha256Cmd.Flags().StringP("str", "S", "", "Operated Str")
	sha256Cmd.MarkFlagRequired("str")
}
