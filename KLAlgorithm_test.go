package klpartitinlin

import (
	"fmt"
	"testing"

	graphlib "github.com/Rakiiii/goGraph"
)

//Benchmark benchmark fast kernigan lin algorithm
func Benchmark(b *testing.B) {
	var parser = new(graphlib.Parser)
	var g, err = parser.ParseUnweightedUndirectedGraphFromFile("test_gr3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("graph parsed")

	result := Result{Matrix: nil, Value: -1}

	for i := 0; i < b.N; i++ {
		result, err = KLPartitionigAlgorithm(g, result.Matrix)
		fmt.Println("graph parted", result.Value)
	}
}

//BenchmarkClssic benchmark classi kernigan lin algorithm
func BenchmarkClassic(b *testing.B) {
	var parser = new(graphlib.Parser)
	var g, err = parser.ParseUnweightedUndirectedGraphFromFile("test_gr3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("graph parsed")

	result := Result{Matrix: nil, Value: -1}

	for i := 0; i < b.N; i++ {
		result, err = KLPartitionigAlgorithmClassic(g, result.Matrix)
		fmt.Println("graph parted", result.Value)
	}
}
