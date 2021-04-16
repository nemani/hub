package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sentinel-official/hub/v0.5/x/node"
	"github.com/sentinel-official/hub/v0.5/x/plan"
	"github.com/sentinel-official/hub/v0.5/x/provider"
	"github.com/sentinel-official/hub/v0.5/x/session"
	"github.com/sentinel-official/hub/v0.5/x/subscription"
	"github.com/sentinel-official/hub/v0.5/x/vpn/keeper"
	"github.com/sentinel-official/hub/v0.5/x/vpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case provider.MsgRegister:
			return provider.HandleRegister(ctx, k.Provider, msg)
		case provider.MsgUpdate:
			return provider.HandleUpdate(ctx, k.Provider, msg)

		case node.MsgRegister:
			return node.HandleRegister(ctx, k.Node, msg)
		case node.MsgUpdate:
			return node.HandleUpdate(ctx, k.Node, msg)
		case node.MsgSetStatus:
			return node.HandleSetStatus(ctx, k.Node, msg)

		case plan.MsgAdd:
			return plan.HandleAdd(ctx, k.Plan, msg)
		case plan.MsgSetStatus:
			return plan.HandleSetStatus(ctx, k.Plan, msg)
		case plan.MsgAddNode:
			return plan.HandleAddNode(ctx, k.Plan, msg)
		case plan.MsgRemoveNode:
			return plan.HandleRemoveNode(ctx, k.Plan, msg)

		case subscription.MsgSubscribeToPlan:
			return subscription.HandleSubscribeToPlan(ctx, k.Subscription, msg)
		case subscription.MsgSubscribeToNode:
			return subscription.HandleSubscribeToNode(ctx, k.Subscription, msg)
		case subscription.MsgCancel:
			return subscription.HandleCancel(ctx, k.Subscription, msg)
		case subscription.MsgAddQuota:
			return subscription.HandleAddQuota(ctx, k.Subscription, msg)
		case subscription.MsgUpdateQuota:
			return subscription.HandleUpdateQuota(ctx, k.Subscription, msg)

		case session.MsgUpsert:
			return session.HandleUpsert(ctx, k.Session, msg)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownMsgType, "%s", msg.Type())
		}
	}
}
