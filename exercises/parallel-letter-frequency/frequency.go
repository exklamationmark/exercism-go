package letter

import (
	"sync"
)

// FreqMap stores rune frequency.
type FreqMap map[rune]int

func ConcurrentFrequencyDevillexio(s []string) FreqMap {
	f := make(chan FreqMap, len(s))

	for _, e := range s {
		go func(i string) { f <- Frequency(i) }(e)
	}

	results := FreqMap{}
	for i := 0; i < len(s); i++ {
		temp := <-f
		for r := range temp {
			results[r] += temp[r]
		}
	}

	return results
}

func ConcurrentFrequencyDevillexioUnbuffered(s []string) FreqMap {
	f := make(chan FreqMap)

	for _, e := range s {
		go func(i string) { f <- Frequency(i) }(e)
	}

	results := FreqMap{}
	for i := 0; i < len(s); i++ {
		temp := <-f
		for r := range temp {
			results[r] += temp[r]
		}
	}

	return results
}

func ConcurrentFrequency(strs []string) FreqMap {
	return ConcurrentFrequencyExklamationmark(strs)
}

// ConcurrentFrequency counts rune frequency in parallels.
func ConcurrentFrequencyExklamationmark(strs []string) FreqMap {
	var mu sync.Mutex
	m := FreqMap{}

	// NOTE: running in parallel this way is not always the faster solution.
	// For example, on my Linux VM, the sequential count is faster.
	// Tuning is required to find the appropriate workload, so the cost of
	// counting in parallel + synchronization >> cost of sequenctial run with caches.

	// NOTE: IRL, we need to tune the number of strings to count/goroutine to get a balanced workload.
	// I only take a simple approach & spin up 1 counting goroutine/element in strs.
	//
	// Performance will be drastically different for each of these cases (each process 10^12 runes in total)
	// 1. 1 element in strs with 10^12 runes
	// 2. 10^3 elements in strs, each have 10^9 runes
	// 3. 10^6 elements in strs, each have 10^6 runes
	// 4. 10^9 elements in strs, each have 10^3 runes
	var wg sync.WaitGroup
	wg.Add(len(strs))
	for _, str := range strs {
		go countRuneFreq([][]rune{[]rune(str)}, &m, &mu, &wg)
	}
	wg.Wait()

	return m
}

// countRuneFreq counts rune frequency, then atomically accumulates the result.
func countRuneFreq(work [][]rune, acc *FreqMap, mu *sync.Mutex, wg *sync.WaitGroup) {
	freq := FreqMap{} // this worker's count

	for _, runes := range work {
		for _, r := range runes {
			freq[r]++
		}
	}

	// atomic update
	mu.Lock()
	for r, count := range freq {
		(*acc)[r] += count
	}
	mu.Unlock()

	wg.Done()
}

func SequentialFrequency(strs []string) FreqMap {
	freq := FreqMap{}

	for _, str := range strs {
		local := Frequency(str)

		for k, v := range local {
			freq[k] += v
		}
	}

	return freq
}

// Frequency counts rune frequency sequenctially.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
