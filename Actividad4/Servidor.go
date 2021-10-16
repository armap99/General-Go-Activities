package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

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

type Server struct{}

var studentsData []StudentData
var subjects []SubjectScore

//////////////////////////////////////////////////////////
//Servidor
//////////////////////////////////////////////////////////
func (this *Server) AddStudentData(data StudentDataRecive, reply *string) error {
	if data.Name != " " && data.Subject != " " && data.Score != 0 {

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

				*reply = " Agregado con exito"
			} else {
				*reply = "Materia " + data.Subject + " registrada para " + data.Name
			}
		} else {
			studentsData = append(studentsData, student)
			if subjectExist(student) {
				addNewStudentToSubject(student)
			} else {

				addNewSubjectToSubjects(student)
			}
			*reply = "El registro se ha agregado con exito"
		}
	} else {
		*reply = "No se completo la operacion por falta de informacion"
	}
	fmt.Println(studentsData)
	fmt.Println(subjects)
	return nil
}

func (this *Server) GetStudentAverage(name string, reply *string) error {
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
			*reply = "El promedio de " + name + " es: " + strconv.FormatFloat(averageStudent, 'f', -1, 64)
		} else {
			*reply = "NO existe el registro ingresado"
		}
	} else {
		*reply = "Aun no hay alumnos registrados"
	}
	return nil
}

func (this *Server) GetGeneralAverageByStudents(petition string, reply *string) error {
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
		*reply = "El promedio general de los estudiantes es: " + strconv.FormatFloat(totalAverage, 'f', -1, 64)
	} else {
		*reply = "Aun no hay alumnos registrados"
	}
	return nil
}

func (this *Server) GetAverageBySubject(subjectName string, reply *string) error {
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
			*reply = "El promedio de la materia " + subjectName + " es: " + strconv.FormatFloat(totalAvarage, 'f', -1, 64)
		} else {
			*reply = "La materia " + subjectName + " no existe en los registros"
		}
	} else {
		*reply = "Aun no hay materias registradas"
	}
	return nil
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

/////////////////////////////////////////////////////////////////////////////////////////////
//main

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Corriendo servidor...")
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
