//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// Result of the request to list Marketplace operations. It contains a list of operations and a URL link to get the next set
	// of results.
	OperationListResult
}

// PrivateStoreClientAcknowledgeOfferNotificationResponse contains the response from method PrivateStoreClient.AcknowledgeOfferNotification.
type PrivateStoreClientAcknowledgeOfferNotificationResponse struct {
	// placeholder for future response values
}

// PrivateStoreClientAdminRequestApprovalsListResponse contains the response from method PrivateStoreClient.AdminRequestApprovalsList.
type PrivateStoreClientAdminRequestApprovalsListResponse struct {
	// List of admin request approval resources
	AdminRequestApprovalsList
}

// PrivateStoreClientBillingAccountsResponse contains the response from method PrivateStoreClient.BillingAccounts.
type PrivateStoreClientBillingAccountsResponse struct {
	// Billing accounts response object
	BillingAccountsResponse
}

// PrivateStoreClientBulkCollectionsActionResponse contains the response from method PrivateStoreClient.BulkCollectionsAction.
type PrivateStoreClientBulkCollectionsActionResponse struct {
	// The bulk collections response. The response contains two lists that indicate for each collection whether the operation
	// succeeded or failed
	BulkCollectionsResponse
}

// PrivateStoreClientCollectionsToSubscriptionsMappingResponse contains the response from method PrivateStoreClient.CollectionsToSubscriptionsMapping.
type PrivateStoreClientCollectionsToSubscriptionsMappingResponse struct {
	// A map of collections subscriptions details
	CollectionsToSubscriptionsMappingResponse
}

// PrivateStoreClientCreateApprovalRequestResponse contains the response from method PrivateStoreClient.CreateApprovalRequest.
type PrivateStoreClientCreateApprovalRequestResponse struct {
	// Request approval resource.
	RequestApprovalResource
}

// PrivateStoreClientCreateOrUpdateResponse contains the response from method PrivateStoreClient.CreateOrUpdate.
type PrivateStoreClientCreateOrUpdateResponse struct {
	// placeholder for future response values
}

// PrivateStoreClientDeleteResponse contains the response from method PrivateStoreClient.Delete.
type PrivateStoreClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateStoreClientFetchAllSubscriptionsInTenantResponse contains the response from method PrivateStoreClient.FetchAllSubscriptionsInTenant.
type PrivateStoreClientFetchAllSubscriptionsInTenantResponse struct {
	// Subscription list operation response.
	SubscriptionsResponse
}

// PrivateStoreClientGetAdminRequestApprovalResponse contains the response from method PrivateStoreClient.GetAdminRequestApproval.
type PrivateStoreClientGetAdminRequestApprovalResponse struct {
	// Admin request approval resource.
	AdminRequestApprovalsResource
}

// PrivateStoreClientGetApprovalRequestsListResponse contains the response from method PrivateStoreClient.GetApprovalRequestsList.
type PrivateStoreClientGetApprovalRequestsListResponse struct {
	// List of admin request approval resources
	RequestApprovalsList
}

// PrivateStoreClientGetRequestApprovalResponse contains the response from method PrivateStoreClient.GetRequestApproval.
type PrivateStoreClientGetRequestApprovalResponse struct {
	// Request approval resource.
	RequestApprovalResource
}

// PrivateStoreClientGetResponse contains the response from method PrivateStoreClient.Get.
type PrivateStoreClientGetResponse struct {
	// The PrivateStore data structure.
	PrivateStore
}

// PrivateStoreClientListNewPlansNotificationsResponse contains the response from method PrivateStoreClient.ListNewPlansNotifications.
type PrivateStoreClientListNewPlansNotificationsResponse struct {
	// List of all new plans notifications for public offers
	NewPlansNotificationsList
}

// PrivateStoreClientListResponse contains the response from method PrivateStoreClient.NewListPager.
type PrivateStoreClientListResponse struct {
	// Describes the json payload for the list of available private stores (between zero and one, inclusive)
	PrivateStoreList
}

// PrivateStoreClientListStopSellOffersPlansNotificationsResponse contains the response from method PrivateStoreClient.ListStopSellOffersPlansNotifications.
type PrivateStoreClientListStopSellOffersPlansNotificationsResponse struct {
	// List of stop sell offers and plans notifications.
	StopSellOffersPlansNotificationsList
}

