package main

import (
	"info-loker/config"
	"info-loker/config/migration"

	"info-loker/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi koneksi ke database dan melakukan auto migrate
	config.InitDatabase()
	migration.Migration()

	app := fiber.New()

	routes.SetupRoutes(app)


	err := app.Listen(":8080")
	if err != nil {
	panic("Failed to start server")
}
}
