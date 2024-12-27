package oracle

type RateClient interface {
	Rate() (uint64, error)
}

type FixedRate struct {
	Value uint64
}

func (f FixedRate) Rate() (uint64, error) {
	return f.Value, nil
}
