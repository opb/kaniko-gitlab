package main

import (
	"errors"

	"os"

	"github.com/docker/docker-credential-helpers/credentials"
)

var notImplemented = errors.New("not implemented")

// ensure GitlabHelper adheres to the credentials.Helper interface
type GitlabHelper struct{}

func main() {
	credentials.Serve(GitlabHelper{})
}

func (GitlabHelper) Add(creds *credentials.Credentials) error {
	// This does not seem to get called
	return notImplemented
}

func (GitlabHelper) Delete(serverURL string) error {
	// This does not seem to get called
	return notImplemented
}

func (GitlabHelper) Get(serverURL string) (string, string, error) {
	password := os.Getenv("CI_BUILD_TOKEN")

	if password == "" {
		return "", "", errors.New("expected ENV var CI_BUILD_TOKEN to be set")
	}

	return "gitlab-ci-token", password, nil
}

func (GitlabHelper) List() (map[string]string, error) {

	registry := os.Getenv("CI_REGISTRY")

	if registry == "" {
		return nil, errors.New("expected ENV var CI_REGISTRY to be set")
	}

	result := map[string]string{}

	result[registry] = "gitlab-ci-token"

	return result, nil
}
