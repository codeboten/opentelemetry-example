# opentelemetry-example

Just some code I'm playing around with to take opentelemetry's go library for a test drive. Most of the code here was published in this [Getting Started](https://lightstep.com/blog/getting-started-with-opentelemetry-alphas-golang/) guide, with some minor changes to accomodate for the changes in the library.

##### Dependencies

* Docker
* Go

##### Testing it out

The following will spin up a jaeger all-in-one container, build the binary and send some traces to the local jaeger. Once it's run, open http://localhost:16686 in your browser to search the traces.

```bash
make all
```

