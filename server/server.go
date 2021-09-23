package main

import (
	"context"
	db "first-app/database"
	pb "first-app/proto"
	"log"
	"net"

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

func (server *EmployeeManagementServer) GetAllEmployees(ctx context.Context, getAllEmployeesParam *pb.GetAllEmployeesParams) (*pb.EmployeeList, error) {
	server.employees = db.GetAllEmployeesData()
	return server.employees, nil
}

func (server *EmployeeManagementServer) InsertEmployee(ctx context.Context, newEmployee *pb.InsertEmployeeParams) (*pb.Response, error) {
	return &pb.Response{
		Message: db.InsertEmployee(newEmployee.GetName(), int(newEmployee.GetAge()), int(newEmployee.GetSalary())),
	}, nil
}

func (server *EmployeeManagementServer) UpdateEmployee(ctx context.Context, newEmployee *pb.Employee) (*pb.Response, error) {
	return &pb.Response{Message: db.UpdateEmployee(*newEmployee)}, nil
}

func (server *EmployeeManagementServer) DeleteEmployee(ctx context.Context, deleteEmployeeParams *pb.DeleteEmployeeParams) (*pb.Response, error) {
	return &pb.Response{Message: db.DeleteEmployee(int(deleteEmployeeParams.Id))}, nil
}

func main() {
	var employeeManagementServer *EmployeeManagementServer = NewEmployeeManagementServer()
	if err := employeeManagementServer.RunEmployeeManagementServer(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
