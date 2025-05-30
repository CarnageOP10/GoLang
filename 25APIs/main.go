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

// model for courses
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

// main
func main() {
	fmt.Println("welcome to APIs")

	r := mux.NewRouter()

	//seeding part
	courses = append(courses, Course{
		CourseID:    "1",
		CourseName:  "Go Basics",
		CoursePrice: 299,
		Author: &Author{
			FullName: "John Doe",
			Website:  "https://johndoe.dev",
		},
	})

	courses = append(courses, Course{
		CourseID:    "2",
		CourseName:  "Advanced Go",
		CoursePrice: 499,
		Author: &Author{
			FullName: "Jane Smith",
			Website:  "https://janesmith.dev",
		},
	})

	courses = append(courses, Course{
		CourseID:    "3",
		CourseName:  "more Advanced Go",
		CoursePrice: 599,
		Author: &Author{
			FullName: "jade Smith",
			Website:  "https://jadesmith.dev",
		},
	})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/corses", getAllCourses).Methods("GET")
	r.HandleFunc("/onecourse/{id}", getOneCourses).Methods("GET") //params["id"]
	r.HandleFunc("/createone", createOneCourse).Methods("POST")
	r.HandleFunc("/updateone/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/removeone/{id}", removeOneCourse).Methods("DELETE")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello</h1>"))
}

// get the courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-type", "application.json")
	json.NewEncoder(w).Encode(courses)
}

// get one course
func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get 1 courses")
	w.Header().Set("Content-type", "application.json")

	//grab id from request
	params := mux.Vars(r)

	//loop through the course to find matching id
	for _, course := range courses {
		if params["id"] == course.CourseID {
			json.NewEncoder(w).Encode(course)
		}
	}

	json.NewEncoder(w).Encode("no course found for this id")
	return
}

// create 1 course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create 1 courses")
	w.Header().Set("Content-type", "application.json")

	//what if entire body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("pls add something")
	}

	//what if {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside teh json")
		return
	}

	//now push
	// generate a new id and add the course

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

// update a course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update 1 courses")
	w.Header().Set("Content-type", "application.json")

	params := mux.Vars(r)

	//loop, id, remove, add with my id
	for index, course := range courses {
		if params["id"] == course.CourseID {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

// delete a course
func removeOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("remove 1 courses")
	w.Header().Set("Content-type", "application.json")

	params := mux.Vars(r)

	//loop, id, remove
	for index, course := range courses {
		if params["id"] == course.CourseID {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("removed the course" + params["id"])
}
