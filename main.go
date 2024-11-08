package main

import (
	"silburyslot/rng"
)

const (
	CHERRY = iota
	ORANGE
	PLUM
	BELL
	SEVEN
	CHOCOLATE
	BAR
)

func main() {
	loadEnvFile(".env")
	symbols := []rng.Symbol{
		{SymbolId: CHERRY, WeightsPerReel: []int{1, 0, 0}, BetMultiplier: 2},
		{SymbolId: ORANGE, WeightsPerReel: []int{3, 0, 0}, BetMultiplier: 4},
		{SymbolId: PLUM, WeightsPerReel: []int{0, 0, 0}, BetMultiplier: 6},
		{SymbolId: BELL, WeightsPerReel: []int{0, 0, 0}, BetMultiplier: 8},
		{SymbolId: SEVEN, WeightsPerReel: []int{0, 0, 0}, BetMultiplier: 10},
		{SymbolId: CHOCOLATE, WeightsPerReel: []int{0, 1, 0}, BetMultiplier: 12},
		{SymbolId: BAR, WeightsPerReel: []int{0, 0, 1}, BetMultiplier: 20},
	}

	slotMachine := []rng.Reel{}
	for i := range len(symbols[0].WeightsPerReel) {
		newReel := rng.BuildReel(symbols, i)
		slotMachine = append(slotMachine, newReel)
	}

	// for _, reel := range slotMachine {
	// 	for _, sym := range reel {
	// 		fmt.Printf("%d ", sym.SymbolId)
	// 	}
	// 	fmt.Println()
	// }
}
