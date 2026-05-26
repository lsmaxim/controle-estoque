package controllers

import (
	"net/http"
	// LISTAR AGENDAMENTOS
	"fmt"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

func ListarAgendamentos(c *gin.Context) {

	rows, err := database.DB.Query(`

	SELECT
		id,
		titulo,
		data_agendamento,
		status,
		IFNULL(observacao, '') as observacao

	FROM agendamentos

	ORDER BY
		CASE
			WHEN status = 'PENDENTE' THEN 1
			WHEN status = 'FINALIZADO' THEN 2
			ELSE 3
		END,

		data_agendamento ASC
`)
	if err != nil {

		fmt.Println("ERRO SQL:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var agendamentos []models.Agendamento

	for rows.Next() {

		var a models.Agendamento

		err := rows.Scan(
			&a.ID,
			&a.Titulo,
			&a.DataAgendamento,
			&a.Status,
			&a.Observacao,
		)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})

			return
		}

		agendamentos = append(agendamentos, a)
	}

	c.JSON(http.StatusOK, agendamentos)
}

// CRIAR AGENDAMENTO
func CriarAgendamento(c *gin.Context) {

	var agendamento models.Agendamento

	if err := c.ShouldBindJSON(&agendamento); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO agendamentos (
			titulo,
			data_agendamento
		)

		VALUES (?, ?)
	`,
		agendamento.Titulo,
		agendamento.DataAgendamento,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Agendamento criado com sucesso",
	})
}

// FINALIZAR AGENDAMENTO
func FinalizarAgendamento(c *gin.Context) {

	id := c.Param("id")

	var dados struct {
		Observacao string `json:"observacao"`
	}

	if err := c.ShouldBindJSON(&dados); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return
	}

	_, err := database.DB.Exec(`
		UPDATE agendamentos
		SET
			status = 'FINALIZADO',
			observacao = ?
		WHERE id = ?
	`,
		dados.Observacao,
		id,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Agendamento finalizado com sucesso",
	})
}
