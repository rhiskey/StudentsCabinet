package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func table(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This will be table data")
	checkError(err)
}

func graph1(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This will be grap1 data")
	checkError(err)
}

func graph2(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This will be grap2 data")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/", table)
	http.HandleFunc("/", graph1)
	http.HandleFunc("/", graph2)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
	//TODO: Connect to pg
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
}
