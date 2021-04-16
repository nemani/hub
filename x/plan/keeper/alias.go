package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/v0.5/types"
	node "github.com/sentinel-official/hub/v0.5/x/node/types"
)

func (k Keeper) HasProvider(ctx sdk.Context, address hub.ProvAddress) bool {
	return k.provider.HasProvider(ctx, address)
}

func (k Keeper) GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool) {
	return k.node.GetNode(ctx, address)
}
