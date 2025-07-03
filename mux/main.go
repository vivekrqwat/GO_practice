package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// file
type Course struct {
	Coursename  string  `json:"fullname"`
	CourseId    string  `json:"courseid"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var course []Course

// middleware

func (c *Course) isempty() bool {
	return c.Coursename == ""
}

func main() {
	course = append(course,
		Course{
			Coursename:  "Go Programming",
			CourseId:    "101",
			CoursePrice: 500,
			Author:      &Author{Fullname: "Alice Smith", Website: "alice.dev"},
		},
		Course{
			Coursename:  "Web Development",
			CourseId:    "102",
			CoursePrice: 700,
			Author:      &Author{Fullname: "Bob Johnson", Website: "bobweb.com"},
		},
		Course{
			Coursename:  "Data Structures",
			CourseId:    "103",
			CoursePrice: 600,
			Author:      &Author{Fullname: "Carol Lee", Website: "carol.codes"},
		},
	)
	router := mux.NewRouter()
	router.HandleFunc("/course", getall).Methods("GET")
	router.HandleFunc("/course/{id}", getone).Methods("GET")
	router.HandleFunc("/course", createone).Methods("POST")
	router.HandleFunc("/course/{id}", update).Methods("POST")
	router.HandleFunc("/course/{id}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func set(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello</h2>"))

}

func sethome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello worlf</h1>"))
}

func getall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)

}
func getone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prams := mux.Vars(r)
	for _, item := range course {
		if item.CourseId == prams["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("NO COURSE FOUND OF GIVVEN ID")

}

func createone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("EMPTY_DATA")
		return
	}

	var newd Course
	_ = json.NewDecoder(r.Body).Decode(&newd)
	if newd.isempty() {
		json.NewEncoder(w).Encode("no courseid")
		return
	}
	//generate id
	rand.Seed(time.Now().UnixNano())
	newd.CourseId = strconv.Itoa(rand.Intn(100))
	course = append(course, newd)
	json.NewEncoder(w).Encode(course)

}
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range course {

		if item.CourseId == param["id"] {
			course = append(course[:index], course[:index+1]...)

			var newd Course
			json.NewDecoder(r.Body).Decode(&newd)
			course = append(course, newd)
			json.NewEncoder(w).Encode(course)

		}
	}

}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range course {
		if item.CourseId == param["id"] {
			course = append(course[:index], course[:index+1]...)
			json.NewEncoder(w).Encode(course)
			return

		}

	}
	json.NewEncoder(w).Encode("NOT_FOUND")

}
