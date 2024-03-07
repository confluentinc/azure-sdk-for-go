//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package backup

// BeginFullBackupOptions contains the optional parameters for the Client.BeginFullBackup method.
type BeginFullBackupOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BeginFullRestoreOptions contains the optional parameters for the Client.BeginFullRestore method.
type BeginFullRestoreOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BeginSelectiveKeyRestoreOptions contains the optional parameters for the Client.BeginSelectiveKeyRestore method.
type BeginSelectiveKeyRestoreOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}