//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_ListByFactory.json
func ExampleDatasetsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatasetsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.DatasetListResponse = armdatafactory.DatasetListResponse{
		// 	Value: []*armdatafactory.DatasetResource{
		// 		{
		// 			Name: to.Ptr("exampleDataset"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/datasets"),
		// 			Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
		// 			Properties: &armdatafactory.AzureBlobDataset{
		// 				Type: to.Ptr("AzureBlob"),
		// 				Description: to.Ptr("Example description"),
		// 				LinkedServiceName: &armdatafactory.LinkedServiceReference{
		// 					Type: to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
		// 					ReferenceName: to.Ptr("exampleLinkedService"),
		// 				},
		// 				Parameters: map[string]*armdatafactory.ParameterSpecification{
		// 					"MyFileName": &armdatafactory.ParameterSpecification{
		// 						Type: to.Ptr(armdatafactory.ParameterTypeString),
		// 					},
		// 					"MyFolderPath": &armdatafactory.ParameterSpecification{
		// 						Type: to.Ptr(armdatafactory.ParameterTypeString),
		// 					},
		// 				},
		// 				TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
		// 					Format: &armdatafactory.TextFormat{
		// 						Type: to.Ptr("TextFormat"),
		// 					},
		// 					FileName: map[string]any{
		// 						"type": "Expression",
		// 						"value": "@dataset().MyFileName",
		// 					},
		// 					FolderPath: map[string]any{
		// 						"type": "Expression",
		// 						"value": "@dataset().MyFolderPath",
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Create.json
func ExampleDatasetsClient_CreateOrUpdate_datasetsCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatasetsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", armdatafactory.DatasetResource{
		Properties: &armdatafactory.AzureBlobDataset{
			Type: to.Ptr("AzureBlob"),
			LinkedServiceName: &armdatafactory.LinkedServiceReference{
				Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
				ReferenceName: to.Ptr("exampleLinkedService"),
			},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"MyFileName": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
				"MyFolderPath": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
			},
			TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
				Format: &armdatafactory.TextFormat{
					Type: to.Ptr("TextFormat"),
				},
				FileName: map[string]any{
					"type":  "Expression",
					"value": "@dataset().MyFileName",
				},
				FolderPath: map[string]any{
					"type":  "Expression",
					"value": "@dataset().MyFolderPath",
				},
			},
		},
	}, &armdatafactory.DatasetsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatasetResource = armdatafactory.DatasetResource{
	// 	Name: to.Ptr("exampleDataset"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/datasets"),
	// 	Etag: to.Ptr("0a0066d4-0000-0000-0000-5b245bd20000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.AzureBlobDataset{
	// 		Type: to.Ptr("AzureBlob"),
	// 		Schema: []any{
	// 			map[string]any{
	// 				"name": "col1",
	// 				"type": "INT_32",
	// 			},
	// 			map[string]any{
	// 				"name": "col2",
	// 				"type": "Decimal",
	// 				"precision": "38",
	// 				"scale": "2",
	// 			},
	// 		},
	// 		LinkedServiceName: &armdatafactory.LinkedServiceReference{
	// 			Type: to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
	// 			ReferenceName: to.Ptr("exampleLinkedService"),
	// 		},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"MyFileName": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 			"MyFolderPath": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 		},
	// 		TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
	// 			Format: &armdatafactory.TextFormat{
	// 				Type: to.Ptr("TextFormat"),
	// 			},
	// 			FileName: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFileName",
	// 			},
	// 			FolderPath: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFolderPath",
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Update.json
func ExampleDatasetsClient_CreateOrUpdate_datasetsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatasetsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", armdatafactory.DatasetResource{
		Properties: &armdatafactory.AzureBlobDataset{
			Type:        to.Ptr("AzureBlob"),
			Description: to.Ptr("Example description"),
			LinkedServiceName: &armdatafactory.LinkedServiceReference{
				Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
				ReferenceName: to.Ptr("exampleLinkedService"),
			},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"MyFileName": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
				"MyFolderPath": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
			},
			TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
				Format: &armdatafactory.TextFormat{
					Type: to.Ptr("TextFormat"),
				},
				FileName: map[string]any{
					"type":  "Expression",
					"value": "@dataset().MyFileName",
				},
				FolderPath: map[string]any{
					"type":  "Expression",
					"value": "@dataset().MyFolderPath",
				},
			},
		},
	}, &armdatafactory.DatasetsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatasetResource = armdatafactory.DatasetResource{
	// 	Name: to.Ptr("exampleDataset"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/datasets"),
	// 	Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.AzureBlobDataset{
	// 		Type: to.Ptr("AzureBlob"),
	// 		Description: to.Ptr("Example description"),
	// 		LinkedServiceName: &armdatafactory.LinkedServiceReference{
	// 			Type: to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
	// 			ReferenceName: to.Ptr("exampleLinkedService"),
	// 		},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"MyFileName": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 			"MyFolderPath": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 		},
	// 		TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
	// 			Format: &armdatafactory.TextFormat{
	// 				Type: to.Ptr("TextFormat"),
	// 			},
	// 			FileName: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFileName",
	// 			},
	// 			FolderPath: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFolderPath",
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Get.json
func ExampleDatasetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatasetsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", &armdatafactory.DatasetsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatasetResource = armdatafactory.DatasetResource{
	// 	Name: to.Ptr("exampleDataset"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/datasets"),
	// 	Etag: to.Ptr("15004c4f-0000-0200-0000-5cbe090e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.AzureBlobDataset{
	// 		Type: to.Ptr("AzureBlob"),
	// 		Description: to.Ptr("Example description"),
	// 		LinkedServiceName: &armdatafactory.LinkedServiceReference{
	// 			Type: to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
	// 			ReferenceName: to.Ptr("exampleLinkedService"),
	// 		},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"MyFileName": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 			"MyFolderPath": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 		},
	// 		TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
	// 			Format: &armdatafactory.TextFormat{
	// 				Type: to.Ptr("TextFormat"),
	// 			},
	// 			FileName: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFileName",
	// 			},
	// 			FolderPath: map[string]any{
	// 				"type": "Expression",
	// 				"value": "@dataset().MyFolderPath",
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Delete.json
func ExampleDatasetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDatasetsClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
