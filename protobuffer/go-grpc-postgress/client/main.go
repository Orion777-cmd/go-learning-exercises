package main

import (
	"context"
	pb "github.com/orion777-cmd/Go-LEARNING-EXCERCISES/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could not connect to grpc server: %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	newTodos := []*pb.NewTodo{
		&pb.NewTodo{Name: "Buy Pizza", Description: "Buy pizza for dinner", Done: false},
		&pb.NewTodo{Name: "Buy Milk", Description: "Buy milk for breakfast", Done: false},
	}

	for _, newTodo := range newTodos{
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: newTodo.Name, Description: newTodo.Description, Done: newTodo.Done})
		if err != nil{
			log.Fatalf("Could not create todo: %v", err)
		}

		log.Printf(`
			ID: %s
			Name: %s
			Description: %s
			Done: %t
		`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}
}