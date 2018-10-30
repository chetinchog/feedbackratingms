package main

import (
	"fmt"
	"os"
	"time"

	"github.com/chetinchog/feedbackratingms/routes"
	"github.com/chetinchog/feedbackratingms/tools/env"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func main() {
	fmt.Printf("$ Starting FeedbackRatingMS...\n\n")

	if len(os.Args) > 1 {
		env.Load(os.Args[1])
	}

	server := gin.Default()
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	server.Use(static.Serve("/", static.LocalFile(env.Get().WWWWPath, true)))

	routes.Route(server)

	fmt.Printf("\n$ Configuration Loaded.\n\n")

	server.Run(fmt.Sprintf(":%d", env.Get().Port))
}
