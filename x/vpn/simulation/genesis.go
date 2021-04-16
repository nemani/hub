package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	deposit "github.com/sentinel-official/hub/v0.5/x/deposit/simulation"
	node "github.com/sentinel-official/hub/v0.5/x/node/simulation"
	plan "github.com/sentinel-official/hub/v0.5/x/plan/simulation"
	provider "github.com/sentinel-official/hub/v0.5/x/provider/simulation"
	session "github.com/sentinel-official/hub/v0.5/x/session/simulation"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/simulation"
	"github.com/sentinel-official/hub/v0.5/x/vpn/types"
)

func RandomizedGenesisState(state *module.SimulationState) {
	state.GenState[types.ModuleName] = state.Cdc.MustMarshalJSON(
		types.NewGenesisState(
			deposit.RandomizedGenesisState(state.Cdc),
			provider.RandomizedGenesisState(state.Cdc),
			node.RandomizedGenesisState(state.Cdc),
			plan.RandomizedGenesisState(state.Cdc),
			subscription.RandomizedGenesisState(state.Cdc),
			session.RandomizedGenesisState(state.Cdc),
		),
	)
}
