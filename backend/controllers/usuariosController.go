package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
)

func ListarUsuarios(c *gin.Context) {

	rows, err := database.DB.Query(`
		SELECT
			id,
			nome,
			email,
			nivel
		FROM usuarios
		ORDER BY nome
	`)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var usuarios []gin.H

	for rows.Next() {

		var id int
		var nome string
		var email string
		var nivel string

		err := rows.Scan(
			&id,
			&nome,
			&email,
			&nivel,
		)

		if err != nil {
			continue
		}

		usuarios = append(usuarios, gin.H{
			"id":    id,
			"nome":  nome,
			"email": email,
			"nivel": nivel,
		})
	}

	c.JSON(http.StatusOK, usuarios)
}
