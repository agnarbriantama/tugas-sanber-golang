package controller

import (
	"errors"
	"info-loker/config"
	"info-loker/models"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// @Summary Apply Job Handler
// @Description melakukan apply job
// @Tags ApplyJob
// @Produce json
// @Success 200 {array} models.ApplyJob
// @Router /applyjob/:id_jobvacancy [post]
func ApplyJobHandler(c *fiber.Ctx) error {
    if err := protectWithJWT(c, "admin", "jobseeker"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
	// Mendapatkan id_jobvacancies dari URL parameter
	idJobVacancies, err := strconv.ParseUint(c.Params("id_jobvacancy"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_jobvacancy"})
	}

	// Mengambil id_user dari token JWT atau sesuai kebutuhan
	idUser, err := GetUserIdFromToken(c)
    if err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}

	// Memeriksa apakah job vacancy dengan ID tertentu ada
	var jobVacancy models.Jobvacancy
	if err := config.DB.First(&jobVacancy, idJobVacancies).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job vacancy not found"})
	}

	// Memeriksa status lowongan pekerjaan
	if jobVacancy.CompanyStatus == 2 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "JLowongan Kerja Sudah Tutup"})
	}

	// Memeriksa apakah job vacancy dengan ID tertentu ada
	if err := config.DB.First(&jobVacancy, idJobVacancies).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job vacancy not found"})
	}

	// Menambahkan data ke tabel apply_jobs
	applyJob := models.ApplyJob{
		User_Id: idUser,
		JobID:   uint(idJobVacancies),
		Status:  "pending", // Sesuaikan status sesuai kebutuhan
	}

	if err := config.DB.Create(&applyJob).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to apply for job"})
	}

	return c.JSON(fiber.Map{"message": "Job application successful"})
}

// @Summary Get List Apply Job 
// @Description menampilkan semuda data list apply job
// @Tags ApplyJob
// @Produce json
// @Success 200 {array} models.ApplyJob
// @Router /applyjob [get]
func GetAllApplyJobsHandler(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    jobID, err := strconv.ParseUint(c.Params("job_id"), 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid job_id"})
    }
    var apply_jobs []models.ApplyJob

    // Mengambil semua data aplikasi dari basis data
    if err := config.DB.Where("job_id = ?", jobID).Find(&apply_jobs).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve applications"})
    }

    // Mengembalikan data aplikasi sebagai respons
    return c.JSON(apply_jobs)
}


// @Summary Confirm Status Apply Job 
// @Description update status pending to applied or rejected
// @Tags ApplyJob
// @Produce json
// @Success 200 {array} models.ApplyJob
// @Router /applyjob/:id_jobvacancy/:id_apply [put]
func ApplyStatusHandler(c *fiber.Ctx) error {
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
	// Mendapatkan id_jobvacancies dan id_apply dari URL parameter
	idJobVacancies, err := strconv.ParseUint(c.Params("id_jobvacancy"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_jobvacancy"})
	}

	idApply, err := strconv.ParseUint(c.Params("id_apply"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_apply"})
	}

	// Mengambil id_user dari token JWT
	idUser, err := GetUserIdFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Memeriksa apakah job vacancy dengan ID tertentu ada
	var jobVacancy models.Jobvacancy
	if err := config.DB.First(&jobVacancy, idJobVacancies).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job vacancy not found"})
	}

	// Memeriksa apakah aplikasi dengan ID tertentu dan User_ID yang sesuai ada
	var applyJob models.ApplyJob
	if err := config.DB.First(&applyJob, idApply).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Application not found"})
	}

	// Memeriksa apakah pengguna yang meminta edit status adalah pemilik aplikasi
	if applyJob.User_Id != idUser {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

    var requestBody map[string]string
    if err := c.BodyParser(&requestBody); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
    }

    newStatus, exists := requestBody["status"]
    if !exists || (newStatus != "applied" && newStatus != "rejected") {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status parameter"})
    }

	applyJob.Status = newStatus

	// Menyimpan perubahan ke dalam database
	if err := config.DB.Save(&applyJob).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update application status"})
	}

	return c.JSON(fiber.Map{"message": "Application status updated successfully"})
}


func DeleteApplyHandler(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
   // Mendapatkan id_user dari token JWT
    idUser, err := GetUserIdFromToken(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Mendapatkan id_apply dari URL
    idApply, err := strconv.ParseUint(c.Params("id_apply"), 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_apply parameter"})
    }

    // Menghapus data apply_job secara logis berdasarkan id_apply dan user_id
    if err := config.DB.Unscoped().Where("id_apply = ? AND user_id = ?", idApply, idUser).Delete(&models.ApplyJob{}).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete application"})
    }

    // Memberikan respons sukses
    return c.JSON(fiber.Map{"message": "Application deleted successfully"})
}


// Fungsi untuk mendapatkan id_user dari token JWT
func GetUserIdFromToken(c *fiber.Ctx) (uint, error) {
	token := c.Get("Authorization")

	// Mendapatkan token tanpa "Bearer "
	jwtToken := strings.TrimPrefix(token, "Bearer ")

	// Parse token JWT
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	// Mendapatkan id_user dari claim JWT
	idUserFloat64, ok := claims["id_user"].(float64)
	if !ok {
		return 0, errors.New("Invalid id_user claim in token")
	}

	idUser := uint(idUserFloat64)
	return idUser, nil
}



func protectWithJWT(c *fiber.Ctx, allowedRoles ...string) error {
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
    for _, allowedRole := range allowedRoles {
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
