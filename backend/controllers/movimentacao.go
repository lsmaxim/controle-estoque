package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

func CriarMovimentacao(c *gin.Context) {

	var movimentacao models.Movimentacao

	// RECEBE JSON
	if err := c.ShouldBindJSON(&movimentacao); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// SALVA MOVIMENTAÇÃO
	_, err := database.DB.Exec(

		`INSERT INTO movimentacoes
		(
			produto_id,
			tipo,
			quantidade,
			observacao
		)
		VALUES (?, ?, ?, ?)`,

		movimentacao.ProdutoID,
		movimentacao.Tipo,
		movimentacao.Quantidade,
		movimentacao.Observacao,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// ENTRADA
	if movimentacao.Tipo == "ENTRADA" {

		_, err = database.DB.Exec(

			`UPDATE produtos
			SET quantidade = quantidade + ?
			WHERE id = ?`,

			movimentacao.Quantidade,
			movimentacao.ProdutoID,
		)

	} else {

		// SAÍDA
		_, err = database.DB.Exec(

			`UPDATE produtos
			SET quantidade = quantidade - ?
			WHERE id = ?`,

			movimentacao.Quantidade,
			movimentacao.ProdutoID,
		)
	}

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Movimentação realizada com sucesso",
	})
}
func ListarMovimentacoes(c *gin.Context) {

	rows, err := database.DB.Query(`

		SELECT

			m.id,

			p.nome,

			m.tipo,

			m.quantidade,

			IFNULL(m.observacao, ''),

			DATE_FORMAT(
				m.data_movimentacao,
				'%d/%m/%Y %H:%i'
			)

		FROM movimentacoes m

		INNER JOIN produtos p
			ON p.id = m.produto_id

		ORDER BY m.data_movimentacao DESC

		LIMIT 20
	`)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var movimentacoes []gin.H

	for rows.Next() {

		var id int
		var produto string
		var tipo string
		var quantidade int
		var observacao string
		var data []byte
		err := rows.Scan(

			&id,

			&produto,

			&tipo,

			&quantidade,

			&observacao,

			&data,
		)

		if err != nil {

			continue
		}

		movimentacoes = append(

			movimentacoes,

			gin.H{

				"id": id,

				"produto": produto,

				"tipo": tipo,

				"quantidade": quantidade,

				"observacao": observacao,

				"data_movimentacao": string(data),
			},
		)
	}

	c.JSON(http.StatusOK, movimentacoes)
}
