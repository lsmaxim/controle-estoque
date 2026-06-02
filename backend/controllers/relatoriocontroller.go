package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"

	"backend/database"
)

func GerarRelatorioEstoque(c *gin.Context) {

	rows, err := database.DB.Query(`
		SELECT
			nome,
			marca,
			modelo,
			quantidade,
			quantidade_minima
		FROM produtos
		ORDER BY nome
	`)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "Relatorio de Estoque")

	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 10)

	pdf.Cell(60, 8, "Produto")
	pdf.Cell(40, 8, "Marca")
	pdf.Cell(40, 8, "Modelo")
	pdf.Cell(25, 8, "Qtd")
	pdf.Cell(25, 8, "Min")

	pdf.Ln(-1)

	for rows.Next() {

		var nome string
		var marca string
		var modelo string
		var quantidade int
		var minima int

		err := rows.Scan(
			&nome,
			&marca,
			&modelo,
			&quantidade,
			&minima,
		)

		if err != nil {
			continue
		}

		pdf.SetFont("Arial", "", 12)

		pdf.Cell(60, 8, nome)
		pdf.Cell(40, 8, marca)
		pdf.Cell(40, 8, modelo)
		pdf.Cell(25, 8, fmt.Sprintf("%d", quantidade))
		pdf.Cell(25, 8, fmt.Sprintf("%d", minima))

		pdf.Ln(-1)
	}

	var buffer bytes.Buffer

	pdf.Output(&buffer)

	c.Header(
		"Content-Disposition",
		"attachment; filename=relatorio_estoque.pdf",
	)

	c.Data(
		http.StatusOK,
		"application/pdf",
		buffer.Bytes(),
	)
}
