package checkpoint

import (
	"context"
	"math/big"
	"testing"

	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/common"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"

	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/storage"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/txrelayer"
)

type mockConsensusContext struct{}

func (ctx *mockConsensusContext) Backend() sdk.Backend {
	return nil
}

func (ctx *mockConsensusContext) Context() context.Context {
	return context.Background()
}

func (ctx *mockConsensusContext) Epoch() uint64 {
	return 0
}

func (ctx *mockConsensusContext) View() uint64 {
	return 0
}

func (ctx *mockConsensusContext) BlockIndex() uint32 {
	return 0
}

func (ctx *mockConsensusContext) Header() *coretypes.Header {
	return &coretypes.Header{}
}

func (ctx *mockConsensusContext) StateDB() sdk.StateDBReader {
	return &state.StateDB{}
}

func (ctx *mockConsensusContext) ParentStateDB() sdk.StateDBReader {
	return &state.StateDB{}
}

func (ctx *mockConsensusContext) IsConsensusNode() bool {
	return false
}

func (ctx *mockConsensusContext) IsProposer() bool {
	return false
}

func (ctx *mockConsensusContext) NumberValidators(epoch uint64) int {
	return 0
}

type mockStaking struct{}

func (s *mockStaking) IsEndOfRound(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	return blockNumber%250 == 0
}

func (s *mockStaking) GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return &cbfttypes.Validators{}, nil
}

func (s *mockStaking) BlocksOfRound(ctx sdk.ConsensusContext) uint64 {
	return 250
}

func encodeExitEvent(id *big.Int, from, to common.Address, data []byte) ([]common.Hash, []byte) {
	eventAbi := contractsapi.L2StateSenderABI.Events["L2StateSynced"]

	bs, _ := eventAbi.Inputs.NonIndexed().Pack(data)
	topics, _ := abi.PackTopics(eventAbi.Inputs, interface{}(id), interface{}(from), interface{}(to))

	return append([]common.Hash{eventAbi.ID}, topics...), bs
}

func encodeExitEventLog(topics []common.Hash, data []byte, blockNumber uint64) *coretypes.Log {
	return &coretypes.Log{
		Address:     constants.StateSenderAddress,
		Topics:      topics,
		Data:        data,
		BlockNumber: blockNumber,
	}
}

func TestProcessLog(t *testing.T) {
	store := memorydb.New()
	m := &Module{
		logger:  log.New("module", "checkpoint"),
		store:   storage.NewStorage(store),
		staking: &mockStaking{},
	}

	from := common.Address{0x11}
	to := common.Address{0x12}

	topics, data := encodeExitEvent(big.NewInt(0), from, to, []byte{})
	err := m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 12,
		},
		encodeExitEventLog(topics, data, 12),
	)
	require.NoError(t, err)

	ev, err := m.store.GetExitEvent(0)
	require.NoError(t, err)
	require.Equal(t, ev.Epoch, uint64(0))
	require.Equal(t, ev.BlockNumber, uint64(12))
	require.Equal(t, ev.L2StateSyncedEvent.Id, big.NewInt(0))
	require.Equal(t, ev.L2StateSyncedEvent.Sender, common.Address{0x11})
	require.Equal(t, ev.L2StateSyncedEvent.Receiver, common.Address{0x12})

	topics0, data0 := encodeExitEvent(big.NewInt(1), from, to, []byte{0x1})
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 247,
		},
		encodeExitEventLog(topics0, data0, 247),
	)
	require.NoError(t, err)

	topics1, data1 := encodeExitEvent(big.NewInt(2), from, to, []byte{})
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 248,
		},
		encodeExitEventLog(topics1, data1, 248),
	)
	require.NoError(t, err)

	topics2, data2 := encodeExitEvent(big.NewInt(3), from, to, []byte{})
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 249,
		},
		encodeExitEventLog(topics2, data2, 249),
	)
	require.NoError(t, err)

	topics3, data3 := encodeExitEvent(big.NewInt(4), from, to, []byte{})
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 250,
		},
		encodeExitEventLog(topics3, data3, 250),
	)
	require.NoError(t, err)

	evs, err := m.store.GetExitEventsByEpoch(0)
	require.NoError(t, err)
	require.Equal(t, len(evs), 2)

	evs, err = m.store.GetExitEventsByEpoch(1)
	require.NoError(t, err)
	require.Equal(t, len(evs), 3)
}

