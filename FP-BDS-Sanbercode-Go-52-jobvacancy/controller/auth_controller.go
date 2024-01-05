package controller

import (
	"database/sql"
	"log"
	"time"

	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/config"
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/models"
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/models/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description Login with JWT
// @Tags Authentication
// @Produce json
// @Success 200 {array} models.Users
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var loginRequest request.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Query user from the database
	var user models.Users
	err := config.DB.QueryRow("SELECT id_user, username, password, role FROM users WHERE username = ?", loginRequest.Username).Scan(&user.Id_User, &user.Username, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	} else if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	log.Printf("User: %+v\n", user)

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.Id_User
	claims["username"] = user.Username
    claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

	tokenString, err := token.SignedString([]byte("Agnar123"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

// @Summary Register
// @Description Register Accout
// @Tags Authentication
// @Produce json
// @Success 200 {array} models.Users
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var registerRequest request.RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Insert user into the database
	_, err = config.DB.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", registerRequest.Username, registerRequest.Email,string(hashedPassword), registerRequest.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @Summary Change Password
// @Description Change password by id and role user
// @Tags Authentication
// @Produce json
// @Success 200 {array} models.Users
// @Router /change-password [post]
func ChangePassword(c *fiber.Ctx) error {
	// Pastikan pengguna sudah login
	if err := protectWithJWT(c, "admin", "guest"); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Bind request payload ke struct
	var changePasswordRequest request.ChangePasswordRequest
	if err := c.BodyParser(&changePasswordRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Dapatkan ID pengguna dari token JWT
	userClaims, ok := c.Locals("user").(jwt.MapClaims)
	if !ok || userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID := int(userClaims["userID"].(float64))
	userRole, ok := userClaims["role"].(string)
	if !ok || userRole == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Query pengguna dari database
	var user models.Users
	err := config.DB.QueryRow("SELECT id_user, password FROM users WHERE id_user = ?", userID).Scan(&user.Id_User, &user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Jika rolenya adalah "admin", izinkan mengubah semua kata sandi
	if userRole == "admin" {
		// Periksa kecocokan kata sandi lama
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePasswordRequest.OldPassword)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid old password"})
		}

		// Hash kata sandi baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}

		// Update kata sandi di database
		_, err = config.DB.Exec("UPDATE users SET password = ? WHERE id_user = ?", string(hashedPassword), userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}
	} else if userRole == "guest" && userID == user.Id_User {
		// Jika rolenya adalah "guest" dan ID pengguna sesuai, izinkan mengubah kata sandi
		// Periksa kecocokan kata sandi lama
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePasswordRequest.OldPassword)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid old password"})
		}

		// Hash kata sandi baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}

		// Update kata sandi di database
		_, err = config.DB.Exec("UPDATE users SET password = ? WHERE id_user = ?", string(hashedPassword), userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}
	} else {
		// Jika rolenya bukan "admin" dan bukan "guest" atau ID pengguna tidak sesuai, tolak akses
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}