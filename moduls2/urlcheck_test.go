package main

import (
	"testing"
)

func Benchmark_fetcher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fetchSequencial(urls)
	}
}

func Benchmark_fetcherConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fetchConcurrent(urls)
	}
}
