package database

import (
	"database/sql"
	"first-app/models"
	pb "first-app/proto"
	"fmt"
	"log"
	"os"
	"strconv"
)

type PostgreDatabase struct {
	PostgreClient *sql.DB
}

func CreateConnection() *sql.DB {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	db, err := sql.Open(dialect, dbURI)
	fmt.Println(dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Success")
	}

	return db
}

func InsertEmployee(employeeName string, employeeAge int, employeeSalary int) string {

	db := CreateConnection()

	sqlStatement := "INSERT INTO employee (name, age, salary) VALUES ($1, $2, $3) RETURNING id"

	var id int

	err := db.QueryRow(sqlStatement, employeeName, employeeAge, employeeSalary).Scan(&id)

	HandleError(fmt.Sprintf("Failed to insert employee data with id %v", id), err)

	redisClient := CreateRedisDatabase()
	key := fmt.Sprintf("employee:%v", id)

	redisClient.Client.Do("SADD", "ids", key)
	redisClient.Client.Do("HSET", key, "id", id, "name", employeeName, "age",
		employeeAge, "salary", employeeSalary)

	return fmt.Sprintf("Employee data with id %v inserted successfully", id)
}

func GetAllEmployeesData() *pb.EmployeeList {
	var redisClient *RedisDatabase = CreateRedisDatabase()
	var employeeList *pb.EmployeeList = &pb.EmployeeList{}

	var employeeIds []string

	employeeIds, err := redisClient.Client.SMembers("ids").Result()

	HandleError("Cannot Get Members for Employee Id", err)

	if len(employeeIds) > 0 {
		for _, id := range employeeIds {

			currentEmployee, err := redisClient.Client.HGetAll(string(id)).Result()
			HandleError("Cannot retrieve data from redis", err)

			employeeId, err := strconv.Atoi(currentEmployee["id"])

			HandleError("Cannot convert employee id to integer", err)

			employeeAge, err := strconv.Atoi(currentEmployee["age"])

			HandleError("Cannot convert employee age to integer", err)

			employeeSalary, err := strconv.Atoi(currentEmployee["salary"])

			HandleError("Cannot convert employee salary to integer", err)

			employee := &pb.Employee{Id: int32(employeeId), Name: currentEmployee["name"], Age: int32(employeeAge), Salary: int32(employeeSalary)}

			employeeList.Employees = append(employeeList.Employees, employee)
		}

		return employeeList
	}

	db := CreateConnection()

	defer db.Close()

	sqlStatement := "SELECT * FROM employee"

	rows, err := db.Query(sqlStatement)

	HandleError("Cannot Executed Get Query from Database", err)

	defer rows.Close()

	for rows.Next() {
		var currentEmployee models.Employee

		err = rows.Scan(&currentEmployee.Id, &currentEmployee.Name, &currentEmployee.Age, &currentEmployee.Salary)

		key := fmt.Sprintf("employee:%v", currentEmployee.Id)
		redisClient.Client.Do("SADD", "ids", key)
		redisClient.Client.Do("HSET", key, "id", currentEmployee.Id, "name", currentEmployee.Name, "age",
			currentEmployee.Age, "salary", currentEmployee.Salary)

		HandleError("Cannot Traverse Data in Employee Table", err)
		employee := &pb.Employee{Id: int32(currentEmployee.Id), Name: currentEmployee.Name, Age: int32(currentEmployee.Age), Salary: int32(currentEmployee.Salary)}
		employeeList.Employees = append(employeeList.Employees, employee)
	}

	return employeeList
}

func UpdateEmployee(newEmployee pb.Employee) string {
	db := CreateConnection()
	redisClient := CreateRedisDatabase()
	defer db.Close()

	sqlStatement := "UPDATE employee SET name=$1, age=$2, salary=$3 WHERE id=$4"
	response, err := db.Exec(sqlStatement, newEmployee.Name, newEmployee.Age, newEmployee.Salary, newEmployee.Id)

	HandleError("Failed to update the data", err)

	rowsAffected, err := response.RowsAffected()

	HandleError("Error while checking the affected data", err)
	key := fmt.Sprintf("employee:%v", newEmployee.Id)

	updateResponse := fmt.Sprintf("Total row / record affected: %v", rowsAffected)
	redisClient.Client.Do("HSET", key, "id", newEmployee.Id, "name", newEmployee.Name, "age",
		newEmployee.Age, "salary", newEmployee.Salary)

	return updateResponse
}

func DeleteEmployee(employeeId int) string {
	db := CreateConnection()
	redisClient := CreateRedisDatabase()

	sqlStatement := "DELETE FROM employee WHERE id=$1"

	response, err := db.Exec(sqlStatement, employeeId)

	HandleError("Failed to delete the data", err)

	rowsAffected, err := response.RowsAffected()

	HandleError("Error while checking the data", err)

	deleteResponse := fmt.Sprintf("Total rows affected : %v", rowsAffected)

	key := fmt.Sprintf("employee:%v", employeeId)

	redisClient.Client.Do("DEL", key)

	redisClient.Client.Do("SREM", "ids", key)

	return deleteResponse
}
