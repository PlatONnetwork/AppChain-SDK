package permit

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/test"
	"github.com/PlatONnetwork/AppChain-SDK/x/votetoken/contracts/eip712"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	evm            *vm.EVM
	contract       *vm.Contract
	eip712Contract *eip712.EIP712
	amount         = big.NewInt(100)
	permit         *ERC20Permit
)

func init() {
	initContract()
}

type Erc20 struct {
}

func (e Erc20) ApproveFrom(owner, spender common.Address, amount *big.Int) {

}

func initContract() {
	evm = test.NewEVM(test.NewMemoryStateDB(), vm.Config{}, vm.TxContext{}, test.NewBlockContext(), nil)
	evm.ChainConfig().ChainID = big.NewInt(31337)
	contract = vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")), big.NewInt(0), 10000000)
	eip712Contract, _ = eip712.NewEIP712(evm, contract, false)
	eip712Contract.Init("My Token", "1")
	permit, _ = NewERC20Permit(evm, contract, false, eip712Contract, &Erc20{})
}

func TestPermit(t *testing.T) {
	deadline, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	err := permit.Permit(common.HexToAddress("0x6ae6b1428c3b58d4d39dc20123bafe8467780cae"),
		common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
		big.NewInt(42),
		deadline,
		27,
		common.HexToHash("2ce96a6c7598ba914d11fe0adbe1e1797690c9fa9ac4b41baeb9abc31bd22e4c"),
		common.HexToHash("3005a586bf9590b96964b3898b4f3fa71d1d2408af6c6958a4010f064eaeaef1"),
	)
	require.Nil(t, err)
}
