package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Course Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
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
	return c.CourseName == ""
	//return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Simple API Running...")
	r := mux.NewRouter()

	// seeding course
	courses = append(courses, Course{
		CourseId: "1", CourseName: "GO", CoursePrice: 200,
		Author: &Author{Fullname: "James bond", Website: "dev.go"},
	})
	courses = append(courses, Course{
		CourseId: "2", CourseName: "Vue js", CoursePrice: 200,
		Author: &Author{Fullname: "Vue Author", Website: "dev.vue"},
	})
	courses = append(courses, Course{
		CourseId: "3", CourseName: "React js", CoursePrice: 200,
		Author: &Author{Fullname: "React Author", Website: "dev.react"},
	})

	// routes call
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getALlCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4040", r))
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

// add course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	w.Header().Set("Content-Type", "application/json")

	// empty body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	// empty object/json
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data found inside of JSON/request")
		return
	}
	// id generate
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	// Added slices
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

// update course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

// delete course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course delete successfully")
			break
		}
	}
}
