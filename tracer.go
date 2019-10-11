package main

import (
	"go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/exporter/trace/jaeger"
	sdk "go.opentelemetry.io/sdk/trace"
)

func setupTracer() (trace.Tracer, *jaeger.Exporter, error) {
	// Register installs a new global tracer instance.
	tracer := sdk.Register()

	// Construct and register an export pipeline using the Jaeger
	// exporter and a span processor.
	exporter, err := jaeger.NewExporter(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "opentelemetry-demo",
		}),
	)
	if err != nil {
		return nil, nil, err
	}

	// A simple span processor calls through to the exporter
	// without buffering.
	ssp := sdk.NewSimpleSpanProcessor(exporter)
	sdk.RegisterSpanProcessor(ssp)

	// Use sdk.AlwaysSample sampler to send all spans.
	sdk.ApplyConfig(
		sdk.Config{
			DefaultSampler: sdk.AlwaysSample(),
		},
	)

	return tracer, exporter, nil
}
