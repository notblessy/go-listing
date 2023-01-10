package console

import (
	"os"

	"github.com/notblessy/go-listing/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-listing",
	Short: "app root",
	Long:  `root for running servers`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	config.LoadENV()
}
