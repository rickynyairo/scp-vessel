package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	pb "github.com/rickynyairo/scp-vessel/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	srv := k8s.NewService(
		// the name should equal the package name provided in the proto definition
		micro.Name("vessel"),
		micro.Version("latest")
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shipping_container_platform").Collection("vessel")
	repository := &VesselRepository{
		vesselCollection,
	}

	// Register our implementation with
	pb.RegisterVesselsHandler(srv.Server(), &Handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
