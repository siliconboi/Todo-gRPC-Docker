package controllers

import (
	pb "Appointy/Tasks/grpc-crud/protobuf"
	"context"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func GetAll(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Second)
	defer cancel()
	res, err := c.GetTasks(ctx, &pb.TaskRequest{})
	if err != nil {
		w.Write([]byte("Error"))
	}
	p, err := proto.Marshal(res)
	w.Write(p)
}

func AddEntry(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	w.Write([]byte("Add entry"))
}

func UpdateEntry(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	w.Write([]byte("Update entry"))
}
func DeleteEntry(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	w.Write([]byte("Delete entry"))
}
func PatchEntry(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	w.Write([]byte("Patch entry"))
}
