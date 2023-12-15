package main 

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/orion777-cmd/go-learning-exersices/grpc/chat"
)

func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil{
		log.Fatalf("Error when calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", response.Body)
}