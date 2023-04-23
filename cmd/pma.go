/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/spf13/cobra"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// pmaCmd represents the pma command
var pmaCmd = &cobra.Command{
	Use:   "pma",
	Short: "open phpMyAdmin",
	Long:  `open phpMyAdmin`,
	Run: func(cmd *cobra.Command, args []string) {
		operator := exec.Command("/bin/bash", "-c", `/opt/lampp/lampp startapache`)

		//创建获取命令输出管道
		stdout, err := operator.StdoutPipe()
		if err != nil {
			fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
			return
		}

		//执行命令
		if err := operator.Start(); err != nil {
			fmt.Println("Error:The command is err,", err)
			return
		}

		//读取所有输出
		bytes, err := io.ReadAll(stdout)
		if err != nil {
			fmt.Println("ReadAll Stdout:", err.Error())
			return
		}

		if err := operator.Wait(); err != nil {
			fmt.Println("wait:", err.Error())
			return
		}
		fmt.Printf("stdout:\n\n %s", bytes)

	},
}

func init() {
	rootCmd.AddCommand(pmaCmd)
}
