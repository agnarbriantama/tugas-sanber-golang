// utils/jwt.go

package utils

import (
	"info-loker/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// SecretKey digunakan untuk menandatangani token JWT
var SecretKey = []byte("Agnar123")

// GenerateJWTToken membuat token JWT berdasarkan informasi pengguna
func GenerateJWTToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id_user"] = user.Id_User
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// middleware/middleware.go

func ProtectWithJWT(c *fiber.Ctx) error {
	// Mendapatkan token dari header Authorization
	tokenString := c.Get("Authorization")

	// Memeriksa keberadaan token
	if tokenString == "" {
		return fiber.ErrUnauthorized
	}

	// Validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil // Ganti dengan secret key yang sesuai
	})

	if err != nil || !token.Valid {
		return fiber.ErrUnauthorized
	}

	// Mendapatkan role dari token
	userClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || userClaims == nil {
		return fiber.ErrUnauthorized
	}

	userRole, ok := userClaims["role"].(string)
	if !ok || userRole == "" {
		return fiber.ErrUnauthorized
	}

	// Memeriksa apakah role pengguna termasuk dalam role yang diizinkan
	roleAllowed := false
	for _, allowedRole := range []string{"jobseeker"} {
		if userRole == allowedRole {
			roleAllowed = true
			break
		}
	}

	if !roleAllowed {
		return fiber.ErrUnauthorized
	}

	// Menambahkan informasi token ke konteks dengan kunci "user"
	c.Locals("user", userClaims)

	return nil
}

