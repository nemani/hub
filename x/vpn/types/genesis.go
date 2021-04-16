package types

import (
	"github.com/sentinel-official/hub/v0.5/x/deposit"
	"github.com/sentinel-official/hub/v0.5/x/node"
	"github.com/sentinel-official/hub/v0.5/x/plan"
	"github.com/sentinel-official/hub/v0.5/x/provider"
	"github.com/sentinel-official/hub/v0.5/x/session"
	"github.com/sentinel-official/hub/v0.5/x/subscription"
)

type GenesisState struct {
	Deposits      deposit.GenesisState      `json:"deposits"`
	Providers     provider.GenesisState     `json:"providers"`
	Nodes         node.GenesisState         `json:"nodes"`
	Plans         plan.GenesisState         `json:"plans"`
	Subscriptions subscription.GenesisState `json:"subscriptions"`
	Sessions      session.GenesisState      `json:"sessions"`
}

func NewGenesisState(deposits deposit.GenesisState, providers provider.GenesisState, nodes node.GenesisState,
	plans plan.GenesisState, subscriptions subscription.GenesisState, sessions session.GenesisState) GenesisState {
	return GenesisState{
		Deposits:      deposits,
		Providers:     providers,
		Nodes:         nodes,
		Plans:         plans,
		Subscriptions: subscriptions,
		Sessions:      sessions,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Deposits:      deposit.DefaultGenesisState(),
		Providers:     provider.DefaultGenesisState(),
		Nodes:         node.DefaultGenesisState(),
		Plans:         plan.DefaultGenesisState(),
		Subscriptions: subscription.DefaultGenesisState(),
		Sessions:      session.DefaultGenesisState(),
	}
}
