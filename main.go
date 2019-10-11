package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

func startHTTPServer() *http.Server {
	srv := &http.Server{Addr: ":7777"}

	http.HandleFunc("/hello", helloHandler)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	return srv
}

func main() {
	// Setup tracing and get a Tracer instance.  We'll use the
	// exporter to flush before exiting.
	tracer, exporter, err := setupTracer()

	if err != nil {
		log.Fatal("Could not initialize tracing: ", err)
	}

	// Tracing uses the standard context for propagation, we'll
	// start with a background context.
	ctx := context.Background()

	_ = tracer.WithSpan(ctx, "foo",
		func(ctx context.Context) error {
			tracer.WithSpan(ctx, "bar",
				func(ctx context.Context) error {
					tracer.WithSpan(ctx, "baz",
						func(ctx context.Context) error {
							return nil
						},
					)
					return nil
				},
			)
			return nil
		},
	)

	sayHello()
	sayHello2()

	srv := startHTTPServer()
	sayHTTPHello(context.Background())
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
	// The Jaeger exporter will have buffered spans at this point, send them.
	exporter.Flush()
	os.Exit(0)
}
