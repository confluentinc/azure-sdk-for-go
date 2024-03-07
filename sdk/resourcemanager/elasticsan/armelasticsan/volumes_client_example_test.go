//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Create_MaximumSet_Gen.json
func ExampleVolumesClient_BeginCreate_volumesCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.Volume{
		Properties: &armelasticsan.VolumeProperties{
			CreationData: &armelasticsan.SourceCreationData{
				CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
				SourceID:     to.Ptr("ARM Id of Resource"),
			},
			ManagedBy: &armelasticsan.ManagedByInfo{
				ResourceID: to.Ptr("mtkeip"),
			},
			SizeGiB: to.Ptr[int64](9),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Create_MinimumSet_Gen.json
func ExampleVolumesClient_BeginCreate_volumesCreateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.Volume{
		Properties: &armelasticsan.VolumeProperties{
			SizeGiB: to.Ptr[int64](9),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Update_MaximumSet_Gen.json
func ExampleVolumesClient_BeginUpdate_volumesUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.VolumeUpdate{
		Properties: &armelasticsan.VolumeUpdateProperties{
			SizeGiB: to.Ptr[int64](11),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Update_MinimumSet_Gen.json
func ExampleVolumesClient_BeginUpdate_volumesUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.VolumeUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Delete_MaximumSet_Gen.json
func ExampleVolumesClient_BeginDelete_volumesDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", &armelasticsan.VolumesClientBeginDeleteOptions{XMSDeleteSnapshots: to.Ptr(armelasticsan.XMSDeleteSnapshotsTrue),
		XMSForceDelete: to.Ptr(armelasticsan.XMSForceDeleteTrue),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Delete_MinimumSet_Gen.json
func ExampleVolumesClient_BeginDelete_volumesDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", &armelasticsan.VolumesClientBeginDeleteOptions{XMSDeleteSnapshots: nil,
		XMSForceDelete: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Get_MaximumSet_Gen.json
func ExampleVolumesClient_Get_volumesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumesClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Get_MinimumSet_Gen.json
func ExampleVolumesClient_Get_volumesGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumesClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_ListByVolumeGroup_MaximumSet_Gen.json
func ExampleVolumesClient_NewListByVolumeGroupPager_volumesListByVolumeGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListByVolumeGroupPager("resourcegroupname", "elasticsanname", "volumegroupname", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VolumeList = armelasticsan.VolumeList{
		// 	Value: []*armelasticsan.Volume{
		// 		{
		// 			Name: to.Ptr("o"),
		// 			Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
		// 			ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
		// 			SystemData: &armelasticsan.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
		// 				CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
		// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("bcclmbseed"),
		// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 			},
		// 			Properties: &armelasticsan.VolumeProperties{
		// 				CreationData: &armelasticsan.SourceCreationData{
		// 					CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
		// 					SourceID: to.Ptr("ARM Id of Resource"),
		// 				},
		// 				ManagedBy: &armelasticsan.ManagedByInfo{
		// 					ResourceID: to.Ptr("mtkeip"),
		// 				},
		// 				ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 				SizeGiB: to.Ptr[int64](9),
		// 				StorageTarget: &armelasticsan.IscsiTargetInfo{
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 					Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
		// 					TargetIqn: to.Ptr("izdwogzjedsfug"),
		// 					TargetPortalHostname: to.Ptr("wyfbjobugmad"),
		// 					TargetPortalPort: to.Ptr[int32](21),
		// 				},
		// 				VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_ListByVolumeGroup_MinimumSet_Gen.json
func ExampleVolumesClient_NewListByVolumeGroupPager_volumesListByVolumeGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListByVolumeGroupPager("resourcegroupname", "elasticsanname", "volumegroupname", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VolumeList = armelasticsan.VolumeList{
		// }
	}
}