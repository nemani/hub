package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/v0.5/types"
	"github.com/sentinel-official/hub/v0.5/x/session/types"
)

func (k Keeper) SetChannel(ctx sdk.Context, address sdk.AccAddress, subscription uint64, node hub.NodeAddress, channel uint64) {
	key := types.ChannelKey(address, subscription, node)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(channel)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetChannel(ctx sdk.Context, address sdk.AccAddress, subscription uint64, node hub.NodeAddress) (channel uint64) {
	store := k.Store(ctx)

	key := types.ChannelKey(address, subscription, node)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &channel)
	return channel
}
