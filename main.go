package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math"
)

type GridRequest struct {
	Grid      uint64  `json:"grid" xml:"grid" form:"grid"`
	Grade     uint64  `json:"grade" xml:"grade" form:"grade"`
	Longitude float64 `json:"longitude" xml:"longitude" form:"longitude"`
	Latitude  float64 `json:"latitude" xml:"latitude" form:"latitude"`
}
type GridResponse struct {
	Result      uint64  `json:"result" xml:"result" form:"result"`
	Capacity    uint64  `json:"capacity" xml:"capacity" form:"capacity"`
	Applied     uint64  `json:"applied" xml:"applied" form:"applied"`
	Population  uint64  `json:"population" xml:"population" form:"population"`
	Distance    float64 `json:"distance" xml:"distance" form:"distance"`
	CompanyName string  `json:"company_name" xml:"company_name" form:"company_name"`
}

type ClientInfo struct {
	ClientId    uint64
	ClientName  string
	CompanyName string
	Longitude   float64
	Latitude    float64
}

type ClientGrid struct {
	ClientId uint64
	GridId   uint64
}

type GridInfo struct {
	GridId     uint64
	GridName   string
	Applied    uint64
	Capacity   uint64
	Population uint64
}

func getDistance(lat1, lon1, lat2, lon2 float64) (distance float64) {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lon1 = lon1 * rad
	lat2 = lat2 * rad
	lon2 = lon2 * rad
	theta := lon2 - lon1
	distance = math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	distance *= radius
	return
}

func handleGrid(db *gorm.DB, req *GridRequest) (resp GridResponse) {
	var cgs []ClientGrid
	db.Table("client_grid").Find(&cgs, "grid_id = ?", req.Grid)
	if len(cgs) == 0 {
		fmt.Printf("gird no found %d\n", req.Grid)
		return
	}
	var clientIds []uint64
	for _, cg := range cgs {
		clientIds = append(clientIds, cg.ClientId)
	}

	var clients []ClientInfo
	db.Table("client_info").Find(&clients, clientIds)
	var minClient ClientInfo = clients[0]
	minDistance := getDistance(req.Latitude, req.Longitude, minClient.Latitude, minClient.Longitude)
	for _, client := range clients {
		dis := getDistance(req.Latitude, req.Longitude, client.Longitude, client.Latitude)
		if dis < minDistance {
			minClient = client
			minDistance = dis
		}
	}

	fmt.Printf("min client id=%d company=%s\n", minClient.ClientId, minClient.CompanyName)
	var gridInfo GridInfo
	db.Table("grid_info").Find(&gridInfo, "grid_id = ?", req.Grid)

	resp = GridResponse{Result: 0, Capacity: gridInfo.Capacity, Applied: gridInfo.Applied,
		Population: gridInfo.Population, Distance: minDistance, CompanyName: minClient.CompanyName}
	jsonStr, _ := json.Marshal(resp)
	fmt.Printf("grids resp=%s\n", jsonStr)
	return
}

type resultRequest struct {
	Business  uint    `json:"business" xml:"business" form:"business"`
	RoomNum   uint    `json:"boomNum" xml:"boomNum" form:"boomNum"`
	Grid      uint    `json:"grid" xml:"grid" form:"grid"`
	Grade     uint    `json:"grade" xml:"grade" form:"grade"`
	Latitude  float64 `json:"latitude" xml:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" xml:"longitude" form:"longitude"`
}
type resultResponse struct {
	Result uint   `json:"result" xml:"result" form:"result"`
	Msg    string `json:"msg" xml:"msg" form:"msg"`
}

func handleResult(req *resultRequest) (resp resultResponse) {
	resp = resultResponse{Result: 0, Msg: "ok"}
	jsonStr, _ := json.Marshal(resp)
	fmt.Printf("result end resp=%s\n", jsonStr)
	return
}

var (
	dbUrl string
	dbPwd string
)

func init() {
	flag.StringVar(&dbUrl, "dbUrl", "", "data base url")
	flag.StringVar(&dbPwd, "dbPwd", "", "data base password")
}

func main() {
	flag.Parse()
	defer glog.Flush()
	app := fiber.New()
	fmt.Printf("start")

	dsn := "root:" + dbPwd + "@tcp(" + dbUrl + ":3306)/Tobacco?charset=utf8&parseTime=True&loc=Local"
	fmt.Printf("connect info=%s\n", dsn)
	db := new(gorm.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database \n")
		return
	}

	app.Post("/tobacco/v1/wx/grids/submit", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		fmt.Printf("grids req=%s\n", string(c.Request().Body()))
		req := new(GridRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		resp := handleGrid(db, req)
		return c.JSON(resp)
	})

	app.Post("/tobacco/v1/wx/result/submit", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		fmt.Printf("result req=%s\n", string(c.Request().Body()))
		req := new(resultRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		resp := handleResult(req)
		return c.JSON(resp)
	})

	app.ListenTLS(":443", "./6023104_www."+dbUrl+".pem", "./6023104_www."+dbUrl+".key")
}
