package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(client)
	if err != nil {
		log.Fatal(err)
	}

	snapshot := &compute.Snapshot{
		Description: "Snapshot of my-disk",
		Name:        "my-disk-snapshot",
	}
	resp, err := computeService.Disks.CreateSnapshot("my-project", "us-central1-a", "my-disk", snapshot).Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Operation: %s", resp.SelfLink)
}
package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(client)
	if err != nil {
		log.Fatal(err)
	}

	snapshot := &compute.Snapshot{
		Description: "Snapshot of my-disk",
		Name:        "my-disk-snapshot",
	}
	resp, err := computeService.Disks.CreateSnapshot("my-project", "us-central1-a", "my-disk", snapshot).Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Operation: %s", resp.SelfLink)
}
