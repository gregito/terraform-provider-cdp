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

// ListVwDiagnosticDataJobsResponse Response object for the listVwDiagnosticDataJobsRequest method.
//
// swagger:model ListVwDiagnosticDataJobsResponse
type ListVwDiagnosticDataJobsResponse struct {

	// The list of the jobs belonging to the particular Virtual Warehouse.
	Jobs []*VwDiagnosticDataJob `json:"jobs"`

	// The token to use when requesting the next set of results. If there are no additional results, the string is empty.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this list vw diagnostic data jobs response
func (m *ListVwDiagnosticDataJobsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVwDiagnosticDataJobsResponse) validateJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list vw diagnostic data jobs response based on the context it is used
func (m *ListVwDiagnosticDataJobsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVwDiagnosticDataJobsResponse) contextValidateJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Jobs); i++ {

		if m.Jobs[i] != nil {

			if swag.IsZero(m.Jobs[i]) { // not required
				return nil
			}

			if err := m.Jobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListVwDiagnosticDataJobsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListVwDiagnosticDataJobsResponse) UnmarshalBinary(b []byte) error {
	var res ListVwDiagnosticDataJobsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
