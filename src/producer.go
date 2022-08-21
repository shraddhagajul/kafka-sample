package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
	"time"
)

const (
	topic   = "message-log"
	broker1 = "localhost:9093"
	broker2 = "localhost:9094"
	broker3 = "localhost:9095"
)

func Producer(ctx context.Context) {
	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1, broker2, broker3},
		Topic:   topic,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message " + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes: ", i)
		i++
		time.Sleep(time.Second)
	}
}

