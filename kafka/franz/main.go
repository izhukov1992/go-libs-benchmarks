package franz

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
)

func TestFranzComplex(generate bool, process bool, size int, batch_size int) {
	data := make([]byte, 512)
	for i := range 512 {
		data[i] = 'b'
	}
	seeds := []string{"localhost:9092"}
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("franz"),
		kgo.ConsumeTopics("input"),
		kgo.MaxBufferedRecords(batch_size),
		kgo.MaxBufferedBytes(batch_size*1024),
		kgo.MetadataMaxAge(60*time.Second),
		kgo.DefaultProduceTopic("output"),
		kgo.ProducerBatchCompression(kgo.NoCompression()),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()

	var (
		wg      sync.WaitGroup
		produce time.Duration
		complex time.Duration
		promise = func(r *kgo.Record, err error) {
			wg.Done()
		}
	)

	ctx := context.Background()

	if generate {
		// Produce test data
		start := time.Now()
		wg.Add(size)
		for range size {
			cl.Produce(ctx, &kgo.Record{Topic: "input", Value: data}, promise)
		}
		wg.Wait()
		produce = time.Since(start)
	}

	if process {
		// Process data
		counter := 0
		start := time.Now()
		for counter < size {
			fetches := cl.PollRecords(ctx, batch_size)
			counter += fetches.NumRecords()
			wg.Add(fetches.NumRecords())
			for _, record := range fetches.Records() {
				cl.Produce(ctx, &kgo.Record{Value: record.Value}, promise)
			}
			wg.Wait()
			cl.CommitUncommittedOffsets(ctx)
		}
		complex = time.Since(start)
	}

	logrus.Errorf("produce: %v, complex: %v", produce, complex)
}
