package main

import (
	"context"
	"kafka-sample/src"
)

func main() {
	ctx := context.Background()
	go src.Producer(ctx)
	src.Consumer(ctx)
}
