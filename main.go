package main

import (
	"log"
	"os"

	"github.com/black-desk/gh-worktree/internal/worktree"
)

func main() {
	cmd, err := worktree.New()
	if err != nil {
                log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
                os.Exit(1)
	}
}
