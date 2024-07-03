package main

import (
	"fmt"

	"github.com/HaDiizz/database"
	m "github.com/HaDiizz/models"
	"github.com/HaDiizz/routes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Dogs{}, &m.Companies{}, &m.UserProfiles{})
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.ExampleRoutes(app)

	app.Listen(":3000")
}
