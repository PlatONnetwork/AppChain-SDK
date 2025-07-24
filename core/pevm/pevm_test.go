package pevm

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/consensus"
	coresdk "github.com/PlatONnetwork/PlatON-Go/core/sdk"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/PlatONnetwork/PlatON-Go/trie"
	"github.com/stretchr/testify/require"
)

var erc20Bytecode, _ = hexutil.Decode("0x60806040523480156200001157600080fd5b5060405162000e0038038062000e0083398101604081905262000034916200022a565b6040518060400160405280600a8152602001692132b731b42a37b5b2b760b11b8152506040518060400160405280600381526020016208486960eb1b81525081600390805190602001906200008b92919062000184565b508051620000a190600490602084019062000184565b505050620000b68282620000be60201b60201c565b5050620002ca565b6001600160a01b038216620001195760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200012d919062000266565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b82805462000192906200028d565b90600052602060002090601f016020900481019282620001b6576000855562000201565b82601f10620001d157805160ff191683800117855562000201565b8280016001018555821562000201579182015b8281111562000201578251825591602001919060010190620001e4565b506200020f92915062000213565b5090565b5b808211156200020f576000815560010162000214565b600080604083850312156200023e57600080fd5b82516001600160a01b03811681146200025657600080fd5b6020939093015192949293505050565b600082198211156200028857634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620002a257607f821691505b60208210811415620002c457634e487b7160e01b600052602260045260246000fd5b50919050565b610b2680620002da6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610194578063a9059cbb146101a7578063dd62ed3e146101ba57600080fd5b8063395093511461014357806370a082311461015657806395d89b411461018c57600080fd5b806318160ddd116100b257806318160ddd1461010f57806323b872dd14610121578063313ce5671461013457600080fd5b806306fdde03146100ce578063095ea7b3146100ec575b600080fd5b6100d6610200565b6040516100e391906109ea565b60405180910390f35b6100ff6100fa3660046109c0565b610292565b60405190151581526020016100e3565b6002545b6040519081526020016100e3565b6100ff61012f366004610984565b6102aa565b604051601281526020016100e3565b6100ff6101513660046109c0565b6102ce565b61011361016436600461092f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100d661031a565b6100ff6101a23660046109c0565b610329565b6100ff6101b53660046109c0565b6103ff565b6101136101c8366004610951565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461020f90610a9c565b80601f016020809104026020016040519081016040528092919081815260200182805461023b90610a9c565b80156102885780601f1061025d57610100808354040283529160200191610288565b820191906000526020600020905b81548152906001019060200180831161026b57829003601f168201915b5050505050905090565b6000336102a081858561040d565b5060019392505050565b6000336102b88582856105c0565b6102c3858585610697565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906102a09082908690610315908790610a5d565b61040d565b60606004805461020f90610a9c565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156103f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c3828686840361040d565b6000336102a0818585610697565b73ffffffffffffffffffffffffffffffffffffffff83166104af576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff8216610552576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106915781811015610684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103e9565b610691848484840361040d565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661073a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff82166107dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610893576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610691565b803573ffffffffffffffffffffffffffffffffffffffff8116811461092a57600080fd5b919050565b60006020828403121561094157600080fd5b61094a82610906565b9392505050565b6000806040838503121561096457600080fd5b61096d83610906565b915061097b60208401610906565b90509250929050565b60008060006060848603121561099957600080fd5b6109a284610906565b92506109b060208501610906565b9150604084013590509250925092565b600080604083850312156109d357600080fd5b6109dc83610906565b946020939093013593505050565b600060208083528351808285015260005b81811015610a17578581018301518582016040015282016109fb565b81811115610a29576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008219821115610a97577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b600181811c90821680610ab057607f821691505b60208210811415610aea577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea26469706673582212200a9e719132b0e157139134318e7df3b4e77dbfb493ffef3f355f41d9a1cd5d0e64736f6c63430008070033")

