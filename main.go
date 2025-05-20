package main

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/izhukov1992/go-kafka-benchmarks/franz"
)

func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	size := parser.Int("s", "size", &argparse.Options{Required: true, Help: "Total size of dataset"})
	parser.Parse(os.Args)

	// Without compression 24Gb+24Gb
	franz.TestFranzComplex(true, true, *size, 1000000) // produce: 33.89, complex: 35.95
	franz.TestFranzComplex(true, true, *size, 100000)  // produce: 33.36, complex: 35.03
	franz.TestFranzComplex(true, true, *size, 10000)   // produce: 32.28, complex: 34.07
	franz.TestFranzComplex(true, true, *size, 1000)    // produce: 38.25, complex: 41.71
	franz.TestFranzComplex(true, true, *size, 100)     // produce: 85.21, complex: 92.94

	franz.TestFranzComplex(true, true, *size, 1000)    // produce: 38.47, complex: 40.29
	franz.TestFranzComplex(true, true, *size, 100)     // produce: 85.67, complex: 94.19
	franz.TestFranzComplex(true, true, *size, 1000000) // produce: 32.46, complex: 34.89
	franz.TestFranzComplex(true, true, *size, 100000)  // produce: 32.69, complex: 34.42
	franz.TestFranzComplex(true, true, *size, 10000)   // produce: 32.19, complex: 35.54

	// With compression 1Gb+1Gb
	// franz.TestFranzComplex(true, true, *size, 1000000) // produce: 17.24, complex: 17.65
	// franz.TestFranzComplex(true, true, *size, 100000)  // produce: 17.35, complex: 17.48
	// franz.TestFranzComplex(true, true, *size, 10000)   // produce: 17.28, complex: 17.72
	// franz.TestFranzComplex(true, true, *size, 1000)    // produce: 23.20, complex: 25.49
	// franz.TestFranzComplex(true, true, *size, 100)     // produce: 79.64, complex: 78.32

	// franz.TestFranzComplex(true, true, *size, 1000)    // produce: 22.81, complex: 25.18
	// franz.TestFranzComplex(true, true, *size, 100)     // produce: 79.02, complex: 78.85
	// franz.TestFranzComplex(true, true, *size, 1000000) // produce: 17.52, complex: 17.55
	// franz.TestFranzComplex(true, true, *size, 100000)  // produce: 17.69, complex: 17.49
	// franz.TestFranzComplex(true, true, *size, 10000)   // produce: 17.53, complex: 17.67
}
