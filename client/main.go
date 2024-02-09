package main

import (
	"Appointy/Tasks/grpc-crud/routers"
	"log"
	"net/http"

	pb "Appointy/Tasks/grpc-crud/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("grpc-first:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect to grpc server")
	}
	defer conn.Close()
	c := pb.NewRouteGuideClient(conn)

	router := routers.CreateRouter(c)
	log.Fatal(http.ListenAndServe(":8080", router))
}
