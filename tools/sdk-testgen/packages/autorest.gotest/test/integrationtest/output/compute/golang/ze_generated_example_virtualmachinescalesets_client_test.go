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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListVirtualMachineScaleSetsInASubscriptionByLocation.json
func ExampleVirtualMachineScaleSetsClient_ListByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	pager := client.ListByLocation("<location>",
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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateACustomImageScaleSetFromAnUnmanagedGeneralizedOsImage.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		golang.VirtualMachineScaleSet{
			Location: to.StringPtr("<location>"),
			Properties: &golang.VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &golang.UpgradePolicy{
					Mode: golang.UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &golang.VirtualMachineScaleSetVMProfile{
					NetworkProfile: &golang.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*golang.VirtualMachineScaleSetNetworkConfiguration{
							{
								Name: to.StringPtr("<name>"),
								Properties: &golang.VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*golang.VirtualMachineScaleSetIPConfiguration{
										{
											Name: to.StringPtr("<name>"),
											Properties: &golang.VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &golang.APIEntityReference{
													ID: to.StringPtr("<id>"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &golang.VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("<admin-password>"),
						AdminUsername:      to.StringPtr("<admin-username>"),
						ComputerNamePrefix: to.StringPtr("<computer-name-prefix>"),
					},
					StorageProfile: &golang.VirtualMachineScaleSetStorageProfile{
						OSDisk: &golang.VirtualMachineScaleSetOSDisk{
							Name:         to.StringPtr("<name>"),
							Caching:      golang.CachingTypesReadWrite.ToPtr(),
							CreateOption: golang.DiskCreateOptionTypes("FromImage").ToPtr(),
							Image: &golang.VirtualHardDisk{
								URI: to.StringPtr("<uri>"),
							},
						},
					},
				},
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
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ForceDeleteVirtualMachineScaleSets.json
func ExampleVirtualMachineScaleSetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&golang.VirtualMachineScaleSetsBeginDeleteOptions{ForceDeletion: to.BoolPtr(true)})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetVirtualMachineScaleSetAutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachineScaleSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&golang.VirtualMachineScaleSetsGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientGetResult)
}
