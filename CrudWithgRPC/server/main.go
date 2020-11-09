package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting server on port :50051...")

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	srv := &BlogServiceServer{}

	blogpb.RegisterBlogServiceServer(s, srv)

	fmt.Println("Connecting to MySQL...")
	sqlCtx = context.Background()
	db, err = sql.Open("mysql", "root:root123@tcp(localhost)/blog?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(sqlCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v\n", err)
	} else {
		fmt.Println("Connected to MySQL")
	}

	blogdb = db.Database("blog").Collection("blog")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

}
