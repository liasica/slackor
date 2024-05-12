// Copyright (C) slackor. 2024-present.
//
// Created at 2024-05-12, by liasica

package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"

	"github.com/liasica/slackor"
)

func Upload() *cobra.Command {
	var (
		channel string
		title   string
		comment string
	)

	cmd := &cobra.Command{
		Use:               "upload <path>",
		Short:             "Upload files to slack",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Args:              cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			p := args[0]
			stat, err := os.Stat(p)
			if err != nil {
				slackor.Exit(fmt.Sprintf("File read error: %s\n", p))
			}

			var f *os.File
			f, err = os.Open(p)
			if err != nil {
				slackor.Exit(fmt.Sprintf("File open error: %s\n", p))
			}

			api := slack.New(slackor.GetToken())
			var summary *slack.FileSummary
			summary, err = api.UploadFileV2(slack.UploadFileV2Parameters{
				File:           p,
				FileSize:       int(stat.Size()),
				Reader:         f,
				Filename:       path.Base(p),
				Title:          title,
				Channel:        channel,
				InitialComment: comment,
			})
			if err != nil {
				slackor.Exit(fmt.Sprintf("Upload file failed: %s\n", err))
			}

			fmt.Printf("Upload file success, file id: %s\n", summary.ID)
		},
	}

	cmd.Flags().StringVar(&channel, "channel", "", "Channel IDs, eg: C12345678,C23456789")
	cmd.Flags().StringVar(&title, "title", "", "Title of file.")
	cmd.Flags().StringVar(&comment, "comment", "", "The message text introducing the file in specified channels.")

	return cmd
}
