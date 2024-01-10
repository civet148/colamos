package keeper_test

import (
	"testing"

	testkeeper "github.com/civet148/colamos/testutil/keeper"
	"github.com/civet148/colamos/x/colamos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ColamosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
