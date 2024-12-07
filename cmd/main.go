package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
	pb "prajithtk/doctor-sevice/proto"
)

type server struct {
    pb.UnimplementedDoctorServiceServer // Embed the unimplemented server for forward compatibility.
}

// Implement methods defined in doctor.proto here...

func main() {
    lis, err := net.Listen("tcp", ":50052") // Choose appropriate port for Doctor Service.
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    pb.RegisterDoctorServiceServer(grpcServer, &server{}) // Register the service with the server

    log.Println("Starting Doctor Service on port 50052...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}