/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sddev12/gotionary/dictionarydefinition"
	"github.com/sddev12/gotionary/getword"
	"github.com/spf13/cobra"
)

var Word string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotionary",
	Short: "Gets the dictionary definition of the provided word",
	Run: func(cmd *cobra.Command, args []string) {
		if Word == "" {
			fmt.Println("Error: Invalid use of gotionary. Use gotionary --help for usage")
			return
		}
		uri := "https://api.dictionaryapi.dev/api/v2/entries/en/" + Word
		body := getword.GetWord(uri)
		definition := []dictionarydefinition.Word{}
		if err := json.Unmarshal(body, &definition); err != nil {
			fmt.Printf("Failed to unmarshal response: %v", err)
		}

		fmt.Println(definition[0].Word + "\n")

		for _, phonetic := range definition[0].Phonetics {
			if len(phonetic.Text) > 0 {
				fmt.Printf("Phonetic: %v \n", phonetic.Text)
			}

			if len(phonetic.Audio) > 0 {
				fmt.Printf("Audio: %v \n", phonetic.Audio)
			}

			fmt.Println("")
		}

		for _, meaning := range definition[0].Meanings {
			if len(meaning.PartOfSpeech) > 0 {
				fmt.Printf("Part of speech: %v \n", meaning.PartOfSpeech)
			}

			if len(meaning.Definitions) > 0 {
				for _, definition := range meaning.Definitions {
					fmt.Printf("Definition: %v \n", definition.Definition)
					fmt.Printf("Example: %v \n \n", definition.Example)
				}
			}

			fmt.Println("")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringVarP(&Word, "word", "w", "", "Word to search")
}
