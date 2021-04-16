package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	deposit "github.com/sentinel-official/hub/v0.5/x/deposit/client/cli"
	node "github.com/sentinel-official/hub/v0.5/x/node/client/cli"
	plan "github.com/sentinel-official/hub/v0.5/x/plan/client/cli"
	provider "github.com/sentinel-official/hub/v0.5/x/provider/client/cli"
	session "github.com/sentinel-official/hub/v0.5/x/session/client/cli"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/client/cli"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the VPN module",
	}

	cmd.AddCommand(deposit.GetQueryCommands(cdc)...)
	cmd.AddCommand(provider.GetQueryCommands(cdc)...)
	cmd.AddCommand(node.GetQueryCommands(cdc)...)
	cmd.AddCommand(plan.GetQueryCommands(cdc)...)
	cmd.AddCommand(subscription.GetQueryCommands(cdc)...)
	cmd.AddCommand(session.GetQueryCommands(cdc)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}

	cmd.AddCommand(provider.GetTxCommands(cdc)...)
	cmd.AddCommand(node.GetTxCommands(cdc)...)
	cmd.AddCommand(plan.GetTxCommands(cdc)...)
	cmd.AddCommand(subscription.GetTxCommands(cdc)...)
	cmd.AddCommand(session.GetTxCommands(cdc)...)

	return cmd
}
