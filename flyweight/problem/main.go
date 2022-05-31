package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/arthurshafikov/patterns/flyweight/problem/tree"
)

func main() {
	res := testing.Benchmark(BenchmarkMain)
	fmt.Printf("Memory allocations : %d \n", res.MemAllocs)
	fmt.Printf("Memory bytes : %d \n", res.MemBytes)
	fmt.Printf("Time taken: %s \n", res.T)
	/*
		Memory allocations : 10 003 738
		Memory bytes : 1 263 430 016
		Time taken: 1.700592369s
	*/
}

func BenchmarkMain(b *testing.B) {
	for j := 0; j < 100; j++ {
		CreateForest()
	}
}

func CreateForest() {
	rand.Seed(time.Now().Unix())
	forest := tree.NewForest()

	treeTypes := []string{
		"birch",
		"spruce",
		"oak",
	}

	treeColor := []string{
		"white",
		"brown",
	}

	treeTextures := [][]byte{}

	for i := 0; i < 5; i++ {
		newTexture := make([]byte, 1000*100) // 100 kilobyte
		rand.Read(newTexture)                // some jpeg or png

		treeTextures = append(treeTextures, newTexture)
	}

	for i := 0; i < 100_000; i++ {
		forest.PlantTree(
			rand.Intn(1000),
			rand.Intn(1000),
			treeTypes[rand.Intn(len(treeTypes))],
			treeColor[rand.Intn(len(treeColor))],
			treeTextures[rand.Intn(len(treeTextures))],
		)
	}

	forest.Draw()
}
