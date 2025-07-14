package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("CHANNEL_ID", "")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelId := os.Getenv("CHANNEL_ID")

	fileArr := []string{"siemenscover.pdf"}

	for _, filePath := range fileArr {
		file, err := os.Open(filePath)
		fmt.Println("UploadingFile", filePath)
		if err != nil {
			fmt.Println("Error opening file %s: %v\n", filePath, err)
		}
		defer file.Close()

		params := slack.UploadFileV2Parameters{
			Channel:  channelId,
			File:     filePath,
			Reader:   file,
			Filename: "demo",
			FileSize: 90,
		}
		UploadFileV2, err := api.UploadFileV2(params)
		// fmt.Println(":::::::::::::::::::::", UploadFileV2.ID)
		if err != nil {
			fmt.Printf("::::::::::::::%s\n", err)
			return
		}
		fmt.Printf("ID : %s, Title : %s\n", UploadFileV2.ID, UploadFileV2.Title)
	}
}
