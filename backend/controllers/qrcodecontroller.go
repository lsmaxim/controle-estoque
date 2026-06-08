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

	// ==========================
	// URL QUE O QR VAI ABRIR
	// ==========================
	url := fmt.Sprintf(
		"http://192.168.0.7:5173/equipamento/%s",
		id,
	)

	// ==========================
	// NOME DO ARQUIVO
	// ==========================
	nomeArquivo := fmt.Sprintf(
		"equipamento_%s.png",
		id,
	)

	// ==========================
	// PASTA DOS QRCODES
	// ==========================
	err := os.MkdirAll(
		"uploads/qrcodes",
		os.ModePerm,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// ==========================
	// CAMINHO COMPLETO
	// ==========================
	caminho := filepath.Join(
		"uploads",
		"qrcodes",
		nomeArquivo,
	)

	// ==========================
	// GERA QR CODE
	// ==========================
	err = qrcode.WriteFile(
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

	// ==========================
	// RETORNO
	// ==========================
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "QR Code gerado com sucesso",
		"arquivo":  nomeArquivo,
		"url_qr":   url,
		"imagem":   "/uploads/qrcodes/" + nomeArquivo,
	})
}
