package main

import (
    "log"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/gofiber/fiber/v2"
)

type Dog struct {
    gorm.Model
    Name      string `json:"name"`
    Breed     string `json:"breed"`
    Age       int    `json:"age"`
    IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
}

var Database *gorm.DB
var dns string = "host=localhost user=root password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func Connect() error {
    var err error
    
    Database, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&Dog{})

    return nil
}

func GetDogs(c *fiber.Ctx) error {
    var dogs []Dog

    Database.Find(&dogs)
    return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
    id := c.Params("id")
    var dog Dog

    result := Database.Find(&dog, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
    dog := new(Dog)

    if err := c.BodyParser(dog); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    Database.Create(&dog)
    return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
    dog := new(Dog)
    id := c.Params("id")

    if err := c.BodyParser(dog); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    Database.Where("id = ?", id).Updates(&dog)
    return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
    id := c.Params("id")
    var dog Dog

    result := Database.Delete(&dog, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}

func main() {
  app := fiber.New()

  Connect()

  app.Get("/dogs", GetDogs)
  app.Get("/dogs/:id", GetDog)
  app.Post("/dogs", AddDog)
  app.Put("/dogs/:id", UpdateDog)
  app.Delete("/dogs/:id", RemoveDog)
  // app.Get("/", GetDogs)

  log.Fatal(app.Listen(":3000"))
}
