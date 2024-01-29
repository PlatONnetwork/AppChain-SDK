# BeginBlocker and EndBlocker

开发者可以选择实现`BeginBlockerModule`和`EndBlockerModule`接口，它们分别在区块开始和结束时触发。

## BeginBlockerModule

!!! info "BeginBlock"

    `BeginBlock(ctx sdk.WorkerContext) error`

    参数：

    - **ctx** AppChain SDK worker context.

    返回值：

    - 成功返回`nil`。
    - 失败返回具体错误。

```go title="x/staking/module.go"
func (s *StakeModule) BeginBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// increase the number of validator blocks generated from the previous block
	parentBlock := currentBlock - 1
	if parentBlock != 0 {
		parentHeader := ctx.ParentBlock().Header()
		s.logger.Debug("Start call setNumberOfBlocksForRoundValidator", "currentBlock", currentBlock, "parentBlock", parentHeader.Number.Uint64())
		if err := s.setNumberOfBlocksForRoundValidator(ctx.StateDB(), parentHeader); nil != err {
			return fmt.Errorf("can not set number of blocks for round validators, %s, parentBlock: %d", err, parentBlock)
		}
	}

	if s.stageModule.IsBeginOfCurrentRound(ctx.StateDB(), currentBlock) {
		currentRound := s.stageModule.GetCurrentRound(ctx.StateDB())
		minRoundValidatorBlockNumber := s.GetMinRoundValidatorBlockNumber(ctx.StateDB())
		// check low blocks validators
		lowBlocksValidatorAddrQueue := db.CheckLowBlocksValidatorForPreviousRound(ctx.StateDB(), s.Address(), currentRound, minRoundValidatorBlockNumber)
		// update validator status
		for _, validatorAddr := range lowBlocksValidatorAddrQueue {
			if err := s.updateValidatorStatus(ctx.StateDB(), validatorAddr, staketypes.Invalided|staketypes.LowBlocks); nil != err {
				return fmt.Errorf("can not update validator status to [lowBlocks], %s, validator: %s", err, validatorAddr.Hex())
			}
		}
	}
	return nil
}
```

## EndBlockerModule

!!! info "EndBlock"

    `EndBlock(ctx sdk.WorkerContext) error`

    参数：

    - **ctx** AppChain SDK worker context.

    返回值：

    - 成功返回`nil`。
    - 失败返回具体错误。

``` go title="x/staking/module.go"
func (s *StakeModule) EndBlock(ctx sdk.WorkerContext) error {
	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// election next round validators (at cuurent round electionBlock)
	if s.stageModule.IsElectionBlockOnCurrentRound(ctx.StateDB(), currentBlock) {
		if err := s.electionRoundValidators(ctx, currentBlock); nil != err {
			return fmt.Errorf("can not elected round validators, %s", err)
		}
	}

	// election next epoch validators (at current epoch endBlock)
	// and store next epochItem
	if s.stageModule.IsEndOfCurrentEpoch(ctx.StateDB(), currentBlock) {
		if err := s.electionEpochValidators(ctx, currentBlock); nil != err {
			return fmt.Errorf("can not elected epoch validators, %s", err)
		}
	}
	return nil
}
```
