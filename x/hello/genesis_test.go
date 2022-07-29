package hello_test

import (
	"testing"

	keepertest "github.com/cosmos/interchain-security/testutil/keeper"
	"github.com/cosmos/interchain-security/testutil/nullify"

	"github.com/cosmos/interchain-security/x/hello"

	"github.com/cosmos/interchain-security/x/hello/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloKeeper(t)
	hello.InitGenesis(ctx, *k, genesisState)
	got := hello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}