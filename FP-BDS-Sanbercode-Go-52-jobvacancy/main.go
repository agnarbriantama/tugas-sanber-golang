package main

import (
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/config"
	_ "github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/docs"
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
// Inisialisasi database
config.InitDatabase()

app := fiber.New()


// Middleware Logger
app.Use(logger.New())

routes.SetupRoutes(app)


// Menjalankan server di port 8080
err := app.Listen(":8080")
if err != nil {
	panic("Failed to start server")
}
}
