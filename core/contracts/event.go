package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

func PackEventTopics(id common.Hash, inputs abi.Arguments, query ...interface{}) ([]common.Hash, error) {
	hashes, err := abi.PackTopics(inputs, query...)
	if err != nil {
		return nil, err
	}
	return append([]common.Hash{id}, hashes...), nil
}

func PackEventData(inputs abi.Arguments, args ...interface{}) ([]byte, error) {
	var needInputs abi.Arguments
	var needArgs []interface{}
	for i, p := range inputs {
		if !p.Indexed {
			needInputs = append(needInputs, p)
			needArgs = append(needArgs, args[i])
		}
	}
	return needInputs.Pack(needArgs...)
}
