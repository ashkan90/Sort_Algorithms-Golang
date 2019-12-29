package main

import (
	"data_structures/algo/benchmarks"
	"data_structures/misc"
	"fmt"
	"os"
	"testing"
)

func main() {
	f, err := os.Create("results.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	bench := []func(b *testing.B){
		benchmarks.BenchmarkBubbleSortMIXED,
		benchmarks.BenchmarkBubbleSortBACKSORTED,
		benchmarks.BenchmarkBubbleSortPRESORTED,
		benchmarks.BenchmarkQuickSortMIXED,
		benchmarks.BenchmarkQuickSortBACKSORTED,
		benchmarks.BenchmarkQuickSortPRESORTED,
		benchmarks.BenchmarkMergeSortMIXED,
		benchmarks.BenchmarkMergeSortBACKSORTED,
		benchmarks.BenchmarkMergeSortPRESORTED,
		benchmarks.BenchmarkSelectionSortMIXED,
		benchmarks.BenchmarkSelectionSortBACKSORTED,
		benchmarks.BenchmarkSelectionSortPRESORTED,
	}

	for k, fn := range bench {
		res := fmt.Sprintf("Slice Type: %s, Algorithm: %s, Information: %s",
			benchmarks.On,
			misc.GetFunctionName(fn),
			testing.Benchmark(fn).String(),
		)

		if k%3 == 0 {
			_, _ = fmt.Fprintln(f)
		}

		_, err = fmt.Fprintln(f, res)
		if err != nil {
			_ = f.Close()
			panic(err)
		}
	}
}
