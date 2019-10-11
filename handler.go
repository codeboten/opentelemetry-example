package main

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"go.opentelemetry.io/api/distributedcontext"
	"go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/plugin/httptrace"
	"google.golang.org/grpc/codes"
)

func sayHTTPHello(ctx context.Context) {
	client := http.DefaultClient
	tracer := trace.GlobalTracer()

	tracer.WithSpan(ctx, "client-call",
		func(ctx context.Context) error {
			req, _ := http.NewRequest("GET", "http://localhost:7777/hello", nil)

			ctx, req = httptrace.W3C(ctx, req)
			httptrace.Inject(ctx, req)

			res, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			_, err = ioutil.ReadAll(res.Body)
			res.Body.Close()
			trace.CurrentSpan(ctx).SetStatus(codes.OK)

			return err
		})
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	tracer := trace.GlobalTracer()

	// Extracts the conventional HTTP span attributes,
	// distributed context tags, and a span context for
	// tracing this request.
	attrs, tags, spanCtx := httptrace.Extract(req.Context(), req)

	// Apply the distributed context tags to the request
	// context.
	req = req.WithContext(distributedcontext.WithMap(req.Context(), distributedcontext.NewMap(distributedcontext.MapUpdate{
		MultiKV: tags,
	})))

	// Start the server-side span, passing the remote
	// child span context explicitly.
	_, span := tracer.Start(
		req.Context(),
		"hello",
		trace.WithAttributes(attrs...),
		trace.ChildOf(spanCtx),
	)
	defer span.End()

	_, _ = io.WriteString(w, "Hello, world!\n")
}
