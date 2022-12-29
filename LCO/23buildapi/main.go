package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("Golang API")

	router := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{
		Fullname: "Sukanta Nandi", Website: "LCO.in",
	}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Python", CoursePrice: 399, Author: &Author{
		Fullname: "Sukanta Nandi", Website: "python.in",
	}})

	// routing
	router.HandleFunc("/", servHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port and server
	log.Fatal(http.ListenAndServe(":3000", router))

}

// controllers - file

// server home route

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang API HomePage</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) // send json
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given Id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is missing
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// if body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("CourseName is empty")
		return
	}

	// check if title is duplicate

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn((100)))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// get course id from request
	params := mux.Vars(r)

	// loop , check id, remove from db, add with new id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // ... denotes unpack the slice and append
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	// Send a response when id is not found
	json.NewEncoder(w).Encode("No Course found with given Id")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// get course id from request
	params := mux.Vars(r)

	// loop , check id, remove from db, add with new id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // ... denotes unpack the slice and append
			json.NewEncoder(w).Encode("Deleted the requested course")
			return
		}

	}
	// Send a response when id is not found
	json.NewEncoder(w).Encode("No Course found with given Id")
	return

}
