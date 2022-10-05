package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
)

//dbConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

const (
	DB_USER     = "user"
	DB_PASSWORD = "pass"
	DB_NAME     = "testdb_task"
	DB_PORT     = "5432"
	DB_HOST     = "postgres"
)

type User struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

func GetDB() *sql.DB {
	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
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

/*
func GetForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
*/

func PostDB(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
	fs := http.FileServer(http.Dir("./static"))
	//router.HandleFunc("/get", GetForm).Methods("GET")
	router.HandleFunc("/post", PostDB).Methods("POST", "OPTIONS")
	router.PathPrefix("/").Handler(fs)
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("Server Error!")
	}
}
