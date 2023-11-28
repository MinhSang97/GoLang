package main

import (
	"app/dbutil"
	"app/payload"
	"app/repo/mysql"
	"context"
	"fmt"
	"os"
)

func main() {
	db := dbutil.ConnectDB()

	eric := payload.Student{}

	JSONString := `{"FirstName":"Eric","LastName":"Nguyen","tuoi":0,"Grade":0,"ClassName":""}`
	eric.FromJson(JSONString)

	repo := mysql.NewStudentRepository(db)
	student := eric.ToModel()

	err := repo.InsertOne(context.Background(), &student)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(student)

	students, err := repo.GetOneByID(context.Background(), 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(students)

	studentsall, err := repo.GetAll(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(studentsall)

	// Update a student
	updatedStudent := students
	updatedStudent.LastName = "LastName"
	updatedStudent.FirstName = "FirstName"
	err = repo.UpdateOne(context.Background(), 3, &updatedStudent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Verify that the student has been updated
	updatedStudent, err = repo.GetOneByID(context.Background(), 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(updatedStudent)

	// Delete a student
	// err = repo.DeleteOne(context.Background(), 6)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)

	// }
	// fmt.Println("Delete student by id", 6)

	// // Verify that the student has been deleted
	// student, err := repo.GetOneByID(context.Background(), 6)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(student)

}
