package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

)

var jwtKey = []byte("SEGREDO_SUPER_FORTE")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {

	var dados struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&dados); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})

		return
	}

	// LOGIN FIXO
	if dados.Email != "informatica@xamegobom.com.br" ||
		dados.Senha != "#@control@#" {

		c.JSON(http.StatusUnauthorized, gin.H{
			"erro": "Usuário ou senha inválidos",
		})

		return
	}

	expirationTime :=
		time.Now().Add(2 * time.Hour)

	claims := &Claims{
		Email: dados.Email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err :=
		token.SignedString(jwtKey)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao gerar token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
