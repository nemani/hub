package plan

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/v0.5/types"
	"github.com/sentinel-official/hub/v0.5/x/plan/keeper"
	"github.com/sentinel-official/hub/v0.5/x/plan/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state {
		k.SetPlan(ctx, item.Plan)

		if item.Plan.Status.Equal(hub.StatusActive) {
			k.SetActivePlan(ctx, item.Plan.ID)
			k.SetActivePlanForProvider(ctx, item.Plan.Provider, item.Plan.ID)
		} else {
			k.SetInactivePlan(ctx, item.Plan.ID)
			k.SetInactivePlanForProvider(ctx, item.Plan.Provider, item.Plan.ID)
		}

		for _, node := range item.Nodes {
			k.SetNodeForPlan(ctx, item.Plan.ID, node)
		}
	}

	k.SetCount(ctx, uint64(len(state)))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	plans := k.GetPlans(ctx, 0, 0)

	items := make(types.GenesisPlans, 0, len(plans))
	for _, plan := range plans {
		item := types.GenesisPlan{
			Plan:  plan,
			Nodes: nil,
		}

		nodes := k.GetNodesForPlan(ctx, plan.ID, 0, 0)
		for _, node := range nodes {
			item.Nodes = append(item.Nodes, node.Address)
		}

		items = append(items, item)
	}

	return types.NewGenesisState(items)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, item := range state {
		if err := item.Plan.Validate(); err != nil {
			return err
		}
	}

	plans := make(map[uint64]bool)
	for _, item := range state {
		id := item.Plan.ID
		if plans[id] {
			return fmt.Errorf("duplicate plan id %d", id)
		}

		plans[id] = true
	}

	for _, item := range state {
		nodes := make(map[string]bool)
		for _, node := range item.Nodes {
			address := node.String()
			if nodes[address] {
				return fmt.Errorf("duplicate node for plan %d", item.Plan.ID)
			}

			nodes[address] = true
		}
	}

	return nil
}
