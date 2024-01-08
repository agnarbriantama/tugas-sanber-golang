package controller

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"info-loker/config"
	"info-loker/models"
	"info-loker/models/request"
	"info-loker/utils"
)

// @Summary Register
// @Description Regisrasi akun dengan role admin dan jobseeker
// @Tags Authentication
// @Produce json
// @Success 200 {array} models.User
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validasi apakah email atau username sudah digunakan
	if isDuplicate := isDuplicateUser(user.Email, user.Username); isDuplicate {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email or username already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Set hashed password
	user.Password = string(hashedPassword)

	// Set default role jika tidak diisi
	if user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role is required"})
	}

	// Cek apakah nilai role yang diinginkan adalah "admin" atau "jobseeker"
	if user.Role != "admin" && user.Role != "jobseeker" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role value"})
	}

	// Set waktu pembuatan dan pembaruan
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert user into the database
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @Summary Login
// @Description Login akun dengan auth JWT
// @Tags Authentication
// @Produce json
// @Success 200 {array} models.User
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var loginRequest request.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Query user from the database by username
	var user models.User
	if err := config.DB.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWTToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Return JWT token in the response
	return c.JSON(fiber.Map{"token": token})
}

// @Summary Change Password
// @Description Role admin bisa mengganti password semua user, role job seeker by id
// @Tags Authentication
// @Produce json
// @Success 200 {array} models.User
// @Router /change-password [post]
func ChangePassword(c *fiber.Ctx) error {
	// Mendapatkan role dari token JWT
	role, err := GetRoleFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Mendapatkan id_user dari token JWT
	idUser, err := GetUserIdFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Mendapatkan id_user yang akan diubah dari parameter URL
	idUserToChange := c.Params("id_user")

	// Validasi apakah role adalah "admin" atau "jobseeker"
	if role == "admin" || role == "jobseeker" {
		if role == "jobseeker" && strconv.FormatUint(uint64(idUser), 10) != idUserToChange {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		oldPassword := c.FormValue("old_password")
		newPassword := c.FormValue("new_password")

		var user models.User
		if err := config.DB.Where("id_user = ?", idUserToChange).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// Membandingkan password yang dimasukkan dengan hash yang tersimpan
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid old password"})
		}

		// Hash password baru sebelum disimpan
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash new password"})
		}

		// Update password baru ke basis data
		user.Password = string(hashedPassword)
		if err := config.DB.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to change password"})
		}

		return c.JSON(fiber.Map{"message": "Password changed successfully"})
	}

	// Jika role tidak sesuai, beri respons Unauthorized
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}
func isDuplicateUser(email, username string) bool {
	var existingUser models.User
	// Cek apakah email sudah ada
	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return true
	}
	// Cek apakah username sudah ada
	if err := config.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return true
	}
	return false
}

func GetRoleFromToken(c *fiber.Ctx) (string, error) {
	// Mendapatkan token dari header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is missing")
	}

	// Mengambil nilai token dari header Authorization
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("Failed to parse token: %v", err)
	}

	// Mendapatkan nilai peran (role) dari claim token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid token claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("Role claim is missing or invalid")
	}

	return role, nil
}