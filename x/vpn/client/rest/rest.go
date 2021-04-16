package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	deposit "github.com/sentinel-official/hub/v0.5/x/deposit/client/rest"
	node "github.com/sentinel-official/hub/v0.5/x/node/client/rest"
	plan "github.com/sentinel-official/hub/v0.5/x/plan/client/rest"
	provider "github.com/sentinel-official/hub/v0.5/x/provider/client/rest"
	session "github.com/sentinel-official/hub/v0.5/x/session/client/rest"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/client/rest"
)

func RegisterRoutes(context context.CLIContext, router *mux.Router) {
	deposit.RegisterRoutes(context, router)
	provider.RegisterRoutes(context, router)
	node.RegisterRoutes(context, router)
	plan.RegisterRoutes(context, router)
	subscription.RegisterRoutes(context, router)
	session.RegisterRoutes(context, router)
}
