package main

import (
	"fmt"

	"github.com/google/uuid"
)

// type ConfigList struct {
// 	Port      int
// 	DbName    string
// 	SQLDriver string
// }

// var Config ConfigList

// func init() {
// 	cfg, _ := ini.Load("config.ini")

// 	Config = ConfigList{
// 		Port: cfg.Section("web").Key("port").MustInt(8080),

// 		DbName: cfg.Section("db").Key("name").MustString("example.sql"),

// 		SQLDriver: cfg.Section("db").Key("driver").String(),
// 	}
// }

func main() {

	uuidObj, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println(" ", uuidObj2.String())

	// fmt.Printf("Port = %v\n", Config.Port)
	// fmt.Printf("DbName = %v\n", Config.DbName)
	// fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)

}
