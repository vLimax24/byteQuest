package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the game",
	Long: `Long Description Placeholder`,
	Run: func(cmd *cobra.Command, args []string) {
		class, errClass := selectClass()
		name, errName := inputName()
		if errClass != nil {
			fmt.Printf("Klassenwahl Fehlgeschlagen: %v\n", errClass)
			return
		}
		if errName != nil {
			fmt.Printf("Namenseingabe Fehlgeschlagen: %v\n", errName)
		}
		fmt.Printf("Hallo %s!\n", name)
		fmt.Printf("Du hast dir die %s Rolle ausgesucht.\n", class)
	},
}

func selectClass() (string, error) {
	classes := []string{"Frontend Dev", "Backend Dev", "DevOps"}
	
	prompt := promptui.Select{
		Label: "Suche dir eine Rolle heraus",
		Items: classes,
	}

	_, result, err := prompt.Run()
	if err != nil {
		
		return "", err
	}

	return result, nil
}

func inputName() (string, error) {
	inputPrompt := promptui.Prompt{
		Label: "Bitte gebe deinen Namen ein",
	}

	result, err := inputPrompt.Run()
	if err != nil {
		return "", err
	}

	confirmPrompt := promptui.Prompt{
		Label:     fmt.Sprintf("Ist '%s' korrekt?", result),
		IsConfirm: true,
		Default:   "y",
	}

	_, err = confirmPrompt.Run()
	if err != nil {
		// If the error is not ErrAbort, it means the user confirmed (either by 'y' or just pressing enter)
		if err != promptui.ErrAbort {
			return result, nil
		}
		// If it's ErrAbort, ask for the name again
		return inputName()
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
