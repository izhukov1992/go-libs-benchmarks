package avro

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/izhukov1992/go-libs-benchmarks/avro/goavrolib"
)

func TestAvro() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	size := parser.Int("s", "size", &argparse.Options{Required: true, Help: "Total size of dataset"})
	parser.Parse(os.Args)

	goavrolib.TestGoavro(*size)
}
