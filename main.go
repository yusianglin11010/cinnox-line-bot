package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yusianglin11010/cinnox-line-bot/internal/cmd"
	"github.com/yusianglin11010/cinnox-line-bot/internal/config"
	"github.com/yusianglin11010/cinnox-line-bot/internal/database"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Hello Cinnox!")
	logger, _ := zap.NewProduction()
	dbConfig := config.NewMongoConfig(logger)

	// initialize mongo DB
	database.Initialize(dbConfig)
	defer database.Close()

	rootCmd := &cobra.Command{Use: "./app-name"}
	rootCmd.AddCommand(cmd.CreateCollCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
