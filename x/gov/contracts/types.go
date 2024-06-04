package contracts

import "math/big"

const (
	Pending = uint8(iota)
	Active
	Canceled
	Defeated
	Succeeded
	Queued
	Expired
	Executed
)
const (
	Against = uint8(iota)
	For
	Abstain
)

func StateToString(state uint8) string {
	str := "Unknown"
	switch state {
	case Pending:
		str = "Pending"
	case Active:
		str = "Active"
	case Canceled:
		str = "Canceled"
	case Defeated:
		str = "Defeated"
	case Succeeded:
		str = "Succeeded"
	case Queued:
		str = "Queued"
	case Expired:
		str = "Expired"
	case Executed:
		str = "Executed"
	}
	return str
}

func VoteToString(vote uint8) string {
	str := "Unknown"
	switch vote {
	case Against:
		str = "Against"
	case For:
		str = "For"
	case Abstain:
		str = "Abstain"
	}
	return str
}

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
	AgainstVotes big.Int
	ForVotes     big.Int
	AbstainVotes big.Int
}
