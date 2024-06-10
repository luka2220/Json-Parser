package cmd

import (
	"github.com/luka2220/json-parser/utils"
	"github.com/spf13/cobra"
)

// tokensCmd represents the tokens command
var tokensCmd = &cobra.Command{
	Use:   "tokens [FILE]",
	Short: "Displays the tokens for the JSON file passed.",
	Long: `Pass in a path to a JSON file. Must be a valid path 
to a JSON file for the tokenizer to run. The tokenizer will not 
indicate if the JSON fileis valid or not, you can inspect the
tokens to view if any of the are invalid.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.ExecuteTokenizer(args[0])
	},
}

func init() {
	rootCmd.AddCommand(tokensCmd)
}
