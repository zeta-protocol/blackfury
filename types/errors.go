// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "blackfury"

// ErrInvalidChainID returns an error resulting from an invalid chain ID.
var ErrInvalidChainID = errorsmod.Register(RootCodespace, 3, "invalid chain ID")
