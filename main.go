package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type DBvars struct {
	user string
	pass string
	name string
	port string
	host string
}

type User struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

func GetDB() *sql.DB {
	var dbvars DBvars
	dbvars.user = viper.GetString("db.user")
	dbvars.pass = viper.GetString("db.pass")
	dbvars.name = viper.GetString("db.name")
	dbvars.port = viper.GetString("db.port")
	dbvars.host = viper.GetString("db.host")

	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbvars.user, dbvars.pass, dbvars.host, dbvars.port, dbvars.name)
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
	w.Write([]byte("OK"))

	defer db.Close()
	fmt.Printf("Name: %v, Login: %v, Pass: %v \n", user.Name, user.Mail, user.Pass)
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("init config error!")
	}

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	router.HandleFunc("/post", PostDB).Methods("POST")
	router.PathPrefix("/").Handler(fs)
	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":"+viper.GetString("port"), handler); err != nil {
		log.Fatal("Server Error!")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configGO")
	return viper.ReadInConfig()
}
