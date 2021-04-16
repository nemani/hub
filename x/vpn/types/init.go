package types

import (
	"fmt"

	deposit "github.com/sentinel-official/hub/v0.5/x/deposit/types"
	node "github.com/sentinel-official/hub/v0.5/x/node/types"
	plan "github.com/sentinel-official/hub/v0.5/x/plan/types"
	provider "github.com/sentinel-official/hub/v0.5/x/provider/types"
	session "github.com/sentinel-official/hub/v0.5/x/session/types"
	subscription "github.com/sentinel-official/hub/v0.5/x/subscription/types"
)

func init() {
	node.ParamsSubspace = fmt.Sprintf("%s/%s", ModuleName, node.ModuleName)
	subscription.ParamsSubspace = fmt.Sprintf("%s/%s", ModuleName, subscription.ModuleName)
	session.ParamsSubspace = fmt.Sprintf("%s/%s", ModuleName, session.ModuleName)

	deposit.RouterKey = ModuleName
	provider.RouterKey = ModuleName
	node.RouterKey = ModuleName
	plan.RouterKey = ModuleName
	subscription.RouterKey = ModuleName
	session.RouterKey = ModuleName

	deposit.StoreKey = ModuleName
	provider.StoreKey = ModuleName
	node.StoreKey = ModuleName
	plan.StoreKey = ModuleName
	subscription.StoreKey = ModuleName
	session.StoreKey = ModuleName
}
