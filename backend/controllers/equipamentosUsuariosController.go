package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

func VincularEquipamentoUsuario(c *gin.Context) {

	var vinculo models.EquipamentoUsuario

	if err := c.ShouldBindJSON(&vinculo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO equipamentos_usuarios (
			produto_id,
			usuario_id
		)
		VALUES (?, ?)
	`,
		vinculo.ProdutoID,
		vinculo.UsuarioID,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Equipamento vinculado com sucesso",
	})
}
func ListarEquipamentosUsuario(c *gin.Context) {

	id := c.Param("id")

	rows, err := database.DB.Query(`
		SELECT
			p.id,
			p.nome
		FROM equipamentos_usuarios eu
		INNER JOIN produtos p
			ON p.id = eu.produto_id
		WHERE eu.usuario_id = ?
	`, id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var equipamentos []gin.H

	for rows.Next() {

		var produtoID int
		var nome string

		rows.Scan(
			&produtoID,
			&nome,
		)

		equipamentos = append(equipamentos, gin.H{
			"id":   produtoID,
			"nome": nome,
		})
	}

	c.JSON(http.StatusOK, equipamentos)
}
