package main

//import (
//	"context"
//	"github.com/jay-wlj/go-metric/instrumentation/github.com/gorm.io/gorm/otelgorm"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm/logger"
//	"log"
//	"time"
//
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
//	"go.opentelemetry.io/otel/sdk/resource"
//	sdktrace "go.opentelemetry.io/otel/sdk/trace"
//	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
//	"gorm.io/gorm"
//)
//
//// 初始化OpenTelemetry
//func initTracer() func() {
//	// 创建一个新的trace导出器，用于将追踪信息写入到控制台
//	exporter, err := stdouttrace.New(
//		stdouttrace.WithPrettyPrint(),
//	)
//	if err != nil {
//		log.Fatalf("创建导出器失败: %v", err)
//	}
//
//	// 创建一个新的trace提供器
//	tp := sdktrace.NewTracerProvider(
//		// 设置追踪提供器使用的导出器
//		sdktrace.WithBatcher(exporter),
//		// 配置资源信息，这里可以设置服务名等属性
//		sdktrace.WithResource(resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceNameKey.String("example-service"),
//		)),
//	)
//
//	// 将创建的trace提供器设置为全局提供器
//	otel.SetTracerProvider(tp)
//
//	// 返回一个函数，用于在程序结束时关闭trace提供器
//	return func() {
//		if err := tp.Shutdown(context.Background()); err != nil {
//			log.Fatalf("关闭trace提供器失败: %v", err)
//		}
//	}
//}
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
////func InitTestDataServiceDB() (*gorm.DB, error) {
////	dsn := "souti_rw:p6jYjyoOEsJXhOBp@tcp(rm-2zeue9603s75870g0wo.mysql.rds.aliyuncs.com:3306)/tool_content_check?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
////	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
////}
//
//func main() {
//	// 初始化OpenTelemetry
//	shutdown := initTracer()
//	defer shutdown()
//
//	// 初始化数据库连接
//	db, err := InitTestDataServiceDB()
//	if err != nil {
//		return
//	}
//
//	// 为GORM启用OpenTelemetry插件
//	if err := db.Use(otelgorm.NewPlugin()); err != nil {
//		panic("启用OpenTelemetry插件失败")
//	}
//
//	// 这里可以定义你的数据库模型
//	type Product struct {
//		gorm.Model
//		Code  string
//		Price uint
//	}
//
//	// 自动迁移模式，创建数据库表
//	db.AutoMigrate(&Product{})
//
//	// 创建一个新追踪器
//	tr := otel.Tracer("example-tracer")
//
//	// 创建一个带有追踪的上下文
//	ctx, span := tr.Start(context.Background(), "operation-name")
//	defer span.End()
//
//	// 使用带有追踪的上下文执行数据库查询
//	var product Product
//	result := db.WithContext(ctx).First(&product, 1)
//
//	// 检查查询结果
//	if result.Error != nil {
//		log.Printf("查询失败: %v", result.Error)
//	} else {
//		log.Printf("查询成功: %#v", product)
//	}
//
//	// 模拟一些额外的操作，以便在追踪中看到更多的信息
//	time.Sleep(100 * time.Millisecond)
//
//	// 你可以添加额外的追踪信息
//	span.AddEvent("模拟事件")
//
//	// 结束追踪
//	span.End()
//}
