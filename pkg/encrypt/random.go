package encrypt

import (
	"math/rand"
)

var memoizedRandomNumberGenerators map[int64]*rand.Rand
var RAND_MAX = 10000

func init() {
	memoizedRandomNumberGenerators = make(map[int64]*rand.Rand)
}

func getRandomNumberGenerator(seed int64) *rand.Rand {
	if _, ok := memoizedRandomNumberGenerators[seed]; !ok {
		// Initialize the random number generator with the seed
		memoizedRandomNumberGenerators[seed] = rand.New(rand.NewSource(seed))
	}

	return memoizedRandomNumberGenerators[seed]
}

func generateRandomNumbers(rng *rand.Rand, count int) []int {
	randomNumbers := make([]int, count)
	for i := 0; i < count; i++ {
		randomNumbers[i] = rng.Intn(10000)
	}

	return randomNumbers
}

func generateRandomNumber(rng *rand.Rand) int {
	return rng.Intn(RAND_MAX)
}

func GetNextRandomNumber(seed int64) int {
	return generateRandomNumber(getRandomNumberGenerator(seed))
}
