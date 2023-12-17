package main

import (
	"context"
	"log"
	"time"
	pb "go-grpc-postgress/proto"
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
		{name: "Buy Pizza", description: "Buy pizza for dinner", done: false},
		{name: "Buy Milk", description: "Buy milk for breakfast", done: false},
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