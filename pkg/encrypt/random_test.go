package encrypt

import (
	"math/rand"
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	seed := int64(0)
	rng := getRandomNumberGenerator(seed)
	count := 5
	zero_seed_randoms := generateRandomNumbers(rng, count)

	rng2 := rand.New(rand.NewSource(seed))

	for i := 0; i < count; i++ {
		rnd2 := rng2.Intn(RAND_MAX)
		if zero_seed_randoms[i] != rnd2 {
			t.Errorf("%d: %d != %d", i, zero_seed_randoms[i], rnd2)
		}
	}
}
