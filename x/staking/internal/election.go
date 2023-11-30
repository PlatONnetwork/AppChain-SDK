package internal

import (
	"encoding/hex"
	"github.com/AlayaNetwork/Alaya-Go/common"
	stakecommon "github.com/PlatONnetwork/AppChain-SDK/x/staking/common"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	vrfInternal "github.com/PlatONnetwork/AppChain-SDK/x/vrf/internal"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/PlatONnetwork/PlatON-Go/x/handler"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
	"math/rand"
	"strconv"
)

type sortValidator struct {
	v       *staketypes.ValidatorSortSnapshot
	x       int64
	weights int64
}

type sortValidatorQueue []*sortValidator

func (svs sortValidatorQueue) Len() int {
	return len(svs)
}

func (svs sortValidatorQueue) Less(i, j int) bool {
	return svs[i].x > svs[j].x
}

func (svs sortValidatorQueue) Swap(i, j int) {
	svs[i], svs[j] = svs[j], svs[i]
}

func ElectionValidatorByVRF(db sdk.StateDB, address basecommon.Address, validatorSnapshotQueue staketypes.ValidatorSharesSnapshotQueue, shiftSize int, nonce []byte, parentHash common.Hash, blockNumber uint64) (staketypes.ValidatorSharesSnapshotQueue, error) {

	// ### NOTE ###
	//
	// At any time, it is necessary to ensure that the number of blocks in the round is greater than the number of validators in the epoch, otherwise VRF elections will have insufficient historical VRF nonces, leading to election failure (especially during parameter governance)
	preNonces, err := vrfInternal.GetNonceQueueUtil(db, address, blockNumber, stakecommon.MAX_EPOCH_VALIDATORS_SIZE)
	if nil != err {
		return nil, err
	}
	if len(preNonces) < len(validatorSnapshotQueue) {
		log.Error("Failed to ElectionValidatorByVRF on Election", "blockNumber", blockNumber, "validatorListSize", len(validatorSnapshotQueue),
			"nonceSize", len(nonce), "preNoncesSize", len(preNonces), "parentHash", hex.EncodeToString(parentHash.Bytes()))
		return nil, staking.ErrWrongFuncParams
	}
	if len(preNonces) > len(validatorSnapshotQueue) {
		preNonces = preNonces[len(preNonces)-len(validatorSnapshotQueue):]
	}
	return probabilityElection(validatorSnapshotQueue, shiftSize, vrf.ProofToHash(nonce), preNonces, blockNumber, copernicus)
}

func probabilityElection(validatorList staking.ValidatorQueue, shiftLen int, currentNonce []byte, preNonces [][]byte, blockNumber uint64, copernicus bool) (staking.ValidatorQueue, error) {
	if len(currentNonce) == 0 || len(preNonces) == 0 || len(validatorList) != len(preNonces) {
		log.Error("Failed to probabilityElection", "blockNumber", blockNumber, "copernicus", copernicus, "validators Size", len(validatorList),
			"currentNonceSize", len(currentNonce), "preNoncesSize", len(preNonces))
		return nil, staking.ErrWrongFuncParams
	}
	totalWeights := new(big.Int)
	totalSqrtWeights := new(big.Int)
	svList := make(sortValidatorQueue, 0)
	for _, val := range validatorList {

		weights := new(big.Int).Div(val.Shares, new(big.Int).SetUint64(1e18))
		totalWeights.Add(totalWeights, weights)
		weights = new(big.Int).Sqrt(weights)
		totalSqrtWeights.Add(totalSqrtWeights, weights)

		sv := &sortValidator{
			v:           val,
			weights:     int64(weights.Uint64()),
			version:     val.ProgramVersion,
			blockNumber: val.StakingBlockNum,
			txIndex:     val.StakingTxIndex,
		}
		svList = append(svList, sv)
	}
	var maxValue float64 = (1 << 256) - 1
	totalWeightsFloat, err := strconv.ParseFloat(totalWeights.Text(10), 64)
	if nil != err {
		return nil, err
	}
	totalSqrtWeightsFloat, err := strconv.ParseFloat(totalSqrtWeights.Text(10), 64)
	if nil != err {
		return nil, err
	}

	var p float64
	if copernicus {
		p = xcom.CalcPV110(totalSqrtWeightsFloat)
	} else {
		p = xcom.CalcP(totalWeightsFloat, totalSqrtWeightsFloat)
	}

	shuffleSeed := new(big.Int).SetBytes(preNonces[0]).Int64()
	log.Debug("Call probabilityElection Basic parameter on Election", "blockNumber", blockNumber, "copernicus", copernicus, "validatorListSize", len(validatorList),
		"p", p, "totalWeights", totalWeightsFloat, "totalSqrtWeightsFloat", totalSqrtWeightsFloat, "shiftValidatorNum", shiftLen, "shuffleSeed", shuffleSeed)

	if copernicus {
		rd := rand.New(rand.NewSource(shuffleSeed))
		rd.Shuffle(len(svList), func(i, j int) {
			svList[i], svList[j] = svList[j], svList[i]
		})
	}

	for index, sv := range svList {
		resultStr := new(big.Int).Xor(new(big.Int).SetBytes(currentNonce), new(big.Int).SetBytes(preNonces[index])).Text(10)
		target, err := strconv.ParseFloat(resultStr, 64)
		if nil != err {
			return nil, err
		}
		targetP := target / maxValue
		bd := math.NewBinomialDistribution(sv.weights, p)
		x, err := bd.InverseCumulativeProbability(targetP)
		if nil != err {
			return nil, err
		}
		sv.x = x

		log.Debug("Call probabilityElection, calculated probability on Election", "nodeId", sv.v.NodeId.TerminalString(),
			"addr", sv.v.NodeAddress.Hex(), "index", index, "currentNonce",
			hex.EncodeToString(currentNonce), "preNonce", hex.EncodeToString(preNonces[index]),
			"target", target, "targetP", targetP, "weight", sv.weights, "x", x, "version", sv.version,
			"blockNumber", sv.blockNumber, "txIndex", sv.txIndex)
	}

	vrfQueue := make(staking.ValidatorQueue, shiftLen)

	log.Debug("Call probabilityElection, sort probability queue", "blockNumber", blockNumber, "copernicus", copernicus, "list", svList)

	if copernicus {
		sort.Sort(newSortValidatorQueue(svList))
	} else {
		sort.Sort(svList)
	}
	for index, sv := range svList {
		if index == shiftLen {
			break
		}
		vrfQueue[index] = sv.v
	}

	log.Debug("Call probabilityElection finished", "blockNumber", blockNumber, "copernicus", copernicus, "vrfQueue", vrfQueue)

	return vrfQueue, nil
}

