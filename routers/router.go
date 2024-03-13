package routers

import (
	"Hero/Tasks/grpc-crud/handlers"
	pb "Hero/Tasks/grpc-crud/protobuf"
	"net/http"
)

func CreateRouter(c pb.RouteGuideClient) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGet(w, r, c)
	})
	router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleAdd(w, r, c)
	})
	router.HandleFunc("/:id", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleChange(w, r, c)
	})
	return router
}
