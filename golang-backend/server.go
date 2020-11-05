package main

import (
	"container/list"
	"database/sql"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

type ParticipantesResult struct {
	id                         int    `json:"id"`
	memberprofileurl           string `json:"member_profile_url"`
	examurl                    string `json:"exam_url"`
	studygroupid               int    `json:"study_group_id"`
	typeofeducationalprogram   string `json:"type_of_educational_program"`
	demoexamtype               string `json:"demo_exam_type"`
	typeofcertification        string `json:"type_of_certification"`
	subjectrfooparticipant     string `json:"subject_rf_oo_participant"`
	ooparticipant              string `json:"oo_participant"`
	typeooparticipant          string `json:"type_oo_participant"`
	oomemberurl                string `json:"oo_member_url"`
	cpde                       string `json:"cpde"`
	idcpde                     int    `json:"id_cpde"`
	professionspecialtycode    int    `json:"profession_specialty_code"`
	nameofprofessionspecialist string `json:"name_of_profession_specialist"`
	onthebasis911gradesentered string `json:"on_the_basis_911_grades_entered"`
	formofstudy                string `json:"form_of_study"`
	trainigcourse              string `json:"trainig_course"`
	departicipantstatus        string `json:"de_participant_status"`
	citizenship                string `json:"citizenship"`
	dateofbirth                string `json:"date_of_birth"`
	email                      string `json:"email"`
	phone                      string `json:"phone"`
	placeofresidencesubject    string `json:"place_of_residence_subject"`
	placeofresidencelocally    string `json:"place_of_residence_locally"`
	gender                     string `json:"gender"`
	snils                      bool   `json:"snils"`
	personwithdisabilities     bool   `json:"person_with_disabilities"`
	surname                    string `json:"surname"`
	name                       string `json:"name"`
	middlename                 string `json:"middlename"`
	participantstatusintheexam string `json:"participant_status_in_the_exam"`
	competence                 string `json:"competence"`
	cod                        string `json:"cod"`
	maximumscore               int    `json:"maximum_score"`
	result                     string `json:"result"`
	dateofpreparatorydayde     string `json:"date_of_preparatory_day_de"`
	destartdate                string `json:"de_start_date"`
	deenddate                  string `json:"de_end_date"`
	nok                        string `json:"nok"`
	qualificationname          string `json:"qualification_name"`
	skilllevel                 string `json:"skill_level"`
}
type DemonstrationExam struct {
	// Country       string
	// City          sql.NullString
	// TelephoneCode int `db:"telcode"`
	gorm.Model
	id                                int    `json:"id"`
	dename                            string `json:"de_name"`
	demoexamtype                      string `json:"demo_exam_type"`
	fgosstatus                        string `json:"fgos_status"`
	programtype                       string `json:"program_type"`
	typeofcertification               string `json:"type_of_certification"`
	nocstatus                         string `json:"noc_status"`
	typeofstudent                     string `json:"type_of_student"`
	examgroupid                       int    `json:"exam_group_id"`
	rfsubjectcode                     int    `json:"rf_subject_code"`
	subjectofrf                       string `json:"subject_of_rf"`
	competence                        string `json:"competence"`
	studygroupid                      int    `json:"study_group_id"`
	studygroupnumber                  int    `json:"study_group_number"`
	professionspecialtycode           int    `json:"profession_specialty_code"`
	nameofprofessionspecialist        string `json:"name_of_profession_specialist"`
	academicyear                      int    `json:"academic_year"`
	formofstudy                       string `json:"form_of_study"`
	calendarstartyear                 int    `json:"calendar_start_year"`
	calendaryearendedlearning         int    `json:"calendar_year_ended_learning"`
	courseofstudycurrent              int    `json:"course_of_study_current"`
	courseofstudytotal                int    `json:"course_of_study_total"`
	idcpde                            int    `json:"id_cpde"`
	nameoftheoowhereitpasses          string `json:"name_of_the_oo_where_it_passes"`
	cpdeaddress                       string `json:"cpde_address"`
	numberofjobs                      int    `json:"number_of_jobs"`
	nameoftheoowhopasses              string `json:"name_of_the_oo_who_passes"`
	dateofticketcreationinthecp       string `json:"date_of_ticket_creation_in_the_cp"`
	dateofissueofthetask              string `json:"date_of_issue_of_the_task"`
	dateofs1                          string `json:"date_of_s1"`
	monthc1                           int    `json:"month_c1"`
	destartdate                       string `json:"de_start_date"`
	deenddate                         string `json:"de_end_date"`
	startdateofthedegroup             string `json:"start_date_of_the_de_group"`
	enddateofthedegroup               string `json:"end_date_of_the_de_group"`
	numberoftheshiftofthede           int    `json:"number_of_the_shift_of_the_de"`
	dateofnoctesting                  string `json:"date_of_noc_testing"`
	numberofstudentsplan              int    `json:"number_of_students_plan"`
	numberofnoctraineesplan           int    `json:"number_of_noc_trainees_plan"`
	qualification                     string `json:"qualification"`
	skilllevel                        string `json:"skill_level"`
	numberofstudentsactuial           int    `json:"number_of_students_actuial"`
	nameofthechefexpert               string `json:"name_of_the_chef_expert"`
	email                             string `json:"email"`
	expertphone                       string `json:"expert_phone"`
	region                            string `json:"region"`
	gestatus                          string `json:"ge_status"`
	dateofappointmentofthechiefexpert string `json:"date_of_appointmentof_the_chief_expert"`
	dateofrequestofthechiefexpert     string `json:"date_of_request_of_the_chief_expert"`
	gecertificatenumber               string `json:"ge_certificate_number"`
	certificateenddate                string `json:"certificate_end_date"`
	rfcertificatenumber               string `json:"rf_certificate_number"`
	enddateofcertificate              string `json:"end_date_of_certificate"`
	passportstatus                    string `json:"passport_status"`
	dateofreceiptofthecertificate     string `json:"date_of_receipt_of_the_certificate"`
	examstatus                        string `json:"exam_status"`
	actofreadiness                    string `json:"act_of_readiness"`
	postingreport                     string `json:"posting_report"`
	cpexamurl                         string `json:"cp_exam_url"`
	studygroupurl                     string `json:"study_group_url"`
	numberoflineexperts               int    `json:"number_of_line_experts"`
	innnameofoowholeases              string `json:"inn_name_of_oo_who_leases"`
	kppnameofoowhopasses              string `json:"kpp_name_of_oo_who_passes"`
	cod                               string `json:"cod"`
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
	// DemonstrationExams =[]DemonstrationExam
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
	// var paricipantes ParticipantesResult
	// var cars []Car
	// db.First(&paricipantes, params["id"])
	rows, err := db.Query("SELECT * FROM public.participantesresults")
	// rows, err := db.Query("SELECT * FROM public.participantesresults")
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

	mylist := list.New()

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
			mylist.PushBack(v)
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
	if err := json.NewEncoder(w).Encode(valuePtrs); err != nil {
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
