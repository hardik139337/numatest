package main

import (
	"context"
	"log"

	"github.com/numaproj/numaflow-go/pkg/mapper"
)

type EvenOdd struct {
}

func (e *EvenOdd) Map(ctx context.Context, keys []string, d mapper.Datum) mapper.Messages {
	// msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	// If msg is not an integer, drop it, otherwise return it with "even" or "odd" key.

	return mapper.MessagesBuilder().Append(mapper.NewMessage([]byte("teset")))

}

func main() {
	err := mapper.NewServer(&EvenOdd{}).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map function server: ", err)
	}
}
