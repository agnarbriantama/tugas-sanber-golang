package controller

import (
	"log"
	"strconv"

	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/config"
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// @Summary Post Apply Job
// @Description Post apply job
// @Tags ApplyJob
// @Produce json
// @Success 200 {array} models.Apply
// @Router /apply_job/:id [post]
func PostApplyJob(c *fiber.Ctx) error {
    // Pastikan pengguna sudah login
    if err := protectWithJWT(c, "guest"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Dapatkan ID pekerjaan dari parameter atau payload (sesuai dengan kebutuhan Anda)
    jobID := c.Params("id")

    // Dapatkan ID pengguna dari token JWT
    userClaims, ok := c.Locals("user").(jwt.MapClaims)
    if !ok || userClaims == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    userID := int(userClaims["userID"].(float64))

    // Dapatkan company_status dari tb_listjob
    var companyStatus int
    err := config.DB.QueryRow("SELECT company_status FROM tb_listjob WHERE id = ?", jobID).Scan(&companyStatus)
    if err != nil {
        log.Fatal(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
    }

    // Periksa company_status dan tambahkan entri baru ke tabel tb_apply jika sesuai
    if companyStatus == 1 {
        _, err := config.DB.Exec("INSERT INTO tb_apply (id, id_user, status_lamaran) VALUES (?, ?, 'pending')", jobID, userID)
        if err != nil {
            log.Fatal(err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
        }

        return c.JSON(fiber.Map{"message": "Job applied successfully"})
    } else if companyStatus == 2 {
        return c.JSON(fiber.Map{"message": "Lowongan sudah ditutup"})
    } else {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Status perusahaan tidak valid"})
    }
}

// @Summary Get All Apply Job
// @Description Display all applied data
// @Tags ApplyJob
// @Produce json
// @Success 200 {array} models.Apply
// @Router /all_apply [get]
func GetAllApplyJob(c *fiber.Ctx) error {
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    rows, err := config.DB.Query(`
        SELECT 
            tb_apply.id_apply, 
            tb_apply.id,
            tb_apply.id_user,
            tb_listjob.id AS job_id, 
            tb_listjob.title AS job_title, 
            tb_listjob.company_name, 
            tb_listjob.company_desc, 
            tb_listjob.company_salary, 
            tb_listjob.company_status, 
            users.id_user, 
            users.username 
        FROM 
            tb_apply
        JOIN 
            tb_listjob ON tb_apply.id = tb_listjob.id
        JOIN 
            users ON tb_apply.id_user = users.id_user;
    `)

    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }
    defer rows.Close()

    var applyjob []models.Apply
    for rows.Next() {
        var app models.Apply
        if err := rows.Scan(
            &app.IDApply,
            &app.ID,
            &app.IDUser,
            &app.JobID,
            &app.JobTitle,
            &app.CompanyName,
            &app.CompanyDesc,
            &app.CompanySalary,
            &app.CompanyStatus,
            &app.UserID,
            &app.Username,
        ); err != nil {
            log.Fatal(err)
            return c.SendStatus(fiber.StatusInternalServerError)
        }
        applyjob = append(applyjob, app)
    }

    return c.JSON(applyjob)
}

func GetApplyJobByUserID(c *fiber.Ctx) error {
    if err := protectWithJWT(c, "guest"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Dapatkan id_user dari token JWT
    userClaims, ok := c.Locals("user").(jwt.MapClaims)
    if !ok || userClaims == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    userIDInToken := int(userClaims["userID"].(float64))

    // Dapatkan id_user dari parameter URL
    idUserStr := c.Params("id_user")
    idUser, err := strconv.Atoi(idUserStr)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
    }

    // Periksa apakah pengguna memiliki hak akses ke data aplikasi
    if userIDInToken != idUser {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    rows, err := config.DB.Query(`
        SELECT 
            tb_apply.id_apply, 
            tb_apply.id,
            tb_apply.id_user,
            tb_listjob.id AS job_id, 
            tb_listjob.title AS job_title, 
            tb_listjob.company_name, 
            tb_listjob.company_desc, 
            tb_listjob.company_salary, 
            tb_listjob.company_status, 
            users.id_user, 
            users.username 
        FROM 
            tb_apply
        JOIN 
            tb_listjob ON tb_apply.id = tb_listjob.id
        JOIN 
            users ON tb_apply.id_user = users.id_user
        WHERE 
            users.id_user = ?;
    `, idUser)

    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }
    defer rows.Close()

    var applyjob []models.Apply
    for rows.Next() {
        var app models.Apply
        if err := rows.Scan(
            &app.IDApply,
            &app.ID,
            &app.IDUser,
            &app.JobID,
            &app.JobTitle,
            &app.CompanyName,
            &app.CompanyDesc,
            &app.CompanySalary,
            &app.CompanyStatus,
            &app.UserID,
            &app.Username,
        ); err != nil {
            log.Fatal(err)
            return c.SendStatus(fiber.StatusInternalServerError)
        }
        applyjob = append(applyjob, app)
    }

    return c.JSON(applyjob)
}

// @Summary Delete Job Apply
// @Description Delete a job 
// @Tags ApplyJob
// @Produce json
// @Success 200 {array} models.Apply
// @Router /apply_job/:id_apply [delete]
func DeleteJobApply(c *fiber.Ctx) error {
    // Pastikan pengguna sudah login
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Dapatkan ID aplikasi dari parameter URL
    idApplyStr := c.Params("id_apply")
    idApply, err := strconv.Atoi(idApplyStr)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    // Hapus entri aplikasi dari database
    _, err = config.DB.Exec("DELETE FROM tb_apply WHERE id_apply = ?", idApply)
    if err != nil {
        log.Fatal(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
    }

    return c.JSON(fiber.Map{"message": "Application deleted successfully"})
}

func UpdateJobApply(c *fiber.Ctx) error {
    // Pastikan pengguna sudah login dan memiliki akses sebagai admin
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Dapatkan ID aplikasi dari parameter URL
    idApplyStr := c.Params("id_apply")
    idApply, err := strconv.Atoi(idApplyStr)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    // Bind request payload ke struct ApplyStatus
    var updatedApplyStatus models.ApplyStatus
    if err := c.BodyParser(&updatedApplyStatus); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    // Perbarui data aplikasi di database
    _, err = config.DB.Exec("UPDATE tb_apply SET status_lamaran = ? WHERE id_apply = ?",
        updatedApplyStatus.Status, idApply)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
    }

    return c.JSON(fiber.Map{"message": "Application status updated successfully"})
}

func GetApplyJobByJobID(c *fiber.Ctx) error {
    // Pastikan pengguna sudah login
    if err := protectWithJWT(c, "admin"); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Dapatkan id_user dari token JWT
    userClaims, ok := c.Locals("user").(jwt.MapClaims)
    if !ok || userClaims == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Dapatkan id_job dari parameter URL
    idJobStr := c.Params("id_job")
    idJob, err := strconv.Atoi(idJobStr)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Job ID"})
    }

    // Query SQL untuk mendapatkan data aplikasi berdasarkan ID tb_listjob
    rows, err := config.DB.Query(`
        SELECT 
            tb_apply.id_apply, 
            tb_apply.id,
            tb_apply.id_user,
            tb_listjob.id AS job_id, 
            tb_listjob.title AS job_title, 
            tb_listjob.company_name, 
            tb_listjob.company_desc, 
            tb_listjob.company_salary, 
            tb_listjob.company_status, 
            users.id_user, 
            users.username 
        FROM 
            tb_apply
        JOIN 
            tb_listjob ON tb_apply.id = tb_listjob.id
        JOIN 
            users ON tb_apply.id_user = users.id_user
        WHERE 
            tb_listjob.id = ?;
    `, idJob)

    if err != nil {
        log.Fatal(err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }
    defer rows.Close()

    var applyjob []models.Apply
    for rows.Next() {
        var app models.Apply
        if err := rows.Scan(
            &app.IDApply,
            &app.ID,
            &app.IDUser,
            &app.JobID,
            &app.JobTitle,
            &app.CompanyName,
            &app.CompanyDesc,
            &app.CompanySalary,
            &app.CompanyStatus,
            &app.UserID,
            &app.Username,
        ); err != nil {
            log.Fatal(err)
            return c.SendStatus(fiber.StatusInternalServerError)
        }
        applyjob = append(applyjob, app)
    }

	 // Periksa apakah tidak ada data ditemukan
	 if len(applyjob) == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tidak ada job yang apply"})
    }

    return c.JSON(applyjob)
}

