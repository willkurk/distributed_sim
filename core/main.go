package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/willkurk/distgame/protos"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

var world []*protos.World

type ActorData struct {
	Explore float32
	Spent   int
}

var lock sync.Mutex

type Point struct {
	X int32
	Y int32
}

var actorMap map[string]*ActorData

var rate time.Duration

type worldServer struct {
	World *protos.World
}

func (s *worldServer) WorldStartListen(req *protos.WorldRequest, stream protos.WorldListen_WorldStartListenServer) error {
	log.Output(1, "World Listen initiated")
	return s.clientStream(stream)
}

func (s *worldServer) clientStream(stream protos.WorldListen_WorldStartListenServer) error {
	for {
		//log.Print("stream now")
		lock.Lock()
		if err := stream.Send(world[0]); err != nil {
			log.Print(err)
			return err
		}
		lock.Unlock()
		time.Sleep(rate)
	}
	return nil
}

func runWorld() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	point := Point{X: 0, Y: 0}

	for {
		if point.X < 1000 && point.Y == 0 {
			point.X = point.X + 20
		} else if point.X == 1000 && point.Y < 1000 {
			point.Y = point.Y + 20
		} else if point.X > 0 && point.Y == 1000 {
			point.X = point.X - 20
		} else if point.X == 0 && point.Y > 0 {
			point.Y = point.Y - 20
		}
		collided := make(map[string]int)
		for _, element := range world[0].Actor {
			if element.Type == "villager" {
				if actorMap[element.Id].Spent == 0 {
					for _, element2 := range world[0].Actor {
						if element.Type == "villager" && element.Id != element2.Id {
							_, ok := collided[element.Id+element2.Id]
							if !ok {
								collided[element.Id+element2.Id] = 0
								if element.Area.X >= element2.Area.X && element.Area.X <= element2.Area.X+element2.Area.Width {
									if element.Area.Y >= element2.Area.Y && element.Area.Y <= element2.Area.Y+element2.Area.Height {
										id := uuid.New().String()
										lock.Lock()
										actorMap[id] = &ActorData{Explore: (actorMap[element.Id].Explore + actorMap[element2.Id].Explore) / float32(2.0), Spent: 120}
										world[0].Actor = append(world[0].Actor, &protos.Actor{Id: id, Type: "villager", Modifier: actorMap[id].Explore, Area: &protos.Rect2D{X: int32(r1.Intn(1000)), Y: int32(r1.Intn(1000)), Width: 50, Height: 50}})
										lock.Unlock()
										actorMap[element.Id].Spent = 60
										actorMap[element2.Id].Spent = 60
									}
								}
							}
						}
					}
				} else {
					actorMap[element.Id].Spent = actorMap[element.Id].Spent - 1
				}
			}
		}
		lock.Lock()
		for index, element := range world[0].Actor {
			//log.Print("Processing id")
			//log.Print(element.Id)
			if element.Type == "villager" {
				var data *ActorData
				data, ok := actorMap[element.Id]
				if ok == false {
					actorMap[element.Id] = &ActorData{Explore: 0.5, Spent: 0}
					data = actorMap[element.Id]
				}
				state := r1.Intn(4) + 1
				move := int32(float32(r1.Intn(20)) * data.Explore)
				factor := int32(1)
				if state == 1 {
					if point.X > element.Area.X {
						factor = 3
					}
					element.Area.X = element.Area.X + move*factor
				} else if state == 2 {
					if point.Y > element.Area.Y {
						factor = 3
					}
					element.Area.Y = element.Area.Y + move*factor
				} else if state == 3 {
					if point.X < element.Area.X {
						factor = 3
					}
					element.Area.X = element.Area.X - move*factor
				} else if state == 4 {
					if point.Y < element.Area.Y {
						factor = 3
					}
					element.Area.Y = element.Area.Y - move*factor
				}
				world[0].Actor[index] = element
			}
		}
		lock.Unlock()
		time.Sleep(rate)

	}
}

// our main function
func main() {
	world1 := &protos.World{Id: 1, Terrain: nil, Actor: nil}
	world1.Terrain = make([]*protos.Terrain, 0)
	world1.Terrain = append(world1.Terrain, &protos.Terrain{Type: "grass", Area: &protos.Rect2D{X: 0, Y: 0, Width: 1000, Height: 1000}})
	world1.Actor = make([]*protos.Actor, 0)
	world1.Actor = append(world1.Actor, &protos.Actor{Id: "1", Type: "villager", Modifier: 0.9, Area: &protos.Rect2D{X: 100, Y: 100, Width: 50, Height: 50}})
	world1.Actor = append(world1.Actor, &protos.Actor{Id: "2", Type: "villager", Modifier: 0.1, Area: &protos.Rect2D{X: 900, Y: 100, Width: 50, Height: 50}})

	world = append(world, world1)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 55500))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	actorMap = make(map[string]*ActorData)
	rate = time.Millisecond * 150

	actorMap["1"] = &ActorData{Explore: 0.9, Spent: 0}
	actorMap["2"] = &ActorData{Explore: 0.1, Spent: 0}
	go runWorld()

	grpcServer := grpc.NewServer()
	protos.RegisterWorldListenServer(grpcServer, &worldServer{})
	grpcServer.Serve(lis)
}
