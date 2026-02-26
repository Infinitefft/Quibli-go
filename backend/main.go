package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// 测试数据库连接
	db, err := sql.Open("postgres", "host=localhost port=5433 user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("failed opening connection to postgres: ", err)
	}
	defer db.Close()

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("failed connecting to postgres: ", err)
	}

	fmt.Println("Database connected successfully!")

	h := server.Default()

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.GET("/db-test", func(ctx context.Context, c *app.RequestContext) {
		if err := db.Ping(); err != nil {
			c.JSON(consts.StatusInternalServerError, utils.H{"error": "database connection failed"})
			return
		}
		c.JSON(consts.StatusOK, utils.H{"message": "database connected successfully"})
	})

	h.Spin()
}
