package provider

import (
	"errors"
	"fmt"
)

type Resolver struct {
	NameToProvider map[string]Provider
}

func NewResolver() Resolver {
	resolver := Resolver{NameToProvider: map[string]Provider{}}
	resolver.Add(&InConfigProvider{})

	return resolver
}

func (it *Resolver) Add(provider Provider) {
	it.NameToProvider[provider.Name()] = provider
}

func (it *Resolver) Resolve(credentials Config) (Provider, error) {
	provider, ok := it.NameToProvider[credentials.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("auth provider %s not found", credentials.Type))
	}

	err := provider.SetCredentials(credentials)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
