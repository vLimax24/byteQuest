/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
	"gorm.io/gorm/logger"
)

func getPlayerData(db *gorm.DB, playerName string) (*models.Player, error) {
    var player models.Player
    result := db.Where("name = ?", playerName).First(&player)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("player not found")
        }
        return nil, fmt.Errorf("database error: %v", result.Error)
    }

    return &player, nil
}

func enterUsername() (string, error) {
	inputPrompt := promptui.Prompt{
		Label: "Bitte gebe deinen Benutzernamen ein",
	}

	result, err := inputPrompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func getDatabase() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(sqlite.Open("players.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database")
	}

	db.AutoMigrate()

	return db, nil
}

func validatePIN(db *gorm.DB, player *models.Player) error {
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
		Label:    "Bitte gib deine PIN ein",
		Mask:     '*',
		Validate: validate,
	}

	result, err := inputPrompt.Run()
	if err != nil {
		return err
	}

	inputPIN, err := strconv.Atoi(result)
	if err != nil {
		return fmt.Errorf("ungültige PIN: %v", err)
	}

	if inputPIN != player.PIN {
		fmt.Println("PIN ist falsch!")
		validatePIN(db, player)
	}

	return nil
}

// continueCmd represents the continue command
var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "Spiele mit einem Account weiter",
	Long: `Long command`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := getDatabase()
		if err != nil {
			fmt.Printf("Datenbankverbindung fehlgeschlagen: %v\n", err)
			return
		}

		name, err := enterUsername()

		if err != nil {
			fmt.Printf("Namenseingabe Fehlgeschlagen: %v\n", err)
			return
		}

		player, err := getPlayerData(db, name)

		if err != nil {
			fmt.Printf("Spieler konnte nicht gefunden werden! \n")
			name, err := enterUsername()

			if err != nil {
				fmt.Printf("Namenseingabe Fehlgeschlagen: %v\n", err)
				return
			}

			player, err := getPlayerData(db, name)

			if err != nil {
				fmt.Printf("Spieler konnte nicht gefunden werden! \n")
				return
			}

			err = validatePIN(db, player)

			if err != nil {
				fmt.Printf("PIN konnte nicht validiert werden: %v\n", err)
				return
			}
		}

		err = validatePIN(db, player)

		if err != nil {
			fmt.Printf("PIN konnte nicht validiert werden: %v\n", err)

			return
			
		}
			
		fmt.Printf("Du bist der %s %s.\n", player.Class, player.Name)
	},
}
func init() {
	rootCmd.AddCommand(continueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// continueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// continueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
