package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID		int 	`json:"id"`
	Name	string 	`json:"name"`
	Age		int 	`json:"age"`
	Grade	string 	`json:"grade"`
}

// Inisiasi database
var students []Student

// Mengambil semua data yang ada dan mengembalikan HTTP status
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

// Mengambil data berdasarkan ID dan mengembalikan HTTP status
func getStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}

// Menambahkan data dan mengembalikan HTTP status
func createStudent(c echo.Context) error {
	// Inisiasi struct Student
	student := new(Student)
	
	// Jika data yang dikirim tidak sesuai dengan struct Student, maka akan mengembalikan HTTP status
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	// Menambahkan data
	student.ID = len(students) + 1
	students = append(students, *student)
	
	return c.JSON(http.StatusCreated, student)
}

// Update data berdasarkan ID dan mengembalikan HTTP status
func updateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, student := range students {
		if student.ID == id {
			// Inisiasi struct Student
			updatedStudent := new(Student)
			
			// Jika data yang dikirim tidak sesuai dengan struct Student, maka akan mengembalikan HTTP status
			if err := c.Bind(updatedStudent); err != nil {
				return c.JSON(http.StatusBadRequest, nil)
			}

			// Update data 
			students[index].Name = updatedStudent.Name
			students[index].Age = updatedStudent.Age
			students[index].Grade = updatedStudent.Grade

			return c.JSON(http.StatusOK, students[index]) 
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}

// Delete data berdasarkan ID dan mengembalikan HTTP status
func deleteStudent(c echo.Context) error {	
	id, _ := strconv.Atoi(c.Param("id"))
	for index, student := range students {
		if student.ID == id {
			students = append(students[:index], students[index+1:]...)
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}

func main() {
	// Inisiasi echo
	e := echo.New()

	// Routing
	e.GET("/students", getStudents) // Mengambil semua data
	e.GET("/students/:id", getStudent) // Mengambil data berdasarkan ID
	e.POST("/students", createStudent) // Menambahkan data
	e.PUT("/students/:id", updateStudent) // Update data berdasarkan ID
	e.DELETE("/students/:id", deleteStudent) // Delete data berdasarkan ID

	// coba
	e.HEAD("/students", getStudents) // Mengambil semua data (header)		
	e.OPTIONS("/students", getStudents) // Mengambil semua data (options)		

	// Run server
	e.Logger.Fatal(e.Start(":8080"))
}