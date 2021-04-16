package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/v0.5/types"
	"github.com/sentinel-official/hub/v0.5/x/subscription/keeper"
	"github.com/sentinel-official/hub/v0.5/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	var (
		log = k.Logger(ctx)
		end = ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	)

	k.IterateInactiveSubscriptions(ctx, end, func(_ int, item types.Subscription) bool {
		log.Info("Inactive subscription", "id", item.ID,
			"owner", item.Owner, "plan", item.Plan, "node", item.Node)

		if item.Plan == 0 {
			consumed := sdk.ZeroInt()
			k.IterateQuotas(ctx, item.ID, func(_ int, quota types.Quota) bool {
				consumed = consumed.Add(quota.Consumed)
				return false
			})

			amount := item.Deposit.Sub(item.Amount(consumed))
			log.Info("", "price", item.Price,
				"deposit", item.Deposit, "consumed", consumed, "amount", amount)

			if err := k.SubtractDeposit(ctx, item.Owner, amount); err != nil {
				panic(err)
			}

			k.DeleteInactiveSubscriptionAt(ctx, item.StatusAt.Add(k.InactiveDuration(ctx)), item.ID)
		} else {
			k.DeleteInactiveSubscriptionAt(ctx, item.Expiry, item.ID)
		}

		k.IterateQuotas(ctx, item.ID, func(_ int, quota types.Quota) bool {
			k.DeleteActiveSubscriptionForAddress(ctx, quota.Address, item.ID)
			k.SetInactiveSubscriptionForAddress(ctx, quota.Address, item.ID)
			return false
		})

		item.Status = hub.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSubscription(ctx, item)

		return false
	})

	return nil
}
