package main

import (
	"fmt"

	"github.com/gk/go-fiber-crm/database"
	"github.com/gk/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func InitDatabase() {

	var err error

	database.DBConn,err = gorm.Open("sqlite3","leads.db")
	if err != nil {
		 panic("Failed to connect Database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database migrated")
}

func main() {

	app := fiber.New()
	InitDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
