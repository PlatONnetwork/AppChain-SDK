package contractsapi

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
	"strings"

	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi/checkpoint_manager"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi/l2_state_sender"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	ethgoabi "github.com/umbracle/ethgo/abi"
)

type Validator struct {
	Address common.Address `abi:"_address"`
	BlsKey  [2]*big.Int    `abi:"blskey"`
}

var (
	ValidatorABIType        = ethgoabi.MustNewType("tuple(address _address,uint256[2] blskey)")
	CheckpointManagerABI, _ = abi.JSON(strings.NewReader(checkpoint_manager.CheckpointManagerMetaData.ABI))
	L2StateSenderABI, _     = abi.JSON(strings.NewReader(l2_state_sender.L2StateSenderMetaData.ABI))
)

func (v *Validator) EncodeAbi() ([]byte, error) {
	return ValidatorABIType.Encode(v)
}

type ExitEvent struct {
	BlockNumber        uint64
	L2StateSyncedEvent *l2_state_sender.L2StateSenderL2StateSynced
}

func (e *ExitEvent) Encode() ([]byte, error) {
	ev := e.L2StateSyncedEvent
	return L2StateSenderABI.Events["L2StateSynced"].Inputs.Pack(ev.Id, ev.Sender, ev.Receiver, ev.Data)
}

func DecodeExitEvent(log *coretypes.Log, number uint64) (*ExitEvent, error) {
	exitEvent := &ExitEvent{
		BlockNumber:        number,
		L2StateSyncedEvent: new(l2_state_sender.L2StateSenderL2StateSynced),
	}

	if err := UnpackLog(&L2StateSenderABI, exitEvent.L2StateSyncedEvent, "L2StateSynced", log); err != nil {
		return nil, err
	}
	return exitEvent, nil
}

func UnpackLog(contractAbi *abi.ABI, out interface{}, event string, log *coretypes.Log) error {
	if log.Topics[0] != contractAbi.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := contractAbi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range contractAbi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
