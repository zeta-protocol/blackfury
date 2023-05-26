package ante_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-protocol/blackfury/v13/app"
	"github.com/zeta-protocol/blackfury/v13/crypto/ethsecp256k1"
	"github.com/zeta-protocol/blackfury/v13/encoding"
	"github.com/zeta-protocol/blackfury/v13/testutil"
	"github.com/zeta-protocol/blackfury/v13/utils"
	feemarkettypes "github.com/zeta-protocol/blackfury/v13/x/feemarket/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var s *AnteTestSuite

type AnteTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	clientCtx client.Context
	app       *app.Blackfury
	denom     string
}

func (suite *AnteTestSuite) SetupTest() {
	t := suite.T()
	privCons, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	consAddress := sdk.ConsAddress(privCons.PubKey().Address())

	isCheckTx := false
	suite.app = app.Setup(isCheckTx, feemarkettypes.DefaultGenesisState())
	suite.Require().NotNil(suite.app.AppCodec())

	header := testutil.NewHeader(
		1, time.Now().UTC(), utils.MainnetChainID+"-1", consAddress, nil, nil)
	suite.ctx = suite.app.BaseApp.NewContext(isCheckTx, header)

	suite.denom = utils.BaseDenom
	evmParams := suite.app.EvmKeeper.GetParams(suite.ctx)
	evmParams.EvmDenom = suite.denom
	_ = suite.app.EvmKeeper.SetParams(suite.ctx, evmParams)

	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
}

func TestAnteTestSuite(t *testing.T) {
	s = new(AnteTestSuite)
	suite.Run(t, s)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Run AnteHandler Integration Tests")
}
