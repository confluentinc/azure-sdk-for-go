//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbillingbenefits

import "time"

// AppliedScopeProperties - Properties specific to applied scope type. Not required if not applicable.
type AppliedScopeProperties struct {
	// Display name
	DisplayName *string

	// Fully-qualified identifier of the management group where the benefit must be applied.
	ManagementGroupID *string

	// Fully-qualified identifier of the resource group.
	ResourceGroupID *string

	// Fully-qualified identifier of the subscription.
	SubscriptionID *string

	// Tenant ID where the benefit is applied.
	TenantID *string
}

// BillingPlanInformation - Information describing the type of billing plan for this savings plan.
type BillingPlanInformation struct {
	// For recurring billing plans, indicates the date when next payment will be processed. Null when total is paid off.
	NextPaymentDueDate *time.Time

	// Amount of money to be paid for the Order. Tax is not included.
	PricingCurrencyTotal *Price

	// Date when the billing plan has started.
	StartDate    *time.Time
	Transactions []*PaymentDetail
}

// Commitment towards the benefit.
type Commitment struct {
	Amount *float64

	// The ISO 4217 3-letter currency code for the currency used by this purchase record.
	CurrencyCode *string

	// Commitment grain.
	Grain *CommitmentGrain
}

type ExtendedStatusInfo struct {
	// The message giving detailed information about the status code.
	Message *string

	// Status code providing additional information.
	StatusCode *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// PaymentDetail - Information about payment related to a savings plan order.
type PaymentDetail struct {
	// Billing account
	BillingAccount *string

	// Amount charged in Billing currency. Tax not included. Is null for future payments
	BillingCurrencyTotal *Price

	// Date when the payment needs to be done.
	DueDate *time.Time

	// Date when the transaction is completed. Is null when it is scheduled.
	PaymentDate *time.Time

	// Amount in pricing currency. Tax not included.
	PricingCurrencyTotal *Price

	// Describes whether the payment is completed, failed, cancelled or scheduled in the future.
	Status *PaymentStatus

	// READ-ONLY
	ExtendedStatusInfo *ExtendedStatusInfo
}

type Price struct {
	Amount *float64

	// The ISO 4217 3-letter currency code for the currency used by this purchase record.
	CurrencyCode *string
}

type PurchaseRequest struct {
	Properties *PurchaseRequestProperties

	// The SKU to be applied for this resource
	SKU *SKU
}

type PurchaseRequestProperties struct {
	// Properties specific to applied scope type. Not required if not applicable.
	AppliedScopeProperties *AppliedScopeProperties

	// Type of the Applied Scope.
	AppliedScopeType *AppliedScopeType

	// Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
	BillingPlan *BillingPlan

	// Subscription that will be charged for purchasing the benefit
	BillingScopeID *string

	// Commitment towards the benefit.
	Commitment *Commitment

	// Friendly name of the savings plan
	DisplayName *string

	// Setting this to true will automatically purchase a new benefit on the expiration date time.
	Renew *bool

	// Represent benefit term in ISO 8601 format.
	Term *Term

	// READ-ONLY; DateTime of the savings plan starts providing benefit from.
	EffectiveDateTime *time.Time
}

type RenewProperties struct {
	PurchaseProperties *PurchaseRequest
}

// ReservationOrderAliasRequest - Reservation order alias
type ReservationOrderAliasRequest struct {
	// REQUIRED; Reservation order SKU
	SKU *SKU

	// The Azure Region where the reservation benefits are applied to.
	Location *string

	// Reservation order alias request properties
	Properties *ReservationOrderAliasRequestProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ReservationOrderAliasRequestProperties - Reservation properties
type ReservationOrderAliasRequestProperties struct {
	// Properties specific to applied scope type. Not required if not applicable.
	AppliedScopeProperties *AppliedScopeProperties

	// Type of the Applied Scope.
	AppliedScopeType *AppliedScopeType

	// Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
	BillingPlan *BillingPlan

	// Subscription that will be charged for purchasing the benefit
	BillingScopeID *string

	// Display name
	DisplayName *string

	// Total Quantity of the SKUs purchased in the Reservation.
	Quantity *int32

	// Setting this to true will automatically purchase a new benefit on the expiration date time.
	Renew *bool

	// Properties specific to each reserved resource type. Not required if not applicable.
	ReservedResourceProperties *ReservationOrderAliasRequestPropertiesReservedResourceProperties

	// The type of the resource that is being reserved.
	ReservedResourceType *ReservedResourceType

	// This is the date-time when the Azure Hybrid Benefit needs to be reviewed.
	ReviewDateTime *time.Time

	// Represent benefit term in ISO 8601 format.
	Term *Term
}

// ReservationOrderAliasRequestPropertiesReservedResourceProperties - Properties specific to each reserved resource type.
// Not required if not applicable.
type ReservationOrderAliasRequestPropertiesReservedResourceProperties struct {
	// Turning this on will apply the reservation discount to other VMs in the same VM size group.
	InstanceFlexibility *InstanceFlexibility
}

// ReservationOrderAliasResponse - Reservation order alias
type ReservationOrderAliasResponse struct {
	// REQUIRED; Reservation order SKU
	SKU *SKU

	// The Azure Region where the reserved resource lives.
	Location *string

	// Reservation order alias response properties
	Properties *ReservationOrderAliasResponseProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ReservationOrderAliasResponseProperties - Reservation properties
type ReservationOrderAliasResponseProperties struct {
	// Properties specific to applied scope type. Not required if not applicable.
	AppliedScopeProperties *AppliedScopeProperties

	// Type of the Applied Scope.
	AppliedScopeType *AppliedScopeType

	// Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
	BillingPlan *BillingPlan

	// Subscription that will be charged for purchasing the benefit
	BillingScopeID *string

	// Display name
	DisplayName *string

	// Total Quantity of the SKUs purchased in the Reservation.
	Quantity *int32

	// Setting this to true will automatically purchase a new benefit on the expiration date time.
	Renew *bool

	// Properties specific to each reserved resource type. Not required if not applicable.
	ReservedResourceProperties *ReservationOrderAliasResponsePropertiesReservedResourceProperties

	// The type of the resource that is being reserved.
	ReservedResourceType *ReservedResourceType

	// This is the date-time when the Reservation needs to be reviewed.
	ReviewDateTime *time.Time

	// Represent benefit term in ISO 8601 format.
	Term *Term

	// READ-ONLY; Provisioning state
	ProvisioningState *ProvisioningState

	// READ-ONLY; Identifier of the reservation order created
	ReservationOrderID *string
}

// ReservationOrderAliasResponsePropertiesReservedResourceProperties - Properties specific to each reserved resource type.
// Not required if not applicable.
type ReservationOrderAliasResponsePropertiesReservedResourceProperties struct {
	// Turning this on will apply the reservation discount to other VMs in the same VM size group.
	InstanceFlexibility *InstanceFlexibility
}

// RoleAssignmentEntity - Role assignment entity
type RoleAssignmentEntity struct {
	// Role assignment entity id
	ID *string

	// Role assignment entity name
	Name *string

	// Role assignment entity properties
	Properties *RoleAssignmentEntityProperties
}

// RoleAssignmentEntityProperties - Role assignment entity properties
type RoleAssignmentEntityProperties struct {
	// Principal Id
	PrincipalID *string

	// Role definition id
	RoleDefinitionID *string

	// Scope of the role assignment entity
	Scope *string
}

// SKU - The SKU to be applied for this resource
type SKU struct {
	// Name of the SKU to be applied
	Name *string
}

// SavingsPlanModel - Savings plan
type SavingsPlanModel struct {
	// REQUIRED; Savings plan SKU
	SKU *SKU

	// Savings plan properties
	Properties *SavingsPlanModelProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type SavingsPlanModelList struct {
	// Url to get the next page.
	NextLink *string
	Value    []*SavingsPlanModel
}

type SavingsPlanModelListResult struct {
	// READ-ONLY; The roll out count summary of the savings plans
	AdditionalProperties []*SavingsPlanSummary

	// READ-ONLY; Url to get the next page.
	NextLink *string

	// READ-ONLY; The list of savings plans.
	Value []*SavingsPlanModel
}

// SavingsPlanModelProperties - Savings plan properties
type SavingsPlanModelProperties struct {
	// Properties specific to applied scope type. Not required if not applicable.
	AppliedScopeProperties *AppliedScopeProperties

	// Type of the Applied Scope.
	AppliedScopeType *AppliedScopeType

	// This is the DateTime when the savings plan benefit started.
	BenefitStartTime *time.Time

	// Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
	BillingPlan *BillingPlan

	// Subscription that will be charged for purchasing the benefit
	BillingScopeID *string

	// Commitment towards the benefit.
	Commitment *Commitment

	// Display name
	DisplayName *string

	// Setting this to true will automatically purchase a new benefit on the expiration date time.
	Renew *bool

	// SavingsPlan Id of the SavingsPlan which is purchased because of renew.
	RenewDestination *string
	RenewProperties  *RenewProperties

	// SavingsPlan Id of the SavingsPlan from which this SavingsPlan is renewed.
	RenewSource *string

	// Represent benefit term in ISO 8601 format.
	Term *Term

	// READ-ONLY; Fully-qualified identifier of the billing account where the savings plan is applied. Present only for Enterprise
	// Agreement customers.
	BillingAccountID *string

	// READ-ONLY; Fully-qualified identifier of the billing profile where the savings plan is applied. Present only for Field-led
	// or Customer-led customers.
	BillingProfileID *string

	// READ-ONLY; Fully-qualified identifier of the customer where the savings plan is applied. Present only for Partner-led customers.
	CustomerID *string

	// READ-ONLY; The provisioning state of the savings plan for display, e.g. Succeeded
	DisplayProvisioningState *string

	// READ-ONLY; DateTime of the savings plan starts providing benefit from.
	EffectiveDateTime *time.Time

	// READ-ONLY; Expiry date time
	ExpiryDateTime *time.Time

	// READ-ONLY
	ExtendedStatusInfo *ExtendedStatusInfo

	// READ-ONLY; Provisioning state
	ProvisioningState *ProvisioningState

	// READ-ONLY; Date time when the savings plan was purchased
	PurchaseDateTime *time.Time

	// READ-ONLY; The applied scope type of the savings plan for display, e.g. Shared
	UserFriendlyAppliedScopeType *string

	// READ-ONLY; Savings plan utilization
	Utilization *Utilization
}

// SavingsPlanOrderAliasModel - Savings plan order alias
type SavingsPlanOrderAliasModel struct {
	// REQUIRED; Savings plan SKU
	SKU *SKU

	// Resource provider kind
	Kind *string

	// Savings plan order alias properties
	Properties *SavingsPlanOrderAliasProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SavingsPlanOrderAliasProperties - Savings plan properties
type SavingsPlanOrderAliasProperties struct {
	// Properties specific to applied scope type. Not required if not applicable.
	AppliedScopeProperties *AppliedScopeProperties

	// Type of the Applied Scope.
	AppliedScopeType *AppliedScopeType

	// Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
	BillingPlan *BillingPlan

	// Subscription that will be charged for purchasing the benefit
	BillingScopeID *string

	// Commitment towards the benefit.
	Commitment *Commitment

	// Display name
	DisplayName *string

	// Represent benefit term in ISO 8601 format.
	Term *Term

	// READ-ONLY; Provisioning state
	ProvisioningState *ProvisioningState

	// READ-ONLY; Identifier of the savings plan created
	SavingsPlanOrderID *string
}

// SavingsPlanOrderModel - Savings plan order
type SavingsPlanOrderModel struct {
	// REQUIRED; Savings plan SKU
	SKU *SKU

	// Savings plan order properties
	Properties *SavingsPlanOrderModelProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type SavingsPlanOrderModelList struct {
	// Url to get the next page.
	NextLink *string
	Value    []*SavingsPlanOrderModel
}

// SavingsPlanOrderModelProperties - Savings plan order properties
type SavingsPlanOrderModelProperties struct {
	// This is the DateTime when the savings plan benefit started.
	BenefitStartTime *time.Time

	// Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
	BillingPlan *BillingPlan

	// Subscription that will be charged for purchasing the benefit
	BillingScopeID *string

	// Display name
	DisplayName *string

	// Information describing the type of billing plan for this savings plan.
	PlanInformation *BillingPlanInformation
	SavingsPlans    []*string

	// Represent benefit term in ISO 8601 format.
	Term *Term

	// READ-ONLY; Fully-qualified identifier of the billing account where the savings plan is applied. Present only for Enterprise
	// Agreement customers.
	BillingAccountID *string

	// READ-ONLY; Fully-qualified identifier of the billing profile where the savings plan is applied. Present only for Field-led
	// or Customer-led customers.
	BillingProfileID *string

	// READ-ONLY; Fully-qualified identifier of the customer where the savings plan is applied. Present only for Partner-led customers.
	CustomerID *string

	// READ-ONLY; Expiry date time
	ExpiryDateTime *time.Time

	// READ-ONLY
	ExtendedStatusInfo *ExtendedStatusInfo

	// READ-ONLY; Provisioning state
	ProvisioningState *ProvisioningState
}

type SavingsPlanPurchaseValidateRequest struct {
	Benefits []*SavingsPlanOrderAliasModel
}

// SavingsPlanSummary - Savings plans list summary
type SavingsPlanSummary struct {
	// The roll up count summary of savings plans in each state
	Value *SavingsPlanSummaryCount

	// READ-ONLY; This property has value 'summary'
	Name *string
}

// SavingsPlanSummaryCount - The roll up count summary of savings plans in each state
type SavingsPlanSummaryCount struct {
	// READ-ONLY; The number of savings plans in Cancelled state
	CancelledCount *float32

	// READ-ONLY; The number of savings plans in Expired state
	ExpiredCount *float32

	// READ-ONLY; The number of savings plans in Expiring state
	ExpiringCount *float32

	// READ-ONLY; The number of savings plans in Failed state
	FailedCount *float32

	// READ-ONLY; The number of savings plans in No Benefit state
	NoBenefitCount *float32

	// READ-ONLY; The number of savings plans in Pending state
	PendingCount *float32

	// READ-ONLY; The number of savings plans in Processing state
	ProcessingCount *float32

	// READ-ONLY; The number of savings plans in Succeeded state
	SucceededCount *float32

	// READ-ONLY; The number of savings plans in Warning state
	WarningCount *float32
}

// SavingsPlanUpdateRequest - Savings plan patch request
type SavingsPlanUpdateRequest struct {
	// Savings plan patch request
	Properties *SavingsPlanUpdateRequestProperties
}

// SavingsPlanUpdateRequestProperties - Savings plan patch request
type SavingsPlanUpdateRequestProperties struct {
	// Properties specific to applied scope type. Not required if not applicable.
	AppliedScopeProperties *AppliedScopeProperties

	// Type of the Applied Scope.
	AppliedScopeType *AppliedScopeType

	// Display name
	DisplayName *string

	// Setting this to true will automatically purchase a new benefit on the expiration date time.
	Renew           *bool
	RenewProperties *RenewProperties
}

type SavingsPlanUpdateValidateRequest struct {
	Benefits []*SavingsPlanUpdateRequestProperties
}

// SavingsPlanValidResponseProperty - Benefit scope response property
type SavingsPlanValidResponseProperty struct {
	// Failure reason if the provided input was invalid
	Reason *string

	// Failure reason code if the provided input was invalid
	ReasonCode *string

	// Indicates if the provided input was valid
	Valid *bool
}

type SavingsPlanValidateResponse struct {
	Benefits []*SavingsPlanValidResponseProperty

	// Url to get the next page.
	NextLink *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// Utilization - Savings plan utilization
type Utilization struct {
	// The array of aggregates of a savings plan's utilization
	Aggregates []*UtilizationAggregates

	// READ-ONLY; The number of days trend for a savings plan
	Trend *string
}

// UtilizationAggregates - The aggregate values of savings plan utilization
type UtilizationAggregates struct {
	// READ-ONLY; The grain of the aggregate
	Grain *float32

	// READ-ONLY; The grain unit of the aggregate
	GrainUnit *string

	// READ-ONLY; The aggregate value
	Value *float32

	// READ-ONLY; The aggregate value unit
	ValueUnit *string
}