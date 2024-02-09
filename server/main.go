package main

import (
	"Appointy/Tasks/grpc-crud/db"
	pb "Appointy/Tasks/grpc-crud/protobuf"
	"context"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRouteGuideServer
	h  *db.DBHandler
	mu sync.Mutex
}

func (s *server) GetTasks(ctx context.Context, in *pb.TaskRequest) (*pb.TaskListResponse, error) {
	task := &pb.Task{}
	tx := s.h.DB.Find(task)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &pb.TaskListResponse{Tasks: []*pb.Task{task}}, nil
}

func (s *server) AddTask(ctx context.Context, in *pb.AddTaskRequest) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}

func (s *server) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}

func (s *server) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	DB, err := db.CreateDB()
	if err != nil {
		panic("failed to connect database")
	}
	pb.RegisterRouteGuideServer(grpcServer, &server{h: db.New(DB)})
	lis, _ := net.Listen("tcp", ":8081")
	grpcServer.Serve(lis)
}
