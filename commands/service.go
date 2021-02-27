package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"

	"strings"
	"time"
	//"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type patient struct {
	id         int
	first_name string
	last_name  string
	email      string
	gender     string
	password   string
	birth_date string
	message    string
}

type Jpatient struct {
	Id          int
	Login       string
	Password    string
	Name        string
	Surname     string
	Timetableld string
}

type Order struct {
}

func bd(id int) {

}

// Обработчик главной странице.
func home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	var e Jpatient
	err1 := json.NewDecoder(r.Body).Decode(&e)
	if err1 != nil {
		panic(err1)
	}
	fmt.Fprintf(w, "Person: %+v", e)
	fmt.Println(e)

	rows, err := db.Query("select first_name from patient where id = " + fmt.Sprint(id))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	patients := []patient{}
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.first_name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
	}
	for _, p := range patients {
		fmt.Fprintf(w, string(p.first_name))
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	var e Jpatient
	err1 := json.NewDecoder(r.Body).Decode(&e)
	if err1 != nil {
		panic(err1)
	}

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into patient (first_name, last_name, email, password) values ($1, $2, $3, $4)",
		e.Name, e.Surname, e.Login, e.Password)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func adddoctor(w http.ResponseWriter, r *http.Request) {
	var e Jpatient
	err1 := json.NewDecoder(r.Body).Decode(&e)
	if err1 != nil {
		panic(err1)
	}

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=127.0.0.1"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into doctor (first_name, last_name) values ($1, $2)",
		e.Name, e.Surname)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	rows, err := db.Query("select id from doctor ORDER BY id DESC LIMIT 1")
	// fmt.Println(rows[0])
	// rows.Scan(&e.Id)
	patients := []patient{}
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
	}

	e.Id = patients[0].id
	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}

func changedoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println(url.Parse(fmt.Sprint(r.URL)))
	str := strings.Split(fmt.Sprint(r.URL), "B")
	str1 := strings.Split(str[1], "%")
	id, err := strconv.Atoi(str1[0])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	var e Jpatient
	err1 := json.NewDecoder(r.Body).Decode(&e)
	if err1 != nil {
		panic(err1)
	}

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("UPDATE doctor SET first_name = $1 , last_name = $2 where id = $3",
		e.Name, e.Surname, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func deletedoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println(url.Parse(fmt.Sprint(r.URL)))
	str := strings.Split(fmt.Sprint(r.URL), "B")
	str1 := strings.Split(str[1], "%")
	id, err := strconv.Atoi(str1[0])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM doctor WHERE id = $1",
		id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func getAllDoctors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("popa")
	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("popa")
	rows, err := db.Query("select id,first_name, last_name from doctor")
	// fmt.Println(rows[0])
	// rows.Scan(&e.Id)
	patients := []patient{}
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.id, &p.first_name, &p.last_name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
		fmt.Fprintf(w, fmt.Sprint(p.id))
		fmt.Fprintf(w, " "+p.first_name)
		fmt.Fprintf(w, " "+p.last_name)
		fmt.Fprintf(w, "\n")

	}

}

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("popa")
	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select doctor.first_name,doctor.last_name,date_time  from registration LEFT JOIN doctor ON registration.id_doctor = doctor.id where id_patient is NULL")
	// fmt.Println(rows[0])
	// rows.Scan(&e.Id)
	patients := []patient{}
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.first_name, &p.last_name, &p.birth_date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
		fmt.Fprintf(w, p.first_name)
		fmt.Fprintf(w, " "+p.last_name)
		fmt.Fprintf(w, " "+p.birth_date)
		fmt.Fprintf(w, "\n")

	}

}

func setme(w http.ResponseWriter, r *http.Request) {

	var e Jpatient
	err1 := json.NewDecoder(r.Body).Decode(&e)
	if err1 != nil {
		panic(err1)
	}

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id from patient where email = $1 and password = $2", e.Login, e.Password)
	if err != nil {
		fmt.Println("jopa")
		panic(err)
	}

	patients := []patient{}
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
	}

	e.Id = patients[0].id

	result, err := db.Exec("UPDATE registration SET id_patient = $1 where id = (SELECT id from registration where date_time = $2 ORDER BY id LIMIT 1)",
		e.Id, e.Timetableld)
	if err != nil {
		fmt.Println(result)
		panic(err)
	}
	fmt.Fprintf(w, "login: "+e.Login)
	fmt.Fprintf(w, " TimeTableld: "+e.Timetableld)
	fmt.Fprintf(w, " TimeNow: "+fmt.Sprint(time.Now()))

	// fmt.Println(rows[0])
	// rows.Scan(&e.Id)

}

func getAllLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(url.Parse(fmt.Sprint(r.URL)))
	str := strings.Split(fmt.Sprint(r.URL), "B")
	str1 := strings.Split(str[1], "%")[0]

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT doctor.first_name,doctor.last_name, date_time FROM registration  LEFT JOIN doctor ON doctor.id = registration.id_doctor where id_patient = (SELECT id from patient where email = $1)", str1)
	if err != nil {
		fmt.Println(rows)
		panic(err)
	}
	patients := []patient{}
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.first_name, &p.last_name, &p.birth_date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
		fmt.Fprintf(w, p.first_name)
		fmt.Fprintf(w, " "+p.last_name)
		fmt.Fprintf(w, " "+p.birth_date)
		fmt.Fprintf(w, "\n")

	}

}

