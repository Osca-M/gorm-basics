package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type user struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Endpoint Hit")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))

}

func initialMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	//defer db.Close()
	db.AutoMigrate(&user{})
}

func main() {
	fmt.Println("Hello, World!")
	initialMigration()
	handleRequests()
}
