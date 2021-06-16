package main

import (
	"gorm-basics/user"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", user.AllUsers).Methods("GET")
	//router.HandleFunc("/user/{name}", user.DeleteUser).Methods("DELETE")
	//router.HandleFunc("/user/{name}/{email}", user.UpdateUser).Methods("PUT")
	//router.HandleFunc("/user/{name}/{email}", user.NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))

}

func initialMigration() {
	_, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	//defer db.Close()
	//db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Hello, World!")
	initialMigration()
	handleRequests()
}
