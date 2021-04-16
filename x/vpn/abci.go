package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/v0.5/x/node"
	"github.com/sentinel-official/hub/v0.5/x/session"
	"github.com/sentinel-official/hub/v0.5/x/subscription"
	"github.com/sentinel-official/hub/v0.5/x/vpn/keeper"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) abci.ValidatorUpdates {
	ctx, write := ctx.CacheContext()
	defer write()

	node.EndBlock(ctx, k.Node)
	session.EndBlock(ctx, k.Session)
	subscription.EndBlock(ctx, k.Subscription)

	return nil
}
