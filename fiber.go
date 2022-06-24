package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}

func GetHandler(c *fiber.Ctx) error {
	return  c.SendString("welcome home:" + c.Params("name"))
}

func PostHandler(c *fiber.Ctx) error {
	person:= new(Person) 

	if err:= c.BodyParser(person); err != nil {
		return err
	}
	fmt.Println(person.Name, person.Email)
	return nil
}

func main() {
	app:= fiber.New()  

	app.Get("/api/:name", GetHandler)

	log.Fatal(app.Listen(":3001"))

}