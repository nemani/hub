package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/v0.5/x/provider/expected"
	"github.com/sentinel-official/hub/v0.5/x/provider/types"
)

type Keeper struct {
	cdc          *codec.Codec
	key          sdk.StoreKey
	params       params.Subspace
	distribution expected.DistributionKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, params params.Subspace) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
	}
}

func (k *Keeper) WithDistributionKeeper(keeper expected.DistributionKeeper) {
	k.distribution = keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
