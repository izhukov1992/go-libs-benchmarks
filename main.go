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
	franz.TestFranzComplex(true, true, *size, 1000000) // produce: 19.33, complex: 29.30
	franz.TestFranzComplex(true, true, *size, 100000)  // produce: 18.88, complex: 29.01
	franz.TestFranzComplex(true, true, *size, 10000)   // produce: 18.55, complex: 28.74
	franz.TestFranzComplex(true, true, *size, 1000)    // produce: 18.54, complex: 34.47
	franz.TestFranzComplex(true, true, *size, 100)     // produce: 39.54, complex: 84.32

	franz.TestFranzComplex(true, true, *size, 1000)    // produce: 18.65, complex: 34.26
	franz.TestFranzComplex(true, true, *size, 100)     // produce: 39.47, complex: 84.48
	franz.TestFranzComplex(true, true, *size, 1000000) // produce: 18.67, complex: 29.45
	franz.TestFranzComplex(true, true, *size, 100000)  // produce: 18.64, complex: 29.64
	franz.TestFranzComplex(true, true, *size, 10000)   // produce: 18.58, complex: 29.63

	// With enabled compression 1Gb+1Gb
	// franz.TestFranzComplex(true, true, *size, 1000000) // produce: 1.56, complex: 4.49
	// franz.TestFranzComplex(true, true, *size, 100000)  // produce: 1.56, complex: 3.81
	// franz.TestFranzComplex(true, true, *size, 10000)   // produce: 2.07, complex: 4.86
	// franz.TestFranzComplex(true, true, *size, 1000)    // produce: 4.30, complex: 9.39
	// franz.TestFranzComplex(true, true, *size, 100)     // produce: 18.47, complex: 42.40

	// franz.TestFranzComplex(true, true, *size, 1000)    // produce: 4.36, complex: 9.50
	// franz.TestFranzComplex(true, true, *size, 100)     // produce: 18.57, complex: 43.45
	// franz.TestFranzComplex(true, true, *size, 1000000) // produce: 1.49, complex: 3.87
	// franz.TestFranzComplex(true, true, *size, 100000)  // produce: 1.61, complex: 3.84
	// franz.TestFranzComplex(true, true, *size, 10000)   // produce: 2.05, complex: 5.01
}
