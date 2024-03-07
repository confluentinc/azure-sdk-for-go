//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package rbac

// CreateOrUpdateRoleDefinitionResponse contains the response from method Client.CreateOrUpdateRoleDefinition.
type CreateOrUpdateRoleDefinitionResponse struct {
	// Role definition.
	RoleDefinition
}

// CreateRoleAssignmentResponse contains the response from method Client.CreateRoleAssignment.
type CreateRoleAssignmentResponse struct {
	// Role Assignments
	RoleAssignment
}

// DeleteRoleAssignmentResponse contains the response from method Client.DeleteRoleAssignment.
type DeleteRoleAssignmentResponse struct {
	// Role Assignments
	RoleAssignment
}

// DeleteRoleDefinitionResponse contains the response from method Client.DeleteRoleDefinition.
type DeleteRoleDefinitionResponse struct {
	// Role definition.
	RoleDefinition
}

// GetRoleAssignmentResponse contains the response from method Client.GetRoleAssignment.
type GetRoleAssignmentResponse struct {
	// Role Assignments
	RoleAssignment
}

// GetRoleDefinitionResponse contains the response from method Client.GetRoleDefinition.
type GetRoleDefinitionResponse struct {
	// Role definition.
	RoleDefinition
}

// ListRoleAssignmentsResponse contains the response from method Client.NewListRoleAssignmentsPager.
type ListRoleAssignmentsResponse struct {
	// Role assignment list operation result.
	RoleAssignmentListResult
}

// ListRoleDefinitionsResponse contains the response from method Client.NewListRoleDefinitionsPager.
type ListRoleDefinitionsResponse struct {
	// Role definition list operation result.
	RoleDefinitionListResult
}