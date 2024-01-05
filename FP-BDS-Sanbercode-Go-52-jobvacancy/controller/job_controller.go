package controller

import (
	"log"

	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/config"
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all job vacancies
// @Description Get a list of all job vacancies
// @Tags Job Vacancy
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy [get]
func GetAllJobvacancy(c *fiber.Ctx) error {
	rows, err := config.DB.Query("SELECT * FROM tb_listjob")
	if err != nil {
		log.Fatal(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	var job []models.Jobvacancy
	for rows.Next() {
		var m models.Jobvacancy
		if err := rows.Scan(&m.ID, &m.Title, &m.CompanyName, &m.CompanyDesc, &m.CompanySalary, &m.CompanyStatus); err != nil {
			log.Fatal(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		job = append(job, m)
	}

	return c.JSON(job)
}

// @Summary Create job vacancies
// @Description Create a new job
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy [post]
func CreateJobvacancy(c *fiber.Ctx) error {
	if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

	var job models.Jobvacancy
    if err := c.BodyParser(&job); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    result, err := config.DB.Exec("INSERT INTO tb_listjob (title, company_name, company_desc, company_salary, company_status) VALUES (?, ?, ?, ?, ?)",
        job.Title, job.CompanyName, job.CompanyDesc, job.CompanySalary, job.CompanyStatus)
    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    id, err := result.LastInsertId()
    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    job.ID = int(id)

    return c.JSON(job)
}

// @Summary Update job vacancy
// @Description Update job
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy/:id [put]
func UpdateJobvacancy(c *fiber.Ctx) error {
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    var updatedJob models.Jobvacancy
    if err := c.BodyParser(&updatedJob); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }


    jobID := c.Params("id")
    
    _, err := config.DB.Exec("UPDATE tb_listjob SET title = ?, company_name = ?, company_desc = ?, company_salary = ?, company_status = ? WHERE id = ?",
        updatedJob.Title, updatedJob.CompanyName, updatedJob.CompanyDesc, updatedJob.CompanySalary, updatedJob.CompanyStatus, jobID)
    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(updatedJob)
}

// @Summary Detail Job Vacancy
// @Description Detail Job Vacancy
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy/:id [get]
func GetJobVacancyDetails(c *fiber.Ctx) error {

    if err := protectWithJWT(c, "admin", "guest"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Harus Login Dulu"})
    }

    jobID := c.Params("id")

    var job models.Jobvacancy

    // Fetch job details from the database
    err := config.DB.QueryRow("SELECT * FROM tb_listjob WHERE id = ?", jobID).
        Scan(&job.ID, &job.Title, &job.CompanyName, &job.CompanyDesc, &job.CompanySalary, &job.CompanyStatus)

    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(job)
}

// @Summary Delete Job Vacancy
// @Description Delete a job from the database
// @Tags Job Vacancy 
// @Produce json
// @Success 200 {array} models.Jobvacancy
// @Router /job-vacancy/:id [delete]
func DeleteJobvacancy(c *fiber.Ctx) error {
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    jobID := c.Params("id")

    _, err := config.DB.Exec("DELETE FROM tb_listjob WHERE id = ?", jobID)
    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.SendStatus(fiber.StatusOK)
}


// Middleware JWT untuk melindungi endpoint
func protectWithJWT(c *fiber.Ctx, allowedRoles ...string) error {
    // Mendapatkan token dari header Authorization
    tokenString := c.Get("Authorization")

    // Memeriksa keberadaan token
    if tokenString == "" {
        return fiber.ErrUnauthorized
    }

    // Validasi token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("Agnar123"), nil // Ganti dengan secret key yang sesuai
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


