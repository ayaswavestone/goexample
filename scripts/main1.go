package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func main1() {
	ctx := context.Background()

	// Replace with your project ID and JSON key file path.
	projectID := "peaceful-nation-305119"
	keyfilePath := "../peaceful-nation-305119-3afbc04ee210.json"

	computeService, err := compute.NewService(ctx, option.WithCredentialsFile(keyfilePath))
	if err != nil {
		log.Fatalf("Failed to create compute service: %v", err)
	}

	zone := "us-central1-a"
	instanceName := "my-instance"

	config := &compute.Instance{
		Name:        instanceName,
		MachineType: fmt.Sprintf("zones/%s/machineTypes/n1-standard-1", zone),
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    fmt.Sprintf("%s-root", instanceName),
					SourceImage: "projects/debian-cloud/global/images/family/debian-10",
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				AccessConfigs: []*compute.AccessConfig{
					{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network: "global/networks/default",
			},
		},
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Email: "default",
				Scopes: []string{
					compute.ComputeReadonlyScope,
				},
			},
		},
	}

	operation, err := computeService.Instances.Insert(projectID, zone, config).Do()
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	log.Printf("Instance creation started. Operation ID: %s", operation.Name)
}
