package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/grafana/hg-snippets/config"
	"github.com/grafana/hg-snippets/snippet"
	"github.com/spf13/cobra"
)

const (
	column = 40
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all snippets",
	Long:  `Show all snippets`,
	RunE:  list,
}

func list(cmd *cobra.Command, args []string) error {
	var snippets snippet.Snippets
	if err := snippets.Load(); err != nil {
		return err
	}

	col := config.Conf.General.Column
	if col == 0 {
		col = column
	}

	for _, snippet := range snippets.Snippets {
		// TODO - Do we need this?
		//if config.Flag.OneLine {
		//	description := runewidth.FillRight(runewidth.Truncate(snippet.Description, col, "..."), col)
		//	commands := runewidth.Truncate(snippet.Commands, 100-4-col, "...")
		//	make sure multiline command printed as oneline
		//commands = strings.Replace(command, "\n", "\\n", -1)
		//fmt.Fprintf(color.Output, "%s : %s\n",
		//	color.GreenString(description), color.YellowString(command))
		//} else {
		fmt.Fprintf(color.Output, "%12s %s\n",
			color.GreenString("Description:"), snippet.Description)
		// TODO - Do we need this?
		//if strings.Contains(snippet.Commands, "\n") {
		//	lines := strings.Split(snippet.Commands, "\n")
		//	firstLine, restLines := lines[0], lines[1:]
		//	fmt.Fprintf(color.Output, "%12s %s\n",
		//		color.YellowString("    Command:"), firstLine)
		//	for _, line := range restLines {
		//		fmt.Fprintf(color.Output, "%12s %s\n",
		//			" ", line)
		//	}
		//} else {
		//}
		fmt.Fprintf(color.Output, "%12s \n", color.YellowString("Example Commands:"))
		for _, command := range snippet.Commands {
			//fmt.Fprintf(color.Output, "%12s %s\n", color.YellowString("    Commands:"), command)
			fmt.Println("$", command)
		}
		if snippet.Tags != nil {
			tags := ""
			for _, tag := range snippet.Tags {
				tags += fmt.Sprintf(" #%s", tag)
			}
			tags = strings.TrimSpace(tags)
			fmt.Fprintf(color.Output, "%12s %s\n",
				color.CyanString("Tags:"), tags)
		}
		if snippet.Output != "" {
			output := strings.Replace(snippet.Output, "\n", "\n             ", -1)
			fmt.Fprintf(color.Output, "%12s %s\n",
				color.RedString("Output:"), output)
		}
		fmt.Println(strings.Repeat("-", 30))
	}
	//}
	return nil
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&config.Flag.OneLine, "oneline", "", false,
		`Display snippets in one line`)
}
