// Copyright (C) slackor. 2024-present.
//
// Created at 2024-05-12, by liasica

package main

import (
	"github.com/spf13/cobra"

	"github.com/liasica/slackor"
	"github.com/liasica/slackor/cmd/slackor/commands"
)

func main() {
	var (
		token string
	)

	cmd := &cobra.Command{
		Use:               "slackor <command> [args]",
		Short:             "Slack tools",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			slackor.SetToken(token)
		},
	}

	cmd.PersistentFlags().StringVar(&token, "token", "", "Slack token")
	_ = cmd.MarkFlagRequired("token")

	cmd.AddCommand(
		commands.Upload(),
	)

	_ = cmd.Execute()
}
