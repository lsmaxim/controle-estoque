package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/routes"

)

func main() {

	database.Conectar()

	r := gin.Default()

	r.Use(cors.Default())

	routes.ConfigurarRotas(r)

	r.Run(":8080")
}