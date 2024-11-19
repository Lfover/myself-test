package main

//
//import (
//	"context"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm/logger"
//	"log"
//
//	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
//	"github.com/uptrace/opentelemetry-go-extra/otelgorsql"
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/attribute"
//	"go.opentelemetry.io/otel/sdk/resource"
//	sdktrace "go.opentelemetry.io/otel/sdk/trace"
//	"go.opentelemetry.io/otel/semconv/v1.4.0"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//)
//
//func InitTestDataServiceDB() (*gorm.DB, error) {
//	// 请确保你的DSN是正确的，并且你有权限访问数据库
//	dsn := "username:password@tcp(your-database-host:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Silent),
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	// 启用OpenTelemetry插件
//	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithTracerProvider(otel.GetTracerProvider()))); err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}
//
//func main() {
//	// 设置链路追踪提供者
//	tp := sdktrace.NewTracerProvider(
//		sdktrace.WithSampler(sdktrace.AlwaysSample()),
//		sdktrace.WithResource(resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceNameKey.String("your-service-name"),
//			attribute.String("environment", "development"),
//		)),
//	)
//	otel.SetTracerProvider(tp)
//	defer func() { _ = tp.Shutdown(context.Background()) }()
//
//	// 使用 otelsql 包装 SQLite 驱动程序
//	driverName, err := otelgorsql.Register("sqlite3", sqlite.Open("test.db"), otelgorsql.WithAllTraceOptions())
//	if err != nil {
//		log.Fatalf("failed to register otelsql driver: %v", err)
//	}
//
//	// 设置 GORM，告诉它使用 otelsql 包装过的驱动程序
//	db, err := InitTestDataServiceDB()
//	if err != nil {
//		log.Fatalf("failed to open database: %v", err)
//	}
//
//	// 应用 otelgorm 插件
//	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithTracerProvider(tp))); err != nil {
//		log.Fatalf("failed to apply otelgorm plugin: %v", err)
//	}
//
//	// 这里开始你的数据库操作
//	// ...
//}
