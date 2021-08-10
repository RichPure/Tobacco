package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/glog"
)

//ccc

func main() {
	flag.Parse()
	defer glog.Flush()

	app := fiber.New()
	fmt.Printf("start")
	type SubmitRequest struct {
		Grid      uint    `json:"grid" xml:"grid" form:"grid"`
		Grade     uint    `json:"grade" xml:"grade" form:"grade"`
		Latitude  float64 `json:"latitude" xml:"latitude" form:"latitude"`
		Longitude float64 `json:"longitude" xml:"longitude" form:"longitude"`
	}
	type SubmitResponse struct {
		Result      uint    `json:"result" xml:"result" form:"result"`
		Capacity    uint    `json:"capacity" xml:"capacity" form:"capacity"`
		Applied     uint    `json:"applied" xml:"applied" form:"applied"`
		Population  uint64  `json:"population" xml:"population" form:"population"`
		Distance    float64 `json:"distance" xml:"distance" form:"distance"`
		CompanyName string  `json:"company_name" xml:"company_name" form:"company_name"`
	}

	app.Post("/tobacco/v1/wx/grids/submit", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		fmt.Printf("grids req=%s\n", string(c.Request().Body()))
		req := new(SubmitRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		resp := SubmitResponse{Result: 0, Capacity: 10, Applied: 2, Population: 688, Distance: 123.456, CompanyName: "哈哈哈哈哈哈"}
		jsonStr, _ := json.Marshal(resp)
		fmt.Printf("grids resp=%s\n", jsonStr)
		return c.JSON(resp)
	})

	type RegisterRequest struct {
		Business  uint    `json:"business" xml:"business" form:"business"`
		RoomNum   uint    `json:"boomNum" xml:"boomNum" form:"boomNum"`
		Grid      uint    `json:"grid" xml:"grid" form:"grid"`
		Grade     uint    `json:"grade" xml:"grade" form:"grade"`
		Latitude  float64 `json:"latitude" xml:"latitude" form:"latitude"`
		Longitude float64 `json:"longitude" xml:"longitude" form:"longitude"`
	}
	type RegisterResponse struct {
		Result uint   `json:"result" xml:"result" form:"result"`
		Msg    string `json:"msg" xml:"msg" form:"msg"`
	}

	app.Post("/tobacco/v1/wx/result/submit", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		fmt.Printf("result req=%s\n", string(c.Request().Body()))
		req := new(RegisterRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		resp := RegisterResponse{Result: 0, Msg: "ok"}
		jsonStr, _ := json.Marshal(resp)
		fmt.Printf("result end resp=%s\n", jsonStr)
		return c.JSON(resp)
	})

	app.ListenTLS(":443", "./6023104_www.jivvsvy.cn.pem", "./6023104_www.jivvsvy.cn.key")
}
