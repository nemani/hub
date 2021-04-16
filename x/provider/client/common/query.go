package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	hub "github.com/sentinel-official/hub/v0.5/types"
	"github.com/sentinel-official/hub/v0.5/x/provider/types"
)

func QueryProvider(ctx context.CLIContext, address hub.ProvAddress) (*types.Provider, error) {
	params := types.NewQueryProviderParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryProvider)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no provider found")
	}

	var provider types.Provider
	if err := ctx.Codec.UnmarshalJSON(res, &provider); err != nil {
		return nil, err
	}

	return &provider, nil
}

func QueryProviders(ctx context.CLIContext, skip, limit int) (types.Providers, error) {
	params := types.NewQueryProvidersParams(skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryProviders)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no providers found")
	}

	var providers types.Providers
	if err := ctx.Codec.UnmarshalJSON(res, &providers); err != nil {
		return nil, err
	}

	return providers, nil
}
