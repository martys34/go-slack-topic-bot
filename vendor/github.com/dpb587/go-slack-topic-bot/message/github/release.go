package github

import (
	"net/http"
	"context"
	"fmt"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/pkg/errors"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Release struct {
	Token string
	Alias string
	Owner string
	Name  string
}

var _ message.Messager = &Release{}

func (m Release) Message() (string, error) {
	ctx := context.Background()

	var client = http.DefaultClient

	if m.Token != "" {
		client = oauth2.NewClient(
			ctx,
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: m.Token}),
		)
	}

	gh := github.NewClient(client)

	// list all repositories for the authenticated user

	rels, _, err := gh.Repositories.ListReleases(ctx, m.Owner, m.Name, nil)
	if err != nil {
		return "", errors.Wrap(err, "listing releases")
	}

  for _, rel := range rels {
	  return fmt.Sprintf("%s/%s", m.Alias, *rel.Name), nil
  }

  return "", nil
}
