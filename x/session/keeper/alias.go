package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	hub "github.com/sentinel-official/hub/v0.5/types"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/types"
)

func (k Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) exported.Account {
	return k.account.GetAccount(ctx, address)
}

func (k Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool {
	return k.plan.HasNodeForPlan(ctx, id, address)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool) {
	return k.subscription.GetSubscription(ctx, id)
}

func (k Keeper) HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool {
	return k.subscription.HasSubscriptionForNode(ctx, address, id)
}

func (k Keeper) SetQuota(ctx sdk.Context, id uint64, quota subscription.Quota) {
	k.subscription.SetQuota(ctx, id, quota)
}

func (k Keeper) GetQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) (subscription.Quota, bool) {
	return k.subscription.GetQuota(ctx, id, address)
}

func (k Keeper) HasQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) bool {
	return k.subscription.HasQuota(ctx, id, address)
}

func (k Keeper) SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, from, to, sdk.NewCoins(coin))
}
