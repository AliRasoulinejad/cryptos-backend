package app

import (
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

var Tracer trace.Tracer

func (application *Application) WithTracer() {
	configs := config.C.TraceProvider

	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost(configs.AgentHost),
		jaeger.WithAgentPort(configs.AgentPort)))
	if err != nil {
		log.Logger.WithError(err).Fatal(fmt.Errorf("starting jaeger exporter failed"))
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Logger.WithError(err).Fatal("error in getting hostname from OS")
	}

	traceProvider := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(exporter),
		sdkTrace.WithSampler(sdkTrace.TraceIDRatioBased(configs.SamplerRatio)),
		sdkTrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceInstanceIDKey.String(hostname),
			semconv.ServiceNameKey.String(AppName),
			semconv.ServiceVersionKey.String(GitTag),
			semconv.DeploymentEnvironmentKey.String(configs.DeploymentEnvironment),
		)),
	)

	otel.SetTracerProvider(traceProvider)
	Tracer = traceProvider.Tracer(configs.ServiceName)
}
