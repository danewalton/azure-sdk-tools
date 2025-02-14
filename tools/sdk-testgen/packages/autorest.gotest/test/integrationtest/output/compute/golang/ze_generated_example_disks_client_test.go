//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package golang_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/CreateAManagedDiskWithDiskAccess.json
func ExampleDisksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDisksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<disk-name>",
		golang.Disk{
			Location: to.StringPtr("<location>"),
			Properties: &golang.DiskProperties{
				CreationData: &golang.CreationData{
					CreateOption: golang.DiskCreateOption("Empty").ToPtr(),
				},
				DiskAccessID:        to.StringPtr("<disk-access-id>"),
				DiskSizeGB:          to.Int32Ptr(200),
				NetworkAccessPolicy: golang.NetworkAccessPolicy("AllowPrivate").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DisksClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/CreateOrUpdateABurstingEnabledManagedDisk.json
func ExampleDisksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDisksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<disk-name>",
		golang.DiskUpdate{
			Properties: &golang.DiskUpdateProperties{
				BurstingEnabled: to.BoolPtr(true),
				DiskSizeGB:      to.Int32Ptr(1024),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DisksClientUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/GetInformationAboutAManagedDisk.json
func ExampleDisksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDisksClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<disk-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DisksClientGetResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/ListManagedDisksInAResourceGroup.json
func ExampleDisksClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDisksClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/ListManagedDisksInASubscription.json
func ExampleDisksClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDisksClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
