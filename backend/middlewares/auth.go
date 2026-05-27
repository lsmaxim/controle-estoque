package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SEGREDO_SUPER_FORTE")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader :=
			c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"erro": "Token não enviado",
			})

			c.Abort()

			return
		}

		tokenString :=
			strings.Replace(
				authHeader,
				"Bearer ",
				"",
				1,
			)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				return jwtKey, nil
			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"erro": "Token inválido",
			})

			c.Abort()

			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok {

			c.Set("email", claims["email"])
		}
		c.Next()
	}
}
