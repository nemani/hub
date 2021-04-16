package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	hub "github.com/sentinel-official/hub/v0.5/types"
	node "github.com/sentinel-official/hub/v0.5/x/node/types"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) exported.Account
}

type DepositKeeper interface {
	SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error
}

type PlanKeeper interface {
	HasNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool
	GetNodesForPlan(ctx sdk.Context, id uint64, skip, limit int) node.Nodes
}

type NodeKeeper interface {
	GetNodes(ctx sdk.Context, skip, limit int) node.Nodes
}

type SubscriptionKeeper interface {
	GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool)

	GetSubscriptionsForNode(ctx sdk.Context, address hub.NodeAddress, skip, limit int) subscription.Subscriptions
	HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool

	GetActiveSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int) subscription.Subscriptions

	SetQuota(ctx sdk.Context, id uint64, quota subscription.Quota)
	GetQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) (subscription.Quota, bool)
	HasQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) bool
}
