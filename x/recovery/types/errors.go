// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrBlockedAddress = errorsmod.Register(ModuleName, 2, "blocked address")
)
