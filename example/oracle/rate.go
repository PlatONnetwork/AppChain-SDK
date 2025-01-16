package oracle

import "time"

type CurrencyPair int

var (
	USDCNY = CurrencyPair(0)
)

type RateMarketClient interface {
	Rate(CurrencyPair) (uint64, error)
}

type FixedRate struct {
	Value uint64
}

func (f FixedRate) Rate(CurrencyPair) (uint64, error) {
	return f.Value, nil
}

type TimerRate struct {
	interval time.Duration
	start    time.Time
	Value    uint64
}

func NewTimerRate(interval time.Duration) *TimerRate {
	return &TimerRate{
		interval: interval,
		start:    time.Now(),
		Value:    0,
	}

}

func (f *TimerRate) Rate(CurrencyPair) (uint64, error) {
	now := time.Now()
	if now.Sub(f.start) > f.interval {
		f.Value++
		f.start = now
	}
	return f.Value, nil
}
