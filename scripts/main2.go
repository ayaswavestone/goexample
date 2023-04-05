package main

import (
	"context"
	"fmt"
	"io"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/protobuf/proto"
)

// createWndowsServerInstanceExternalIP creates a new Windows Server instance
// that has an external IP address.
func createWndowsServerInstanceExternalIP(
	w io.Writer,
	projectID, zone, instanceName, machineType, sourceImageFamily string,
) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"
	// machineType := "n1-standard-1"
	// sourceImageFamily := "windows-2012-r2"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %v", err)
	}
	defer instancesClient.Close()

	disk := &computepb.AttachedDisk{
		// Describe the size and source image of the boot disk to attach to the instance.
		InitializeParams: &computepb.AttachedDiskInitializeParams{
			DiskSizeGb: proto.Int64(64),
			SourceImage: proto.String(
				fmt.Sprintf(
					"projects/windows-cloud/global/images/family/%s",
					sourceImageFamily,
				),
			),
		},
		AutoDelete: proto.Bool(true),
		Boot:       proto.Bool(true),
	}

	network := &computepb.NetworkInterface{
		// If you are using a custom VPC network it must be configured
		// to allow access to kms.windows.googlecloud.com.
		// https://cloud.google.com/compute/docs/instances/windows/creating-managing-windows-instances#kms-server.
		Name: proto.String("global/networks/default"),
		AccessConfigs: []*computepb.AccessConfig{
			{
				Type: proto.String("ONE_TO_ONE_NAT"),
				Name: proto.String("External NAT"),
			},
		},
	}

	inst := &computepb.Instance{
		Name: proto.String(instanceName),
		Disks: []*computepb.AttachedDisk{
			disk,
		},
		MachineType: proto.String(fmt.Sprintf("zones/%s/machineTypes/%s", zone, machineType)),
		NetworkInterfaces: []*computepb.NetworkInterface{
			network,
		},
		// If you chose an image that supports Shielded VM,
		// you can optionally change the instance's Shielded VM settings.
		// ShieldedInstanceConfig: &computepb.ShieldedInstanceConfig{
		//      EnableSecureBoot: proto.Bool(true),
		//      EnableVtpm: proto.Bool(true),
		//      EnableIntegrityMonitoring: proto.Bool(true),
		// },
	}

	req := &computepb.InsertInstanceRequest{
		Project:          projectID,
		Zone:             zone,
		InstanceResource: inst,
	}

	op, err := instancesClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to create instance: %v", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %v", err)
	}

	fmt.Fprintf(w, "Instance created\n")

	return nil
}

func main2() {
	var w io.Writer
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"
	// machineType := "n1-standard-1"
	// sourceImageFamily := "windows-2012-r2"
	// Creation and request for a new Windows Server Instance
	createWndowsServerInstanceExternalIP(w, "peaceful-nation-305119", "europe-central2-b", "VM-03-2023", "n1-standard-1", "windows-2022-r2")
	fmt.Println("Instance lancé")
}
