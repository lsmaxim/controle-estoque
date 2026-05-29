package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func GerarQRCode(c *gin.Context) {

	id := c.Param("id")

	// URL que será aberta no QR
	url := fmt.Sprintf(
		"http://localhost:5173/equipamento/%s",
		id,
	)

	// Nome do arquivo
	nomeArquivo := fmt.Sprintf("%s.png", id)

	// Caminho
	caminho := filepath.Join(
		"uploads",
		"qrcodes",
		nomeArquivo,
	)

	// Cria pasta
	os.MkdirAll("uploads/qrcodes", os.ModePerm)

	// Gera QR
	err := qrcode.WriteFile(
		url,
		qrcode.Medium,
		256,
		caminho,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "QR Code gerado",
		"arquivo":  nomeArquivo,
		"url":      "/uploads/qrcodes/" + nomeArquivo,
	})
}
