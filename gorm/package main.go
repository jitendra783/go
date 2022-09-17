package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
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

/*
func allUsers(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    var users []User
    db.Find(&users)
    fmt.Println("{}", users)

    json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("New User Endpoint Hit")

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]
    email := vars["email"]

    db.Create(&User{Name: name, Email: email})
    fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]

    var user User
    db.Where("name = ?", name).Find(&user)
    db.Delete(&user)

    fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]
    email := vars["email"]

    var user User
    db.Where("name = ?", name).Find(&user)

    user.Email = email

    db.Save(&user)
    fmt.Fprintf(w, "Successfully Updated User")
}


func initialMigration() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&User{})
}
*/

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	// Add the call to our new initialMigration function
	// initialMigration()

	handleRequests()
}
