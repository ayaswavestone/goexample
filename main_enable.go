package main

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/servicemanagement/v1"
)

func maiEEn() {
	ctx := context.Background()

	// Replace with your project ID and JSON key file path.
	projectID := "peaceful-nation-305119"
	keyfilePath := "../peaceful-nation-305119-3afbc04ee210.json"

	serviceManagementService, err := servicemanagement.NewService(ctx, option.WithCredentialsFile(keyfilePath))
	if err != nil {
		log.Fatalf("Failed to create service management service: %v", err)
	}

	/*
		App Engine: appengine.googleapis.com
		Compute Engine: compute.googleapis.com
		Google Kubernetes Engine : container.googleapis.com

	*/
	apiToEnable := "container.googleapis.com"
	operation, err := serviceManagementService.Services.Enable(apiToEnable, &servicemanagement.EnableServiceRequest{
		ConsumerId: "project:" + projectID,
	}).Do()

	if err != nil {
		log.Fatalf("Failed to enable Compute Engine API: %v", err)
	}

	log.Printf("API enabling started. Operation ID: %s", operation.Name)
}
