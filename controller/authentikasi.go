package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apiraport/config"
	"github.com/apiraport/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var json model.Requestlgoin
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "000002",
			"status":     "Bad Request",
		})
	} else {

		data := []byte(json.Password)
		b := md5.Sum(data)

		pass := hex.EncodeToString(b[:])
		var count int

		config.DB.Table("tbl_siswa").
			Where("tbl_siswa.niksiswa = ? AND tbl_siswa.password = ?", json.Username, pass).
			Count(&count)
		if count > 0 {
			var jwtToken = createToken(json.Username)
			c.JSON(http.StatusOK, gin.H{
				"error_code": "000000",
				"status":     "Autentikasi Berhasil",
				"token":      jwtToken,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code": "000003",
				"status":     "Username dan Password tidak ditemukan",
			})
		}
	}
}

func createToken(niksiswa string) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"niksiswa": niksiswa,
		"exp":      time.Now().AddDate(0, 0, 7).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