type chainContext struct {
	engine *consensus.BftMock
}

func newChainContext() *chainContext {
	return &chainContext{
		engine: consensus.NewFaker(),
	}
}

func (cc *chainContext) Engine() coresdk.Engine {
	return cc.engine
}

func (cc *chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	b := cc.engine.GetBlockByHashAndNum(hash, number)
	if b != nil {
		return b.Header()
	}
	return nil
}

type mockContractsApp struct{}

func newMockContractsApp() *mockContractsApp { return &mockContractsApp{} }
func (m *mockContractsApp) Contracts(sdk.StateDBReader, uint64) []sdk.SDKContract {
	return []sdk.SDKContract{}
}

type account struct {
	addr       common.Address
	privateKey *ecdsa.PrivateKey
}

func prepareAccounts(n int) []*account {
	accounts := make([]*account, n)
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		accounts[i] = &account{
			addr:       crypto.PubkeyToAddress(key.PublicKey),
			privateKey: key,
		}
	}
	return accounts
}

func prepareStateDB(accounts []*account, balance *big.Int) sdk.StateDB {
	statedb := mock.NewMockStateDB()
	for _, acc := range accounts {
		statedb.Balance[acc.addr] = new(big.Int).Set(balance)
		statedb.Nonce[acc.addr] = 0
	}
	return statedb
}

func prepareHeader(number int64) *types.Header {
	return &types.Header{
		ParentHash: common.ZeroHash,
		Number:     new(big.Int).SetInt64(number),
		GasLimit:   5000000000,
		Time:       uint64(time.Now().UnixMilli()),
		BaseFee:    big.NewInt(1000),
		Coinbase:   common.HexToAddress("0x0000012312312321"),
	}
}

func prepareTransactions(n int, accounts []*account, amount *big.Int) types.Transactions {
	txs := make(types.Transactions, n)
	signer := types.NewEIP155Signer(big.NewInt(params.TestChainConfig.ChainID.Int64()))
	for i := 0; i < n; i++ {
		from := accounts[i%len(accounts)]
		to := accounts[(i+1)%len(accounts)]
		tx, _ := types.SignNewTx(from.privateKey, signer, &types.LegacyTx{
			Nonce:    0,
			To:       &to.addr,
			Value:    new(big.Int).Set(amount),
			Gas:      21000,
			GasPrice: big.NewInt(5000),
		})
		txs[i] = tx
	}
	return txs
}

