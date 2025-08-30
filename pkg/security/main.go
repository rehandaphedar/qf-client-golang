package security

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuth2SecurityProvider struct {
	ClientID    string
	TokenSource oauth2.TokenSource
}

func (o *OAuth2SecurityProvider) Intercept(ctx context.Context, req *http.Request) error {
	token, err := o.TokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("x-auth-token", token.AccessToken)
	req.Header.Set("x-client-id", o.ClientID)

	return nil
}
