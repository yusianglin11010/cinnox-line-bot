package main

import (
	"github.com/spf13/cobra"
	"github.com/yusianglin11010/cinnox-line-bot/internal/cmd"
)

func main() {
	rootCmd := &cobra.Command{Use: "./app-name"}
	rootCmd.AddCommand(cmd.CreateCollCmd)
	rootCmd.AddCommand(cmd.CreateConfig)
	rootCmd.Execute()
}
