// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)

package v13

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmkeeper "github.com/zeta-protocol/blackfury/v13/x/evm/keeper"
	evmtypes "github.com/zeta-protocol/blackfury/v13/x/evm/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v13
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	ek evmkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		if err := setPrecompilesParams(ctx, ek); err != nil {
			logger.Error("error while setting precompiles parameters", "error", err)
		}
		// Leave modules are as-is to avoid running InitGenesis.
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func setPrecompilesParams(ctx sdk.Context, ek evmkeeper.Keeper) error {
	params := ek.GetParams(ctx)
	params.ActivePrecompiles = evmtypes.DefaultActivePrecompiles
	return ek.SetParams(ctx, params)
}
