package hub

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

func init() {
	simapp.GetSimulatorFlags()
}

func TestAppExport(t *testing.T) {
	var (
		db      = tmdb.NewMemDB()
		logger  = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
		app     = NewApp(logger, db, nil, true, map[int64]bool{}, 0)
		genesis = ModuleBasics.DefaultGenesis()
	)

	state, err := codec.MarshalJSONIndent(app.cdc, genesis)
	assert.Nil(t, err, "Marshalling Genesis")

	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: state,
		},
	)
	app.Commit()

	app = NewApp(logger, db, nil, true, map[int64]bool{}, 0)
	_, _, err = app.ExportAppStateAndValidators(false, []string{})
	assert.Nil(t, err, "ExportAppStateAndValidators")
}

func fauxMerkleModeOpt(app *baseapp.BaseApp) {
	app.SetFauxMerkleMode()
}

func TestFullAppSimulation(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("leveldb-simulation", "simulation")
	if skip {
		t.Skip("Skipped App Simulation Setup")
	}
	assert.Nil(t, err, "App Simulation Setup Failed")

	defer func() {
		assert.Nil(t, db.Close(), "Closing DB Connection")
		assert.Nil(t, os.RemoveAll(dir), "Remove Everything in Simulation Path")
	}()

	app := NewApp(logger, db, nil, true, map[int64]bool{}, simapp.FlagPeriodValue, fauxMerkleModeOpt)
	assert.Equal(t, appName, app.Name(), "App Name")

	_, params, err := simulation.SimulateFromSeed(
		t, os.Stdout, app.BaseApp, simapp.AppStateFn(app.Codec(), app.SimulationManager()),
		simapp.SimulationOperations(app, app.Codec(), config),
		app.ModuleAccountAddrs(), config,
	)

	assert.Nil(t, simapp.CheckExportSimulation(app, config, params), "Export Simulation")
	assert.Nil(t, err, "Simulate From Seed")

	if config.Commit {
		simapp.PrintStats(db)
	}
}
