package deploy

import (
	"testing"

	"github.com/bjartek/go-with-the-flow/gwtf"
	util "github.com/flow-usdc/flow-usdc"
	"github.com/stretchr/testify/assert"
)

func TestFiatTokenTotalSupplyInOwnerVault(t *testing.T) {
	g := gwtf.NewGoWithTheFlow("../../../flow.json")
	supply, err := util.GetTotalSupply(g)
	assert.NoError(t, err)
	assert.Equal(t, "10000.00000000", supply.String())

	balance, err := util.GetBalance(g, "owner")
	assert.NoError(t, err)
	assert.Equal(t, "10000.00000000", balance.String())
}
