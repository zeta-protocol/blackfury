package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/zeta-protocol/blackfury/v13/app"
	blackfuryd "github.com/zeta-protocol/blackfury/v13/cmd/blackfuryd"
	"github.com/zeta-protocol/blackfury/v13/utils"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := blackfuryd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"blackfury-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, utils.TestnetChainID+"-1"),
	})

	err := svrcmd.Execute(rootCmd, "blackfuryd", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := blackfuryd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "BLACKFURYD", app.DefaultNodeHome)
	require.Error(t, err)
}
