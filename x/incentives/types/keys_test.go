package types_test

import (
	"testing"

	utiltx "github.com/zeta-protocol/blackfury/v13/testutil/tx"
	"github.com/zeta-protocol/blackfury/v13/x/incentives/types"

	"github.com/stretchr/testify/require"
)

func TestSplitGasMeterKey(t *testing.T) {
	contract := utiltx.GenerateAddress()
	user := utiltx.GenerateAddress()

	key := types.KeyPrefixGasMeter
	key = append(key, contract.Bytes()...)
	key = append(key, user.Bytes()...)

	contract2, user2 := types.SplitGasMeterKey(key)
	require.Equal(t, contract2, contract)
	require.Equal(t, user2, user)
}
