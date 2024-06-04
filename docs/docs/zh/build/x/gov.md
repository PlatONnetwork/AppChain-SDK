# Governance

治理模块是 AppChain-SDK 提供的基础治理模块，通过治理模块可以提交提案、投票等操作。


## 治理参与者

治理参与者是质押及委托的用户，每次的质押及委托都会调用投票代币合约，通过投票代币合约来计算用户的权重。

## 创世配置

```json
"governance":{
    "name": "gov", //治理合约名字
    "govVersion": "1", //治理版本
    "voteDelay": 10, //投票延时，提案发起后，经过多少区块才能进行投票
    "votePeriod": 10, //投票周期
    "quorumNumerator": 1, //最小提案投票比例1-100
    "proposalThreshold": 123333, //提案人投票权重阈值
    "owner": "lat1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpfr7f80", //治理合约拥有者
    "voteToken": "lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqtnsk753" //投票代币合约地址，查询参与用户的投票权重
}
```

## 提案流程

治理分为L1、L2治理，L1的治理，需要通过L2的治理合约进行投票，通过checkpoint机制进行提交到L1


```mermaid
flowchart
    Proposal[发布提案]
    Vote[提案投票]
    VotePass[提案通过]
    IsSendL1{是否发送到 L1}
    Checkpoint[治理事件有 Checkpoint 提交]
    L1Proof[提交提案证明]
    L1Execute[执行提案]
    L2ExecuteProposal[发送交易执行提案]
    UpgradeModule[提案加入到升级模块]
Proposal --> Vote
Vote--> VotePass
VotePass --> IsSendL1
IsSendL1 -->|Yes| Checkpoint
Checkpoint --> L1Proof
L1Proof --> L1Execute
IsSendL1 -->|No| L2ExecuteProposal

```

## 执行流程

### L2 Staking 账户

* L1->L2 

```mermaid
sequenceDiagram
    L1 -->> L2StateSync: Staking Commitment
    L2StateSync -->> L2Staking: stake/addStake/delegate
    L2Staking -->> L2ERC20Vote: mint
    L2ERC20Vote -->> L2ERC20Vote: mint, totalSupply, moveVotingPower
    L2ERC20Vote -->> L2Staking: Return
    L2Staking -->> L2StateSync: Return
```

* L2->L1

```mermaid
sequenceDiagram
    User -->> L2Staking: slash/unstake/undelegate/withdrawUnstake/withdrawUndelegate
    L2Staking -->> L2ERC20Vote: burn
    L2ERC20Vote -->> L2ERC20Vote: burn, totalSupply, moveVotingPower
    L2ERC20Vote -->> L2Staking: Return
```

Staking 账户调用了 `mint/burn`，资金在 L1->L2、L2->L1 互相流转，没有产生 L2->L2 转账的行为

## 治理提案流程

```mermaid
sequenceDiagram 
    Proposer -->> GovContract: send propose tx
    GovContract -->>+ ERC20Vote: getVotes
    ERC20Vote -->>- GovContract: Return
    alt check proposalThreshold success
    GovContract -->> GovContract: Get GovernanceParams
    GovContract -->> GovContract: snapshot = block.number + voteDelay, deadline = snapshot + votePeriod
    GovContract -->> GovContract: set the proposal state to Pending
    GovContract -->> GovContract: store proposal
    GovContract -->> GovContract: emit ProposalCreated
    else
    GovContract -->> Propsoser: Return
    end
```

#### 治理投票流程

```mermaid
sequenceDiagram 
    Proposer -->> GovContract: send castVote tx
    GovContract -->> GovContract: get proposal
    alt block.number > snapshot 
        GovContract -->> GovContract: set the proposal state to Active
    else block.number > deadline
        alt vote is not enough
            GovContract -->> GovContract: set the proposal state to Defeated
        end
    end
    alt had voted
        GovContract -->> Proposer: Return
    end
    alt proposal state is active
        GovContract -->> GovContract: count vote
        alt vote is enough
            GovContract -->> GovContract: set the proposal state to Succeeded
        end
        GovContract -->> GovContract: emit VoteCast
    end
```

#### 治理执行流程

```mermaid
sequenceDiagram 
    Sender -->> GovContract: send execute tx
    GovContract -->> GovContract: get proposal
    
    alt proposal state is Succeeded 
        GovContract -->> GovContract: set the proposal state to Executed
        GovContract -->> GovContract: emit ProposalExecuted
        alt is L1 proposal
            GovContract -->> StateSender: syncState
        else
            GovContract -->> Target: Call Target contract
        end
    else
        GovContract -->> Proposer: Return
    end

```

#### 提案取消流程


```mermaid
sequenceDiagram 
    Sender -->> GovContract: send cancel tx
    GovContract -->> GovContract: get proposal
    GovContract -->> ERC20Vote: get votes
    alt sender is proposer or the vote weight of the proposer less than threshold
        GovContract -->> GovContract: set the proposal state to Canceled
        GovContract -->> GovContract: emit ProposalCanceled
    end
```

