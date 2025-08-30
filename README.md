# Introduction

A golang API client for the [Quran Foundation API](https://api-docs.quran.foundation/).

# Installation

```sh
go get git.sr.ht/~rehandaphedar/qf-client-golang
```

# Usage

The package handles OAuth2. Below is an example:

```golang
package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"git.sr.ht/~rehandaphedar/qf-client-golang/pkg/content"
	"git.sr.ht/~rehandaphedar/qf-client-golang/pkg/oauth2"
	"git.sr.ht/~rehandaphedar/qf-client-golang/pkg/security"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	var clientID string     // Assign this
	var clientSecret string // Assign this

	tokenURL, err := url.JoinPath(oauth2.ServerUrlProductionServer, "/oauth2/token")
	if err != nil {
		log.Fatal(err)
	}

	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"content"},
	}
	tokenSource := config.TokenSource(context.Background())

	provider := &security.OAuth2SecurityProvider{
		ClientID:    clientID,
		TokenSource: tokenSource,
	}

	client, err := content.NewClientWithResponses(content.ServerUrlProductionServer, content.WithRequestEditorFn(provider.Intercept))
	if err != nil {
		log.Fatal(err)
	}

	// Now you can use client normally. Example to fetch a specific chapter below.

	chapter, err := client.ChapterWithResponse(context.Background(), 1, &content.ChapterParams{})
	fmt.Println(chapter)
}
```

You can also just use the types from the `content` package, for example, for unmarshaling API responses in a case where they are cached as local JSON files.

# Generation

```sh
go generate
```

# Considerations

- For now, the schemas are slightly different from the ones in the [official repo](https://github.com/quran/qf-api-docs) - Mainly with regards to `operationId`s. Based off of [my PR](https://github.com/quran/qf-api-docs/pull/16).

- Some schemas are incomplete, so you might have to manually unmarshal some responses.
