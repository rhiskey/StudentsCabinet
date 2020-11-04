package main

import (
	"database/sql"
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
	//decoder := json.NewDecoder(r.Body)

	//var numsData numbers
	//var numsResData numsResponseData

	//decoder.Decode(&numsData)

	//----------------DB----------
	// Connection String to PG
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//fmt.Println(psqlInfo)

	// Opening a connection to PG
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//var p

	//var p []byte
	// rows, err := db.Query("SELECT * FROM demo_exams LIMIT $1", 1000)

	rows, err := db.Query("SELECT * FROM participant_results ")
	fmt.Println(rows)

	// var result []interface{}

	// cols, _ := rows.Columns()
	// pointers := make([]interface{}, len(cols))
	// container := make([]json.RawMessage, len(cols))
	// for i, _ := range pointers {
	// 	pointers[i] = &container[i]
	// }
	// for rows.Next() {
	// 	rows.Scan(pointers...)
	// 	result = append(result, container)
	// }

	// fmt.Println(container)

	//c.JSON(200, container)

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {

		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i, col := range columns {

			var v interface{}

			val := values[i]

			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}

			// TODO: Create Json to pass data in FRONT
			fmt.Println(col, v)
		}
	}

	//----------------DB----------
	//numsResData = process(numsData) // pass data from Database

	//fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rows); err != nil {
		panic(err)
	}
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
	http.HandleFunc("/table", table)
	// http.HandleFunc("/", graph1)
	// http.HandleFunc("/", graph2)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
