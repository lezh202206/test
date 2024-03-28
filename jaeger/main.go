package main

import (
	logger "log"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	// Create Jaeger configuration
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		logger.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return
	}

	// Set the service name
	cfg.ServiceName = "my-go-service"

	// Initialize Jaeger tracer
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		logger.Printf("Could not initialize Jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	// Set the global tracer instance
	opentracing.SetGlobalTracer(tracer)

	// Define a route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		span := opentracing.GlobalTracer().StartSpan("hello")
		defer span.Finish()

		// Set some tags and log
		span.SetTag("hello-to", "world")
		span.LogFields(
			log.String("event", "hello"),
			log.String("message", "Hello, world!"),
		)

		// Your actual application logic here
		w.Write([]byte("Hello, World!"))
	})

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	logger.Printf("Server running on port %s", port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
