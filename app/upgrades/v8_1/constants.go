// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)

package v81

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v8.1.0"
	// MainnetUpgradeHeight defines the Blackfury mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 4_500_000
	// TestnetUpgradeHeight defines the Blackfury testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 5_280_000
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.1.0/blackfury_8.1.0_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.1.0/blackfury_8.1.0_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.1.0/blackfury_8.1.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.1.0/blackfury_8.1.0_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.1.0/blackfury_8.1.0_Windows_x86_64.zip"}}'`
)
