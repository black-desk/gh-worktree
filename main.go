package main

import (
	"log"

	"github.com/black-desk/gh-worktree/internal/worktree"
)

func main() {
	cmd, err := worktree.New()
	if err != nil {
                log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
                log.Fatal(err)
	}
}
