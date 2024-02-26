package wrap

import (
	"fmt"

	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
)

const (
	ElectionBase = 43
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

func ElectionValidatorWithVRF(db sdk.StateDB, vrfModule staketypes.VRFModuler, validatorSnapshotQueue staketypes.ValidatorSortSnapshotQueue, blockNumber, vrfElectionSize uint64) (staketypes.ValidatorSortSnapshotQueue, error) {

	// ### NOTE ###
	//
	// At any time, it is necessary to ensure that the number of blocks in the round is greater than the number of validators in the epoch,
	// otherwise VRF elections will have insufficient historical VRF nonces, leading to election failure (especially during parameter governance)
	//
	// (the validator snapshot queue (validatorSnapshotQueue) is definitely smaller than the number of validators in the epoch)
	//
	// NOTE: the size of diffValidatorSnapshotQueue may be zero
	historyNonceQueue, err := vrfModule.GetNonceQueueFromTail(db, blockNumber-1, uint64(len(validatorSnapshotQueue)))
	if nil != err {
		return nil, err
	}
	if len(historyNonceQueue) != len(validatorSnapshotQueue) {
		return nil, fmt.Errorf("had not enough history vrf nonces")
	}

	currentVRFNonce, err := vrfModule.GetCurrentNonce(db, blockNumber)
	if nil != err {
		return nil, err
	}

	return electionByProbability(validatorSnapshotQueue, currentVRFNonce, historyNonceQueue, blockNumber, vrfElectionSize)
}

func electionByProbability(validatorSnapshotQueue staketypes.ValidatorSortSnapshotQueue, currentVRFNonce common.Hash, historyVRFNonceQueue []common.Hash, blockNumber, vrfElectionSize uint64) (staketypes.ValidatorSortSnapshotQueue, error) {
	if currentVRFNonce == common.ZeroHash || len(historyVRFNonceQueue) == 0 || len(validatorSnapshotQueue) != len(historyVRFNonceQueue) {
		return nil, fmt.Errorf("invalid params")
	}
	totalWeights := common.Big0
	totalSqrtWeights := common.Big0
	svqueue := make(sortValidatorQueue, 0)
	for _, snap := range validatorSnapshotQueue {

		weights := new(big.Int).Div(snap.Shares(), new(big.Int).SetUint64(1e18))
		totalWeights.Add(totalWeights, weights)
		weights = new(big.Int).Sqrt(weights)
		totalSqrtWeights.Add(totalSqrtWeights, weights)

		sv := &sortValidator{
			v:       snap,
			weights: int64(weights.Uint64()),
		}
		svqueue = append(svqueue, sv)
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

	// Generate reference values for binomial curves
	p := float64(ElectionBase) / totalSqrtWeightsFloat

	// Nothing special, just to get a value that everyone is the same
	shuffleSeed := new(big.Int).SetBytes(historyVRFNonceQueue[0].Bytes()).Int64()

	log.Debug("Call electionByProbability Basic parameter", "blockNumber", blockNumber, "validatorSnapshotQueue size", len(validatorSnapshotQueue),
		"p", p, "totalWeights", totalWeightsFloat, "totalSqrtWeightsFloat", totalSqrtWeightsFloat, "vrfElectionSize", vrfElectionSize, "shuffleSeed", shuffleSeed)

	// rand shuffle validator snapshot queue
	rd := rand.New(rand.NewSource(shuffleSeed))
	rd.Shuffle(len(svqueue), func(i, j int) {
		svqueue[i], svqueue[j] = svqueue[j], svqueue[i]
	})

	for index, sv := range svqueue {

		// Take the historical vrf nonces in order from the nearest to the farthest,
		// and perform XOR operations on each current vrf nonce to obtain the pseudo-random value
		// of the validator in the corresponding index's snapshot queue.
		// Based on this value and their respective binomial curves, calculate the target value.
		resultStr := new(big.Int).Xor(new(big.Int).SetBytes(currentVRFNonce.Bytes()), new(big.Int).SetBytes(historyVRFNonceQueue[index].Bytes())).Text(10)
		target, err := strconv.ParseFloat(resultStr, 64)
		if nil != err {
			return nil, err
		}
		// Each verifier has their own target probability
		targetP := target / maxValue
		// Each validator generates their own binomial curve using reference values and their respective weights.
		bd := math.NewBinomialDistribution(sv.weights, p)
		// The difference between the binomial curve and the target probability.
		x, err := bd.InverseCumulativeProbability(targetP)
		if nil != err {
			return nil, err
		}
		sv.x = x

		log.Debug("Call electionByProbability calculated probability", "validatorAddr", sv.v.ValidatorAddr.Hex(), "index", index, "currentVRFNonce",
			currentVRFNonce.Hex(), "previousNonce", historyVRFNonceQueue[index].Hex(),
			"target", target, "targetP", targetP, "weight", sv.weights, "x", x)
	}

	validatorSnapshotVRFQueue := make(staketypes.ValidatorSortSnapshotQueue, vrfElectionSize)

	log.Debug("Call electionByProbability sort probability queue", "blockNumber", blockNumber, "queue", svqueue)

	sort.Sort(svqueue)

	for index, sv := range svqueue {
		if index == int(vrfElectionSize) {
			break
		}
		validatorSnapshotVRFQueue[index] = sv.v
	}

	log.Debug("Finished electionByProbability", "blockNumber", blockNumber, "validatorSnapshotVRFQueue", validatorSnapshotVRFQueue)

	return validatorSnapshotVRFQueue, nil
}

// ----------------------------

func ShuffleQueue(db sdk.StateDB, vrfModule staketypes.VRFModuler, currentRoundValidatorSnapshotQueue, validatorSnapshotVRFQueue staketypes.ValidatorSortSnapshotQueue, blockNumber, maxRoundValidatorsSize uint64) (staketypes.ValidatorSortSnapshotQueue, error) {

	currentSize := uint64(len(currentRoundValidatorSnapshotQueue))
	totalQueue := append(currentRoundValidatorSnapshotQueue, validatorSnapshotVRFQueue...)
	maxRoundShiftValidatorSize := (maxRoundValidatorsSize - 1) / 3

	for currentSize > maxRoundValidatorsSize-maxRoundShiftValidatorSize && uint64(len(totalQueue)) > maxRoundValidatorsSize {
		totalQueue = totalQueue[1:]
		currentSize--
	}

	if uint64(len(totalQueue)) > maxRoundValidatorsSize {
		totalQueue = totalQueue[:maxRoundValidatorsSize]
	}

	nextQueue := make(staketypes.ValidatorSortSnapshotQueue, len(totalQueue))

	copy(nextQueue, totalQueue)

	// Divide all consensus nodes into two groups, the front and back positions of each group are not changed,
	// but random ordering is performed in each group
	// The first group: the first f nodes
	// The second group: the last 2f + 1 nodes
	nextQueue, err := orderValidatorQueueByRandom(db, vrfModule, nextQueue, blockNumber, maxRoundShiftValidatorSize)
	if nil != err {
		return nil, err
	}
	return nextQueue, nil
}

type randomOrderValidator struct {
	validator *staketypes.ValidatorSortSnapshot
	value     *big.Int
}
type randomOrderValidatorQueue []*randomOrderValidator

func (r randomOrderValidatorQueue) Len() int {
	return len(r)
}

func (r randomOrderValidatorQueue) Less(i, j int) bool {
	return r[i].value.Cmp(r[j].value) > 0
}

func (r randomOrderValidatorQueue) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Randomly sort nodes
func orderValidatorQueueByRandom(db sdk.StateDB, vrfModule staketypes.VRFModuler, validatorSnapshotQueue staketypes.ValidatorSortSnapshotQueue, blockNumber, maxRoundShiftValidatorSize uint64) (staketypes.ValidatorSortSnapshotQueue, error) {

	historyNonceQueue, err := vrfModule.GetNonceQueueFromTail(db, blockNumber-1, uint64(len(validatorSnapshotQueue)))
	if nil != err {
		return nil, fmt.Errorf("GetNonceQueueFromTail, %s", err)
	}
	if len(historyNonceQueue) != len(validatorSnapshotQueue) {
		return nil, fmt.Errorf("had not enough history vrf nonces")
	}

	if len(validatorSnapshotQueue) <= int(maxRoundShiftValidatorSize) {
		return validatorSnapshotQueue, nil
	}

	orderQueue := make(randomOrderValidatorQueue, len(validatorSnapshotQueue))
	for i, snap := range validatorSnapshotQueue {
		value := new(big.Int).Xor(new(big.Int).SetBytes(snap.ValidatorAddr.Bytes()), new(big.Int).SetBytes(historyNonceQueue[i][:common.AddressLength]))
		orderQueue[i] = &randomOrderValidator{
			validator: snap,
			value:     value,
		}
		log.Debug("Call orderValidatorQueueByRandom xor", "validatorAddr", snap.ValidatorAddr.Hex(), "vrf nonce", historyNonceQueue[i].Hex(), "xor value", value)
	}

	frontPart := orderQueue[:maxRoundShiftValidatorSize]
	backPart := orderQueue[maxRoundShiftValidatorSize:]

	sort.Sort(frontPart)
	sort.Sort(backPart)

	orderQueue = make(randomOrderValidatorQueue, 0)
	orderQueue = append(orderQueue, frontPart...)
	orderQueue = append(orderQueue, backPart...)

	resultQueue := make(staketypes.ValidatorSortSnapshotQueue, len(orderQueue))
	for i, v := range orderQueue {
		resultQueue[i] = v.validator
	}
	log.Debug("Succeed call orderValidatorQueueByRandom", "blockNumber", blockNumber, "resultQueueSize", len(resultQueue))
	return resultQueue, nil
}
