package main

import (
	"context"
	db "first-app/database"
	pb "first-app/proto/gql"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8001"
)

type EmployeeManagementWithGQLServer struct {
	pb.UnimplementedManageEmployeeWithGQLServer
	employees *pb.EmployeeListGQL
}

func NewEmployeeManagementServerWithGQL() *EmployeeManagementWithGQLServer {
	return &EmployeeManagementWithGQLServer{
		employees: &pb.EmployeeListGQL{},
	}
}

func (employeeManagementServer *EmployeeManagementWithGQLServer) RunEmployeeManagementServerWithGQL() error {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterManageEmployeeWithGQLServer(server, employeeManagementServer)
	reflection.Register(server)
	log.Printf("Server listening at %v", lis.Addr())

	return server.Serve(lis)
}

func (server *EmployeeManagementWithGQLServer) GetAllEmployeesWithGQL(ctx context.Context, getAllEmployeesWithGQLParams *pb.GetAllEmployeesWithGQLParams) (*pb.EmployeeListGQL, error) {
	server.employees = db.GetAllEmployeesDataWithGQL()
	return server.employees, nil
}

func (server *EmployeeManagementWithGQLServer) InsertEmployeeWithGQL(ctx context.Context, insertEmployeeWithGQLParams *pb.InsertEmployeeWithGQLParams) (*pb.ResponseGQL, error) {
	return &pb.ResponseGQL{
		Message: db.InsertEmployee(insertEmployeeWithGQLParams.GetName(), int(insertEmployeeWithGQLParams.Age), int(insertEmployeeWithGQLParams.Salary)),
	}, nil
}

func main() {
	var employeeManagementServer *EmployeeManagementWithGQLServer = NewEmployeeManagementServerWithGQL()
	if err := employeeManagementServer.RunEmployeeManagementServerWithGQL(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
