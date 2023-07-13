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
	"github.com/go-openapi/validate"
)

// ListUserAssignedResourceRolesResponse Response object for a list user assigned roles request.
//
// swagger:model ListUserAssignedResourceRolesResponse
type ListUserAssignedResourceRolesResponse struct {

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`

	// The user's resource assignments.
	// Required: true
	ResourceAssignments []*ResourceAssignment `json:"resourceAssignments"`
}

// Validate validates this list user assigned resource roles response
func (m *ListUserAssignedResourceRolesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUserAssignedResourceRolesResponse) validateResourceAssignments(formats strfmt.Registry) error {

	if err := validate.Required("resourceAssignments", "body", m.ResourceAssignments); err != nil {
		return err
	}

	for i := 0; i < len(m.ResourceAssignments); i++ {
		if swag.IsZero(m.ResourceAssignments[i]) { // not required
			continue
		}

		if m.ResourceAssignments[i] != nil {
			if err := m.ResourceAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list user assigned resource roles response based on the context it is used
func (m *ListUserAssignedResourceRolesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUserAssignedResourceRolesResponse) contextValidateResourceAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceAssignments); i++ {

		if m.ResourceAssignments[i] != nil {

			if swag.IsZero(m.ResourceAssignments[i]) { // not required
				return nil
			}

			if err := m.ResourceAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListUserAssignedResourceRolesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListUserAssignedResourceRolesResponse) UnmarshalBinary(b []byte) error {
	var res ListUserAssignedResourceRolesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
