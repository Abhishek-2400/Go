package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Course struct { // if we want to use this struct in other packages, we need to export it by capitalizing the first letter
	CourseId    string  `json:"courseid"` // json tag to specify the name in JSON format
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // nested struct, we can use it like this
}
type Author struct {
	Fullname string `json:"fullname"` // json tag to specify the name in JSON format
	Website  string `json:"website"`
}

var courses []Course // act like db

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page Hit")
	w.Write([]byte("<h1>Welcome to the Home Page!<h1>")) // Write response to the client in the form of bytes
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting up the response header to application/json
	json.NewEncoder(w).Encode(courses)                 // encoding the courses slice to JSON and writing it to the response writer
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// we can get id from url also but we  will use mux as it offers us that feature as well (a alternative)
	params := mux.Vars(r) // this will give us a map of parameters from the URL
	for _, course := range courses {
		if course.CourseId == params["id"] {
			w.WriteHeader(http.StatusOK)      // setting the status code to 200 OK
			json.NewEncoder(w).Encode(course) // encoding the course to JSON and writing it to the response writer
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)                                                                      // setting the status code to 404 Not Found
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Course not found %s", params["id"])}) // encoding the error message to JSON and writing it to the response writer
	//For formatting a string without printing, use fmt.Sprintf
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decoding the request body to the course struct
	if course.isEmpty() {
		w.WriteHeader(http.StatusBadRequest)                                                   // setting the status code to 400 Bad Request
		json.NewEncoder(w).Encode(map[string]string{"error": "Please provide a valid course"}) // encoding the error message to JSON and writing it to the response writer
		return
	}

	course.CourseId = uuid.New().String()
	courses = append(courses, course) // appending the course to the courses slice
	w.WriteHeader(http.StatusCreated) // setting the status code to 201 Created
	json.NewEncoder(w).Encode(course) // encoding the course to JSON and writing it to the response writer
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // this will give us a map of parameters from the URL
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // removing the course from the slice
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse) // decoding the request body to the updatedCourse struct
			updatedCourse.CourseId = course.CourseId           // setting the CourseId to the old CourseId for precaution
			courses = append(courses, updatedCourse)           // appending the updated course to the courses slice
			w.WriteHeader(http.StatusOK)                       // setting the status code to 200 OK
			json.NewEncoder(w).Encode(updatedCourse)           // encoding the updated course to JSON and writing it to the response writer
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)                                                                      // setting the status code to 404 Not Found
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Course not found %s", params["id"])}) // encoding the error message to JSON and writing it to the response writer
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // this will give us a map of parameters from the URL
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)                                // removing the course from the slice
			w.WriteHeader(http.StatusOK)                                                           // setting the status code to 200 OK
			json.NewEncoder(w).Encode(map[string]string{"message": "Course deleted successfully"}) // encoding the success message to JSON and writing it to the response writer
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) // setting the status code to 404 Not Found
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Course not found %s", params["id"])})
}

func main() {
	r := mux.NewRouter() // this is gorilla mux router

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Go", CoursePrice: 299, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 199, Author: &Author{Fullname: "Jane Doe", Website: "janedoe.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":8080", r))

}
