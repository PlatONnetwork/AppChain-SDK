package main

import "gopkg.in/urfave/cli.v1"

var (
	DeployKeyFlag = cli.StringFlag{
		Name:   "deploykey",
		Usage:  "Deploy contract key",
		EnvVar: "DEPLOY_KEY",
	}
	InitializeKeyFlag = cli.StringFlag{
		Name:   "initializeKey",
		Usage:  "Initialize registry manager key",
		EnvVar: "INITIALIZE_KEY",
	}

	RegistryManagerOwnerFlag = cli.StringFlag{
		Name:   "registrymanagerowner",
		Usage:  "RegistryManager contract owner",
		EnvVar: "REGISTRY_MANAGER_OWNER",
	}
	TemplateResultFileFlag = cli.StringFlag{
		Name:   "templateresultfile",
		Usage:  "Template contracts deploy result file name",
		EnvVar: "TEMPLATE_RESULT_FILE",
		Value:  "template.toml",
	}
	RootChianUrlFlag = cli.StringFlag{
		Name:   "rootchainurl",
		Usage:  "Root chain url",
		EnvVar: "ROOTCHAIN_URL",
	}
	OutputFlag = cli.StringFlag{
		Name:   "output",
		Usage:  "Output result directory",
		EnvVar: "OUTPUT",
		Value:  ".",
	}

	GasLimitFlag = cli.Uint64Flag{
		Name:   "gaslimit",
		EnvVar: "GASLIMIT",
		Usage:  "Transaction gaslimit",
		Value:  8000000,
	}
	GasPriceFlag = cli.Uint64Flag{
		Name:   "gasprice",
		EnvVar: "GASPRICE",
		Usage:  "Transaction gasprice",
		Value:  1000000000,
	}
	ChildChainIdFlag = cli.Uint64Flag{
		Name:   "childchainid",
		EnvVar: "CHILDCHAINID",
		Usage:  "Child chainid",
		Value:  123,
	}
	RootChainIdFlag = cli.Uint64Flag{
		Name:   "rootchainid",
		EnvVar: "ROOTCHAINID",
		Usage:  "Root chainid",
		Value:  123083,
	}
	StakeOwnerFlag = cli.StringFlag{
		Name:   "stakeowner",
		Usage:  "Stake owner",
		EnvVar: "STAKE_OWNER",
		Value:  "0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c",
	}
	StakeAmountFlag = cli.StringFlag{
		Name:   "stakeamount",
		Usage:  "Stake amount",
		EnvVar: "STAKE_AMOUNT",
		Value:  "1000000000",
	}
	CommissionRateFlag = cli.StringFlag{
		Name:   "commissionrate",
		Usage:  "Commission rate",
		EnvVar: "COMMISSION_RATE",
		Value:  "0",
	}
	MinStakeFlag = cli.StringFlag{
		Name:   "minstake",
		EnvVar: "MINSTAKE",
		Usage:  "Child chain min stake",
		Value:  "1000",
	}

	MinDelegateFlag = cli.StringFlag{
		Name:   "mindelegate",
		EnvVar: "MINDELEGATE",
		Usage:  "Child chain min delegate",
		Value:  "10",
	}
	ChildchainOwnerFlag = cli.StringFlag{
		Name:   "childchainowner",
		Usage:  "Child chain owner private key",
		EnvVar: "CHILDCHAIN_OWNER",
	}
	ChildchainTokenFlag = cli.StringFlag{
		Name:   "childchaintoken",
		Usage:  "Child chain owner",
		EnvVar: "CHILDCHAIN_TOKEN",
	}
	ChildchainStakeFlag = cli.StringFlag{
		Name:   "childchainostake",
		Usage:  "Child chain stake contract address",
		EnvVar: "CHILDCHAIN_STAKE",
		Value:  "0x1000000000000000000000000000000000000005",
	}
	ChildchainDepositFlag = cli.StringFlag{
		Name:   "childchaindeposit",
		Usage:  "Child chain deposit contract address",
		EnvVar: "CHILDCHAIN_DEPOSIT",
		Value:  "0x1000000000000000000000000000000000000007",
	}

	ChildchainWithdrawFlag = cli.StringFlag{
		Name:   "childchainwithdraw",
		Usage:  "Child chain withdraw contract address",
		EnvVar: "CHILDCHAIN_WITHDRAW",
		Value:  "0x1000000000000000000000000000000000000008",
	}

	ChildchainGovernanceFlag = cli.StringFlag{
		Name:   "childchaingovernance",
		Usage:  "Child chain governance contract address",
		EnvVar: "CHILDCHAIN_GOVERNANCE",
		Value:  "0x100000000000000000000000000000000000000A",
	}

	ChildchainIssuanceOptionFlag = cli.IntFlag{
		Name:   "childchainissuance",
		Usage:  "Child chain token issuance option",
		EnvVar: "CHILDCHAIN_ISSUANCE",
		Value:  1,
	}

	ChildchainContractFileFlag = cli.StringFlag{
		Name:   "childchaincontractfile",
		Usage:  "Child chain contract file name",
		EnvVar: "CHILDCHAIN_CONTRACT_FILE",
		Value:  "childchaincontract.toml",
	}

	CheckpointKeystoreFileFlag = cli.StringFlag{
		Name:   "checkpointkeystorefilename",
		Usage:  "Checkpoint keystore file name",
		EnvVar: "CHECKPOINT_KEYSTORE_FILE_NAME",
		Value:  "l1checkpointsender",
	}
	L2TxKeystoreFileFlag = cli.StringFlag{
		Name:   "l2txkeystorefilename",
		Usage:  "L2 transaction keystore file name",
		EnvVar: "L2TX_KEYSTORE_FILE_NAME",
		Value:  "l2txsender",
	}

	CheckpointKeystoreFlag = cli.StringFlag{
		Name:   "checkpointkeystore",
		Usage:  "Checkpoint keystore file name",
		EnvVar: "CHECKPOINT_KEYSTORE",
	}
	CheckpointKeystorePasswordFlag = cli.StringFlag{
		Name:   "checkpointkeystorepassword",
		Usage:  "Checkpoint keystore password file name",
		EnvVar: "CHECKPOINT_KEYSTORE_PASSWORD",
	}
	L2TxKeystoreFlag = cli.StringFlag{
		Name:   "l2txkeystore",
		Usage:  "L2 tx keystore file name",
		EnvVar: "L2TX_KEYSTORE",
	}
	L2TxKeystorePasswordFlag = cli.StringFlag{
		Name:   "l2txkeystorepassword",
		Usage:  "L2 tx keystore password file name",
		EnvVar: "L2TX_KEYSTORE_PASSWORD",
	}
	ExtraArgsFlag = cli.StringFlag{
		Name:   "extraargs",
		Usage:  "Childchain start extra args",
		EnvVar: "EXTRA_ARGS",
		Value:  "--asyncblock.computersenderthread 4 --asyncblock.entrysize 1000 --asyncblock.splitthreshold 3000 --asyncblock.concurrency_level 7 --asyncblock.txs_batch 512 --cbft.wal.disabled",
	}

	GenesisCbftPeriodFlag = cli.Uint64Flag{
		Name:   "cbft.period",
		Usage:  "CBFT params period",
		EnvVar: "CBFT_PERIOD",
		Value:  10000,
	}
	GenesisCbftAmountFlag = cli.Uint64Flag{
		Name:   "cbft.amount",
		Usage:  "CBFT params amount",
		EnvVar: "CBFT_amount",
		Value:  10,
	}
	GenesisStageRoundValidatorElectionDistanceFlag = cli.Uint64Flag{
		Name:   "stage.roundvalidatorElectionDistance",
		Usage:  "Stage params roundValidatorElectionDistance",
		EnvVar: "STAGE_ROUND_VALIDATOR_ELECTION_DISTANCE",
		Value:  20,
	}
	GenesisStageRoundSizeFlag = cli.Uint64Flag{
		Name:   "stage.roundsize",
		Usage:  "Stage params round size",
		EnvVar: "STAGE_ROUND_SIZE",
		Value:  250,
	}
	GenesisStageEpochSizeFlag = cli.Uint64Flag{
		Name:   "stage.epochsize",
		Usage:  "Stage params epoch size",
		EnvVar: "STAGE_EPOCH_SIZE",
		Value:  500,
	}
	GenesisVrfNonceFlag = cli.StringFlag{
		Name:   "vrf.nonce",
		Usage:  "VRF params nonce",
		EnvVar: "VRF_NONCE",
		Value:  "0x000000000000000000000000000000000067656e657369735652464e6f6e6365",
	}
	GenesisStakeWithdrawalWaitPeriodFlag = cli.Uint64Flag{
		Name:   "stake.stakewithdrawalwaitperiod",
		Usage:  "Stake params stakeWithdrawalWaitPeriod",
		EnvVar: "STAKE_WITHDRAWAL_WAIT_PERIOD",
		Value:  20,
	}
	GenesisStakeDelegateWithdrawalWaitPeriodFlag = cli.Uint64Flag{
		Name:   "stake.delegatewithdrawalwaitperiod",
		Usage:  "Stake params delegateWithdrawalWaitPeriod",
		EnvVar: "STAKE_DELEGATE_WITHDRAWAL_WAIT_PERIOD",
		Value:  6,
	}

	GenesisStakeSlashingPercentageFlag = cli.Uint64Flag{
		Name:   "stake.slashingpercentage",
		Usage:  "Stake params slashingPercentage",
		EnvVar: "STAKE_SLASHING_PERCENTAGE",
		Value:  50,
	}

	GenesisStakeSlashIncentivePercentageFlag = cli.Uint64Flag{
		Name:   "stake.slashincentivepercentage",
		Usage:  "Stake params slashIncentivePercentage",
		EnvVar: "STAKE_SLASH_INCENTIVE_PERCENTAGE",
		Value:  30,
	}

	GenesisStakeMaxRoundValidatorsSizeFlag = cli.Uint64Flag{
		Name:   "stake.maxroundvalidatorssize",
		Usage:  "Stake params maxRoundValidatorsSize",
		EnvVar: "STAKE_MAX_ROUND_VALIDATORS_SIZE",
		Value:  4,
	}

	GenesisStakeMaxEpochValidatorsSizeFlag = cli.Uint64Flag{
		Name:   "stake.maxepochvalidatorssize",
		Usage:  "Stake params maxEpochValidatorsSize",
		EnvVar: "STAKE_MAX_EPOCH_VALIDATORS_SIZE",
		Value:  201,
	}

	GenesisStakeMinBlocksOfRoundValidatorFlag = cli.Uint64Flag{
		Name:   "stake.minblocksofroundvalidator",
		Usage:  "Stake params minBlocksOfRoundValidator",
		EnvVar: "STAKE_MIN_BLOCKS_OF_ROUND_VALIDATOR",
		Value:  1,
	}
	GenesisRewardPerBlockFlag = cli.StringFlag{
		Name:   "reward.perblock",
		Usage:  "Reward params rewardPerBlock",
		EnvVar: "REWARD_PERBLOCK",
		Value:  "4000000000000000000",
	}
	GenesisRewardPerEpochFlag = cli.StringFlag{
		Name:   "reward.perepoch",
		Usage:  "Reward params rewardPerEpoch",
		EnvVar: "REWARD_PEREPOCH",
		Value:  "6000000000000000000",
	}
	GenesisVoteTokenNameFlag = cli.StringFlag{
		Name:   "votetoken.name",
		Usage:  "VoteToken params name",
		EnvVar: "VOTETOKEN_NAME",
		Value:  "name",
	}
	GenesisVoteTokenSymbolFlag = cli.StringFlag{
		Name:   "votetoken.symbol",
		Usage:  "VoteToken params symbol",
		EnvVar: "VOTETOKEN_SYMBOL",
		Value:  "VOTE",
	}
	GenesisVoteTokenVersionFlag = cli.StringFlag{
		Name:   "votetoken.version",
		Usage:  "VoteToken params version",
		EnvVar: "VOTETOKEN_VERSION",
		Value:  "1",
	}
	GenesisVoteTokenOwnerFlag = cli.StringFlag{
		Name:   "votetoken.owner",
		Usage:  "VoteToken params owner",
		EnvVar: "VOTETOKEN_OWNER",
		Value:  "0x1000000000000000000000000000000000000005",
	}
	GenesisGovNameFlag = cli.StringFlag{
		Name:   "gov.name",
		Usage:  "Gov params owner",
		EnvVar: "GOV_NAME",
		Value:  "gov",
	}
	GenesisGovVersionFlag = cli.StringFlag{
		Name:   "gov.version",
		Usage:  "Gov params name",
		EnvVar: "GOV_VERSION",
		Value:  "1",
	}
	GenesisGovVoteDelayFlag = cli.StringFlag{
		Name:   "gov.votedelay",
		Usage:  "Gov params voteDelay",
		EnvVar: "GOV_VOTE_DELAY",
		Value:  "1",
	}
	GenesisGovVotePeriodFlag = cli.StringFlag{
		Name:   "gov.voteperiod",
		Usage:  "Gov params votePeriod",
		EnvVar: "GOV_VOTE_PERIOD",
		Value:  "1",
	}
	GenesisGovQuorumNumeratorFlag = cli.StringFlag{
		Name:   "gov.quorumnumerator",
		Usage:  "Gov params quorumNumerator",
		EnvVar: "GOV_QUORUM_NUMERATOR",
		Value:  "1",
	}
	GenesisGovProposalThresholdFlag = cli.StringFlag{
		Name:   "gov.proposalthreshold",
		Usage:  "Gov params proposalThreshold",
		EnvVar: "GOV_PROPOSAL_THRESHOLD",
		Value:  "1",
	}
	GenesisGovOwnerFlag = cli.StringFlag{
		Name:   "gov.owner",
		Usage:  "Gov params owner",
		EnvVar: "GOV_OWNER",
		Value:  "0x0000000000000000000000000000000000000065",
	}
)
