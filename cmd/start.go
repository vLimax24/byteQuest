package cmd

import (
	"byteQuest/models"
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the game",
	Long:  `Long Description Placeholder`,
	Run: func(cmd *cobra.Command, args []string) {
		name, errName := inputName()
		class, errClass := selectClass()
		pin, errPIN := createPIN()
		if errPIN != nil {
			fmt.Printf("PIN Fehlgeschlagen: %v\n", errPIN)
			return
		}
		if errClass != nil {
			fmt.Printf("Klassenwahl Fehlgeschlagen: %v\n", errClass)
			return
		}
		if errName != nil {
			fmt.Printf("Namenseingabe Fehlgeschlagen: %v\n", errName)
			return
		}
		fmt.Printf("Hallo %s!\n", name)
		fmt.Printf("Du hast dir die %s Rolle ausgesucht.\n", class)

		// Add player to the database
		db, err := gorm.Open(sqlite.Open("players.db"), &gorm.Config{})
		if err != nil {
			fmt.Printf("Datenbankverbindung fehlgeschlagen: %v\n", err)
			return
		}

		// Add this line to auto-migrate the database
		err = db.AutoMigrate(&models.Player{})
		if err != nil {
			fmt.Printf("Datenbankmigrierung fehlgeschlagen: %v\n", err)
			return
		}

		player := models.Player{
			Name:       name,
			Class:      class,
			PIN:        pin,
			Level:      1,
			Experience: 0,
			Bytes:      0,
		}

		result := db.Create(&player)
		if result.Error != nil {
			fmt.Printf("Spieler konnte nicht zur Datenbank hinzugefügt werden: %v\n", result.Error)
			return
		}

		fmt.Println("Spieler erfolgreich zur Datenbank hinzugefügt!")
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
		Label: "Bitte erstelle einen Benutzernamen",
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

func createPIN() (int, error) {
	validate := func(input string) error {
		if len(input) != 4 {
			return errors.New("PIN muss genau 4 Ziffern lang sein")
		}
		for _, char := range input {
			if char < '0' || char > '9' {
				return errors.New("PIN darf nur Ziffern enthalten")
			}
		}
		return nil
	}

	inputPrompt := promptui.Prompt{
		Label:    "Bitte gib eine 4-stellige PIN ein (Damit du später deinen Account wieder finden kannst)",
		Mask:     '*',
		Validate: validate,
	}

	result, err := inputPrompt.Run()
	if err != nil {
		return 0, err
	}

	pin, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("ungültige PIN: %v", err)
	}

	return pin, nil
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
