package types

type EpochItem struct {
	StartBlock uint64
	EndBlock   uint64
	RoundCount uint64
}

func NewEpochItem(startBlock, endBlock, roundCount uint64) *EpochItem {
	return &EpochItem{
		StartBlock: startBlock,
		EndBlock:   endBlock,
		RoundCount: roundCount,
	}
}

func (item *EpochItem) IsEmpty() bool {
	return nil == item
}

func (item *EpochItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type EpochQueue []*EpochItem

func NewEpochQueue(size uint64) EpochQueue {
	queue := make(EpochQueue, size)
	return queue
}

func (queue EpochQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue EpochQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

type RoundItem struct {
	StartBlock uint64
	EndBlock   uint64
}

func NewRoundItem(startBlock, endBlock uint64) *RoundItem {
	return &RoundItem{
		StartBlock: startBlock,
		EndBlock:   endBlock,
	}
}

func (item *RoundItem) IsEmpty() bool {
	return nil == item
}

func (item *RoundItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type RoundQueue []*RoundItem

func NewRoundQueue(size uint64) RoundQueue {
	queue := make(RoundQueue, size)
	return queue
}

func (queue RoundQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue RoundQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}