// ----------------------------

func shuffleQueue(remainCurrQueue, vrfQueue staking.ValidatorQueue, blockNumber uint64, parentHash common.Hash) (staking.ValidatorQueue, error) {

	remainLen := len(remainCurrQueue)
	totalQueue := append(remainCurrQueue, vrfQueue...)

	for remainLen > int(xcom.MaxConsensusVals()-xcom.ShiftValidatorNum()) && len(totalQueue) > int(xcom.MaxConsensusVals()) {
		totalQueue = totalQueue[1:]
		remainLen--
	}

	if len(totalQueue) > int(xcom.MaxConsensusVals()) {
		totalQueue = totalQueue[:xcom.MaxConsensusVals()]
	}

	next := make(staking.ValidatorQueue, len(totalQueue))

	copy(next, totalQueue)

	// Divide all consensus nodes into two groups, the front and back positions of each group are not changed,
	// but random ordering is performed in each group
	// The first group: the first f nodes
	// The second group: the last 2f + 1 nodes
	next, err := randomOrderValidatorQueue(blockNumber, parentHash, next)
	if nil != err {
		return nil, err
	}
	return next, nil
}

type randomOrderValidator struct {
	validator *staking.Validator
	value     *big.Int
}
type randomOrderValidatorList []*randomOrderValidator

func (r randomOrderValidatorList) Len() int {
	return len(r)
}

func (r randomOrderValidatorList) Less(i, j int) bool {
	return r[i].value.Cmp(r[j].value) > 0
}

func (r randomOrderValidatorList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Randomly sort nodes
func randomOrderValidatorQueue(blockNumber uint64, parentHash common.Hash, queue staking.ValidatorQueue) (staking.ValidatorQueue, error) {
	preNonces, err := handler.GetVrfHandlerInstance().Load(parentHash)
	if nil != err {
		return nil, err
	}
	if len(preNonces) < len(queue) {
		log.Error("Failed to randomOrderValidatorQueue on Election", "blockNumber", blockNumber, "validatorListSize", len(queue),
			"preNoncesSize", len(preNonces), "parentHash", parentHash.TerminalString())
		return nil, staking.ErrWrongFuncParams
	}
	if len(preNonces) > len(queue) {
		preNonces = preNonces[len(preNonces)-len(queue):]
	}

	if len(queue) <= int(xcom.ShiftValidatorNum()) {
		return queue, nil
	}

	orderList := make(randomOrderValidatorList, len(queue))
	for i, v := range queue {
		value := new(big.Int).Xor(new(big.Int).SetBytes(v.NodeAddress.Bytes()), new(big.Int).SetBytes(preNonces[i][:common.AddressLength]))
		orderList[i] = &randomOrderValidator{
			validator: v,
			value:     value,
		}
		log.Debug("Call randomOrderValidatorQueue xor", "nodeId", v.NodeId.TerminalString(), "nodeAddress", v.NodeAddress.Hex(), "nonce", hexutil.Encode(preNonces[i]), "xorValue", value)
	}

	frontPart := orderList[:xcom.ShiftValidatorNum()]
	backPart := orderList[xcom.ShiftValidatorNum():]

	sort.Sort(frontPart)
	sort.Sort(backPart)

	orderList = make(randomOrderValidatorList, 0)
	orderList = append(orderList, frontPart...)
	orderList = append(orderList, backPart...)

	resultQueue := make(staking.ValidatorQueue, len(orderList))
	for i, v := range orderList {
		resultQueue[i] = v.validator
	}
	log.Debug("Call randomOrderValidatorQueue success", "blockNumber", blockNumber, "parentHash", parentHash.TerminalString(), "resultQueueSize", len(resultQueue))
	return resultQueue, nil
}
