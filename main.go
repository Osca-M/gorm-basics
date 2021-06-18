package main

import (
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"gorm-basics/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)


func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", user.AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}", user.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{name}/{email}", user.NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))

}

func initialMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user.User{})
}

func main() {
	fmt.Println("Hello, World!")
	initialMigration()
	handleRequests()
}
