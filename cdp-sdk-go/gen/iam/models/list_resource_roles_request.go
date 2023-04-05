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

// ListResourceRolesRequest Request object for a list resource roles request.
//
// swagger:model ListResourceRolesRequest
type ListResourceRolesRequest struct {

	// The size of each page.
	// Maximum: 100
	// Minimum: 1
	PageSize int32 `json:"pageSize,omitempty"`

	// The resource roles CRNs to retrieve. If empty all resource roles will be returned.
	ResourceRoleNames []string `json:"resourceRoleNames"`

	// A token to specify where to start paginating. This is the nextToken from a previously truncated response.
	StartingToken string `json:"startingToken,omitempty"`
}

// Validate validates this list resource roles request
func (m *ListResourceRolesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListResourceRolesRequest) validatePageSize(formats strfmt.Registry) error {
	if swag.IsZero(m.PageSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("pageSize", "body", int64(m.PageSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "body", int64(m.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list resource roles request based on context it is used
func (m *ListResourceRolesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListResourceRolesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListResourceRolesRequest) UnmarshalBinary(b []byte) error {
	var res ListResourceRolesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
