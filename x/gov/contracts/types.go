package contracts

import "math/big"

const (
	Pending = iota
	Active
	Canceled
	Defeated
	Succeeded
	Queued
	Expired
	Executed
)
const (
	Against = iota
	For
	Abstain
)

type BlockNumber uint64

func (b *BlockNumber) isExpired(number BlockNumber) bool {
	return b.isStarted() && *b <= number
}

func (b BlockNumber) getDeadline() *big.Int {
	return new(big.Int).SetUint64(uint64(b))
}

func (b *BlockNumber) reset() {
	*b = 0
}

func (b *BlockNumber) isStarted() bool {
	return *b > 0
}

func (b *BlockNumber) isPending(number BlockNumber) bool {
	return *b > number
}

type ProposalVote struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}
