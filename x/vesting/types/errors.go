// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrInsufficientVestedCoins   = errorsmod.Register(ModuleName, 2, "insufficient vested coins error")
	ErrVestingLockup             = errorsmod.Register(ModuleName, 3, "vesting lockup error")
	ErrInsufficientUnlockedCoins = errorsmod.Register(ModuleName, 4, "insufficient unlocked coins error")
)
