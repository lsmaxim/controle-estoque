package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
)

type Historico struct {
	ID           int    `json:"id"`
	ProdutoID    int    `json:"produto_id"`
	Usuario      string `json:"usuario"`
	Acao         string `json:"acao"`
	Descricao    string `json:"descricao"`
	DataRegistro string `json:"data_registro"`
}

// =======================
// REGISTRAR HISTÓRICO
// =======================
func RegistrarHistorico(
	produtoID int,
	usuario string,
	acao string,
	descricao string,
) {

	_, err := database.DB.Exec(`
		INSERT INTO historico (
			produto_id,
			usuario,
			acao,
			descricao
		)
		VALUES (?, ?, ?, ?)
	`,
		produtoID,
		usuario,
		acao,
		descricao,
	)

	if err != nil {
		fmt.Println("Erro ao registrar histórico:", err.Error())
	}
}

// =======================
// LISTAR HISTÓRICO
// =======================
func ListarHistorico(c *gin.Context) {

	id := c.Param("id")

	fmt.Println("ID recebido:", id)

	rows, err := database.DB.Query(`
		SELECT
			id,
			produto_id,
			usuario,
			acao,
			descricao,
			data_registro
		FROM historico
		WHERE produto_id = ?
		ORDER BY data_registro DESC
	`, id)

	if err != nil {

		fmt.Println("ERRO QUERY:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var historico []Historico

	for rows.Next() {

		var item Historico

		err := rows.Scan(
			&item.ID,
			&item.ProdutoID,
			&item.Usuario,
			&item.Acao,
			&item.Descricao,
			&item.DataRegistro,
		)

		if err != nil {

			fmt.Println("ERRO SCAN:", err)

			continue
		}

		fmt.Printf("ITEM: %+v\n", item)

		historico = append(historico, item)
	}

	if err := rows.Err(); err != nil {

		fmt.Println("ERRO ROWS:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	fmt.Println("TOTAL:", len(historico))

	c.JSON(http.StatusOK, historico)
}
