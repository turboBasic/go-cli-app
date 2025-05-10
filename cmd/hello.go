/*
Copyright © 2025 Andriy Melnyk
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  `Print greeting for provided list of entities.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Hello, world!")
			return
		}
		for _, name := range args {
			fmt.Printf("Hello, %s!\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
