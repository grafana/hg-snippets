package cmd

import (
	"fmt"
	"strings"

	"github.com/knqyf263/pet/snippet"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var lookupCmd = &cobra.Command{
	Use:     "lookup",
	Short:   "Look up a snippet by description, or a description for a snippet.",
	Long:    "Look up a snippet by description or the reverse. \nIf you provide a description, lookup returns the first example for that snippet. \nThis command is primarily used by this CLI's own fuzzy search tool and may be flaky otherwise.",
	Example: "$pet lookup '[get instance details]'\n$pet lookup '$my-command'",
	RunE:    lookup,
}

func lookup(cmd *cobra.Command, args []string) (err error) {
	input := strings.Join(args, " ")
	if strings.TrimSpace(input) == "" {
		return nil
	}

	// in case this is a command, get rid of any leading characters
	if strings.Index(input, "$ ") == 0 || strings.Index(input, "% ") == 0 {
		input = input[2:]
	}

	// if this is a description, get rid of any brackets
	if input[0] == '[' && strings.Index(input, "]") > 0 {
		input = input[1:strings.LastIndex(input, "]")]
	}

	var snippets snippet.Snippets
	if err := snippets.Load(); err != nil {
		return err
	}

	for _, s := range snippets.Snippets {
		// loop through the descriptions
		if s.Description == input {
			fmt.Printf("Default command: `%s`\n", s.Commands[0])
			return nil
		}

		// if it wasn't the description, check the subcommands
		for _, c := range s.Commands {
			if c == input {
				fmt.Printf("Description: \"%s\"\n", s.Description)
				return nil
			}
		}

	}

	return fmt.Errorf("not found")
}

func init() {
	RootCmd.AddCommand(lookupCmd)
}
