package main

import (
	"context"
	"log"

	"backend/internal/config"
	"backend/internal/ent"
	"backend/internal/handler"
	"backend/internal/router"
	"backend/internal/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/lib/pq"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbDSN := config.GetDBDSN()
	client, err := ent.Open("postgres", dbDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	authService := service.NewAuthService(client)
	authHandler := handler.NewAuthHandler(authService)
	router := router.NewRouter(authHandler)

	h := server.Default(
		server.WithHostPorts(config.AppConfig.ServerPort),
		server.WithMaxRequestBodySize(10<<20),
	)

	h.Use(func(ctx context.Context, c *app.RequestContext) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if string(c.Method()) == "OPTIONS" {
			c.SetStatusCode(204)
			c.Abort()
			return
		}
		c.Next(ctx)
	})

	router.Register(h)

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.Spin()
}
