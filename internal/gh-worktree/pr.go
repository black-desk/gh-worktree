package ghworktree

type PullRequest interface {
	Reference() string
}

type pullRequest struct {
	reference string
}

func (pr *pullRequest) Reference() string {
	return pr.reference
}
