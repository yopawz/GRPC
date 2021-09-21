package main

import (
	"context"
	"database/sql"
	"first-app/models"
	pb "first-app/proto"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"

	reflection "google.golang.org/grpc/reflection"
)

const (
	port = ":8000"
)

type EmployeeManagementServer struct {
	pb.UnimplementedManageEmployeeServer
	employees *pb.EmployeeList
}

func NewEmployeeManagementServer() *EmployeeManagementServer {
	return &EmployeeManagementServer{
		employees: &pb.EmployeeList{},
	}
}

func (employeeManagementServer *EmployeeManagementServer) RunEmployeeManagementServer() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterManageEmployeeServer(server, employeeManagementServer)
	reflection.Register(server)
	log.Printf("Server listening at %v", lis.Addr())

	return server.Serve(lis)
}

func createConnection() *sql.DB {

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

func (server *EmployeeManagementServer) GetAllEmployees(ctx context.Context, getAllEmployeesParam *pb.GetAllEmployeesParams) (*pb.EmployeeList, error) {
	db := createConnection()

	defer db.Close()

	fmt.Println("QUERY")
	sqlStatement := "SELECT * FROM employee"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Println("INVALID")
		log.Fatalf("luar loop %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var currentEmployee models.Employee

		err = rows.Scan(&currentEmployee.Id, &currentEmployee.Name, &currentEmployee.Age, &currentEmployee.Salary)

		if err != nil {
			fmt.Println("KOSONG")
			log.Fatalf("dalem loop %v", err)
		}

		employee := &pb.Employee{Id: int32(currentEmployee.Id), Name: currentEmployee.Name, Age: int32(currentEmployee.Age), Salary: int32(currentEmployee.Salary)}
		fmt.Println(server.employees.Employees)
		fmt.Println("MASUK")
		server.employees.Employees = append(server.employees.Employees, employee)
	}

	return server.employees, nil
}

func (server *EmployeeManagementServer) InsertEmployee(ctx context.Context, newEmployee *pb.InsertEmployeeParams) (*pb.Response, error) {
	log.Printf("Received : %v", newEmployee.GetName())
	db := createConnection()

	defer db.Close()

	sqlStatement := "INSERT INTO employee (name, age, salary) VALUES ($1, $2, $3) RETURNING id"

	var id int

	err := db.QueryRow(sqlStatement, newEmployee.GetName(), newEmployee.GetAge(), newEmployee.GetSalary()).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.Response{Message: "Successfully inserted employee data with id: " + string(id)}, nil
}

func (server *EmployeeManagementServer) UpdateEmployee(ctx context.Context, newEmployee *pb.Employee) (*pb.Response, error) {
	db := createConnection()

	defer db.Close()

	sqlStatement := "UPDATE employee SET name=$1, age=$2, salary=$3 WHERE id=$4"
	response, err := db.Exec(sqlStatement, newEmployee.Name, newEmployee.Age, newEmployee.Salary, newEmployee.Id)

	if err != nil {
		log.Fatalf("Failed to update the data: %v", err)
	}

	rowsAffected, err := response.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected data: %v", err)
	}

	updateResponse := fmt.Sprintf("Total row / record affected: %v", rowsAffected)
	fmt.Println(updateResponse)

	return &pb.Response{Message: updateResponse}, nil
}

func (server *EmployeeManagementServer) DeleteEmployee(ctx context.Context, deleteEmployeeParams *pb.DeleteEmployeeParams) (*pb.Response, error) {
	db := createConnection()

	sqlStatement := "DELETE FROM employee WHERE id=$1"

	response, err := db.Exec(sqlStatement, deleteEmployeeParams.Id)

	if err != nil {
		log.Fatalf("Failed to delete the data: %v", err)
	}

	rowsAffected, err := response.RowsAffected()

	if err != nil {
		log.Fatalf("Errror while checking the data")
	}

	deleteResponse := fmt.Sprintf("Total rows affected : %v", rowsAffected)

	return &pb.Response{Message: deleteResponse}, nil
}

func main() {
	var employeeManagementServer *EmployeeManagementServer = NewEmployeeManagementServer()
	if err := employeeManagementServer.RunEmployeeManagementServer(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
