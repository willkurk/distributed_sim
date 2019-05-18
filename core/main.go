package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/willkurk/distgame/protos"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

type worldServer struct {
	World *protos.World
}

func (s *worldServer) WorldDownload(req *protos.WorldRequest, stream protos.World_WorldDownloadServer) error {
	log.Output(1, "World Start initiated")
	return s.clientStream(stream)
}

// our main function
func main() {
	world1 := &protos.World{Id: 1, Entity: nil}
	world1.Entity = make([]*protos.Entity, 0)
	world1.Entity = append(world1.Entity, &protos.Entity{Type: "grass", Area: &protos.Rect2D{X: 0, Y: 0, Width: 1000, Height: 1000}})
	world1.Entity = append(world1.Entity, &protos.Entity{Type: "villager", Area: &protos.Rect2D{X: 100, Y: 100, Width: 50, Height: 50}})
	world = append(world, world1)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 55600))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterWorldSyncServer(grpcServer, &worldServer{})
	grpcServer.Serve(lis)
}
