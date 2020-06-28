package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/session/keeper"
	"github.com/sentinel-official/hub/x/dvpn/session/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, session := range state {
		k.SetSession(ctx, session)
		if session.Status.Equal(hub.StatusActive) {
			k.SetActiveSessionID(ctx, session.Subscription, session.Node, session.Address, session.ID)
		}

		k.SetSessionsCount(ctx, k.GetSessionsCount(ctx)+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetSessions(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
