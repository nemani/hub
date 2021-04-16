package types

import (
	hub "github.com/sentinel-official/hub/v0.5/types"
)

type GenesisPlan struct {
	Plan  Plan              `json:"_"`
	Nodes []hub.NodeAddress `json:"nodes"`
}

type GenesisPlans []GenesisPlan

type (
	GenesisState = GenesisPlans
)

func NewGenesisState(plans GenesisPlans) GenesisState {
	return plans
}

func DefaultGenesisState() GenesisState {
	return nil
}
