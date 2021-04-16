package types

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/v0.5/x/deposit"
	"github.com/sentinel-official/hub/v0.5/x/node"
	"github.com/sentinel-official/hub/v0.5/x/plan"
	"github.com/sentinel-official/hub/v0.5/x/provider"
	"github.com/sentinel-official/hub/v0.5/x/session"
	"github.com/sentinel-official/hub/v0.5/x/subscription"
)

var (
	ModuleCdc *codec.Codec
)

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	deposit.RegisterCodec(cdc)
	provider.RegisterCodec(cdc)
	node.RegisterCodec(cdc)
	plan.RegisterCodec(cdc)
	subscription.RegisterCodec(cdc)
	session.RegisterCodec(cdc)
}
