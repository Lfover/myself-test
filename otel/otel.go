package main

//
//import (
//	"database/sql"
//	"github.com/XSAM/otelsql"
//)
//
//// 确保你已经设置了 OpenTelemetry Tracer
//
//func main1() {
//	// 包装数据库驱动
//	driverName, err := otelsql.Register("postgres" /* 其他选项 */)
//	if err != nil {
//		// 处理错误
//	}
//
//	// 使用包装后的驱动连接数据库
//	db, err := sql.Open(driverName, "your-dsn")
//	if err != nil {
//		// 处理错误
//	}
//	defer db.Close()
//
//	// 现在，你可以使用 db 变量来执行 SQL 操作，并且这些操作会被 OpenTelemetry 自动追踪
//	err = db.Ping()
//	if err != nil {
//		return
//	}
//}
