package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "policy-checker",
		Short: "Policy checker is a CLI tool for checking policies",
		Long:  `Policy checker is a CLI tool for checking policies based on Open Policy Agent`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := Policy(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error()) // PolicyErrorのErrorメソッドを使用
		os.Exit(1)
	}
}
