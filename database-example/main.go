package main

import (
	"database-example/entity"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode = disable", host, port, user, password, dbname)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "enigmacamp"
)

func main() {
	// student := entity.Student{Id: 9, Name: "Siti", Email: "siti@gmail.com", Address: "Surabaya", BirthDate: time.Date(2000, 11, 20, 0, 0, 0, 0, time.Local), Gender: "F"}

	// addStudent(student)
	// updateStudent(student)
	// deleteStudent("9")

	// students := getAllStudent()
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	// }

	// fmt.Println(getStudentById(7))

	students := searchBy("Ri", "2000-02-02")

	for _, student := range students {
		fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	}
}

func addStudent(student entity.Student) {
	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO MST_STUDENT(id, name, email, address, birth_date, gender) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Insert Data!!")
	}
}

func updateStudent(student entity.Student) {
	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE MST_STUDENT SET name = $2, email = $3, address = $4, birth_date = $5, gender = $6 WHERE id = $1;"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Update Data!!")
	}
}

func deleteStudent(id string) {
	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM MST_STUDENT WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Delete Data!!")
	}
}

func getAllStudent() []entity.Student {
	db := connectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM MST_STUDENT;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func getStudentById(id int) entity.Student {
	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM MST_STUDENT WHERE id = $1;"

	student := entity.Student{}
	err = db.QueryRow(sqlStatement, id).Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)

	if err != nil {
		panic(err)
	}

	return student
}

func searchBy(name string, birthDate string) []entity.Student {
	db := connectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM MST_STUDENT WHERE name LIKE $1 AND birth_date = $2;"

	rows, err := db.Query(sqlStatement, "%"+name+"%", birthDate)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func scanStudent(rows *sql.Rows) []entity.Student {
	students := []entity.Student{}
	var err error

	for rows.Next() {
		student := entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)

		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return students
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")
	return db
}
