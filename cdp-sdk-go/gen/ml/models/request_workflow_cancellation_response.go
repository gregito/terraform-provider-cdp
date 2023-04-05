// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RequestWorkflowCancellationResponse Response object for RequestWorkflowCancellation.
//
// swagger:model RequestWorkflowCancellationResponse
type RequestWorkflowCancellationResponse struct {

	// The list of workflow metedata for cancelled workflows.
	WorkflowMetadata []*WorkflowMetadata `json:"workflowMetadata"`
}

// Validate validates this request workflow cancellation response
func (m *RequestWorkflowCancellationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkflowMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestWorkflowCancellationResponse) validateWorkflowMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowMetadata) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkflowMetadata); i++ {
		if swag.IsZero(m.WorkflowMetadata[i]) { // not required
			continue
		}

		if m.WorkflowMetadata[i] != nil {
			if err := m.WorkflowMetadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workflowMetadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workflowMetadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this request workflow cancellation response based on the context it is used
func (m *RequestWorkflowCancellationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkflowMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestWorkflowCancellationResponse) contextValidateWorkflowMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WorkflowMetadata); i++ {

		if m.WorkflowMetadata[i] != nil {
			if err := m.WorkflowMetadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workflowMetadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workflowMetadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestWorkflowCancellationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestWorkflowCancellationResponse) UnmarshalBinary(b []byte) error {
	var res RequestWorkflowCancellationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
