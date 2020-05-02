package database

import (
	"fmt"
	"github.com/benammann/drinkspot-core/api/app/model"
	"github.com/jinzhu/gorm"
	"os"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Database struct {
	Connection *gorm.DB
}

var Current *Database

func NewDatabase() *Database {

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	connection, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, password, host, "3306", database))

	if err != nil {
		panic(err)
	}

	dbInstance := &Database{
		Connection: connection,
	}

	// initialises the database instance
	// migrations etc
	dbInstance.initialize()

	Current = dbInstance

	return dbInstance

}

func (db *Database) initialize() {
	db.migrateModels()
}

// auto migrates all the model
func (db *Database) migrateModels() {
	db.Connection.AutoMigrate(&model.User{})
}
