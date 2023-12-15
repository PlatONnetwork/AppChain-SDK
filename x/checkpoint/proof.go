package checkpoint

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
)

func (m *Module) generateExitProof(exitID uint64) (types.Proof, error) {
	m.logger.Debug("Generating proof for exit", "exitID", exitID)
	exitEvent, err := m.store.GetExitEvent(exitID)
	if err != nil {
		return types.Proof{}, err
	}

	blockNumber := big.NewInt(0).SetUint64(exitEvent.BlockNumber + types.CheckpointCommitDis)
	isFound, checkpointBlock, err := m.getCheckpointBlock(blockNumber)
	if err != nil {
		return types.Proof{}, err
	}

	if !isFound {
		return types.Proof{}, fmt.Errorf("checkpoint block not found for exit ID %d", exitID)
	}

	exitEventEncoded, err := exitEvent.Encode()
	if err != nil {
		return types.Proof{}, err
	}

	exitEvents, err := m.store.GetExitEventsByEpoch(exitEvent.Epoch)
	if err != nil {
		return types.Proof{}, nil
	}

	tree, err := createExitTree(exitEvents)
	if err != nil {
		return types.Proof{}, err
	}

	proof, err := tree.GenerateProof(exitEventEncoded)
	if err != nil {
		return types.Proof{}, err
	}

	leafIndex, err := tree.LeafIndex(exitEventEncoded)
	if err != nil {
		return types.Proof{}, err
	}

	m.logger.Debug("Generated proof for exit", "exitID", exitID, "leafIndex", leafIndex, "proofLen", len(proof))

	exitEventHex := hex.EncodeToString(exitEventEncoded)

	return types.Proof{
		Data: proof,
		Metadata: map[string]interface{}{
			"LeafIndex":       leafIndex,
			"ExitEvent":       exitEventHex,
			"CheckpointBlock": checkpointBlock,
		},
	}, nil
}

func (m *Module) getCheckpointBlock(blockNumber *big.Int) (bool, *big.Int, error) {
	checkpointABI := contractsapi.CheckpointManagerABI
	method := "getcheckpointBlock"
	input, err := checkpointABI.Pack(method, blockNumber)
	if err != nil {
		return false, nil, err
	}
	output, err := m.txRelayer.Call(m.key.Address, m.checkpointManagerAddr, input)
	if err != nil {
		return false, nil, err
	}
	res, err := checkpointABI.Unpack(method, output)
	if err != nil {
		return false, nil, nil
	}

	out0 := *abi.ConvertType(res[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(res[1], new(big.Int)).(**big.Int)

	return out0, out1, err
}
