package world

import (
    "math/rand"
    "reflect"
    "testing"
)

func TestSeedRandomChangesGeneratorState(t *testing.T) {
    rand.Seed(1)
    before := make([]int, 10)
    for i := range before {
        before[i] = rand.Intn(1_000_000)
    }

    SeedRandom()

    after := make([]int, 10)
    for i := range after {
        after[i] = rand.Intn(1_000_000)
    }

    if reflect.DeepEqual(before, after) {
        t.Fatal("SeedRandom() did not reseed the RNG")
    }
}
