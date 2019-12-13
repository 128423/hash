package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/128423/hash/server/controllers"
	"github.com/128423/hash/server/database"
	"github.com/128423/hash/server/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting in the data base")
	var err error
	ctx, lx := context.WithTimeout(context.Background(), 10*time.Second)
	database.Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error in connect to db : %v", err)
	}
	err = database.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error in connect to db : %v", err)
	}

	fmt.Println("Starting Listener")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterHashGrpcAPiServer(s, &controllers.Server{})

	reflection.Register(s)

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping the server...")
	s.Stop()
	fmt.Println("Closing the listener...")
	lis.Close()
	fmt.Print("Context Cancel")
	lx()
	fmt.Println("END")
}
