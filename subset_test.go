package load_balancing

import (
	"fmt"
	"math/rand"
	"testing"
)

func setupIntSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}

func setupStringSlice(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = fmt.Sprintf("%d", i)
	}
	return slice
}

func calculateDistribution(subsetList [][]string) map[string]int {
	subsetMap := make(map[string]int)
	for _, list := range subsetList {
		for _, v := range list {
			if _, ok := subsetMap[v]; !ok {
				subsetMap[v] = 1
			} else {
				subsetMap[v]++
			}
		}
	}
	return subsetMap
}

func TestSubset(t *testing.T) {
	backends := setupStringSlice(10)
	clientId := setupIntSlice(10)
	subsetSize := 3
	subsetList := [][]string{}

	for _, v := range clientId {
		subset := Subset(backends, v, subsetSize)
		subsetList = append(subsetList, subset)
	}

	for i := 0; i < 5; i++ {
		subsetMap := calculateDistribution(subsetList)
		t.Log(subsetMap)
	}

}

func TestRandomShuffle(t *testing.T) {
	round := 2
	backends := setupStringSlice(10)
	rand.Seed(int64(round))
	rand.Shuffle(len(backends), func(i, j int) {
		backends[i], backends[j] = backends[j], backends[i]
	})
	t.Log(backends)
}
