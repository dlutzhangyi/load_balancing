package load_balancing

import (
	"math/rand"
)

func Subset(backends []string, clientId int,subsetSize int) []string{
	subSetCount := len(backends) / subsetSize

	round := clientId / subSetCount
	rand.Seed(int64(round))
	rand.Shuffle(len(backends), func(i, j int) {
		backends[i],backends[j]= backends[j],backends[i]
	})

	subsetId:= clientId%subSetCount
	start := subsetId * subsetSize

	return backends[start:start+subsetSize]
}