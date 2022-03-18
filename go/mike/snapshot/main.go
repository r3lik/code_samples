package main

import (
	"context"

	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
}
