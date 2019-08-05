package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/willkurk/distgame/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
)

var world []*protos.WorldRender

var (
	coreAddr string
)

type worldSyncRender struct {
	WorldRender *protos.WorldRender
}

func (s *worldSyncRender) WorldStartRender(req *protos.WorldRenderRequest, stream protos.WorldSyncRender_WorldStartRenderServer) error {
	log.Output(1, "World Start initiated")
	return s.clientStream(stream)
}

func (s *worldSyncRender) clientStream(stream protos.WorldSyncRender_WorldStartRenderServer) error {
	log.Println("Connecting to Core API")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(coreAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := protos.NewWorldListenClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	core_stream, err := client.WorldStartListen(ctx, &protos.WorldRequest{})
	if err != nil {
		log.Fatalf("%v.Get World(_) = _, %v", client, err)
	}

	for {
		resp, err := core_stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", core_stream, err, nil)
			continue
		}
		log.Printf("Get World: %v", resp)
		world_received := resp
		//world1 := &protos.World{Id: 1, Terrain: nil, Actor: nil}
		//	world1.Terrain = make([]*protos.Terrain, 0)
		//	world1.Terrain = append(world1.Terrain, &protos.Terrain{Type: "grass", Area: &protos.Rect2D{X: 0, Y: 0, Width: 1000, Height: 1000}})
		//	world1.Actor = make([]*protos.Actor, 0)
		//	world1.Actor = append(world1.Actor, &protos.Actor{Type: "villager", Area: &protos.Rect2D{X: 100, Y: 100, Width: 50, Height: 50}}).
		rendered := &protos.WorldRender{Id: world_received.Id, Entity: nil}
		rendered.Entity = make([]*protos.EntityRender, 0)
		for _, terrain := range world_received.Terrain {
			rendered.Entity = append(rendered.Entity, &protos.EntityRender{Type: terrain.Type, Area: terrain.Area, Color: 1})
		}
		for _, actor := range world_received.Actor {
			rendered.Entity = append(rendered.Entity, &protos.EntityRender{Type: actor.Type, Area: actor.Area, Color: actor.Modifier})
		}
		if err := stream.Send(rendered); err != nil {
			log.Print(err)
			return err
		}

	}
	return nil
}

func GetWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(world)
}

// our main function
func main() {
	world1 := &protos.WorldRender{Id: 1, Entity: nil}
	world1.Entity = make([]*protos.EntityRender, 0)
	world1.Entity = append(world1.Entity, &protos.EntityRender{Type: "grass", Area: &protos.Rect2D{X: 0, Y: 0, Width: 1000, Height: 1000}})
	world1.Entity = append(world1.Entity, &protos.EntityRender{Type: "villager", Area: &protos.Rect2D{X: 100, Y: 100, Width: 50, Height: 50}})
	world = append(world, world1)

	flag.StringVar(&coreAddr, "coreAddr", "localhost:55500", "Address of core serve")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 55600))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterWorldSyncRenderServer(grpcServer, &worldSyncRender{WorldRender: world1})
	grpcServer.Serve(lis)

	router := mux.NewRouter()

	corsObj := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/world", GetWorld).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))
}