var tpl = template.Must(template.ParseFiles("index.html"))

var auth1 = template.Must(template.ParseFiles("message.html"))

var redirect = template.Must(template.ParseFiles("aut.html"))

type ViewData struct {
	First_name string
	Last_name  string
	Message    []string
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("jopa")
	cookie, err := r.Cookie("cookie1")
	fmt.Println("cookie:", cookie, "err", err)
	send1 := strings.Split(fmt.Sprint(cookie), "=")
	if len(send1) == 1 {
		tpl.Execute(w, nil)
	} else {

		connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT first_name, last_name from patient where cookie = $1", send1[1])
		if err != nil {
			fmt.Println(rows)
			panic(err)
		}

		patients := []patient{}
		k := 0
		var name string
		var surname string
		for rows.Next() {
			p := patient{}
			err := rows.Scan(&p.first_name, &p.last_name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			patients = append(patients, p)
			k = 1
			fmt.Print(p.first_name)
			fmt.Print("dasdasdasdasdweq")
			name = p.first_name
			surname = p.last_name
		}
		if k == 1 {

			rows, err := db.Query("SELECT patient.first_name, patient.last_name, message from allmessages LEFT JOIN patient ON allmessages.id_sender = patient.id ORDER BY allmessages.id")
			if err != nil {
				fmt.Println(rows)
				panic(err)
			}
			fmt.Println(rows)
			patients := []patient{}
			var str []string
			for rows.Next() {
				p := patient{}
				err := rows.Scan(&p.first_name, &p.last_name, &p.message)
				if err != nil {
					fmt.Println(err)
					continue
				}
				patients = append(patients, p)
				k = 1

				str = append(str, p.first_name+" "+p.last_name+": "+p.message)
			}
			fmt.Print(str)
			data := ViewData{
				First_name: name,
				Last_name:  surname,
				Message:    str,
			}

			auth1.Execute(w, data)
		} else {
		}
	}
}

func auth(w http.ResponseWriter, r *http.Request) {

	fmt.Print("jopa")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	str := strings.Split(string(body), "=")
	str1 := strings.Split(str[1], "&")
	fmt.Println(str[2])

	log1 := strings.Split(str1[0], "%40")
	if len(log1) == 1 {
		http.NotFound(w, r)
		return
	}
	log2 := log1[0] + "@" + log1[1]
	fmt.Println(log2)
	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,first_name, last_name from patient where email = $1 and password = $2", log2, str[2])
	if err != nil {
		fmt.Println(rows)
		panic(err)
	}

	fmt.Println(rows)
	patients := []patient{}
	k := 0
	var id string
	//var name string
	//var surname string
	for rows.Next() {
		p := patient{}
		err := rows.Scan(&p.email, &p.first_name, &p.last_name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		patients = append(patients, p)
		k = 1
		id = p.email
		//name = p.first_name
		//surname = p .last_name
	}

	if k == 0 {
		http.NotFound(w, r)
		return
	} else {
		cookie, err := r.Cookie("cookie1")
		fmt.Println("cookie:", cookie, "err", err)
		if err != nil {
			fmt.Println("no cookie")
			cookie = &http.Cookie{
				Name:  "cookie1",
				Value: str[2],
				//HttpOnly:	true,
				//MaxAge:		100,
			}
			fmt.Println("popa1")
			http.SetCookie(w, cookie)
			fmt.Println("popa2")
			result, err := db.Exec("UPDATE patient SET cookie = $1 where id = $2", str[2], id)
			if err != nil {
				panic(err)
				fmt.Println(result)
			}

		}

		//data := ViewData{
		//	First_name: name,
		//	Last_name: surname,
		//}

		//http.Redirect(w,r,"localhost:4000/login", 301)
		Login(w, r)
		//http.RedirectHandler("localhost:4000/login", 301)

		//redirect.Execute(w, data)
	}

}

func doctor(w http.ResponseWriter, r *http.Request) {
	met := r.Method
	fmt.Print(met)
	if met == "POST" {
		adddoctor(w, r)
	}
	if met == "GET" {
		getAllDoctors(w, r)
	}
}

func doctorId(w http.ResponseWriter, r *http.Request) {
	met := r.Method
	fmt.Print(met)
	if met == "PUT" {
		changedoctor(w, r)
	}
	if met == "DELETE" {
		deletedoctor(w, r)
	}
}

var mux = http.NewServeMux()

func patientReq(w http.ResponseWriter, r *http.Request) {
	met := r.Method
	fmt.Print(met)
	if met == "POST" {
		setme(w, r)
	}
	if met == "GET" {
		getAll(w, r)
	}
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	fmt.Print("jopa")
	mux.HandleFunc("/", home)
	mux.HandleFunc("/getall/", getAllLogin)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/doctor", doctor)
	mux.HandleFunc("/doctor/", doctorId)
	mux.HandleFunc("/patient", patientReq)
	mux.HandleFunc("/patient/register", register)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	//dsadsad
}
