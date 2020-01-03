package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_customerHttpDeliver "github.com/mr-fame/pd-dome-api/customer/delivery/http"
	_customerRepo "github.com/mr-fame/pd-dome-api/customer/repository"
	_customerUcase "github.com/mr-fame/pd-dome-api/customer/usecase"
	_customerRoomHttpDeliver "github.com/mr-fame/pd-dome-api/customer_room/delivery/http"
	_customerRoomRepo "github.com/mr-fame/pd-dome-api/customer_room/repository"
	_customerRoomUcase "github.com/mr-fame/pd-dome-api/customer_room/usecase"
	"github.com/mr-fame/pd-dome-api/middleware"
	"github.com/mr-fame/pd-dome-api/models"
	_roomHttpDeliver "github.com/mr-fame/pd-dome-api/room/delivery/http"
	_roomRepo "github.com/mr-fame/pd-dome-api/room/repository"
	_roomUcase "github.com/mr-fame/pd-dome-api/room/usecase"
)

var isInitDB = flag.Bool("initDB", false, "initialize data")
var isAutoMigrate = flag.Bool("autoMigrate", false, "auto migration database")
var dbConn *gorm.DB

func init() {
	flag.Parse()

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	dbConn = models.ConnectMysqlDB(*isInitDB, *isAutoMigrate)
}

func main() {
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	cr := _customerRepo.NewMysqlCustomerRepository(dbConn)
	rr := _roomRepo.NewMysqlRoomRepository(dbConn)
	crr := _customerRoomRepo.NewMysqlCustomerRoomRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	cu := _customerUcase.NewCustomerUsecase(cr, timeoutContext)
	_customerHttpDeliver.NewCustomerHandler(e, cu)
	ru := _roomUcase.NewRoomUsecase(rr, timeoutContext)
	_roomHttpDeliver.NewRoomHandler(e, ru)
	cru := _customerRoomUcase.NewCustomerRoomUsecase(crr, timeoutContext)
	_customerRoomHttpDeliver.NewCustomerRoomHandler(e, cru)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
