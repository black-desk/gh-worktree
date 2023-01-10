package ghworktree

import (
	"sync"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

var _once sync.Once
var _client api.RESTClient
var _clientErr error

func client() (api.RESTClient, error) {
	_once.Do(func() {
		_client, _clientErr = gh.RESTClient(nil)
	})
	return _client, _clientErr
}
