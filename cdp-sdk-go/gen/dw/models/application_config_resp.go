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

// ApplicationConfigResp Configurations for an application inside a service.
//
// swagger:model ApplicationConfigResp
type ApplicationConfigResp struct {

	// List of ConfigBlocks for the application.
	ConfigBlocks []*ConfigBlockResp `json:"configBlocks"`
}

// Validate validates this application config resp
func (m *ApplicationConfigResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigBlocks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationConfigResp) validateConfigBlocks(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigBlocks) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigBlocks); i++ {
		if swag.IsZero(m.ConfigBlocks[i]) { // not required
			continue
		}

		if m.ConfigBlocks[i] != nil {
			if err := m.ConfigBlocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configBlocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configBlocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this application config resp based on the context it is used
func (m *ApplicationConfigResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigBlocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationConfigResp) contextValidateConfigBlocks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConfigBlocks); i++ {

		if m.ConfigBlocks[i] != nil {

			if swag.IsZero(m.ConfigBlocks[i]) { // not required
				return nil
			}

			if err := m.ConfigBlocks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configBlocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configBlocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationConfigResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationConfigResp) UnmarshalBinary(b []byte) error {
	var res ApplicationConfigResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
