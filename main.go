package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/speps/go-hashids"
)

//Myurl represents the model for url shortner System
type Myurl struct {
	ID       string `json:"Id"`
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
}

var db *gorm.DB

//Database Connection start
func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connet database")
	}
	db.Exec("CREATE DATABASE Myurl")
	db.Exec("USE Myurl")
	db.AutoMigrate(&Myurl{})
}

//Database Connection end

// createurl function start
func createurl(w http.ResponseWriter, r *http.Request) {
	var url Myurl
	json.NewDecoder(r.Body).Decode(&url)
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	url.ID, _ = h.Encode([]int{int(now.Unix())})
	url.ShortURL = "http://localhost:8080/" + url.ID
	db.Create(&url)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

// createurl function end

// geturls function starts
func geturls(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var urls []Myurl
	db.Find(&urls)
	json.NewEncoder(w).Encode(urls)
}

// geturls function starts

// redirect function starts
func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var url Myurl
	db.Where("id = ?", params["Id"]).First(&url)
	http.Redirect(w, r, url.LongURL, 301)
	json.NewEncoder(w).Encode(url)
}

// redirect function end


// deleteurl function starts
func deleteurl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["Id"]
	var url Myurl
	db.Where("id = ? ", inputOrderID).Delete(&url)
}

// deleteurl function end

func main() {
	//all Routers
	router := mux.NewRouter()
	router.HandleFunc("/create/url", createurl).Methods("PUT")
	router.HandleFunc("/urls", geturls).Methods("GET")
	router.HandleFunc("/{Id}", redirect).Methods("GET")
	router.HandleFunc("/url/{Id}", deleteurl).Methods("DELETE")
	initDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}
