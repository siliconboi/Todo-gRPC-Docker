package handlers

import (
	"Appointy/Tasks/grpc-crud/controllers"
	pb "Appointy/Tasks/grpc-crud/protobuf"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetAll(w, r, c)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleChange(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	switch r.Method {
	case http.MethodDelete:
		controllers.DeleteEntry(w, r, c)
	case http.MethodPut:
		controllers.UpdateEntry(w, r, c)
	case http.MethodPatch:
		controllers.PatchEntry(w, r, c)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleAdd(w http.ResponseWriter, r *http.Request, c pb.RouteGuideClient) {
	switch r.Method {
	case http.MethodPost:
		controllers.AddEntry(w, r, c)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
