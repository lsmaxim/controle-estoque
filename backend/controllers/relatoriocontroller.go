package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/text/encoding/charmap"

	"backend/database"
)

func textoPDF(texto string) string {

	conversor := charmap.Windows1252.NewEncoder()

	resultado, err := conversor.String(texto)

	if err != nil {
		return texto
	}

	return resultado
}

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

	// ==========================
	// CABEÇALHO
	// ==========================
	pdf.SetHeaderFunc(func() {

		pdf.Image(
			"C:/Projetos/controle-estoque/frontend/public/img/LOGOTIPO-XB.png",
			10,
			10,
			25,
			0,
			false,
			"",
			0,
			"",
		)

		pdf.SetFont("Arial", "B", 16)
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(
			190,
			10,
			textoPDF("RELATÓRIO DE ESTOQUE"),
			"",
			1,
			"C",
			false,
			0,
			"",
		)

		pdf.SetFont("Arial", "", 10)

		pdf.CellFormat(
			190,
			6,
			textoPDF(
				"Emitido em: "+
					time.Now().Format("02/01/2006 15:04"),
			),
			"",
			1,
			"R",
			false,
			0,
			"",
		)

		pdf.Ln(8)

		// Cabeçalho da tabela
		pdf.SetFillColor(220, 220, 220)

		pdf.SetFont("Arial", "B", 10)

		pdf.CellFormat(
			60,
			8,
			textoPDF("Produto"),
			"1",
			0,
			"C",
			true,
			0,
			"",
		)

		pdf.CellFormat(
			30,
			8,
			textoPDF("Marca"),
			"1",
			0,
			"C",
			true,
			0,
			"",
		)

		pdf.CellFormat(
			55,
			8,
			textoPDF("Modelo"),
			"1",
			0,
			"C",
			true,
			0,
			"",
		)

		pdf.CellFormat(
			25,
			8,
			textoPDF("Qtd"),
			"1",
			0,
			"C",
			true,
			0,
			"",
		)

		pdf.CellFormat(
			25,
			8,
			textoPDF("Mín"),
			"1",
			1,
			"C",
			true,
			0,
			"",
		)
	})

	// ==========================
	// RODAPÉ
	// ==========================
	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(-15)

		pdf.SetFont("Arial", "I", 8)

		pdf.CellFormat(
			0,
			10,
			fmt.Sprintf(
				textoPDF("Página %d"),

				pdf.PageNo(),
			),
			"",
			0,
			"C",
			false,
			0,
			"",
		)
	})

	pdf.AddPage()

	pdf.SetFont("Arial", "", 10)

	totalProdutos := 0

	// ==========================
	// DADOS
	// ==========================
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

		totalProdutos++

		// Vermelho para estoque crítico
		if quantidade <= minima {

			pdf.SetTextColor(255, 0, 0)

		} else {

			pdf.SetTextColor(0, 0, 0)
		}
		//Colunas da tabela
		pdf.CellFormat(
			60,
			8,
			textoPDF(nome),
			"1",
			0,
			"L",
			false,
			0,
			"",
		)

		pdf.CellFormat(
			30,
			8,
			textoPDF(marca),
			"1",
			0,
			"L",
			false,
			0,
			"",
		)

		pdf.CellFormat(
			55,
			8,
			textoPDF(modelo),
			"1",
			0,
			"L",
			false,
			0,
			"",
		)

		pdf.CellFormat(
			25,
			8,
			fmt.Sprintf("%d", quantidade),
			"1",
			0,
			"C",
			false,
			0,
			"",
		)

		pdf.CellFormat(
			25,
			8,
			fmt.Sprintf("%d", minima),
			"1",
			1,
			"C",
			false,
			0,
			"",
		)
	}

	// ==========================
	// RESUMO
	// ==========================
	pdf.Ln(5)

	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Arial", "B", 11)

	pdf.CellFormat(
		190,
		8,
		textoPDF(
			fmt.Sprintf(
				"Total de produtos: %d",
				totalProdutos,
			),
		),
		"",
		1,
		"R",
		false,
		0,
		"",
	)

	// ==========================
	// GERAR PDF
	// ==========================
	var buffer bytes.Buffer

	if err := pdf.Output(&buffer); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao gerar PDF",
		})

		return
	}

	c.Header(
		"Content-Type",
		"application/pdf",
	)

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
