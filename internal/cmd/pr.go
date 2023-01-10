package cmd

import (
	ghworktree "github.com/black-desk/gh-worktree/internal/gh-worktree"
	"github.com/black-desk/lib/go/errwrap"
	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr <pr number>",
	Short: "worktree from PR",
	Long:  "Create a new worktree from a PR number",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := ghworktree.NewRepo()
		if err != nil {
			return errwrap.Trace(err)
		}

		prNumber := args[0]

		pr, err := repo.PullRequest(prNumber)

		if err != nil {
			return errwrap.Trace(err)
		}

		parentDir, err := cmd.Flags().GetString("parent")
		if err != nil {
			return errwrap.Trace(err)
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return errwrap.Trace(err)
		}

		if path == "" {
			path = pr.Reference()
		}

		builder := ghworktree.NewWorkTreeBuilder().With(
			ghworktree.WorkTreeBuilderOption{
				Parent: parentDir,
			},
		)

		return errwrap.Trace(builder.Build(path, pr.Reference()))
	},
}

func init() {
	rootCmd.AddCommand(prCmd)

	prCmd.Flags().StringP("parent", "p", ".", "Specify the parent directory which new worktree will create in.")
	prCmd.Flags().StringP("path", "P", "", "Specify the directory name of new worktree. [default to branch name]")
}
