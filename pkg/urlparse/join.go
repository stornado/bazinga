package urlparse

import (
	"net/url"
	"path"
)

func URLJoin(server, urlpath string, others ...string) (string, error) {
	serverURL, err := url.Parse(server)
	if err != nil {
		return "", err
	}
	serverURL.Path = path.Join(serverURL.Path, urlpath)
	for _, other := range others {
		serverURL.Path = path.Join(serverURL.Path, other)
	}
	return serverURL.String(), nil
}
