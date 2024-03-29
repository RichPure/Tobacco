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
	Longitude   float64 `json:"longitude" xml:"longitude" form:"longitude"`
	Latitude    float64 `json:"latitude" xml:"latitude" form:"latitude"`
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

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := PI * lat1 / 180
	radlat2 := PI * lat2 / 180

	theta := lng1 - lng2
	radtheta := PI * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func getDistance(lat1, lon1, lat2, lon2 float64) (distance float64) {
	radius := 6378137.0
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
	minDistance := math.MaxFloat64
	if len(clients) != 0 {
		for i := 0; i < len(clients); i++ {
			dis := getDistance(req.Latitude, req.Longitude, clients[i].Latitude, clients[i].Longitude)
			if dis < minDistance {
				minClient = clients[i]
				minDistance = dis
			}
		}
	}

	fmt.Printf("min client id=%d company=%s\n", minClient.ClientId, minClient.CompanyName)
	var gridInfo GridInfo
	db.Table("grid_info").Find(&gridInfo, "grid_id = ?", req.Grid)

	resp = GridResponse{Result: 0, Capacity: gridInfo.Capacity, Applied: gridInfo.Applied,
		Population: gridInfo.Population, Distance: minDistance, CompanyName: minClient.CompanyName,
		Longitude: minClient.Longitude, Latitude: minClient.Latitude}
	jsonStr, _ := json.Marshal(resp)
	fmt.Printf("grids resp=%s\n", jsonStr)
	return
}

type resultRequest struct {
	Business           uint    `json:"business" xml:"business" form:"business"`
	RoomNum            uint    `json:"room_num" xml:"room_num" form:"room_num"`
	Grid               uint    `json:"grid" xml:"grid" form:"grid"`
	Grade              uint    `json:"grade" xml:"grade" form:"grade"`
	Latitude           float64 `json:"latitude" xml:"latitude" form:"latitude"`
	Longitude          float64 `json:"longitude" xml:"longitude" form:"longitude"`
	RegisteName        string  `json:"registe_name: xml:"registe_name" form:"registe_name"`
	RegisteIdNum       string  `json:"registe_id_num: xml:"registe_id_num" form:"registe_id_num"`
	RegisteCompanyName string  `json:"registe_company_name: xml:"registe_company_name" form:"registe_company_name"`
	RegisteAddress     string  `json:"registe_address: xml:"registe_address" form:"registe_address"`
	RegisteBelongGrid  uint    `json:"registe_belong_grid: xml:"registe_belong_grid" form:"registe_belong_grid"`
	RegisteInspector   string  `json:"registe_inspector: xml:"registe_inspector" form:"registe_inspector"`
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
