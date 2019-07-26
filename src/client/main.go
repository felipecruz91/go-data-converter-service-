package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"../proto"
)

type JSONRequest struct {
	Body string `json:"body"`
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewConverterServiceClient(conn)

	g := gin.Default()

	g.GET("/api/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "I'm alive!")
	})

	g.POST("/api/convert", func(ctx *gin.Context) {
		var json JSONRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req := &proto.Request{Body: json.Body}

		if data, err := client.Convert(ctx, req); err == nil {
			ctx.ProtoBuf(http.StatusOK, data)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run API: %v", err)
	}
}
