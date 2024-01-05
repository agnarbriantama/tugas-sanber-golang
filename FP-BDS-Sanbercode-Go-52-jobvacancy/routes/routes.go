package routes

import (
	"github.com/agnarbriantama/tugas-sanber-golang/FP-BDS-Sanbercode-Go-52-jobvacancy/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /job-vacancy-api/v1

func SetupRoutes(app *fiber.App){
	f := fiber.New()
	g:= f.Group("/job-vacancy-api/v1")

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))



	g.Get("/job-vacancy", controller.GetAllJobvacancy)
	app.Post("/job-vacancy",controller.CreateJobvacancy)
	app.Get("/job-vacancy/:id", controller.GetJobVacancyDetails)
	app.Put("/job-vacancy/:id", controller.UpdateJobvacancy)
	app.Delete("/job-vacancy/:id", controller.DeleteJobvacancy)

	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)
	app.Post("/change-password", controller.ChangePassword)

	app.Post("/apply_job/:id", controller.PostApplyJob)
	app.Get("/all_apply", controller.GetAllApplyJob)
	app.Get("/apply_job/:id_user", controller.GetApplyJobByUserID)
	app.Delete("/apply_job/:id_apply", controller.DeleteJobApply)
	app.Get("/apply_by_job/:id_job", controller.GetApplyJobByJobID)

	
}