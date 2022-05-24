package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Course Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// helpers - file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers - file

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api</h1>"))
}

// get all courses
func getALlCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get all courses
func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(courses)

	// grab if from request
	params := mux.Vars(r)

	// loop through courses and find the course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No record found")
	return
}
