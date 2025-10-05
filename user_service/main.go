package main

import (
	"context"
	"fmt"
	"time"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra/mongo"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra/postgres"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/server"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/pkg/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"go.uber.org/zap"
)

func main() {

	appconfig := config.ReadConfig()

	defer zap.L().Sync()

	pgHandler, err := postgres.NewDBRepository(
		appconfig.PgConfig.Host,
		appconfig.PgConfig.Port,
		appconfig.PgConfig.User,
		appconfig.PgConfig.Password,
		appconfig.PgConfig.Name,
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}

	mongoHandler, err := mongo.NewDBRepository(
		appconfig.MongoConfig.User,
		appconfig.MongoConfig.Password,
		appconfig.MongoConfig.AppName,
		appconfig.MongoConfig.DbName,
		appconfig.MongoConfig.Collection,
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}

	server.Start(pgHandler, mongoHandler, appconfig)

}
func startTracing() (*trace.TracerProvider, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	tracerprovider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("cqrs"),
			),
		),
	)

	otel.SetTracerProvider(tracerprovider)

	return tracerprovider, nil
}