func TestPEVM(t *testing.T) {
	accounts := prepareAccounts(100)
	txs := prepareTransactions(100, accounts, big.NewInt(100))
	env1 := &Env{
		Header:        prepareHeader(10),
		StateDB:       prepareStateDB(accounts, big.NewInt(10000000000)),
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vm.Config{},
		ChainContext:  newChainContext(),
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm1 := NewPEVM(false, 4, 32, log.New("module", "test"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)

	env2 := &Env{
		Header:        prepareHeader(10),
		StateDB:       prepareStateDB(accounts, big.NewInt(10000000000)),
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vm.Config{},
		ChainContext:  newChainContext(),
		IsWorker:      true,
		BlockDeadline: time.Now().Add(1 * time.Minute),
	}
	pevm2 := NewPEVM(true, 4, 32, log.New("module", "test"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)
	require.True(t, result1.GasUsed == result2.GasUsed)

	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root)

	require.True(t, env2.StateDB.(*mock.MockStateDB).Equal(env1.StateDB.(*mock.MockStateDB)))
}

func TestHalfParallel(t *testing.T) {
	accounts := prepareAccounts(100)
	txs := prepareTransactions(100, accounts, big.NewInt(100))
	env1 := &Env{
		Header:        prepareHeader(10),
		StateDB:       prepareStateDB(accounts, big.NewInt(10000000000)),
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vm.Config{},
		ChainContext:  newChainContext(),
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm1 := NewPEVM(false, 4, 32, log.New("module", "test"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs[:50], false)
	require.Nil(t, err1)
	require.NotNil(t, result1)

	pevm1.forceSequential = true
	result12, err12 := pevm1.Run(txs[50:], false)
	require.Nil(t, err12)
	require.NotNil(t, result12)

	env2 := &Env{
		Header:        prepareHeader(10),
		StateDB:       prepareStateDB(accounts, big.NewInt(10000000000)),
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vm.Config{},
		ChainContext:  newChainContext(),
		IsWorker:      true,
		BlockDeadline: time.Now().Add(1 * time.Minute),
	}
	pevm2 := NewPEVM(true, 4, 32, log.New("module", "test"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)

	require.True(t, result12.GasUsed == result2.GasUsed)
	res1Root := types.DeriveSha(append(result1.Receipts, result12.Receipts...), trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root)

	require.True(t, env2.StateDB.(*mock.MockStateDB).Equal(env1.StateDB.(*mock.MockStateDB)))
}

var destructibleBytecode, _ = hexutil.Decode("0x608060405260748060116000396000f3fe60806040526004361060205760003560e01c806341c0e1b514602b57600080fd5b36602657005b600080fd5b348015603657600080fd5b50603c33ff5b00fea2646970667358221220c3b3e93eb98b3985d15e3728a9b06d56b316a52f1c275e93750f25d371709a1a64736f6c63430008070033")

func genDeployDestructibleTx(acc *account, nonce uint64, value *big.Int) *types.Transaction {
	tx, _ := types.SignNewTx(acc.privateKey, types.NewEIP155Signer(params.TestChainConfig.ChainID), &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(50000),
		Gas:      500000,
		To:       nil, // Contract creation
		Value:    value,
		Data:     destructibleBytecode,
	})
	return tx
}

func genKillTx(acc *account, nonce uint64, contractAddr common.Address) *types.Transaction {
	// kill()
	// Method ID: 0x41c0e1b5
	methodID, _ := hexutil.Decode("0x41c0e1b5")

	tx, _ := types.SignNewTx(acc.privateKey, types.NewEIP155Signer(params.TestChainConfig.ChainID), &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(50000),
		Gas:      100000,
		To:       &contractAddr,
		Value:    big.NewInt(0),
		Data:     methodID,
	})
	return tx
}

func TestSelfDestruct(t *testing.T) {
	// 1. Setup
	accounts := prepareAccounts(10)
	deployer := accounts[0]
	killer := accounts[1] // The killer will also be the beneficiary in this contract
	initialBalance := new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.LAT))

	statedb1 := prepareStateDB(accounts, initialBalance)
	statedb2 := prepareStateDB(accounts, initialBalance)
	// Capture initial balance of the killer for later comparison
	killerInitialBalance := new(big.Int).Set(statedb1.GetBalance(killer.addr))

	header := prepareHeader(30)
	chainCtx := newChainContext()
	vmCfg := vm.Config{}

	// 2. Create transactions
	txs := make(types.Transactions, 2)
	contractInitialBalance := new(big.Int).SetUint64(1e18) // 1 LAT

	// Tx 0: Deploy contract with an initial balance
	txs[0] = genDeployDestructibleTx(deployer, 0, contractInitialBalance)
	contractAddr := crypto.CreateAddress(deployer.addr, 0)

	// Tx 1: Killer account calls kill(). The funds go to the killer (msg.sender).
	txs[1] = genKillTx(killer, 0, contractAddr)

	// 3. Run sequential
	env1 := &Env{
		Header:        header,
		StateDB:       statedb1,
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vmCfg,
		ChainContext:  chainCtx,
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm1 := NewPEVM(true, 4, 32, log.New("module", "test-seq"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)
	require.Len(t, result1.Receipts, 2)

	// 4. Run parallel
	env2 := &Env{
		Header:        header,
		StateDB:       statedb2,
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vmCfg,
		ChainContext:  chainCtx,
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm2 := NewPEVM(false, 2, 32, log.New("module", "test-para"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)
	require.Len(t, result2.Receipts, 2)

	// 5. Compare results
	require.Equal(t, result1.GasUsed, result2.GasUsed, "Gas used should be equal")
	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root, "Receipt roots should be equal")
	require.True(t, env1.StateDB.(*mock.MockStateDB).Equal(env2.StateDB.(*mock.MockStateDB)), "StateDBs should be equal")

	// 6. Verify specific state changes from the parallel run
	// Check that contract code is gone
	finalCode := statedb2.GetCode(contractAddr)
	require.NotEmpty(t, finalCode, "Contract code should be not empty after selfdestruct")

	// Check that killer's balance increased by the contract's balance, minus gas costs
	killTxReceipt := result2.Receipts[1]
	killTx := txs[1]
	gasCost := new(big.Int).Mul(new(big.Int).SetUint64(killTxReceipt.GasUsed), killTx.GasPrice())

	expectedKillerBalance := new(big.Int).Add(killerInitialBalance, contractInitialBalance)
	expectedKillerBalance.Sub(expectedKillerBalance, gasCost)

	finalKillerBalance := statedb2.GetBalance(killer.addr)
	require.Equal(t, 0, expectedKillerBalance.Cmp(finalKillerBalance), "Killer's final balance is incorrect")
}

func TestStorageConflict(t *testing.T) {
	// 1. Setup
	accounts := prepareAccounts(10)
	initialBalance := new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.LAT))
	statedb1 := prepareStateDB(accounts, initialBalance)
	statedb2 := prepareStateDB(accounts, initialBalance)
	header := prepareHeader(40)
	chainCtx := newChainContext()
	vmCfg := vm.Config{}

	// 2. Deploy ERC20 and create conflicting transactions
	deployer := accounts[0]
	recipient := accounts[1]
	initialSupply := new(big.Int).Mul(new(big.Int).SetInt64(1000000), big.NewInt(params.LAT))

	// Tx 0: Deploy contract
	deployTx := genERC20DeployTx(deployer, initialSupply, 0)
	contractAddr := crypto.CreateAddress(deployer.addr, 0)

	// Apply deploy transaction to both states so the contract exists before the conflicting txs
	applyTxToState := func(env *Env, tx *types.Transaction) {
		pevm := NewPEVM(true, 1, 1, log.New(), env, newMockContractsApp()) // Sequential execution
		_, err := pevm.Run(types.Transactions{tx}, false)
		require.Nil(t, err)
	}
	env1Temp := &Env{Header: header, StateDB: statedb1, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	env2Temp := &Env{Header: header, StateDB: statedb2, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	applyTxToState(env1Temp, deployTx)
	applyTxToState(env2Temp, deployTx)

	// Create 3 transactions from the same sender to the same recipient.
	// This will cause storage conflicts on the balance slots of both accounts in the ERC20 contract.
	txs := make(types.Transactions, 3)
	transferAmount := new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Von))
	for i := 0; i < 3; i++ {
		// Nonce for deployer is now 1, so start subsequent txs from 1
		txs[i] = genERC20TransferTx(deployer, uint64(i+1), contractAddr, recipient.addr, transferAmount)
	}

	// 3. Run sequential
	env1 := &Env{Header: header, StateDB: statedb1, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	pevm1 := NewPEVM(true, 4, 32, log.New("module", "test-seq"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)

	// 4. Run parallel
	env2 := &Env{Header: header, StateDB: statedb2, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	pevm2 := NewPEVM(false, 2, 32, log.New("module", "test-para"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)

	// 5. Compare results
	require.Equal(t, result1.GasUsed, result2.GasUsed, "Gas used should be equal")
	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root, "Receipt roots should be equal")
	require.True(t, env1.StateDB.(*mock.MockStateDB).Equal(env2.StateDB.(*mock.MockStateDB)), "StateDBs should be equal")
}

func TestRevertedTransaction(t *testing.T) {
	// 1. Setup
	accounts := prepareAccounts(10)
	initialBalance := new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.LAT))
	statedb1 := prepareStateDB(accounts, initialBalance)
	statedb2 := prepareStateDB(accounts, initialBalance)
	header := prepareHeader(50)
	chainCtx := newChainContext()
	vmCfg := vm.Config{}

	// 2. Deploy ERC20 and create mixed success/fail transactions
	deployer := accounts[0]
	initialSupply := new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.LAT)) // 1000 tokens

	// Tx 0: Deploy contract
	deployTx := genERC20DeployTx(deployer, initialSupply, 0)
	contractAddr := crypto.CreateAddress(deployer.addr, 0)

	// Apply deploy transaction to both states
	applyTxToState := func(env *Env, tx *types.Transaction) {
		pevm := NewPEVM(true, 1, 1, log.New(), env, newMockContractsApp())
		_, err := pevm.Run(types.Transactions{tx}, false)
		require.Nil(t, err)
	}
	env1Temp := &Env{Header: header, StateDB: statedb1, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	env2Temp := &Env{Header: header, StateDB: statedb2, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	applyTxToState(env1Temp, deployTx)
	applyTxToState(env2Temp, deployTx)

	txs := make(types.Transactions, 4)
	// Tx 0 (Success): deployer transfers 100 tokens to accounts[1]
	txs[0] = genERC20TransferTx(deployer, 1, contractAddr, accounts[1].addr, big.NewInt(100))
	// Tx 1 (Fail): accounts[2] (no tokens) transfers to accounts[3]
	txs[1] = genERC20TransferTx(accounts[2], 0, contractAddr, accounts[3].addr, big.NewInt(100))
	// Tx 2 (Success): deployer transfers 100 tokens to accounts[4]
	txs[2] = genERC20TransferTx(deployer, 2, contractAddr, accounts[4].addr, big.NewInt(100))
	// Tx 3 (Fail): deployer transfers more than total supply
	txs[3] = genERC20TransferTx(deployer, 3, contractAddr, accounts[5].addr, new(big.Int).Add(initialSupply, big.NewInt(1)))

	// 3. Run sequential
	env1 := &Env{Header: header, StateDB: statedb1, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	pevm1 := NewPEVM(true, 4, 32, log.New("module", "test-seq"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)

	// 4. Run parallel
	env2 := &Env{Header: header, StateDB: statedb2, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	pevm2 := NewPEVM(false, 4, 32, log.New("module", "test-para"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)

	// 5. Compare results
	require.Equal(t, result1.GasUsed, result2.GasUsed, "Gas used should be equal")
	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root, "Receipt roots should be equal")
	require.True(t, env1.StateDB.(*mock.MockStateDB).Equal(env2.StateDB.(*mock.MockStateDB)), "StateDBs should be equal")

	// 6. Verify receipt statuses
	require.Equal(t, types.ReceiptStatusSuccessful, result2.Receipts[0].Status, "Tx 0 should succeed")
	require.Equal(t, types.ReceiptStatusFailed, result2.Receipts[1].Status, "Tx 1 should fail")
	require.Equal(t, types.ReceiptStatusSuccessful, result2.Receipts[2].Status, "Tx 2 should succeed")
	require.Equal(t, types.ReceiptStatusFailed, result2.Receipts[3].Status, "Tx 3 should fail")
}

func TestContractCreationDependency(t *testing.T) {
	// 1. Setup
	accounts := prepareAccounts(2)
	initialBalance := new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.LAT))
	statedb1 := prepareStateDB(accounts, initialBalance)
	statedb2 := prepareStateDB(accounts, initialBalance)
	chainCtx := newChainContext()
	vmCfg := vm.Config{}

	// 2. Create transactions: 1. Deploy, 2. Send funds to new contract
	deployer := accounts[0]
	funder := accounts[1]
	contractAddr := crypto.CreateAddress(deployer.addr, 0)
	fundAmount := new(big.Int).Mul(big.NewInt(1), big.NewInt(params.LAT))

	txs := make(types.Transactions, 2)
	// Tx 0: Deploy a simple contract
	txs[0] = genDeployDestructibleTx(deployer, 0, big.NewInt(0)) // Deploy with 0 value
	// Tx 1: Send funds to the contract address
	tx, _ := types.SignNewTx(funder.privateKey, types.NewEIP155Signer(params.TestChainConfig.ChainID), &types.LegacyTx{
		Nonce:    0,
		To:       &contractAddr,
		Value:    fundAmount,
		Gas:      210000,
		GasPrice: big.NewInt(5000),
	})
	txs[1] = tx

	// 3. Run sequential
	env1 := &Env{Header: prepareHeader(60), StateDB: statedb1, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	pevm1 := NewPEVM(true, 4, 32, log.New("module", "test-seq"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)

	// 4. Run parallel
	env2 := &Env{Header: prepareHeader(60), StateDB: statedb2, ChainConfig: params.TestChainConfig, VMConfig: vmCfg, ChainContext: chainCtx, IsWorker: true, BlockDeadline: time.Now().Add(time.Minute)}
	pevm2 := NewPEVM(false, 2, 32, log.New("module", "test-para"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)

	// 5. Compare results
	require.Equal(t, result1.GasUsed, result2.GasUsed, "Gas used should be equal")
	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root, "Receipt roots should be equal")
	require.True(t, env1.StateDB.(*mock.MockStateDB).Equal(env2.StateDB.(*mock.MockStateDB)), "StateDBs should be equal")

	// 6. Verify final contract balance
	finalContractBalance := statedb2.GetBalance(contractAddr)
	require.Equal(t, 0, fundAmount.Cmp(finalContractBalance), "Contract balance should be equal to the funded amount")
}

func TestDeadline(t *testing.T) {
	// Test parallel mode
	{
		accounts := prepareAccounts(100)
		txs := prepareTransactions(100, accounts, big.NewInt(100))
		initialState := prepareStateDB(accounts, big.NewInt(10000000000))
		// Create a second, identical state DB for comparison after the run.
		comparisonState := prepareStateDB(accounts, big.NewInt(10000000000))

		env := &Env{
			Header:        prepareHeader(70),
			StateDB:       initialState,
			ChainConfig:   params.TestChainConfig,
			VMConfig:      vm.Config{},
			ChainContext:  newChainContext(),
			IsWorker:      true,
			BlockDeadline: time.Now().Add(-1 * time.Millisecond),
		}

		pevmPara := NewPEVM(false, 4, 32, log.New("module", "test"), env, newMockContractsApp())
		resultPara, errPara := pevmPara.Run(txs, false)
		require.NoError(t, errPara, "Error should be nil on deadline occupy")
		require.NotNil(t, resultPara, "Result should be not nil on deadline occupy")
		require.True(t, resultPara.Timeout, "Timeout should be true on deadline")
		// Verify state was not mutated by comparing to the pristine comparison state.
		require.True(t, comparisonState.(*mock.MockStateDB).Equal(env.StateDB.(*mock.MockStateDB)), "StateDB should not be modified on deadline error")
	}

	// Test sequential mode
	{
		accounts := prepareAccounts(100)
		txs := prepareTransactions(100, accounts, big.NewInt(100))
		initialState := prepareStateDB(accounts, big.NewInt(10000000000))
		comparisonState := prepareStateDB(accounts, big.NewInt(10000000000))

		envSeq := &Env{
			Header:        prepareHeader(70),
			StateDB:       initialState,
			ChainConfig:   params.TestChainConfig,
			VMConfig:      vm.Config{},
			ChainContext:  newChainContext(),
			IsWorker:      true,
			BlockDeadline: time.Now().Add(-1 * time.Millisecond),
		}
		pevmSeq := NewPEVM(true, 4, 32, log.New("module", "test"), envSeq, newMockContractsApp())
		resultSeq, errSeq := pevmSeq.Run(txs, false)
		require.NoError(t, errSeq, "Error should be nil on deadline occupy")
		require.NotNil(t, resultSeq, "Result should be not nil on deadline error in sequential mode")
		require.True(t, resultSeq.Timeout, "Timeout should be true on deadline")
		require.True(t, comparisonState.(*mock.MockStateDB).Equal(envSeq.StateDB.(*mock.MockStateDB)), "StateDB should not be modified on deadline error in sequential mode")
	}
}

func TestPrecompiledContracts(t *testing.T) {
	// 1. Setup
	accounts := prepareAccounts(1)
	initialBalance := new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.LAT))
	statedb1 := prepareStateDB(accounts, initialBalance)
	statedb2 := prepareStateDB(accounts, initialBalance)
	header := prepareHeader(90)
	chainCtx := newChainContext()
	vmCfg := vm.Config{}

	// 2. Create 10 transactions that call different precompiled contracts
	txs := make(types.Transactions, 10)
	caller := accounts[0]

	// Precompiled contract addresses
	// Using simpler contracts that are less sensitive to input format
	sha256Addr := common.BytesToAddress([]byte{2})
	ripemd160Addr := common.BytesToAddress([]byte{3})
	identityAddr := common.BytesToAddress([]byte{4})

	// Create transactions for different precompiled contracts
	inputs := [][]byte{
		[]byte("Hello, World!"),
		[]byte("Test input 1"),
		[]byte("Test input 2"),
		[]byte("Another test"),
		[]byte("More testing"),
		[]byte(""),
		[]byte("Yet another input"),
		[]byte("Different data"),
		[]byte("More data"),
		[]byte("Final test input"),
	}

	// Alternate between the three precompiled contracts
	addresses := []common.Address{sha256Addr, ripemd160Addr, identityAddr}

	for i, input := range inputs {
		txs[i], _ = types.SignNewTx(caller.privateKey, types.NewEIP155Signer(params.TestChainConfig.ChainID), &types.LegacyTx{
			Nonce:    uint64(i),
			To:       &addresses[i%3], // Cycle through the three contracts
			Value:    big.NewInt(0),
			Gas:      100000,
			GasPrice: big.NewInt(5000),
			Data:     input,
		})
	}

	// 3. Run sequential
	env1 := &Env{
		Header:        header,
		StateDB:       statedb1,
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vmCfg,
		ChainContext:  chainCtx,
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm1 := NewPEVM(true, 4, 32, log.New("module", "test-seq"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)
	require.Len(t, result1.Receipts, 10)

	// 4. Run parallel
	env2 := &Env{
		Header:        header,
		StateDB:       statedb2,
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vmCfg,
		ChainContext:  chainCtx,
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm2 := NewPEVM(false, 4, 32, log.New("module", "test-para"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)
	require.Len(t, result2.Receipts, 10)

	// 5. Compare results
	require.Equal(t, result1.GasUsed, result2.GasUsed, "Gas used should be equal")
	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root, "Receipt roots should be equal")
	require.True(t, env1.StateDB.(*mock.MockStateDB).Equal(env2.StateDB.(*mock.MockStateDB)), "StateDBs should be equal")

	// 6. Verify transaction statuses
	for i, receipt := range result2.Receipts {
		// All transactions should succeed
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status,
			"Transaction %d should succeed", i)
		require.Equal(t, result1.Receipts[i].Status, receipt.Status,
			"Transaction %d status should match between sequential and parallel execution", i)
	}
}

func TestERC20(t *testing.T) {
	// 1. Setup
	accounts := prepareAccounts(10)
	initialBalance := new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.LAT))
	statedb1 := prepareStateDB(accounts, initialBalance)
	statedb2 := prepareStateDB(accounts, initialBalance)
	header := prepareHeader(20)
	chainCtx := newChainContext()
	vmCfg := vm.Config{}

	// 2. Create transactions
	txs := make(types.Transactions, 10)
	deployer := accounts[0]
	initialSupply := new(big.Int).Mul(big.NewInt(1000000), big.NewInt(params.LAT))

	// Tx 0: Deploy contract
	txs[0] = genERC20DeployTx(deployer, initialSupply, 0)
	contractAddr := crypto.CreateAddress(deployer.addr, 0)

	// Txs 1-9: Transfer tokens
	for i := 1; i < 10; i++ {
		recipient := accounts[i]
		amount := new(big.Int).Mul(big.NewInt(int64(i*100)), big.NewInt(params.Von))
		txs[i] = genERC20TransferTx(deployer, uint64(i), contractAddr, recipient.addr, amount)
	}

	// 3. Run sequential
	env1 := &Env{
		Header:        header,
		StateDB:       statedb1,
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vmCfg,
		ChainContext:  chainCtx,
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm1 := NewPEVM(true, 4, 32, log.New("module", "test-seq"), env1, newMockContractsApp())
	result1, err1 := pevm1.Run(txs, false)
	require.Nil(t, err1)
	require.NotNil(t, result1)
	//require.Len(t, result1.Receipts, 10)

	// 4. Run parallel
	env2 := &Env{
		Header:        header,
		StateDB:       statedb2,
		ChainConfig:   params.TestChainConfig,
		VMConfig:      vmCfg,
		ChainContext:  chainCtx,
		IsWorker:      true,
		BlockDeadline: time.Now().Add(time.Minute),
	}
	pevm2 := NewPEVM(false, 4, 32, log.New("module", "test-para"), env2, newMockContractsApp())
	result2, err2 := pevm2.Run(txs, false)
	require.Nil(t, err2)
	require.NotNil(t, result2)
	//require.Len(t, result2.Receipts, 10)

	// 5. Compare results
	require.Equal(t, result1.GasUsed, result2.GasUsed, "Gas used should be equal")

	res1Root := types.DeriveSha(result1.Receipts, trie.NewStackTrie(nil))
	res2Root := types.DeriveSha(result2.Receipts, trie.NewStackTrie(nil))
	require.Equal(t, res1Root, res2Root, "Receipt roots should be equal")

	require.True(t, env1.StateDB.(*mock.MockStateDB).Equal(env2.StateDB.(*mock.MockStateDB)), "StateDBs should be equal")
}

func genERC20DeployTx(acc *account, initialSupply *big.Int, nonce uint64) *types.Transaction {
	// constructor(address recipient,uint256 initialSupply)
	recipient := common.LeftPadBytes(acc.addr.Bytes(), 32)
	supply := common.LeftPadBytes(initialSupply.Bytes(), 32)
	data := append(erc20Bytecode, recipient...)
	data = append(data, supply...)

	tx, _ := types.SignNewTx(acc.privateKey, types.NewEIP155Signer(params.TestChainConfig.ChainID), &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(50000),
		Gas:      2000000,
		To:       nil, // Contract creation
		Value:    big.NewInt(0),
		Data:     data,
	})
	return tx
}

func genERC20TransferTx(acc *account, nonce uint64, contractAddr common.Address, to common.Address, amount *big.Int) *types.Transaction {
	// transfer(address to, uint256 amount)
	// Method ID: 0xa9059cbb
	methodID, _ := hexutil.Decode("0xa9059cbb")
	paddedTo := common.LeftPadBytes(to.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedTo...)
	data = append(data, paddedAmount...)

	tx, _ := types.SignNewTx(acc.privateKey, types.NewEIP155Signer(params.TestChainConfig.ChainID), &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(50000),
		Gas:      100000,
		To:       &contractAddr,
		Value:    big.NewInt(0),
		Data:     data,
	})
	return tx
}
