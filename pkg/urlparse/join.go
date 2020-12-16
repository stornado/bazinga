package urlparse

import (
	"fmt"
	"net/url"
	"path"
)

func URLJoin(server, urlpath string, others ...string) (string, error) {
	serverURL, err := url.Parse(server)
	if err != nil {
		return "", err
	}

	if !isValidPath(urlpath) {
		return "", fmt.Errorf("urlpath must be path, [%s] is not a allowed", urlpath)
	}

	if path.IsAbs(urlpath) {
		serverURL.Path = ""
	}
	serverURL.Path = path.Join(serverURL.Path, urlpath)

	for _, other := range others {
		if !isValidPath(other) {
			return "", fmt.Errorf("other urlpath must be path, [%s] is not allowed", other)
		}
		serverURL.Path = path.Join(serverURL.Path, other)
	}
	return serverURL.String(), nil
}

func isValidPath(urlpath string) bool {
	pathURL, err := url.Parse(urlpath)
	if err == nil && pathURL.Host == "" {
		return true
	}
	return false
}
