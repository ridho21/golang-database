package main

import (
	"database/sql"
	"fmt"
	"golang-transaction/entity"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "enigmacamp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	studentEnrollment := entity.StudentEnrollment{Id: 3, Student_Id: 7, Subject: "Data Structure", Credit: 4}

	enrollSubject(studentEnrollment)
}

func enrollSubject(studentEnrollment entity.StudentEnrollment) {
	db := connectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	insertStudentEnrollment(studentEnrollment, tx)

	takenCredit := getSumCreditOfStudent(studentEnrollment.Student_Id, tx)

	updateStudent(takenCredit, studentEnrollment.Student_Id, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}
}

func insertStudentEnrollment(studentEnrollment entity.StudentEnrollment, tx *sql.Tx) {
	insertStudentEnrollment := "INSERT INTO tx_student_enrollment (id, student_id, subject, credit) VALUES ($1, $2, $3, $4);"

	_, err := tx.Exec(insertStudentEnrollment, studentEnrollment.Id, studentEnrollment.Student_Id, studentEnrollment.Subject, studentEnrollment.Credit)

	validate(err, "Insert", tx)
}

func getSumCreditOfStudent(id int, tx *sql.Tx) int {
	sumCredit := "SELECT SUM(credit) FROM tx_student_enrollment WHERE student_id = $1;"

	takenCredit := 0
	err := tx.QueryRow(sumCredit, id).Scan(&takenCredit)
	validate(err, "SELECT", tx)

	return takenCredit
}

func updateStudent(takenCredit int, studentId int, tx *sql.Tx) {
	updateStudent := "UPDATE MST_STUDENT SET taken_credit = $1 WHERE id = $2;"

	_, err := tx.Exec(updateStudent, takenCredit, studentId)
	validate(err, "UPDATE", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println("Successfully " + message + " data!")
	}
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succefully connected")
	}
	return db

}
