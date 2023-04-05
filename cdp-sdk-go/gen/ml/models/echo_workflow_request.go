// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EchoWorkflowRequest Request object for EchoWorkflow.
//
// swagger:model EchoWorkflowRequest
type EchoWorkflowRequest struct {

	// The message payload.
	// Required: true
	Message *string `json:"message"`

	// The unique identifier of the resource, this will be the business key.
	// Required: true
	ResourceID *string `json:"resourceId"`
}

// Validate validates this echo workflow request
func (m *EchoWorkflowRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EchoWorkflowRequest) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *EchoWorkflowRequest) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resourceId", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this echo workflow request based on context it is used
func (m *EchoWorkflowRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EchoWorkflowRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EchoWorkflowRequest) UnmarshalBinary(b []byte) error {
	var res EchoWorkflowRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
