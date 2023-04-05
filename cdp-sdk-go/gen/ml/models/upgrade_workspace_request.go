// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpgradeWorkspaceRequest Request object for the UpgradeWorkspace method.
//
// swagger:model UpgradeWorkspaceRequest
type UpgradeWorkspaceRequest struct {

	// The environment of the workspace.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The version of Kubernetes to upgrade to.
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`

	// The version of workspace to upgrade to. (Deprecated).
	MlVersion string `json:"mlVersion,omitempty"`

	// The CRN of the workspace. If CRN is specified only the CRN is used for identifying the workspace, environment and name arguments are ignored.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`

	// The name of the workspace.
	WorkspaceName string `json:"workspaceName,omitempty"`

	// The experimental entitlements to pass to CML for testing purpose. This is not meant for external customers in any case
	XEntitlements []string `json:"xEntitlements"`
}

// Validate validates this upgrade workspace request
func (m *UpgradeWorkspaceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upgrade workspace request based on context it is used
func (m *UpgradeWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res UpgradeWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
