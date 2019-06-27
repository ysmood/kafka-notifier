package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	kit "github.com/ysmood/gokit"
)

func main() {
	kit.Tasks().Add(
		kit.Task("serve", "run the service").Init(func(cmd kit.TaskCmd) func() {
			cmd.Default()

			return func() {
				serve()
			}
		}),
	)
}

// Consumer ...
type Consumer struct {
	reader     *kafka.Reader
	config     kafka.ReaderConfig
	initOffset int64
}

// Consume ...
func (c *Consumer) Consume() {
	r := kafka.NewReader(c.config)
	r.SetOffset(c.initOffset)

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message %s\n", string(m.Value))
		r.CommitMessages(ctx, m)
	}
}

func serve() {

}
