package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgo",
	Short: "imgo offers a variety of image processing capabilities",
	Long:  `asd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from cobra")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
