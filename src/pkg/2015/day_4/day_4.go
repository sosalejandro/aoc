package day_4

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/sosalejandro/aoc/src/pkg/helpers"
	"runtime"
	"strings"
	"sync"
)

func Run() {
	input := helpers.ReadInput("./src/pkg/2015/day_4/input.txt")
	input = strings.TrimSpace(input)
	fmt.Println("Day 4, Part 1:", FindAdventCoin(input, Hash.IsAdventCoin))
	fmt.Println("Day 4, Part 2:", FindAdventCoin(input, Hash.IsAdventCoin2))
}

// FindAdventCoin finds the first advent coin for the given input.
// An advent coin is a hash that starts with 5 or 6 zeroes.
// The function uses the given isAdventCoin function to determine if a hash is an advent coin.
// The function uses the given input as a prefix for the hash.
// The function uses the number of available CPUs as the number of workers.
// The function returns the first advent coin found.
// When an advent coin is found, the function cancels the other workers. This is done by using a context.
// Exiting the other goroutines is done by using a context avoiding extra computational work.
func FindAdventCoin(input string, isAdventCoin func(Hash) bool) int {
	numWorkers := runtime.NumCPU()
	results := make(chan int)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			var i int
			for {
				select {
				case <-ctx.Done():
					return
				default:
					hash := md5.New()
					v := fmt.Sprintf("%s%d", input, workerID+numWorkers*i)
					hash.Write([]byte(v))
					hashed := Hash(hash.Sum(nil))
					if isAdventCoin(hashed) {
						results <- workerID + numWorkers*i
						cancel()
						return
					}
					i++
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	result := <-results
	return result
}

// Hash represents a hash.
type Hash []byte

// String returns the string representation of the hash.
func (h Hash) String() string {
	return hex.EncodeToString(h)
}

// IsAdventCoin returns true if the hash is an advent coin with 5 zeros prefixed.
func (h Hash) IsAdventCoin() bool {
	return strings.HasPrefix(h.String(), "00000")
}

// IsAdventCoin2 returns true if the hash is an advent coin with 6 zeros prefixed.
func (h Hash) IsAdventCoin2() bool {
	return strings.HasPrefix(h.String(), "000000")
}
