package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/routes"
)

func main() {

	database.Conectar()

	r := gin.Default()

	// LIBERA PASTA UPLOADS
	r.Static("/uploads", "./uploads")

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://192.168.0.7:5173",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// CARREGA AS ROTAS
	routes.ConfigurarRotas(r)

	// INICIA SERVIDOR
	r.Run(":8080")
}
