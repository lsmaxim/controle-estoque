package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func GerarEtiqueta(c *gin.Context) {

	id := c.Param("id")

	// EXEMPLO TEMPORÁRIO
	// depois vamos buscar do banco
	codigoProduto := "PC-0001"
	equipamento := "COMPUTADOR"
	setor := "TI"

	// CRIA PASTAS
	os.MkdirAll("uploads/etiquetas", os.ModePerm)

	// NOME PDF
	nomeArquivo := fmt.Sprintf(
		"etiqueta_%s.pdf",
		id,
	)

	caminhoPDF := fmt.Sprintf(
		"uploads/etiquetas/%s",
		nomeArquivo,
	)

	// CAMINHO QR CODE
	qrPath := fmt.Sprintf(
		"uploads/qrcodes/%s.png",
		id,
	)

	// PDF PERSONALIZADO
	pdf := gofpdf.NewCustom(
		&gofpdf.InitType{
			UnitStr: "mm",
			Size: gofpdf.SizeType{
				Wd: 80,
				Ht: 50,
			},
		},
	)

	pdf.SetMargins(4, 4, 4)

	pdf.AddPage()

	// TÍTULO
	pdf.SetFont("Arial", "B", 11)

	pdf.CellFormat(
		0,
		5,
		"ETIQUETA PATRIMONIAL",
		"",
		1,
		"C",
		false,
		0,
		"",
	)

	pdf.Ln(7)

	// DADOS
	pdf.SetFont("Arial", "", 9)

	pdf.Cell(0, 4, "Codigo: "+codigoProduto)

	pdf.Ln(4)

	pdf.Cell(0, 4, "Equipamento: "+equipamento)

	pdf.Ln(4)

	pdf.Cell(0, 4, "Setor: "+setor)

	pdf.Ln(5)

	// QR CODE CENTRALIZADO
	if _, err := os.Stat(qrPath); err == nil {

		pdf.Image(
			qrPath,
			25,
			24,
			25,
			25,
			false,
			"",
			0,
			"",
		)
	}

	// CÓDIGO ABAIXO DO QR
	//pdf.SetY(48)

	//pdf.SetFont("Arial", "B", 8)

	//pdf.CellFormat(
	//0,
	//	4,
	//	codigoProduto,
	//	"",
	//	0,
	//	"C",
	//	false,
	//	0,
	//	"",
	//)

	// SALVA PDF
	err := pdf.OutputFileAndClose(caminhoPDF)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// RETORNA PDF
	c.File(caminhoPDF)
}
