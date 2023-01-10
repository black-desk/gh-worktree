package worktree

import (
	"github.com/cli/go-gh"
	"github.com/black-desk/gh-worktree/internal/worktree/commands/pr"
	"github.com/spf13/cobra"
)

func New() (*cobra.Command, error) {
	rest, err := gh.RESTClient(nil)
	if err != nil {
		return nil, err
	}

	var rootCmd = &cobra.Command{
		Use:   "worktree",
		Short: "Git worktrees, dawg",
		Long:  "commands to create and manage git worktrees",
	}

	rootCmd.AddCommand(pr.New(rest))
	return rootCmd, nil
}
