package main

import (
	"log"
	"context"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/MilaSnetkova/grpc/internal/config"

	desc "github.com/MilaSnetkova/gRPC/pkg/note_v1"
) 

const (
	address = "localhost:50051"
	noteID = 12
)

func main() {
	connect, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect server: #{err}")
	}
	defer connect.Close()

	c := desc.NewNoteV1Client(connect)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get note by id: #{err}")
	}

	log.Printf(color.BlueString("Note info:\n"), color.GreenString("%+v", r.GetNote()))
}