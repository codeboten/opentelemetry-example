package main

import (
	"context"

	"go.opentelemetry.io/api/trace"
)

func sayHello() {
	ctx := context.Background()
	tracer := trace.GlobalTracer()

	ctx, trace := tracer.Start(ctx, "say-hello")

	trace.End()
}

func sayHello2() {
	ctx := context.Background()
	tracer := trace.GlobalTracer()

	err := tracer.WithSpan(ctx, "say-hello2", func(ctx context.Context) error {
		// This body is traced, and the span will End() despite panics.
		return nil
	})

	if err != nil {
		// ...
	}
}
