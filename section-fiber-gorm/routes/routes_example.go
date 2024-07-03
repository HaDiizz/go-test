package routes

import (
	c "github.com/HaDiizz/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func ExampleRoutes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	profile := v1.Group("/profile")
	profile.Get("", c.GetProfiles)
	profile.Get("/filter", c.GetProfile)
	profile.Get("/ages", c.GetProfileAnyAges)
	profile.Get("/user", c.SearchData)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}))

	// CRUD profiles
	profile.Post("/", c.AddProfile)
	profile.Put("/:id", c.UpdateProfile)
	profile.Delete("/:id", c.RemoveProfile)

	v3 := api.Group("/v3")

	v1.Get("/hello", c.HelloTest)
	v1.Post("/fact/:num", c.Factorial)
	v1.Post("/register", c.Register)
	v3.Post("/dis", c.AsciiGenerator)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/deleted", c.GetDeletedDogs)
	dog.Get("/range", c.GetDogsRangeCountByDogId)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJsonSummary)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	// dog.Get("/json", c.GetDogsJson)

	// CRUD companies
	company := v1.Group("/company")
	company.Get("", c.GetCompanies)
	company.Get("/filter", c.GetCompany)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)
}
