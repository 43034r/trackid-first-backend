package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
	"monitoriong.wiki/trackid-first-backend/controllers"
	"monitoriong.wiki/trackid-first-backend/database"
)

// For opentelemetry
/*
var (
	serviceName  = os.Getenv("TI_SERVICE_NAME")
	collectorURL = os.Getenv("TI_OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("TI_INSECURE_MODE")
)
*/

var (
	serviceName  = "TI_SERVICE_NAME"
	collectorURL = "192.168.1.150:4317"
	insecure     = "true"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	cleanup := initTracer()
	defer cleanup(context.Background())

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	r.GET("/trackids/:id", controllers.ReadTrackid)
	r.GET("/trackids", controllers.ReadTrackids)
	r.POST("/trackids", controllers.CreateTrackid)
	r.PUT("/trackids/:id", controllers.UpdateTrackid)
	r.DELETE("/trackids/:id", controllers.DeleteTrackid)
	r.Run(":5000")
}

func initTracer() func(context.Context) error {

	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: ", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
