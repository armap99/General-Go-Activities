package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type StudentData struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}

type StudentDataRecive struct {
	ID      uint64  `json:"id"`
	Name    string  `json:"name"`
	Subject string  `json:"subject"`
	Score   float64 `json:"score"`
}

type Subject struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}



var studentsData []StudentData

func student(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		fmt.Println("get case")
		res_json, err := getRecordStudents()
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set(
			"Content-Type",
			"application/json",
		)
		res.Write(res_json)
	case "POST":
		var student StudentDataRecive
		err := json.NewDecoder(req.Body).Decode(&student)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res_json := AddStudentData(student)
		res.Header().Set(
			"Content-Type",
			"application/json",
		)
		res.Write(res_json)
		fmt.Println(studentsData)
	default:
	}
}

func studentId(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseUint(strings.TrimPrefix(req.URL.Path, "/student/"), 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(id)
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		res_json, err := getRecordStudentsID(id)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set(
			"Content-Type",
			"application/json",
		)
		res.Write(res_json)
	case "DELETE":
		res_json, err := deleteRecordStudentsID(id)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set(
			"Content-Type",
			"application/json",
		)
		res.Write(res_json)
	case "PUT":
		var student StudentDataRecive

		err := json.NewDecoder(req.Body).Decode(&student)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res_json := updateStudent(id, student)
		res.Header().Set(
			"Content-Type",
			"application/json",
		)
		res.Write(res_json)
	default:
	}
}

func AddStudentData(data StudentDataRecive) []byte {
	if data.ID != 0 && data.Name != " " && data.Subject != " " && data.Score != 0 {
		subject := Subject{
			Name:  data.Subject,
			Score: data.Score,
		}
		subjectArray := []Subject{
			subject,
		}

		student := StudentData{
			ID:       data.ID,
			Name:     data.Name,
			Subjects: subjectArray,
		}

		if studentExist(data.ID) {
			if !studentSubjectExist(student) {
				addNewSubjectToStudent(student)
				return []byte(`{"code": "Datos del estudiante han sido agregados con exito"}`)
			} else {
				return []byte(`{"code": "La materia ya fue ragistrada para el alumno"}`)
			}
		} else {
			studentsData = append(studentsData, student)
			return []byte(`{"code": "El registro se ha agregado con exito"}`)
		}
	} else {
		return []byte(`{"code": "No se completo la operacion por falta de informacion"}`)
	}
}

func getRecordStudents() ([]byte, error) {
	jsonData, err := json.MarshalIndent(studentsData, "", "	")
	if err != nil {
		return jsonData, nil
	}
	return jsonData, err
}

func getRecordStudentsID(id uint64) ([]byte, error) {
	jsonData := []byte(`{"code":"No exite el alumno con ese ID"}`)
	if studentExist(id) {
		var student StudentData
		for i := 0; i < len(studentsData); i++ {
			if studentsData[i].ID == id {
				student = studentsData[i]
				break
			}
		}
		jsonData, err := json.MarshalIndent(student, "", "	")
		if err != nil {
			return jsonData, err
		}
		return jsonData, nil
	}
	return jsonData, nil
}

func deleteRecordStudentsID(id uint64) ([]byte, error) {
	jsonData := []byte(`{"code":"No exite el alumno con ese ID"}`)
	if studentExist(id) {
		var tempRecord []StudentData
		for i := 0; i < len(studentsData); i++ {
			if studentsData[i].ID != id {
				tempRecord = append(tempRecord, studentsData[i])
			}
		}
		studentsData = tempRecord
		fmt.Println("eliminados: ", studentsData)
		jsonData = []byte(`{"code":"El registro ha sido eliminado con exito"}`)
		return jsonData, nil
	}
	return jsonData, nil
}

func updateStudent(id uint64, student StudentDataRecive) []byte {
	subjectFound := false
	if studentExist(id) {
		for i := 0; i < len(studentsData); i++ {
			if studentsData[i].ID == id {
				for j := 0; j < len(studentsData[i].Subjects); j++ {
					if studentsData[i].Subjects[j].Name == student.Subject {
						studentsData[i].Subjects[j].Score = student.Score
						subjectFound = true
						break
					}
				}
			}
		}
		if subjectFound {
			return []byte(`{"code": "El registro se actualizo exitosamente"}`)
		} else {
			return []byte(`{"code": "La materia recibida no existe"}`)
		}
	}
	return []byte(`{"code": "El ID de alumno no existe"}`)
}

func studentExist(id uint64) bool {
	for i := 0; i < len(studentsData); i++ {
		if studentsData[i].ID == id {
			return true
		}
	}
	return false
}

func studentSubjectExist(student StudentData) bool {
	for i := 0; i < len(studentsData); i++ {
		if studentsData[i].ID == student.ID {
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
		if studentsData[i].ID == studentData.ID {
			
			studentsData[i].Subjects = append(studentsData[i].Subjects, studentData.Subjects[0])
		}
	}
}

func main() {
	http.HandleFunc("/student", student)
	http.HandleFunc("/student/", studentId)
	fmt.Println("API ejecutandose ...")
	http.ListenAndServe(":9000", nil)
}
