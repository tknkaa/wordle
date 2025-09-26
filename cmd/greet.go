/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greets a person by name",
	Long:  `A simple command that takes a name as an argument and prints a friendly greeting.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		isUpperCase, _ := cmd.Flags().GetBool("uppercase")
		name := strings.Join(args, " ")
		greeting := fmt.Sprintf("Hello, %s\n", name)

		if isUpperCase {
			greeting = strings.ToUpper(greeting)
		}

		fmt.Println(greeting)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().BoolP("uppercase", "u", false, "Print the greeting in uppercase")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
