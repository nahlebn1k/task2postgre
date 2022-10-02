package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	DB_USER     = "admin"
	DB_PASSWORD = "pass"
	DB_NAME     = "testdb_task"
)

type User struct {
	Name string
	Mail string
	Pass string
}

func GetDB() *sql.DB {
	dbConnStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func PostDB(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	var user User
	user.Name = r.FormValue("username")
	user.Mail = r.FormValue("user-email")
	user.Pass = r.FormValue("user-pass")

	_, err := db.Exec("INSERT INTO users(username, login, pass) VALUES ($1,$2,$3)", user.Name, user.Mail, user.Pass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
	fmt.Printf("Name: %v, Login: %v, Pass: %v \n", user.Name, user.Mail, user.Pass)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", GetForm).Methods("GET")
	router.HandleFunc("/post", PostDB).Methods("POST")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("Server Error!")
	}
}
