package main

import (
	"flag"
	"fmt"
	"github.com/willkurk/distgame/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

var world []*protos.World

type worldServer struct {
	World *protos.World
}

func (s *worldServer) WorldStartListen(req *protos.WorldRequest, stream protos.WorldListen_WorldStartListenServer) error {
	log.Output(1, "World Listen initiated")
	//return s.clientStream(stream)
	return nil
}

// our main function
func main() {
	world1 := &protos.World{Id: 1, Terrain: nil, Actor: nil}
	world1.Terrain = make([]*protos.Terrain, 0)
	world1.Terrain = append(world1.Terrain, &protos.Terrain{Type: "grass", Area: &protos.Rect2D{X: 0, Y: 0, Width: 1000, Height: 1000}})
	world1.Actor = make([]*protos.Actor, 0)
	world1.Actor = append(world1.Actor, &protos.Actor{Type: "villager", Area: &protos.Rect2D{X: 100, Y: 100, Width: 50, Height: 50}})
	world = append(world, world1)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 55500))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterWorldListenServer(grpcServer, &worldServer{})
	grpcServer.Serve(lis)
}
