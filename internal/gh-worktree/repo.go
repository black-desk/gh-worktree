package ghworktree

import (
	"fmt"

	"github.com/black-desk/lib/go/errwrap"
	"github.com/cli/cli/v2/pkg/cmd/factory"
	"github.com/cli/go-gh/pkg/repository"
)

type Repo interface {
	PullRequest(id string) (PullRequest, error)
}

type repo struct {
	ghRepo repository.Repository
}

type ghRepo struct {
	host  string
	owner string
	name  string
}

// Host implements repository.Repository
func (r *ghRepo) Host() string {
	return r.host
}

// Name implements repository.Repository
func (r *ghRepo) Name() string {
	return r.name
}

// Owner implements repository.Repository
func (r *ghRepo) Owner() string {
	return r.owner
}

var _ repository.Repository = &ghRepo{}

var _ Repo = &repo{}

func NewRepo() (Repo, error) {

	f := factory.New("1") // TODO(black_desk): which version should I use?
	f.BaseRepo = factory.SmartBaseRepoFunc(f)

	ghcliRepo, err := f.BaseRepo()
	if err != nil {
		return nil, errwrap.Trace(err)
	}

	ghRepo := &ghRepo{
		host:  ghcliRepo.RepoHost(),
		owner: ghcliRepo.RepoOwner(),
		name:  ghcliRepo.RepoName(),
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
