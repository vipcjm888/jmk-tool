package cmd

import (
	"fmt"

	"encoding/base64"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 encode & decode",
	Long:  `Base64 encode & decode`,
	Run: func(cmd *cobra.Command, args []string) {
		flag, _ := cmd.Flags().GetInt("flag")
		str, _ := cmd.Flags().GetString("str")
		if flag == 1 {
			res, err := base64.StdEncoding.DecodeString(str)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(res))
			}
		} else {
			fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().IntP("flag", "F", 0, "Is Decode?")
	base64Cmd.Flags().StringP("str", "S", "", "Operated Str")
	base64Cmd.MarkFlagRequired("str")

}
