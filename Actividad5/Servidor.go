package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Message struct {
	Type string
	Text string
}

type StudentDataRecive struct {
	Name    string
	Subject string
	Score   float64
}

type Subject struct {
	Name  string
	Score float64
}

type StudentData struct {
	Name     string
	Subjects []Subject
}

type StudentSubjectScore struct {
	Name  string
	Score float64
}

type SubjectScore struct {
	Name     string
	Students []StudentSubjectScore
}

var studentsData []StudentData
var subjects []SubjectScore

func loadHtml(fileName string) string {
	html, _ := ioutil.ReadFile(fileName)
	return string(html)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHtml("index.html"),
	)
}

func formStudent(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHtml("formStudent.html"),
	)
}

func formGetStudentScore(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHtml("studentScore.html"),
	)
}

func formGetSubjectScore(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		loadHtml("subjectScore.html"),
	)
}

func formGetGeneralSubject(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	messageResponse := getGeneralAverageByStudents()
	showResponseRequest(res, messageResponse)
}

func responseStatusView(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.FormValue("view"))
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	switch req.Method {
	case "POST":
		switch req.FormValue("view") {
		case "add-student":
			score, _ := strconv.ParseFloat(req.FormValue("score"), 32)
			data := StudentDataRecive{
				Name:    req.FormValue("name"),
				Subject: req.FormValue("subject"),
				Score:   score,
			}
			messageResponse := addStudentData(data)
			showResponseRequest(res, messageResponse)
		case "get-student-score":
			messageResponse := getStudentAverage(req.FormValue("name"))
			showResponseRequest(res, messageResponse)
		case "get-subject-score":
			messageResponse := getAverageBySubject(req.FormValue("name"))
			showResponseRequest(res, messageResponse)
		}
	}
}

func showResponseRequest(res http.ResponseWriter, message Message) {
	switch message.Type {
	case "success":
		fmt.Fprintf(
			res,
			loadHtml("successPage.html"),
			message.Text,
		)
	case "error":
		fmt.Fprintf(
			res,
			loadHtml("errorPage.html"),
			message.Text,
		)
	default:
		fmt.Fprintf(
			res,
			loadHtml("errorPage.html"),
			message.Text,
		)
	}
}

func addStudentData(data StudentDataRecive) Message {
	subject := Subject{
		Name:  data.Subject,
		Score: data.Score,
	}
	subjectArray := []Subject{
		subject,
	}

	student := StudentData{
		Name:     data.Name,
		Subjects: subjectArray,
	}

	if studentExist(data.Name) {
		if !studentSubjectExist(student) {
			addNewSubjectToStudent(student)

			if subjectExist(student) {
				addNewStudentToSubject(student)
			} else {
				addNewSubjectToSubjects(student)
			}
			fmt.Println(studentsData)
			fmt.Println(subjects)
			message := Message{"success", "Operacion exitosa"}
			return message
		} else {
			fmt.Println(studentsData)
			fmt.Println(subjects)
			message := Message{"error", "El alumno " + data.Name + " ya tiene registrado la materia de " + data.Subject}
			return message
		}
	} else {

		studentsData = append(studentsData, student)
		if subjectExist(student) {
			addNewStudentToSubject(student)
		} else {
			addNewSubjectToSubjects(student)
		}
		fmt.Println(studentsData)
		fmt.Println(subjects)
		message := Message{"success", "Operacion exitosa"}
		return message
	}
	fmt.Println(studentsData)
	fmt.Println(subjects)

	message := Message{"default", "Error"}
	return message
}

func getStudentAverage(name string) Message {
	if len(studentsData) > 0 {
		if studentExist(name) {
			var averageStudent float64
			var totalSubjectsStudent float64

			for i := 0; i < len(studentsData); i++ {
				if studentsData[i].Name == name {
					for j := 0; j < len(studentsData[i].Subjects); j++ {
						averageStudent += studentsData[i].Subjects[j].Score
						totalSubjectsStudent++
					}
					break
				}
			}

			averageStudent = averageStudent / totalSubjectsStudent
			message := Message{"success", "El promedio del estudiante " + name + " es:" + strconv.FormatFloat(averageStudent, 'f', -1, 64)}
			return message
		} else {
			message := Message{"error", "El estudiante ingresado no existe"}
			return message
		}
	} else {
		message := Message{"error", "No existen estudiantes registrados"}
		return message
	}
}

