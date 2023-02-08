package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yusianglin11010/cinnox-line-bot/internal/database"
)

func createCollection(collName string) {
	mongo := database.GetMongo()
	mongo.InitLineMessage(collName)
}

var CreateCollCmd = &cobra.Command{
	Use:   "create [collection_name]",
	Short: "Create a collection in MongoDB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createCollection(args[0])
		fmt.Printf("Collection %s created\n", args[0])
	},
}
