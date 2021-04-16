package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/v0.5/x/deposit/expected"
	"github.com/sentinel-official/hub/v0.5/x/deposit/types"
)

type Keeper struct {
	key    sdk.StoreKey
	cdc    *codec.Codec
	supply expected.SupplyKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		key: key,
		cdc: cdc,
	}
}

func (k *Keeper) WithSupplyKeeper(keeper expected.SupplyKeeper) {
	k.supply = keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
