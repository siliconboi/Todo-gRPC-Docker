package main

import (
	pb "Hero/Tasks/graphql-crud/protobuf"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	c     pb.RouteGuideClient
	cInit sync.Once
)

func initClient() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect to grpc")
	}
	c = pb.NewRouteGuideClient(conn)
	fmt.Println("grpc client")
}

var taskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"duration": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"tasks": &graphql.Field{
				Type: graphql.NewList(taskType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// ctx := p.Context
					tasks, _ := c.GetTasks(context.Background(), &pb.TaskRequest{})
					fmt.Println("in in gql")
					return tasks.Tasks, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createTask": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					title, _ := params.Args["title"].(string)
					description, _ := params.Args["description"].(string)
					duration, _ := params.Args["duration"].(string)
					task := &pb.AddTaskRequest{
						Title:       title,
						Description: description,
						Duration:    duration,
					}
					createdTask, _ := c.AddTask(context.Background(), task)
					return createdTask.Task, nil
				},
			},
			"updateTask": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					title, _ := params.Args["title"].(string)
					description, _ := params.Args["description"].(string)
					duration, _ := params.Args["duration"].(string)
					task := &pb.UpdateTaskRequest{
						Id:          int32(id),
						Title:       title,
						Description: description,
						Duration:    duration,
					}
					updatedTask, _ := c.UpdateTask(context.Background(), task)
					return updatedTask.Task, nil
				},
			},
			"deleteTask": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					task := &pb.DeleteTaskRequest{
						Id: int32(id),
					}
					deletedTask, _ := c.DeleteTask(context.Background(), task)
					return deletedTask.Task, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	cInit.Do(initClient)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if c == nil {
			http.Error(w, "gRPC client is not initialized yet", http.StatusInternalServerError)
			return
		}
		res := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(res)
	})
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
