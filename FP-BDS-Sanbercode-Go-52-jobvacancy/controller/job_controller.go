package controller

import (
	"info-loker/config"
	"info-loker/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get all job vacancies
// @Description Menampilkan semua list job vacancy
// @Tags Job Vacancy
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy [get]
func GetAllJobVacancy(c *fiber.Ctx) error {
	// Membuat slice untuk menampung hasil query
	var jobVacancies []models.Jobvacancy
	// Melakukan query ke database untuk mengambil semua job vacancy
	if err := config.DB.Find(&jobVacancies).Error; err != nil {
		// Jika terjadi error, mengembalikan response error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	// Mengembalikan response JSON berisi daftar job vacancy
	return c.JSON(fiber.Map{"Data": jobVacancies})
}

// @Summary Detail Job Vacancy
// @Description Detail Job Vacancy
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy/:id [get]
func GetJobVacancyByID(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
	// Mendapatkan ID job vacancy dari parameter URL
	id := c.Params("id")

	// Mencari job vacancy berdasarkan ID
	var jobVacancy models.Jobvacancy
	if err := config.DB.First(&jobVacancy, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job Vacancy not found"})
	}

	// Mengembalikan response JSON dengan detail job vacancy
	return c.JSON(fiber.Map{"jobVacancy": jobVacancy})
}

// @Summary Create job vacancy
// @Description Create a new job
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy [post]
func CreateJobVacancy(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    // Bind request payload ke struct Jobvacancy
    var jobVacancy models.Jobvacancy
    if err := c.BodyParser(&jobVacancy); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    // Melakukan insert data ke dalam database
    if err := config.DB.Create(&jobVacancy).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
    }

    // Mengembalikan response JSON berisi data job vacancy yang telah ditambahkan
    return c.JSON(fiber.Map{"jobVacancy": jobVacancy})
}

// @Summary Update job vacancy
// @Description Update job
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy/:id [put]
func UpdateJobVacancy(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    // Mendapatkan ID job vacancy dari parameter URL
    id := c.Params("id")

    // Mencari job vacancy berdasarkan ID
    var jobVacancy models.Jobvacancy
    if err := config.DB.First(&jobVacancy, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job Vacancy not found"})
    }

    // Bind request payload ke struct Jobvacancy
    if err := c.BodyParser(&jobVacancy); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    // Memperbarui data job vacancy di dalam database
    if err := config.DB.Save(&jobVacancy).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
    }

    // Mengembalikan response JSON berisi data job vacancy yang telah diperbarui
    return c.JSON(fiber.Map{"jobVacancy": jobVacancy, "message" : "Berhasil ditambahkan"})
}

// @Summary Delete Job Vacancy
// @Description Delete a job from the database
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy/:id [delete]
func DeleteJobVacancy(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
	// Mendapatkan ID job vacancy dari parameter URL
	id := c.Params("id")

	// Mencari job vacancy berdasarkan ID
	var jobVacancy models.Jobvacancy
	if err := config.DB.First(&jobVacancy, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job Vacancy not found"})
	}

	// Menghapus data job vacancy dari database
	if err := config.DB.Unscoped().Delete(&jobVacancy).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Mengembalikan response JSON sebagai konfirmasi penghapusan
	return c.JSON(fiber.Map{"message": "Job Vacancy deleted successfully"})
}