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

var world []*protos.WorldRender

type worldSyncRender struct {
	WorldRender *protos.WorldRender
}

func (s *worldSyncRender) clientStream(stream protos.WorldSyncRender_WorldStartRenderServer) error {
	var state int32 = 1
	for {
		log.Print("streaming", state)
		for index, element := range world[0].Entity {
			if element.Type == "villager" {
				if state == 1 {
					state = 2
					element.Area.X = element.Area.X + 10
				} else if state == 2 {
					state = 3
					element.Area.Y = element.Area.Y + 10
				} else if state == 3 {
					state = 4
					element.Area.X = element.Area.X - 10
				} else if state == 4 {
					state = 1
					element.Area.Y = element.Area.Y - 10
				}
				world[0].Entity[index] = element
			}
		}
		log.Print("stream now")
		if err := stream.Send(world[0]); err != nil {
			log.Print(err)
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *worldSyncRender) WorldStartRender(req *protos.WorldRenderRequest, stream protos.WorldSyncRender_WorldStartRenderServer) error {
	log.Output(1, "World Start initiated")
	return s.clientStream(stream)
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
