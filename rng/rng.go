package rng

import (
	"fmt"
	"math"
)

// BuildReel creates a Reel for a given index, by taking the symbols and
// their weights at the given index and converting them into a probability
// distribution. The Reel is then filled according to this distribution.
//
// The process is as follows:
//
//  1. Create an array of the same length as the number of symbols, and fill
//     it with the weights of the symbols at the given reel index.
//  2. Convert the weights into a probability distribution (i.e. all the
//     weights will add up to 100).
//  3. Iterate over the probability distribution and fill the Reel with
//     symbols according to their probabilities.
func BuildReel(symbols []Symbol, reelIndex int) (r Reel) {
	allWeightsOnReel := make([]int, len(symbols))

	for i, symbol := range symbols {
		allWeightsOnReel[i] = symbol.WeightsPerReel[reelIndex]
	}

	probabilities := convertWeightsToProbabilities(allWeightsOnReel)

	fmt.Printf("Converted weights %v to probabilities %v\n", allWeightsOnReel, probabilities)

	writePtr := 0
	for i, probability := range probabilities {
		for range probability {
			if writePtr == len(Reel{}) {
				return r
			}
			r[writePtr] = symbols[i]
			writePtr++
		}
	}

	return r
}

func convertWeightsToProbabilities(weights []int) []int {
	totalWeight := 0

	for _, weight := range weights {
		totalWeight += weight
	}

	probabilities := make([]int, len(weights))

	for i, weight := range weights {
		chance := float64(weight) / float64(totalWeight)
		probabilities[i] = int(math.Round(chance * 100))
	}

	return probabilities
}
