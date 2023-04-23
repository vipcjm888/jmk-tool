/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "md5 encode",
	Long:  `md5 encode`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		hash := md5.Sum([]byte(str))
		hashStr := hex.EncodeToString(hash[:])
		fmt.Println(hashStr)
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)
	md5Cmd.Flags().StringP("str", "S", "", "Operated Str")
	md5Cmd.MarkFlagRequired("str")
}