func getAverageBySubject(subjectName string) Message {
	var totalAvarage float64
	var totalRecords float64

	if len(subjects) > 0 {
		if subjectExistInSubjectRecords(subjectName) {
			for i := 0; i < len(subjects); i++ {
				if subjects[i].Name == subjectName {
					for j := 0; j < len(subjects[i].Students); j++ {
						totalAvarage += subjects[i].Students[j].Score
						totalRecords++
					}
					break
				}
			}
			totalAvarage = totalAvarage / totalRecords
			message := Message{"success", "Promedio de la materia " + subjectName + ": " + strconv.FormatFloat(totalAvarage, 'f', -1, 64)}
			return message
		} else {
			message := Message{"error", "La materia " + subjectName + " no existe"}
			return message
		}
	} else {
		message := Message{"error", "No existen materias registradas"}
		return message
	}
}

func getGeneralAverageByStudents() Message {
	var totalSubjects float64
	var totalAverage float64

	if len(studentsData) > 0 {
		for i := 0; i < len(studentsData); i++ {
			for j := 0; j < len(studentsData[i].Subjects); j++ {
				totalAverage += studentsData[i].Subjects[j].Score
				totalSubjects++
			}
		}
		totalAverage = totalAverage / totalSubjects
		message := Message{"success", "Promedio general:" + strconv.FormatFloat(totalAverage, 'f', -1, 64)}
		return message
	} else {
		message := Message{"error", "No existen estudiantes"}
		return message
	}
}

func subjectExistInSubjectRecords(name string) bool {
	for i := 0; i < len(subjects); i++ {
		if subjects[i].Name == name {
			return true
		}
	}
	return false
}

func studentExist(studentName string) bool {
	for i := 0; i < len(studentsData); i++ {
		if studentsData[i].Name == studentName {
			return true
		}
	}
	return false
}

func studentSubjectExist(student StudentData) bool {
	for i := 0; i < len(studentsData); i++ {
		if studentsData[i].Name == student.Name {
			for j := 0; j < len(studentsData[i].Subjects); j++ {
				if studentsData[i].Subjects[j].Name == student.Subjects[0].Name {
					return true
				}
			}
		}
	}
	return false
}

func addNewSubjectToStudent(studentData StudentData) {
	for i := 0; i < len(studentsData); i++ {
		if studentsData[i].Name == studentData.Name {
			studentsData[i].Subjects = append(studentsData[i].Subjects, studentData.Subjects[0])
		}
	}
}

func subjectExist(student StudentData) bool {
	for i := 0; i < len(subjects); i++ {
		if subjects[i].Name == student.Subjects[0].Name {
			return true
		}
	}
	return false
}

func addNewStudentToSubject(student StudentData) {
	for i := 0; i < len(subjects); i++ {
		if subjects[i].Name == student.Subjects[0].Name {
			newRegister := StudentSubjectScore{student.Name, student.Subjects[0].Score}
			subjects[i].Students = append(subjects[i].Students, newRegister)
		}
	}
}

func addNewSubjectToSubjects(student StudentData) {
	subjectStudent := StudentSubjectScore{
		Name:  student.Name,
		Score: student.Subjects[0].Score,
	}
	studentArray := []StudentSubjectScore{
		subjectStudent,
	}
	newSubject := SubjectScore{
		student.Subjects[0].Name,
		studentArray,
	}
	subjects = append(subjects, newSubject)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/form-student", formStudent)
	http.HandleFunc("/response-status", responseStatusView)
	http.HandleFunc("/studet-score", formGetStudentScore)
	http.HandleFunc("/subject-score", formGetSubjectScore)
	http.HandleFunc("/general-subject", formGetGeneralSubject)
	fmt.Println("Corriendo servidor...")
	http.ListenAndServe(":9000", nil)
}
