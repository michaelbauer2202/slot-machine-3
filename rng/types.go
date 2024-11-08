package rng

import "silburyslot/randomorg"

type Symbol struct {
	SymbolId       int
	WeightsPerReel []int
	BetMultiplier  float64
}

type Reel [100]Symbol

func (reel Reel) PickRandomSymbol() Symbol {
	idx := randomorg.RandomIntegers(1, 0, len(reel)-1)[0]
	return reel[idx]
}

type SlotMachine []Reel
