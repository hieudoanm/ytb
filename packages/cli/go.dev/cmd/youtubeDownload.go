// Package cmd ...
package cmd

import (
	"fmt"
	"youtube-cli/utils/figlet"

	"github.com/spf13/cobra"
)

// youtubeDownloadCmd represents the youtubeDownload command
var youtubeDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()
		// Get URL
		fmt.Print("URL (https://www.youtube.com/watch?v=jNQXAC9IVRw): ")
		var url string
		fmt.Scanln(&url)
		fmt.Println(url)
	},
}

func init() {
	rootCmd.AddCommand(youtubeDownloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// youtubeDownloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// youtubeDownloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
