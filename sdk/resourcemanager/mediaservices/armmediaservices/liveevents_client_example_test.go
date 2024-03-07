//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-list-all.json
func ExampleLiveEventsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLiveEventsClient().NewListPager("mediaresources", "slitestmedia10", nil)
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
		// page.LiveEventListResult = armmediaservices.LiveEventListResult{
		// 	Value: []*armmediaservices.LiveEvent{
		// 		{
		// 			Name: to.Ptr("myLiveEvent1"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
		// 			ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmediaservices.LiveEventProperties{
		// 				Description: to.Ptr("test event 1"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.556Z"); return t}()),
		// 				CrossSiteAccessPolicies: &armmediaservices.CrossSiteAccessPolicies{
		// 				},
		// 				Encoding: &armmediaservices.LiveEventEncoding{
		// 					EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
		// 				},
		// 				Input: &armmediaservices.LiveEventInput{
		// 					AccessToken: to.Ptr("<accessToken>"),
		// 					Endpoints: []*armmediaservices.LiveEventEndpoint{
		// 						{
		// 							URL: to.Ptr("http://clouddeployment.media-test.net/de153bb0814542d9b7e2339ce9430dc4/ingest.isml"),
		// 							Protocol: to.Ptr("FragmentedMP4"),
		// 					}},
		// 					KeyFrameIntervalDuration: to.Ptr("PT6S"),
		// 					StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
		// 				},
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.556Z"); return t}()),
		// 				Preview: &armmediaservices.LiveEventPreview{
		// 					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
		// 						IP: &armmediaservices.IPAccessControl{
		// 							Allow: []*armmediaservices.IPRange{
		// 								{
		// 									Name: to.Ptr("AllowAll"),
		// 									Address: to.Ptr("0.0.0.0"),
		// 									SubnetPrefixLength: to.Ptr[int32](0),
		// 							}},
		// 						},
		// 					},
		// 					Endpoints: []*armmediaservices.LiveEventEndpoint{
		// 						{
		// 							URL: to.Ptr("https://myliveevent1-slitestmedia10.preview-usso.channel.mediaservices.windows.net/a220e223-faf8-4e03-b9a9-2c2432f48025/preview.ism/manifest"),
		// 							Protocol: to.Ptr("FragmentedMP4"),
		// 					}},
		// 					PreviewLocator: to.Ptr("a220e223-faf8-4e03-b9a9-2c2432f48025"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateStopped),
		// 				StreamOptions: []*armmediaservices.StreamOptionsFlag{
		// 				},
		// 				UseStaticHostname: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-list-by-name.json
func ExampleLiveEventsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLiveEventsClient().Get(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LiveEvent = armmediaservices.LiveEvent{
	// 	Name: to.Ptr("myLiveEvent1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmediaservices.LiveEventProperties{
	// 		Description: to.Ptr(""),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 		CrossSiteAccessPolicies: &armmediaservices.CrossSiteAccessPolicies{
	// 			ClientAccessPolicy: to.Ptr("<access-policy><cross-domain-access><policy><allow-from http-methods=\"*\"><domain uri=\"http://*\"/></allow-from><grant-to><resource path=\"/\" include-subpaths=\"true\"/></grant-to></policy></cross-domain-access></access-policy>"),
	// 			CrossDomainPolicy: to.Ptr("<cross-domain-policy><allow-access-from domain=\"*\" secure=\"false\" /></cross-domain-policy>"),
	// 		},
	// 		Encoding: &armmediaservices.LiveEventEncoding{
	// 			EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
	// 		},
	// 		Input: &armmediaservices.LiveEventInput{
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 				{
	// 					URL: to.Ptr("http://clouddeployment.media-test.net/ingest.isml"),
	// 					Protocol: to.Ptr("FragmentedMP4"),
	// 			}},
	// 			KeyFrameIntervalDuration: to.Ptr("PT6S"),
	// 			StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 		Preview: &armmediaservices.LiveEventPreview{
	// 			AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowAll"),
	// 							Address: to.Ptr("0.0.0.0"),
	// 							SubnetPrefixLength: to.Ptr[int32](0),
	// 					}},
	// 				},
	// 			},
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 				{
	// 					URL: to.Ptr("https://testeventopito4idh2r-weibzmedia05.preview-ts051.channel.media-test.windows-int.net/763f3ea4-d94f-441c-a634-c833f61a4e04/preview.ism/manifest"),
	// 					Protocol: to.Ptr("FragmentedMP4"),
	// 			}},
	// 			PreviewLocator: to.Ptr("763f3ea4-d94f-441c-a634-c833f61a4e04"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateStopped),
	// 		StreamOptions: []*armmediaservices.StreamOptionsFlag{
	// 			to.Ptr(armmediaservices.StreamOptionsFlagDefault)},
	// 			UseStaticHostname: to.Ptr(false),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-create.json
func ExampleLiveEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginCreate(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", armmediaservices.LiveEvent{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armmediaservices.LiveEventProperties{
			Description: to.Ptr("test event 1"),
			Input: &armmediaservices.LiveEventInput{
				AccessControl: &armmediaservices.LiveEventInputAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:               to.Ptr("AllowAll"),
								Address:            to.Ptr("0.0.0.0"),
								SubnetPrefixLength: to.Ptr[int32](0),
							}},
					},
				},
				KeyFrameIntervalDuration: to.Ptr("PT6S"),
				StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolRTMP),
			},
			Preview: &armmediaservices.LiveEventPreview{
				AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:               to.Ptr("AllowAll"),
								Address:            to.Ptr("0.0.0.0"),
								SubnetPrefixLength: to.Ptr[int32](0),
							}},
					},
				},
			},
		},
	}, &armmediaservices.LiveEventsClientBeginCreateOptions{AutoStart: nil})
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
	// res.LiveEvent = armmediaservices.LiveEvent{
	// 	Name: to.Ptr("myLiveEvent1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmediaservices.LiveEventProperties{
	// 		Description: to.Ptr("test event 1"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:06.098Z"); return t}()),
	// 		CrossSiteAccessPolicies: &armmediaservices.CrossSiteAccessPolicies{
	// 		},
	// 		Encoding: &armmediaservices.LiveEventEncoding{
	// 			EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
	// 		},
	// 		Input: &armmediaservices.LiveEventInput{
	// 			AccessControl: &armmediaservices.LiveEventInputAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowAll"),
	// 							Address: to.Ptr("0.0.0.0"),
	// 							SubnetPrefixLength: to.Ptr[int32](0),
	// 					}},
	// 				},
	// 			},
	// 			AccessToken: to.Ptr("<accessToken>"),
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 			},
	// 			KeyFrameIntervalDuration: to.Ptr("PT6S"),
	// 			StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:06.098Z"); return t}()),
	// 		Preview: &armmediaservices.LiveEventPreview{
	// 			AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowAll"),
	// 							Address: to.Ptr("0.0.0.0"),
	// 							SubnetPrefixLength: to.Ptr[int32](0),
	// 					}},
	// 				},
	// 			},
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 			},
	// 			PreviewLocator: to.Ptr("c91726b4-880c-4090-94aa-e6ddb1384b37"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateStopped),
	// 		StreamOptions: []*armmediaservices.StreamOptionsFlag{
	// 		},
	// 		UseStaticHostname: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-update.json
func ExampleLiveEventsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginUpdate(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", armmediaservices.LiveEvent{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
			"tag3": to.Ptr("value3"),
		},
		Properties: &armmediaservices.LiveEventProperties{
			Description: to.Ptr("test event updated"),
			Input: &armmediaservices.LiveEventInput{
				AccessControl: &armmediaservices.LiveEventInputAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:    to.Ptr("AllowOne"),
								Address: to.Ptr("192.1.1.0"),
							}},
					},
				},
				KeyFrameIntervalDuration: to.Ptr("PT6S"),
				StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
			},
			Preview: &armmediaservices.LiveEventPreview{
				AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:    to.Ptr("AllowOne"),
								Address: to.Ptr("192.1.1.0"),
							}},
					},
				},
			},
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
	// res.LiveEvent = armmediaservices.LiveEvent{
	// 	Name: to.Ptr("myLiveEvent1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 		"tag3": to.Ptr("value3"),
	// 	},
	// 	Properties: &armmediaservices.LiveEventProperties{
	// 		Description: to.Ptr("test event updated"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		Encoding: &armmediaservices.LiveEventEncoding{
	// 			EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
	// 		},
	// 		Input: &armmediaservices.LiveEventInput{
	// 			AccessControl: &armmediaservices.LiveEventInputAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowOne"),
	// 							Address: to.Ptr("192.1.1.0"),
	// 					}},
	// 				},
	// 			},
	// 			AccessToken: to.Ptr("<accessToken>"),
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 			},
	// 			KeyFrameIntervalDuration: to.Ptr("PT6S"),
	// 			StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		Preview: &armmediaservices.LiveEventPreview{
	// 			AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowOne"),
	// 							Address: to.Ptr("192.1.1.0"),
	// 					}},
	// 				},
	// 			},
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 			},
	// 			PreviewLocator: to.Ptr("c10ea3fc-587f-4daf-b2b2-fa8f647a9ed2"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateRunning),
	// 		StreamOptions: []*armmediaservices.StreamOptionsFlag{
	// 		},
	// 		UseStaticHostname: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-delete.json
func ExampleLiveEventsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginDelete(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-allocate.json
func ExampleLiveEventsClient_BeginAllocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginAllocate(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-start.json
func ExampleLiveEventsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginStart(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-stop.json
func ExampleLiveEventsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginStop(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", armmediaservices.LiveEventActionInput{
		RemoveOutputsOnStop: to.Ptr(false),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-reset.json
func ExampleLiveEventsClient_BeginReset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginReset(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/async-operation-result.json
func ExampleLiveEventsClient_AsyncOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLiveEventsClient().AsyncOperation(ctx, "mediaresources", "slitestmedia10", "62e4d893-d233-4005-988e-a428d9f77076", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AsyncOperationResult = armmediaservices.AsyncOperationResult{
	// 	Name: to.Ptr("62e4d893-d233-4005-988e-a428d9f77076"),
	// 	Error: &armmediaservices.ErrorDetail{
	// 		Code: to.Ptr("None"),
	// 		Target: to.Ptr("d7aea790-8acb-40b9-8f9f-21cc37c9e519"),
	// 	},
	// 	Status: to.Ptr(armmediaservices.AsyncOperationStatusInProgress),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-operation-location.json
func ExampleLiveEventsClient_OperationLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLiveEventsClient().OperationLocation(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", "62e4d893-d233-4005-988e-a428d9f77076", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LiveEvent = armmediaservices.LiveEvent{
	// 	Name: to.Ptr("myLiveEvent1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmediaservices.LiveEventProperties{
	// 		Description: to.Ptr(""),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 		CrossSiteAccessPolicies: &armmediaservices.CrossSiteAccessPolicies{
	// 			ClientAccessPolicy: to.Ptr("<access-policy><cross-domain-access><policy><allow-from http-methods=\"*\"><domain uri=\"http://*\"/></allow-from><grant-to><resource path=\"/\" include-subpaths=\"true\"/></grant-to></policy></cross-domain-access></access-policy>"),
	// 			CrossDomainPolicy: to.Ptr("<cross-domain-policy><allow-access-from domain=\"*\" secure=\"false\" /></cross-domain-policy>"),
	// 		},
	// 		Encoding: &armmediaservices.LiveEventEncoding{
	// 			EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
	// 		},
	// 		Input: &armmediaservices.LiveEventInput{
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 				{
	// 					URL: to.Ptr("http://clouddeployment.media-test.net/ingest.isml"),
	// 					Protocol: to.Ptr("FragmentedMP4"),
	// 			}},
	// 			KeyFrameIntervalDuration: to.Ptr("PT6S"),
	// 			StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 		Preview: &armmediaservices.LiveEventPreview{
	// 			AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowAll"),
	// 							Address: to.Ptr("0.0.0.0"),
	// 							SubnetPrefixLength: to.Ptr[int32](0),
	// 					}},
	// 				},
	// 			},
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 				{
	// 					URL: to.Ptr("https://testeventopito4idh2r-weibzmedia05.preview-ts051.channel.media-test.windows-int.net/763f3ea4-d94f-441c-a634-c833f61a4e04/preview.ism/manifest"),
	// 					Protocol: to.Ptr("FragmentedMP4"),
	// 			}},
	// 			PreviewLocator: to.Ptr("763f3ea4-d94f-441c-a634-c833f61a4e04"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateStopped),
	// 		StreamOptions: []*armmediaservices.StreamOptionsFlag{
	// 			to.Ptr(armmediaservices.StreamOptionsFlagDefault)},
	// 			UseStaticHostname: to.Ptr(false),
	// 		},
	// 		SystemData: &armmediaservices.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 			CreatedBy: to.Ptr("example@microsoft.com"),
	// 			CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("example@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 		},
	// 	}
}