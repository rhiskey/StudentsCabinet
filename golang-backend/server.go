package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

// Config - data structure to work with json
type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
	cfg      map[string]interface{}
)

// pg connection
// const (
// 	host     = "rhiskey.ddns.net"
// 	port     = 2059
// 	user     = "postgres"
// 	password = "6794"
// 	dbname   = "postgres"
// )

// Config
func init() {
	jsonByte, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of config file:", string(jsonByte))

	if json.Unmarshal(jsonByte, &cfg); err != nil {
		panic(err)
	}
	host = cfg["host"].(string)
	port = int(cfg["port"].(float64))
	user = cfg["user"].(string)
	password = cfg["password"].(string)
	dbname = cfg["dbname"].(string)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
}
