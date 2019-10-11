JAEGER_CONTAINER=jaeger
BINARY=example

.PHONY: all start-jaeger stop-jaeger clean

all: start-jaeger build run

build:
	go build -o $(BINARY)

run:
	./$(BINARY)

start-jaeger:
	docker run -d -p 6831:6831/udp -p 16686:16686 -p 14268:14268 --name $(JAEGER_CONTAINER) jaegertracing/all-in-one:latest

stop-jaeger:
	docker stop $(JAEGER_CONTAINER)

clean:
	docker rm -f $(JAEGER_CONTAINER)
	rm $(BINARY)