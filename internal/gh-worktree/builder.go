package ghworktree

import (
	"path"

	"github.com/black-desk/lib/go/errwrap"
)

type WorkTreeBuilder interface {
	Build(pathToWorkTree string, branch string) error
	With(opt WorkTreeBuilderOption) WorkTreeBuilder
}

type WorkTreeBuilderOption struct {
	Parent string
}

type workTreeBuilder struct {
	parent string
}

var _ WorkTreeBuilder = &workTreeBuilder{}

func NewWorkTreeBuilder() WorkTreeBuilder {
	return &workTreeBuilder{}
}

func (b *workTreeBuilder) Build(pathToWorkTree string, branch string) error {
	cmd := []string{
		"git", "worktree", "add",
		path.Join(b.parent, pathToWorkTree), branch,
	}
	return errwrap.Trace(RunCommand(cmd))
}

func (b *workTreeBuilder) With(opt WorkTreeBuilderOption) WorkTreeBuilder {
	b.parent = opt.Parent
	return b
}
