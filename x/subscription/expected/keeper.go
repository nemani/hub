package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	hub "github.com/sentinel-official/hub/v0.5/types"
	node "github.com/sentinel-official/hub/v0.5/x/node/types"
	plan "github.com/sentinel-official/hub/v0.5/x/plan/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) exported.Account
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) error
}

type DepositKeeper interface {
	Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
	GetNodes(ctx sdk.Context, skip, limit int) node.Nodes
	GetActiveNodes(ctx sdk.Context, skip, limit int) node.Nodes
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plan.Plan, bool)
	GetPlans(ctx sdk.Context, skip, limit int) plan.Plans
	GetActivePlans(ctx sdk.Context, skip, limit int) plan.Plans
}
