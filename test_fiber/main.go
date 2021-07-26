package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)


func main() {
	app := fiber.New()

	// Field names should start with an uppercase letter
	type SubmitRequest struct {
		Grid      uint    `json:"grid" xml:"grid" form:"grid"`
		Grade     uint    `json:"grade" xml:"grade" form:"grade"`
		Latitude  float64 `json:"latitude" xml:"latitude" form:"latitude"`
		Longitude float64 `json:"longitude" xml:"longitude" form:"longitude"`
	}
	type SubmitResponse struct {
		 Capacity   uint    `json:"capacity" xml:"capacity" form:"capacity"`
		 Applied    uint    `json:"applied" xml:"applied" form:"applied"`
		 Population uint64  `json:"population" xml:"population" form:"population"`
		 Distance   float64 `json:"distance" xml:"distance" form:"distance"`
	}
	app.Post("/tobacco/v1/wx/submit", func(c *fiber.Ctx) error {
		req := new(SubmitRequest)

		if err := c.BodyParser(req); err != nil {
			return err
		}
		log.Println("submit start")
		log.Println(req.Grid)
		log.Println(req.Grade)
		log.Println(req.Latitude)
		log.Println(req.Longitude)
		log.Println("submit end")

		resp := SubmitResponse{Capacity: 10, Applied: 2, Population: 888, Distance: 123.456}
		jsonStr, _ := json.Marshal(resp)
		return c.SendString(string(jsonStr))
	})

	app.Post("/tobacco/v1/wx/register", func(c *fiber.Ctx) error {
		req := new(RegisterInfo)

		if err := c.BodyParser(req); err != nil {
			return err
		}

		log.Println(req.Name) // john
		log.Println(req.Pass) // doe
	})

	app.Listen(":8888")
}
