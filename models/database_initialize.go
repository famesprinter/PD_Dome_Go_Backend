package models

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

// ConnectMysqlDB connecnt database
func ConnectMysqlDB(isInit bool, isAutoMigrate bool) *gorm.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	var err error
	db, err = gorm.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if isInit {
		autoMigrate()
		addForeignKey()
		initDatabase()
	} else if isAutoMigrate {
		autoMigrate()
	}

	return db
}

func autoMigrate() {
	fmt.Println("Migrating...")

	// Create Table
	// db.Exec(`ALTER DATABASE ` + `pd_dorm_db` + ` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;`)
	// Customer
	db.AutoMigrate(&Customer{})
	// Room
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Rent{})
	db.AutoMigrate(&InsuranceFee{})
	db.AutoMigrate(&Level{})
	// Customer + Room
	db.AutoMigrate(&CustomerRoom{})
	db.AutoMigrate(&CustomerRoomStatus{})
	// Receipt
	db.AutoMigrate(&Receipt{})
	db.AutoMigrate(&ReceiptStatus{})
	db.AutoMigrate(&ElectricUnitCal{})
	db.AutoMigrate(&WaterUnitCal{})
}

func addForeignKey() {
	fmt.Println("Adding ForeignKey...")

	// Room
	db.Model(&Room{}).AddForeignKey("rent_id", "rents(id)", "RESTRICT", "RESTRICT")
	db.Model(&Room{}).AddForeignKey("insurance_fee_id", "insurance_fees(id)", "RESTRICT", "RESTRICT")
	db.Model(&Room{}).AddForeignKey("level_id", "levels(id)", "RESTRICT", "RESTRICT")
	// Customer + Room
	db.Model(&CustomerRoom{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "RESTRICT")
	db.Model(&CustomerRoom{}).AddForeignKey("room_id", "rooms(id)", "RESTRICT", "RESTRICT")
	// Receipt
	db.Model(&Receipt{}).AddForeignKey("receipt_status_id", "receipt_statuses(id)", "RESTRICT", "RESTRICT")
	db.Model(&Receipt{}).AddForeignKey("room_id", "rooms(id)", "RESTRICT", "RESTRICT")
	db.Model(&Receipt{}).AddForeignKey("electric_unit_cal_id", "electric_unit_cals(id)", "RESTRICT", "RESTRICT")
	db.Model(&Receipt{}).AddForeignKey("water_unit_cal_id", "water_unit_cals(id)", "RESTRICT", "RESTRICT")
}

func initDatabase() {
	fmt.Println("Initial Database...")

	// Room
	{
		price := 1500.00
		isActice := true
		rent := Rent{
			ID:       1,
			Price:    &price,
			IsActive: &isActice,
		}
		db.Create(&rent)
	}
	// CustomerRoomStatus
	{
		waitingIn := "waiting_in"
		online := "online"
		waitingOut := "waiting_out"
		offline := "offline"

		waitingInStatus := CustomerRoomStatus{
			ID:   1,
			Name: &waitingIn,
		}
		onlineStatus := CustomerRoomStatus{
			ID:   2,
			Name: &online,
		}
		waitingOutStatus := CustomerRoomStatus{
			ID:   3,
			Name: &waitingOut,
		}
		offlineStatus := CustomerRoomStatus{
			ID:   4,
			Name: &offline,
		}

		db.Create(&waitingInStatus)
		db.Create(&onlineStatus)
		db.Create(&waitingOutStatus)
		db.Create(&offlineStatus)
	}
	// ReceiptStatus
	{
		unpaid := "unpaid"
		paid := "paid"

		receiptUnPaidStatus := ReceiptStatus{
			ID:   1,
			Name: &unpaid,
		}
		receiptPaidStatus := ReceiptStatus{
			ID:   2,
			Name: &paid,
		}
		db.Create(&receiptUnPaidStatus)
		db.Create(&receiptPaidStatus)
	}
}
