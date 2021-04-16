package vpn

import (
	"encoding/json"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	xsimulation "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/v0.5/x/vpn/client/cli"
	"github.com/sentinel-official/hub/v0.5/x/vpn/client/rest"
	"github.com/sentinel-official/hub/v0.5/x/vpn/expected"
	"github.com/sentinel-official/hub/v0.5/x/vpn/keeper"
	"github.com/sentinel-official/hub/v0.5/x/vpn/querier"
	"github.com/sentinel-official/hub/v0.5/x/vpn/simulation"
	"github.com/sentinel-official/hub/v0.5/x/vpn/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (a AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

func (a AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (a AppModuleBasic) ValidateGenesis(data json.RawMessage) error {
	var state types.GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &state)

	return ValidateGenesis(state)
}

func (a AppModuleBasic) RegisterRESTRoutes(context context.CLIContext, router *mux.Router) {
	rest.RegisterRoutes(context, router)
}

func (a AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

func (a AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(cdc)
}

type AppModule struct {
	AppModuleBasic
	ak expected.AccountKeeper
	k  keeper.Keeper
}

func NewAppModule(ak expected.AccountKeeper, k keeper.Keeper) AppModule {
	return AppModule{
		ak: ak,
		k:  k,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var state types.GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &state)
	InitGenesis(ctx, a.k, state)

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(ExportGenesis(ctx, a.k))
}

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (a AppModule) Route() string {
	return types.RouterKey
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.k)
}

func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (a AppModule) NewQuerierHandler() sdk.Querier {
	return querier.NewQuerier(a.k)
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlock(ctx, a.k)
}

func (a AppModule) GenerateGenesisState(state *module.SimulationState) {
	simulation.RandomizedGenesisState(state)
}

func (a AppModule) ProposalContents(_ module.SimulationState) []xsimulation.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(_ *rand.Rand) []xsimulation.ParamChange {
	return simulation.RandomizedParams()
}

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(state module.SimulationState) []xsimulation.WeightedOperation {
	return simulation.WeightedOperations(state.AppParams, state.Cdc, a.ak, a.k)
}
