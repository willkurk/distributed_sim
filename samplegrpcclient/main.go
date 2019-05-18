/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"context"
	"flag"
	pb "github.com/willkurk/distgame/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"log"
	"time"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

/*
func runWorldStart(client pb.WorldSyncClient) {
	log.Printf("Getting World")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	world, err := client.WorldStart(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(world.String())
}


// printFeatures lists all the features within the given bounding Rectangle.
func printFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}
*/
// runRecordRoute sends a sequence of points to server and expects to get a RouteSummary from server.
func runWorldStart(client pb.WorldSyncClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.WorldStart(ctx, &pb.WorldRequest{})
	if err != nil {
		log.Fatalf("%v.Get World(_) = _, %v", client, err)
	}

	reply, err := stream.Recv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Get World: %v", reply)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewWorldSyncClient(conn)

	runWorldStart(client)
}
