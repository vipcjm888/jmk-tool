package cmd

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "open postman",
	Long:  `open postman`,
	Run: func(cmd *cobra.Command, args []string) {
		operator := exec.Command("/bin/bash", "-c", `~/Postman/Postman`)

		stdout, err := operator.StdoutPipe()
		if err != nil {
			fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
			return
		}

		if err := operator.Start(); err != nil {
			fmt.Println("Error:The command is err,", err)
			return
		}

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
	rootCmd.AddCommand(apiCmd)
}
