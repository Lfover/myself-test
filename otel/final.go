package main

import (
	"context"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var tracerExp *otlptrace.Exporter

func retryInitTracer() func() {
	var shutdown func()
	go func() {
		for {
			// otel will reconnected and re-send spans when otel col recover. so, we don't need to re-init tracer exporter.
			if tracerExp == nil {
				shutdown = initTracer()
			} else {
				break
			}
			time.Sleep(time.Minute * 5)
		}
	}()
	return shutdown
}

func initTracer() func() {
	// temporarily set timeout to 10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serviceName, ok := os.LookupEnv("OTEL_SERVICE_NAME")
	if !ok {
		serviceName = "server_name"
		os.Setenv("OTEL_SERVICE_NAME", serviceName)
	}
	otelAgentAddr, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if !ok {
		otelAgentAddr = "http://localhost:4317"
		os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", otelAgentAddr)
	}
	zap.S().Infof("OTLP Trace connect to: %s with service name: %s", otelAgentAddr, serviceName)

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithDialOption(grpc.WithBlock()))
	if err != nil {
		handleErr(err, "OTLP Trace gRPC Creation")
		return nil
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL)))

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracerExp = traceExporter
	return func() {
		// Shutdown will flush any remaining spans and shut down the exporter.
		handleErr(tracerProvider.Shutdown(ctx), "failed to shutdown TracerProvider")
	}
}

func handleErr(err error, message string) {
	if err != nil {
		zap.S().Errorf("%s: %v", message, err)
	}
}

func InitDataServiceDB() (*gorm.DB, error) {
	dsn := "souti_rw:p6jYjyoOEsJXhOBp@tcp(rm-2zeue9603s75870g0wo.mysql.rds.aliyuncs.com:3306)/xiaosi_butterfly_mark?loc=PRC&charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newExporter(url string) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}

func main() {
	// 初始化OpenTelemetry导出器
	ctx := context.Background()
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("创建导出器失败: %v", err)
	}
	//url := "http://localhost:8007"
	//exporter, _ := newExporter(url)

	// 创建TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("your-service-name"),
		)),
	)
	otel.SetTracerProvider(tp)
	defer func() { _ = tp.Shutdown(ctx) }()

	// 连接数据库
	db, err := InitDataServiceDB()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 注册otelgorm插件
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatalf("注册otelgorm插件失败: %v", err)
	}

	// 运行数据库查询
	runDatabaseQueries(ctx, db)

	// 确保所有的追踪数据都被发送到导出器
	time.Sleep(1 * time.Second)
}

func runDatabaseQueries(ctx context.Context, db *gorm.DB) {
	type Team struct {
		ID        int
		Name      string
		Status    int
		Remark    string
		UserCount int
	}

	// 查询
	var teams []Team
	if result := db.WithContext(ctx).Table("team").Find(&teams); result.Error != nil {
		log.Fatalf("查询失败: %v", result.Error)
	}
	log.Printf("查询到的团队数量: %d", len(teams))
}
