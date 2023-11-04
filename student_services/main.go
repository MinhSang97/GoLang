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

	eric := payload.Student{}

	JSONString := `{"FirstName":"Eric","ho":"Nguyen","tuoi":0,"Grade":0,"ClassName":"", "EntranceDate": "2023-11-01 12:33:20"}`
	eric.FromJson(JSONString)

	db := dbutil.ConnectDB()

	repo := mysql.NewStudentRepository(db)
	// student := eric.ToModel()
	// err := repo.InsertOne(context.Background(), &student)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(student)

	students, err := repo.GetOneByID(context.Background(), 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(students)

	// studentsall, err := repo.GetAll(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(studentsall)

	// Update a student
	updatedStudent := students
	updatedStudent.FirstName = "UpdatedName"
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

}
