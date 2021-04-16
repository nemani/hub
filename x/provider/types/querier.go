package types

import (
	hub "github.com/sentinel-official/hub/v0.5/types"
)

const (
	QueryProvider  = "Provider"
	QueryProviders = "Providers"
)

// QueryProviderParams is the request parameters for querying a provider.
type QueryProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryProviderParams(address hub.ProvAddress) QueryProviderParams {
	return QueryProviderParams{
		Address: address,
	}
}

// QueryProvidersParams is the request parameters for querying the providers.
type QueryProvidersParams struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQueryProvidersParams(skip, limit int) QueryProvidersParams {
	return QueryProvidersParams{
		Skip:  skip,
		Limit: limit,
	}
}
