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

// DescribeServerSettingResponse Response object for the describeServerSetting method.
//
// swagger:model DescribeServerSettingResponse
type DescribeServerSettingResponse struct {

	// The list of the server settings.
	Settings []*DescribeServerSettingItem `json:"settings"`
}

// Validate validates this describe server setting response
func (m *DescribeServerSettingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeServerSettingResponse) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	for i := 0; i < len(m.Settings); i++ {
		if swag.IsZero(m.Settings[i]) { // not required
			continue
		}

		if m.Settings[i] != nil {
			if err := m.Settings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this describe server setting response based on the context it is used
func (m *DescribeServerSettingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeServerSettingResponse) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Settings); i++ {

		if m.Settings[i] != nil {

			if swag.IsZero(m.Settings[i]) { // not required
				return nil
			}

			if err := m.Settings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeServerSettingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeServerSettingResponse) UnmarshalBinary(b []byte) error {
	var res DescribeServerSettingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
