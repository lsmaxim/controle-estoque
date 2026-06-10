package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

// =======================
// LISTAR CHAMADOS
// =======================
func ListarChamados(c *gin.Context) {

	rows, err := database.DB.Query(`
		SELECT
			id,
			titulo,
			descricao,
			solicitante,
			setor,
			equipamento_id,
			status,
			prioridade,
			data_abertura,
			data_fechamento
		FROM chamados
		ORDER BY data_abertura DESC
	`)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var chamados []models.Chamado

	for rows.Next() {

		var chamado models.Chamado

		err := rows.Scan(
			&chamado.ID,
			&chamado.Titulo,
			&chamado.Descricao,
			&chamado.Solicitante,
			&chamado.Setor,
			&chamado.EquipamentoID,
			&chamado.Status,
			&chamado.Prioridade,
			&chamado.DataAbertura,
			&chamado.DataFechamento,
		)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})

			return
		}

		chamados = append(chamados, chamado)
	}

	c.JSON(http.StatusOK, chamados)
}

// =======================
// CRIAR CHAMADO
// =======================
func CriarChamado(c *gin.Context) {

	var chamado models.Chamado

	if err := c.ShouldBindJSON(&chamado); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	result, err := database.DB.Exec(`
		INSERT INTO chamados (
			titulo,
			descricao,
			solicitante,
			setor,
			equipamento_id,
			status,
			prioridade
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		chamado.Titulo,
		chamado.Descricao,
		chamado.Solicitante,
		chamado.Setor,
		chamado.EquipamentoID,
		"ABERTO",
		chamado.Prioridade,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	id, _ := result.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"id":       id,
		"mensagem": "Chamado criado com sucesso!",
	})
}

// =======================
// FECHAR CHAMADO
// =======================
func FecharChamado(c *gin.Context) {

	id := c.Param("id")

	var dados struct {
		Solucao string `json:"solucao"`
	}

	if err := c.ShouldBindJSON(&dados); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	_, err := database.DB.Exec(`
		UPDATE chamados
		SET
			status = 'FECHADO',
			solucao = ?,
			data_fechamento = NOW()
		WHERE id = ?
	`,
		dados.Solucao,

		id,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Chamado fechado com sucesso!",
	})
}