// PrivateStoreClientListSubscriptionsContextResponse contains the response from method PrivateStoreClient.ListSubscriptionsContext.
type PrivateStoreClientListSubscriptionsContextResponse struct {
	// List of subscription Ids in the private store
	SubscriptionsContextList
}

// PrivateStoreClientQueryApprovedPlansResponse contains the response from method PrivateStoreClient.QueryApprovedPlans.
type PrivateStoreClientQueryApprovedPlansResponse struct {
	// Query approved plans response
	QueryApprovedPlansResponse
}

// PrivateStoreClientQueryNotificationsStateResponse contains the response from method PrivateStoreClient.QueryNotificationsState.
type PrivateStoreClientQueryNotificationsStateResponse struct {
	// Get private store notifications state
	PrivateStoreNotificationsState
}

// PrivateStoreClientQueryOffersResponse contains the response from method PrivateStoreClient.QueryOffers.
type PrivateStoreClientQueryOffersResponse struct {
	// List of offers
	QueryOffers
}

// PrivateStoreClientQueryRequestApprovalResponse contains the response from method PrivateStoreClient.QueryRequestApproval.
type PrivateStoreClientQueryRequestApprovalResponse struct {
	// Gets the request plans with indication on each plan whether is approved by the admin, has pending request or not requested
	// yet
	QueryRequestApproval
}

// PrivateStoreClientUpdateAdminRequestApprovalResponse contains the response from method PrivateStoreClient.UpdateAdminRequestApproval.
type PrivateStoreClientUpdateAdminRequestApprovalResponse struct {
	// Admin request approval resource.
	AdminRequestApprovalsResource
}

// PrivateStoreClientWithdrawPlanResponse contains the response from method PrivateStoreClient.WithdrawPlan.
type PrivateStoreClientWithdrawPlanResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionClientCreateOrUpdateResponse contains the response from method PrivateStoreCollectionClient.CreateOrUpdate.
type PrivateStoreCollectionClientCreateOrUpdateResponse struct {
	// The Collection data structure.
	Collection
}

// PrivateStoreCollectionClientDeleteResponse contains the response from method PrivateStoreCollectionClient.Delete.
type PrivateStoreCollectionClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionClientGetResponse contains the response from method PrivateStoreCollectionClient.Get.
type PrivateStoreCollectionClientGetResponse struct {
	// The Collection data structure.
	Collection
}

// PrivateStoreCollectionClientListResponse contains the response from method PrivateStoreCollectionClient.List.
type PrivateStoreCollectionClientListResponse struct {
	CollectionsList
}

// PrivateStoreCollectionClientPostResponse contains the response from method PrivateStoreCollectionClient.Post.
type PrivateStoreCollectionClientPostResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionClientTransferOffersResponse contains the response from method PrivateStoreCollectionClient.TransferOffers.
type PrivateStoreCollectionClientTransferOffersResponse struct {
	// The transfer items response. The response contains two lists that indicate for each collection whether the operation succeeded
	// or failed
	TransferOffersResponse
}

// PrivateStoreCollectionOfferClientCreateOrUpdateResponse contains the response from method PrivateStoreCollectionOfferClient.CreateOrUpdate.
type PrivateStoreCollectionOfferClientCreateOrUpdateResponse struct {
	// The privateStore offer data structure.
	Offer
}

// PrivateStoreCollectionOfferClientDeleteResponse contains the response from method PrivateStoreCollectionOfferClient.Delete.
type PrivateStoreCollectionOfferClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionOfferClientGetResponse contains the response from method PrivateStoreCollectionOfferClient.Get.
type PrivateStoreCollectionOfferClientGetResponse struct {
	// The privateStore offer data structure.
	Offer
}

// PrivateStoreCollectionOfferClientListResponse contains the response from method PrivateStoreCollectionOfferClient.NewListPager.
type PrivateStoreCollectionOfferClientListResponse struct {
	OfferListResponse
}

// PrivateStoreCollectionOfferClientPostResponse contains the response from method PrivateStoreCollectionOfferClient.Post.
type PrivateStoreCollectionOfferClientPostResponse struct {
	// placeholder for future response values
}