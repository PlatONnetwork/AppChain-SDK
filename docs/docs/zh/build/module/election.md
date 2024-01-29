# 验证人选举

验证人选举是底层共识引擎所必需的。AppChain-SDK提供了一个实现`staking`模块，对大多数应用来说使用`staking`模块即可，无需额外开发模块。

## ElectionModule

### NewHeader

`staking`模块通过该方法设置区块头的`coinbase`字段，以标识出打包该区块的验证人。

!!! info "NewHeader"
    `NewHeader(ctx sdk.ConsensusContext, header *types.header) error`

    参数：

    - **ctx** AppChain SDK consensus context.
    - **header** 正在打包的区块的区块头

    返回值：

    - 成功返回`nil`。
    - 失败返回具体错误。

```go title="x/staking/module.go"
func (s *StakeModule) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {

	if ctx.IsProposer() {
		currentValidatorAddr := crypto.PubkeyToAddress(s.nodePrivateKey.PublicKey)

		currentValidator := db.GetValidator(ctx.ParentStateDB(), s.Address(), currentValidatorAddr)
		if currentValidator.IsEmpty() {
			return errors.New("not found validator")
		}

		header.Coinbase = currentValidator.Owner
	}
	return nil
}
```

#### GetLastNumber

!!! info "GetLastNumber"

    `GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64`

    参数：

    - **ctx** AppChain SDK consensus context.
    - **blockNumber** 块高。

    返回值：返回`blockNumber`对应共识轮最后一个区块的块高。

```go title="x/staking/module.go"
func (s *StakeModule) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return s.stageModule.GetLastNumber(ctx.StateDB(), blockNumber)
}
```

#### GetValiator

!!! info "GetValidator"

    `GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)`

    参数：

    - **ctx** AppChain SDK consensus context.
    - **blockNumber** 块高。

    返回值：

    - 成功，返回`blockNumber`对应共识轮的验证人列表。
    - 失败返回具体错误。

```go title="x/staking/module.go"
func (s *StakeModule) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return s.GetRoundValidator(ctx, blockNumber)
}
```

#### IsCandidateNode

!!! info "IsCandidateNode"

    `IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool`

    参数：

    - **ctx** AppChain SDK consensus context.
    - **nodeID** 节点ID。

    返回值：`nodeID`对应的节点是候选验证节点，返回true，否则返回false。


```go title="x/staking/module.go"
func (s *StakeModule) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	currentEpoch := s.stageModule.GetCurrentEpoch(ctx.StateDB())
	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentEpoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "epoch", currentEpoch)
		return false
	}

	for _, snap := range validatorSnapQueue {
		v := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsEmpty() {
			continue
		}
		pubkey, _ := v.PubKey.Pubkey()
		if enode.PubkeyToIDV4(pubkey) == nodeID.ID() {
			return true
		}
	}
	return false
}
```
