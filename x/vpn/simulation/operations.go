package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	node "github.com/sentinel-official/hub/v0.5/x/node/simulation"
	plan "github.com/sentinel-official/hub/v0.5/x/plan/simulation"
	provider "github.com/sentinel-official/hub/v0.5/x/provider/simulation"
	session "github.com/sentinel-official/hub/v0.5/x/session/simulation"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/simulation"
	"github.com/sentinel-official/hub/v0.5/x/vpn/expected"
	"github.com/sentinel-official/hub/v0.5/x/vpn/keeper"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, k keeper.Keeper) (operations simulation.WeightedOperations) {
	operations = append(operations, provider.WeightedOperations(params, cdc, ak, k.Provider)...)
	operations = append(operations, node.WeightedOperations(params, cdc, ak, k.Provider, k.Node)...)
	operations = append(operations, plan.WeightedOperations(params, cdc, ak, k.Provider, k.Node, k.Plan)...)
	operations = append(operations, subscription.WeightedOperations(params, cdc, ak, k.Plan, k.Node, k.Subscription)...)
	operations = append(operations, session.WeightedOperations(params, cdc, ak, k.Plan, k.Subscription)...)

	return operations
}
