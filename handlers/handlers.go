package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/gorm"

    "example/hello/database"
    "example/hello/model"
)

var db *gorm.DB

func init() {
    // Initialize the database connection
    var err error
    db, err = database.ConnectDB()
    if err != nil {
        panic(err)
    }
    db.AutoMigrate(&model.Name{})
}

func GetNames(w http.ResponseWriter, r *http.Request) {
    var names []model.Name
    db.Find(&names)
    json.NewEncoder(w).Encode(names)
}

func GetName(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var name model.Name
    db.First(&name, params["id"])
    json.NewEncoder(w).Encode(name)
}

func CreateName(w http.ResponseWriter, r *http.Request) {
    var name model.Name
    json.NewDecoder(r.Body).Decode(&name)
    db.Create(&name)
    json.NewEncoder(w).Encode(name)
}

func UpdateName(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var name model.Name
    db.First(&name, params["id"])
    json.NewDecoder(r.Body).Decode(&name)
    db.Save(&name)
    json.NewEncoder(w).Encode(name)
}

func DeleteName(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var name model.Name
    db.Delete(&name, params["id"])
    json.NewEncoder(w).Encode(name)
}