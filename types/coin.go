// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoBlackfury defines the default coin denomination used in Blackfury in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Blackfury.
	AttoBlackfury string = "ablackfury"

	// BaseDenomUnit defines the base denomination unit for Blackfury.
	// 1 blackfury = 1x10^{BaseDenomUnit} ablackfury
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewBlackfuryCoin is a utility function that returns an "ablackfury" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewBlackfuryCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoBlackfury, amount)
}

// NewBlackfuryDecCoin is a utility function that returns an "ablackfury" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewBlackfuryDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoBlackfury, amount)
}

// NewBlackfuryCoinInt64 is a utility function that returns an "ablackfury" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewBlackfuryCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoBlackfury, amount)
}
