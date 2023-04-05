package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/container/v1"
	"google.golang.org/api/option"
)

func mainDF() {
	projectID := "peaceful-nation-305119"
	keyFilePath := "../peaceful-nation-305119-3afbc04ee210.json"

	ctx := context.Background()

	containerService, err := container.NewService(ctx, option.WithCredentialsFile(keyFilePath))
	if err != nil {
		log.Fatalf("Error creating Compute Engine service client %v", err)
	}

	_, err = containerService.Projects.Locations.Clusters.Create(
		fmt.Sprintf("projects/%S/locations/-", projectID),
		&container.CreateClusterRequest{
			Cluster: &container.Cluster{
				Name: "my-cluster",
			},
		},
	).Do()

	if err != nil {
		log.Fatalf("Error enabling Compute Engine API: %v", err)
	}

	fmt.Println("Kubernetes Engine API enabled successfully")
}
