package main

import (
	"log"
	"main/config"
	"main/handlers"

	"github.com/gofiber/fiber/v2"
)

// var Database *gorm.DB
// var dns string = "host=localhost user=root password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// func Connect() error {
//     var err error

//     Database, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

//     if err != nil {
//         panic(err)
//     }

//     Database.AutoMigrate(&Dog{})

//     return nil
// }

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/dogs", handlers.GetDogs)
	app.Get("/dogs/:id", handlers.GetDog)
	app.Post("/dogs", handlers.AddDog)
	app.Put("/dogs/:id", handlers.UpdateDog)
	app.Delete("/dogs/:id", handlers.RemoveDog)

	// app.Get("/", GetDogs)

	log.Fatal(app.Listen(":3000"))
}
