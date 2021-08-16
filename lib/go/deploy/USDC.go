package deploy

import (
	"encoding/hex"

	util "github.com/flow-usdc/flow-usdc"
	"github.com/onflow/cadence"

	"github.com/bjartek/go-with-the-flow/gwtf"
)

func DeployFiatTokenContract(
	g *gwtf.GoWithTheFlow,
	ownerAcct string) (events []*gwtf.FormatedEvent, err error) {
	contractCode := util.ParseCadenceTemplate("../../contracts/FiatToken.cdc")
	txFilename := "../../transactions/deploy/deploy_contract_with_auth.cdc"
	code := util.ParseCadenceTemplate(txFilename)
	encodedStr := hex.EncodeToString(contractCode)
	g.CreateAccountPrintEvents(
		"vaulted-account",
		"non-vaulted-account",
		"pauser",
		"non-pauser",
		"blocklister",
		"non-blocklister",
		"allowance",
		"non-allowance",
		"minter",
		"non-minter",
		"minterController1",
		"minterController2",
		"w-1000",
		"w-500-1",
		"w-500-2",
		"w-250-1",
		"w-250-2",
		"non-multisig-account",
	)

	pk1000 := g.Accounts[util.Acct1000].PrivateKey.PublicKey().String()
	pk500_1 := g.Accounts[util.Acct500_1].PrivateKey.PublicKey().String()
	pk500_2 := g.Accounts[util.Acct500_2].PrivateKey.PublicKey().String()
	pk250_1 := g.Accounts[util.Acct250_1].PrivateKey.PublicKey().String()
	pk250_2 := g.Accounts[util.Acct250_2].PrivateKey.PublicKey().String()

	w1000, _ := cadence.NewUFix64("1000.0")
	w500, _ := cadence.NewUFix64("500.0")
	w250, _ := cadence.NewUFix64("250.0")

	multiSigPubKeys := []cadence.Value{
		cadence.String(pk1000[2:]),
		cadence.String(pk500_1[2:]),
		cadence.String(pk500_2[2:]),
		cadence.String(pk250_1[2:]),
		cadence.String(pk250_2[2:]),
	}
	multiSigKeyWeights := []cadence.Value{w1000, w500, w500, w250, w250}

	e, err := g.TransactionFromFile(txFilename, code).
		SignProposeAndPayAs(ownerAcct).
		StringArgument("FiatToken").
		StringArgument(encodedStr).
		// Vault
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCVault"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultBalance"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultUUID"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultAllowance"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultReceiver"}).
		// Blocklist executor
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCBlocklistExe"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCBlocklistExeCap"}).
		// Blocklister
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCBlocklister"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCBlocklisterCapReceiver"}).
		// Pause executor
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCPauseExe"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCPauseExeCap"}).
		// Pauser
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCPauser"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCPauserCapReceiver"}).
		// Owner
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCOwner"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCOwnerCap"}).
		// Masterminter
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCMasterMinter"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCMasterMinterCap"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCMasterMinterPublicSigner"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCMasterMinterUUID"}).
		// Minter Controller
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCMinterController"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCMinterControllerUUID"}).
		// Minter
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCMinter"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCMinterUUID"}).
		StringArgument("USDC").
		UFix64Argument("10000.00000000").
		BooleanArgument(false).
		Argument(cadence.NewArray(multiSigPubKeys)).
		Argument(cadence.NewArray(multiSigKeyWeights)).
		Run()
	gwtf.PrintEvents(e, map[string][]string{})
	events = util.ParseTestEvents(e)

	return
}
