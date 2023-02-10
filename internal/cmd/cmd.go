package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yusianglin11010/cinnox-line-bot/internal/config"
	"github.com/yusianglin11010/cinnox-line-bot/internal/database"
)

var CreateCollCmd = &cobra.Command{
	Use:   "create [collection_name]",
	Short: "Create a collection in MongoDB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mongo := database.GetMongo()
		err := mongo.InitLineMessage(args[0])
		if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Printf("Collection %s created\n", args[0])
		}
	},
}

var CreateConfig = &cobra.Command{
	Use:   "config",
	Short: "Create a config from environment variables",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		err := config.WriteConfig()
		if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Printf("Config created\n")
		}

	},
}
