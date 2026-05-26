package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

func ListarCategorias(c *gin.Context) {

	rows, erro := database.DB.Query(
		"SELECT id, nome, data_cadastro FROM categorias",
	)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	defer rows.Close()

	categorias := []models.Categoria{}

	for rows.Next() {

		var categoria models.Categoria

		rows.Scan(
			&categoria.ID,
			&categoria.Nome,
			&categoria.DataCadastro,
		)

		categorias = append(categorias, categoria)
	}

	c.JSON(http.StatusOK, categorias)
}

func CriarCategoria(c *gin.Context) {

	var categoria models.Categoria

	if erro := c.ShouldBindJSON(&categoria); erro != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})

		return
	}

	insert, erro := database.DB.Exec(
		"INSERT INTO categorias(nome) VALUES(?)",
		categoria.Nome,
	)

	if erro != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})

		return
	}

	id, _ := insert.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
		"mensagem": "Categoria criada com sucesso!",
	})
}