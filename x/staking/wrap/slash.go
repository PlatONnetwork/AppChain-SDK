package wrap

import (
	"errors"
	"fmt"
	stakedb "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	basetypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func SetNumberOfBlocksForRoundValidator(db sdk.StateDB, header *basetypes.Header) error {
	// Extract the validator public key of the build block based on the signature in the block header
	sign := header.Signature()
	sealhash := header.SealHash().Bytes()
	pk, err := crypto.SigToPub(sealhash, sign)
	if err != nil {
		panic(fmt.Sprintf("can not sigToPub, blockNumber %d, %s", header.Number, err))
	}
	currentRound := stakedb.GetCurrentRound(db)
	currentRoundItem := stakedb.GetRoundItem(db, currentRound)

	var round uint64
	if currentRoundItem.StartBlock <= header.Number.Uint64() && currentRoundItem.EndBlock >= header.Number.Uint64() {
		round = currentRound
	} else {
		previousRoundItem := stakedb.GetRoundItem(db, currentRound-1)
		if previousRoundItem.StartBlock <= header.Number.Uint64() && previousRoundItem.EndBlock >= header.Number.Uint64() {
			round = currentRound - 1
		} else {
			return errors.New("not found round")
		}
	}
	stakedb.IncrementNumberOfBlocksForRoundValidator(db, crypto.PubkeyToAddress(*pk), round, 1)
	return nil
}
