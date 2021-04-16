package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/v0.5/types"
	"github.com/sentinel-official/hub/v0.5/x/plan/client/common"
	"github.com/sentinel-official/hub/v0.5/x/plan/types"
)

func queryPlan(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "plan",
		Short: "Query a plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			plan, err := common.QueryPlan(ctx, id)
			if err != nil {
				return err
			}

			fmt.Println(plan)
			return nil
		},
	}
}

func queryPlans(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plans",
		Short: "Query plans",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			provider, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			s, err := cmd.Flags().GetString(flagStatus)
			if err != nil {
				return err
			}

			skip, err := cmd.Flags().GetInt(flagSkip)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			var (
				address hub.ProvAddress
				plans   types.Plans
				status  = hub.StatusFromString(s)
			)

			if len(provider) > 0 {
				address, err = hub.ProvAddressFromBech32(provider)
				if err != nil {
					return err
				}

				plans, err = common.QueryPlansForProvider(ctx, address, status, skip, limit)
			} else {
				plans, err = common.QueryPlans(ctx, status, skip, limit)
			}

			if err != nil {
				return err
			}

			for _, plan := range plans {
				fmt.Printf("%s\n\n", plan)
			}

			return nil
		},
	}

	cmd.Flags().String(flagProvider, "", "provider address")
	cmd.Flags().String(flagStatus, "", "status")
	cmd.Flags().Int(flagSkip, 0, "skip")
	cmd.Flags().Int(flagLimit, 25, "limit")

	return cmd
}
