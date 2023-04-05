// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateModelRegistryRequest Request object for creating model registry.
//
// swagger:model CreateModelRegistryRequest
type CreateModelRegistryRequest struct {

	// The create workspace request on which model registry helm chart is deployed.
	CreateWorkspacePayload *CreateWorkspaceRequest `json:"createWorkspacePayload,omitempty"`

	// The creator of model registry.
	CreatorCrn string `json:"creatorCrn,omitempty"`

	// The environment CRN of model registry.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The environment name.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The namespace of the model registry.
	Namespace string `json:"namespace,omitempty"`

	// The S3 access key of Ozone.
	S3AccessKey string `json:"s3AccessKey,omitempty"`

	// The s3Bucket of Ozone.
	S3Bucket string `json:"s3Bucket,omitempty"`

	// The endpoint of Ozone.
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// The S3 secret key of Ozone.
	S3SecretKey string `json:"s3SecretKey,omitempty"`
}

// Validate validates this create model registry request
func (m *CreateModelRegistryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateWorkspacePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateModelRegistryRequest) validateCreateWorkspacePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateWorkspacePayload) { // not required
		return nil
	}

	if m.CreateWorkspacePayload != nil {
		if err := m.CreateWorkspacePayload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createWorkspacePayload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createWorkspacePayload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create model registry request based on the context it is used
func (m *CreateModelRegistryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreateWorkspacePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateModelRegistryRequest) contextValidateCreateWorkspacePayload(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateWorkspacePayload != nil {
		if err := m.CreateWorkspacePayload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createWorkspacePayload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createWorkspacePayload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateModelRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateModelRegistryRequest) UnmarshalBinary(b []byte) error {
	var res CreateModelRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
