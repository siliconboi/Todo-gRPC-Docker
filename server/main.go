package main

import (
	"Hero/Tasks/graphql-crud/db"
	pb "Hero/Tasks/graphql-crud/protobuf"
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRouteGuideServer
	mu sync.Mutex
	h  *db.DBHandler
}

func (s *server) GetTasks(ctx context.Context, in *pb.TaskRequest) (*pb.TaskListResponse, error) {
	fmt.Println("in in grpc")
	tasks := []*pb.Task{}
	tx := s.h.DB.Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &pb.TaskListResponse{Tasks: tasks}, nil
}

func (s *server) AddTask(ctx context.Context, in *pb.AddTaskRequest) (*pb.TaskResponse, error) {
	fmt.Println("in in grpc")
	task := pb.Task{Title: in.Title, Description: in.Description, Duration: in.Duration}
	tx := s.h.DB.Create(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &pb.TaskResponse{Task: &task}, nil
}

func (s *server) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	fmt.Println("in in grpc")
	task := pb.Task{Id: in.Id, Title: in.Title, Description: in.Description, Duration: in.Duration}
	tx := s.h.DB.Model(&pb.Task{}).Where("id = ?", in.Id).Updates(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &pb.TaskResponse{Task: &task}, nil
}

func (s *server) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.TaskResponse, error) {
	fmt.Println("in in grpc")
	tx := s.h.DB.Delete(&pb.Task{}, in.Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &pb.TaskResponse{}, nil
}

func main() {
	DB, _ := db.CreateDB()
	h := db.New(DB)
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &server{h: h})
	list, _ := net.Listen("tcp", ":8081")
	grpcServer.Serve(list)
}
