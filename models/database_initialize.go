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

	// Customer
	db.AutoMigrate(&Customer{})
	// Room
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Rent{})
	db.AutoMigrate(&InsuranceFee{})
	db.AutoMigrate(&Level{})
}

func addForeignKey() {
	fmt.Println("Adding ForeignKey...")

	// Room
	db.Model(&Room{}).AddForeignKey("rent_id", "rents(id)", "RESTRICT", "RESTRICT")
	db.Model(&Room{}).AddForeignKey("insurance_fee_id", "insurance_fees(id)", "RESTRICT", "RESTRICT")
	db.Model(&Room{}).AddForeignKey("level_id", "levels(id)", "RESTRICT", "RESTRICT")
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
}
