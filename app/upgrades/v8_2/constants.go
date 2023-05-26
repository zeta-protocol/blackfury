// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/zeta-protocol/blackfury/blob/main/LICENSE)

package v82

const (
	// UpgradeName is the shared upgrade plan name for mainnet
	UpgradeName = "v8.2.0"
	// MainnetUpgradeHeight defines the Blackfury mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 4_888_000
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/amd64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.2.0/blackfury_8.2.0_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.2.0/blackfury_8.2.0_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.2.0/blackfury_8.2.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.2.0/blackfury_8.2.0_Linux_amd64.tar.gz","windows/x86_64":"https://github.com/zeta-protocol/blackfury/releases/download/v8.2.0/blackfury_8.2.0_Windows_x86_64.zip"}}'`
)
