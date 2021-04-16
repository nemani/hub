package simulation

import (
	"github.com/cosmos/cosmos-sdk/x/simulation"

	node "github.com/sentinel-official/hub/v0.5/x/node/simulation"
	session "github.com/sentinel-official/hub/v0.5/x/session/simulation"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/simulation"
)

func RandomizedParams() (params []simulation.ParamChange) {
	params = append(params, node.RandomizedParams()...)
	params = append(params, subscription.RandomizedParams()...)
	params = append(params, session.RandomizedParams()...)

	return params
}
