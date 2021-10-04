package server

import (
	"basic_archetech/middlewares/rdbms"
	"basic_archetech/utils"
	"log"

	"gorm.io/gorm"
)

func makeRDBMSConnection() *gorm.DB {
	dbc := &rdbms.ConnectDB_Configure{
		Host:     "localhost",
		Port:     "3306",
		User:     "test",
		Password: "test",
		Database: "test",
	}

	db, err := dbc.ConnectDB()

	if err != nil {
		log.Panicf("Can not create RDBMS connection. Please check information :\n%s", utils.PrettyPrint(dbc))
		return nil
	}

	return db
}
