package ghworktree

import (
	"fmt"

	"github.com/black-desk/lib/go/errwrap"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
)

type Repo interface {
	PullRequest(id string) (PullRequest, error)
}

type repo struct {
	ghRepo repository.Repository
}

var _ Repo = &repo{}

func NewRepo() (Repo, error) {
	ghRepo, err := gh.CurrentRepository()
	if err != nil {
		return nil, errwrap.Trace(err)
	}

	log.Infof(
		"Resolve pull request/issue id in \"%v/%v/%v\"",
		ghRepo.Host(), ghRepo.Owner(), ghRepo.Name(),
	)

	return &repo{ghRepo: ghRepo}, nil
}

func (r *repo) PullRequest(id string) (PullRequest, error) {
	url := fmt.Sprintf(
		"repos/%s/%s/pulls/%s",
		r.ghRepo.Owner(),
		r.ghRepo.Name(),
		id,
	)

	client, err := client()
	if err != nil {
		return nil, errwrap.Trace(err)
	}

	var response = struct {
		Head struct {
			Ref string
		}
		State string
	}{}

	err = client.Get(url, &response)
	if err != nil {
		return nil, err
	}

	if response.State != "open" {
		log.Warn("This pr is NOT open, remote branch might has gone")
	}

	return &pullRequest{reference: response.Head.Ref}, nil
}
