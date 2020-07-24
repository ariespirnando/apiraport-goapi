package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) > 10 {
			token, _ := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("jwt_niksiswa", claims["niksiswa"])
			} else {
				c.JSON(422, gin.H{
					"error_code": "000004",
					"status":     "Invalid Token",
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(422, gin.H{
				"error_code": "000005",
				"status":     "Authorization token not provided",
			})
			c.Abort()
			return
		}
	}
}
