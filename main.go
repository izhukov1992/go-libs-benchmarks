package main

import (
	"github.com/izhukov1992/go-libs-benchmarks/avro"
	"github.com/izhukov1992/go-libs-benchmarks/kafka"
)

func main() {
	kafka.TestKafka()
	avro.TestAvro()
}