func TestBuildRoot(t *testing.T) {
	store := memorydb.New()
	m := &Module{
		logger:  log.New("module", "checkpoint"),
		store:   storage.NewStorage(store),
		staking: &mockStaking{},
	}

	root, err := m.BuildEventRoot(0)
	require.NoError(t, err)
	require.Equal(t, root, common.ZeroHash)

	from := common.Address{0x11}
	to := common.Address{0x12}

	topics, data := encodeExitEvent(big.NewInt(0), from, to, []byte{})
	evLog := encodeExitEventLog(topics, data, 12)
	exitEv, err := contractsapi.DecodeExitEvent(evLog, 0, 12)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 12,
		},
		evLog,
	)
	require.NoError(t, err)

	topics0, data0 := encodeExitEvent(big.NewInt(1), from, to, []byte{0x1})
	evLog0 := encodeExitEventLog(topics0, data0, 247)
	exitEv0, err := contractsapi.DecodeExitEvent(evLog0, 0, 247)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 247,
		},
		evLog0,
	)
	require.NoError(t, err)

	topics1, data1 := encodeExitEvent(big.NewInt(2), from, to, []byte{})
	evLog1 := encodeExitEventLog(topics1, data1, 248)
	exitEv1, err := contractsapi.DecodeExitEvent(evLog1, 1, 251)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 248,
		},
		evLog1,
	)
	require.NoError(t, err)

	topics2, data2 := encodeExitEvent(big.NewInt(3), from, to, []byte{})
	evLog2 := encodeExitEventLog(topics2, data2, 248)
	exitEv2, err := contractsapi.DecodeExitEvent(evLog2, 1, 251)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 249,
		},
		evLog2,
	)
	require.NoError(t, err)

	topics3, data3 := encodeExitEvent(big.NewInt(4), from, to, []byte{})
	evLog3 := encodeExitEventLog(topics3, data3, 248)
	exitEv3, err := contractsapi.DecodeExitEvent(evLog3, 1, 251)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 250,
		},
		evLog3,
	)
	require.NoError(t, err)

	root0, err := m.BuildEventRoot(0)
	require.NoError(t, err)
	tree, err := createExitTree([]*contractsapi.ExitEvent{exitEv, exitEv0})
	require.NoError(t, err)
	require.Equal(t, root0, tree.Hash())

	root1, err := m.BuildEventRoot(1)
	require.NoError(t, err)
	tree0, err := createExitTree([]*contractsapi.ExitEvent{exitEv1, exitEv2, exitEv3})
	require.NoError(t, err)
	require.Equal(t, root1, tree0.Hash())
}

func TestGenerateExitProof(t *testing.T) {
	store := memorydb.New()
	txRelayer := txrelayer.NewModule("https://devnet2openapi2.platon.network/rpc", txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)
	err := txRelayer.Init(nil)
	require.NoError(t, err)
	m := &Module{
		logger:                log.New("module", "checkpoint"),
		store:                 storage.NewStorage(store),
		staking:               &mockStaking{},
		key:                   &keystore.Key{Address: common.HexToAddress("0xD4eBFb2355dfc3c662ae748EF4c1bE3258F9363c")},
		checkpointManagerAddr: common.HexToAddress("0456455348563f0d2894ADCdb4a7b4E794BC98d6"),
		txRelayer:             txRelayer,
	}

	from := common.Address{0x11}
	to := common.Address{0x12}

	topics, data := encodeExitEvent(big.NewInt(0), from, to, []byte{})
	evLog := encodeExitEventLog(topics, data, 12)
	exitEv, err := contractsapi.DecodeExitEvent(evLog, 0, 12)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 12,
		},
		evLog,
	)
	require.NoError(t, err)

	topics0, data0 := encodeExitEvent(big.NewInt(1), from, to, []byte{0x1})
	evLog0 := encodeExitEventLog(topics0, data0, 247)
	exitEv0, err := contractsapi.DecodeExitEvent(evLog0, 0, 247)
	require.NoError(t, err)
	err = m.ProcessLog(&mockConsensusContext{},
		&coretypes.Header{},
		&ctypes.QuorumCert{
			Epoch:       0,
			BlockNumber: 247,
		},
		evLog0,
	)
	require.NoError(t, err)

	tree, err := createExitTree([]*contractsapi.ExitEvent{exitEv, exitEv0})
	require.NoError(t, err)

	proof, err := m.generateExitProof(0)
	require.NoError(t, err)

	exitEvEncoded, err := exitEv.Encode()
	require.NoError(t, err)
	proof0, err := tree.GenerateProof(exitEvEncoded)
	leafIndex, err := tree.LeafIndex(exitEvEncoded)

	require.Equal(t, proof.Data, proof0)
	require.Equal(t, proof.Metadata["LeafIndex"], leafIndex)
}
