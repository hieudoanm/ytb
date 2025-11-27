// Package cmd ...
package cmd

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"youtube-cli/utils/figlet"

	youtube "github.com/kkdai/youtube/v2"
	"github.com/spf13/cobra"
)

func GetYouTubeID(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Normal format: https://www.youtube.com/watch?v=ID
	q := u.Query().Get("v")
	if q != "" {
		return q, nil
	}

	// Short format: https://youtu.be/ID
	if u.Host == "youtu.be" {
		return u.Path[1:], nil
	}

	return "", fmt.Errorf("no video ID found")
}

func GetExtension(mime string) string {
	if mime == "" {
		return "mp4"
	}

	// mime looks like: "video/mp4; codecs=..."
	parts := strings.Split(mime, "/")
	if len(parts) < 2 {
		return "mp4"
	}

	ext := parts[1]
	if i := strings.Index(ext, ";"); i != -1 {
		ext = ext[:i]
	}

	return ext
}

// downloadCmd represents the youtubeDownload command
var downloadCmd = &cobra.Command{
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

		client := youtube.Client{}

		videoID, _ := GetYouTubeID(url)

		video, err := client.GetVideo(videoID)
		if err != nil {
			fmt.Printf("Error getting video info: %v\n", err)
			return
		}

		// Choose a format to download (e.g., best quality)
		formats := video.Formats.WithAudioChannels() // Filter for formats with audio
		if len(formats) == 0 {
			fmt.Println("No formats with audio found.")
			return
		}
		format := formats[0] // Select the first available format (often highest quality)

		fmt.Printf("Downloading %s in %s format...\n", video.Title, format.QualityLabel)

		stream, _, err := client.GetStream(video, &format)
		if err != nil {
			fmt.Printf("Error getting video stream: %v\n", err)
			return
		}
		defer stream.Close()

		ext := GetExtension(format.MimeType)
		file, err := os.Create(video.Title + "." + ext)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			fmt.Printf("Error saving video: %v\n", err)
			return
		}

		fmt.Printf("Downloaded %s successfully!\n", file.Name())
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
