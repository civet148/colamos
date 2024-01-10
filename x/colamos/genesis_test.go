package colamos_test

import (
	"testing"

	keepertest "github.com/civet148/colamos/testutil/keeper"
	"github.com/civet148/colamos/testutil/nullify"
	"github.com/civet148/colamos/x/colamos"
	"github.com/civet148/colamos/x/colamos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ColamosKeeper(t)
	colamos.InitGenesis(ctx, *k, genesisState)
	got := colamos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
