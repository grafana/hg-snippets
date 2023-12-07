package cmd

import (
	"github.com/grafana/hg-snippets/config"
	petSync "github.com/grafana/hg-snippets/sync"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync snippets",
	Long:  `Sync snippets with gist/gitlab`,
	RunE:  sync,
}

func sync(cmd *cobra.Command, args []string) (err error) {
	return petSync.AutoSync(config.Conf.General.SnippetFile)
}

// TODO figure out how we want to use the sync command. Only for Gists?
func init() {
	//RootCmd.AddCommand(syncCmd)
}
