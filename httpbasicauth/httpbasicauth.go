package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	USERNAME = "student1"
	PASSWORD = "secret"
)

type Student struct {
	Id    string
	Name  string
	Grade int32
}

var students = []*Student{}

func init() {
	students = append(students, &Student{"1", "Budi", 1}, &Student{"2", "Ari", 2}, &Student{"3", "Adit", 3})
	fmt.Println(students)
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Fprintf(w, "something went wrong")
		return false
	}
	isValid := u == USERNAME && p == PASSWORD
	if !isValid {
		fmt.Fprintf(w, "username or paswword incorrect")
		return false
	}

	return true
}

func AllowGET(w http.ResponseWriter, r *http.Request) bool {
	switch r.Method {
	case "GET":
		return true
	default:
		fmt.Fprintf(w, "method not allowed")
		return false
	}
}

func OutputJSON(w http.ResponseWriter, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, student := range students {
		if student.Id == id {
			return student
		}
	}

	return nil
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !AllowGET(w, r) {
		return
	}
	if !Auth(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	s := new(http.Server)
	s.Addr = "localhost:80"
	log.Fatalln(s.ListenAndServe())
}
