/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the game",
	Long: `Long Description Placeholder`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		age, _ := cmd.Flags().GetInt("age")

		if name != "" {
			fmt.Printf("Hello %s!\n", name)
		}

		if age > 0 {
			fmt.Printf("You are %d years old.\n", age)
		}

		class, err := selectClass()
		if err != nil {
			fmt.Printf("Class selection failed: %v\n", err)
			return
		}
		fmt.Printf("You have chosen the %s class.\n", class)
	},
}

func selectClass() (string, error) {
	classes := []string{"Frontend Dev", "Backend Dev", "DevOps"}
	
	prompt := promptui.Select{
		Label: "Select your class",
		Items: classes,
	}

	_, result, err := prompt.Run()
	if err != nil {
		
		return "", err
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("name", "n", "", "Name to greet")
	startCmd.Flags().IntP("age", "a", 1, "Age of yourself")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
