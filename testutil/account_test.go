package testutil

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccount(t *testing.T) {
	genesisJson, err := GenerateGenesis(DefaultAccount, nil)
	require.Nil(t, err)
	fmt.Println(string(genesisJson))
}
