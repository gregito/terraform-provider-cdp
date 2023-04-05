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

// Image Basic information about an image.
//
// swagger:model Image
type Image struct {

	// Indicates the visibility of the imgae in the catalog.
	Advertised *bool `json:"advertised,omitempty"`

	// Image creation timestamp.
	Created int64 `json:"created,omitempty"`

	// The date when the image was created.
	Date string `json:"date,omitempty"`

	// Description of the image.
	Description string `json:"description,omitempty"`

	// Available regions and images for each cloud provider.
	Images *ImageReferenceSet `json:"images,omitempty"`

	// Installed OS of the image.
	Os string `json:"os,omitempty"`

	// The distribution family of OS installed on the image.
	OsType string `json:"osType,omitempty"`

	// Package versions.
	PackageVersions map[string]string `json:"packageVersions,omitempty"`

	// The UUID of the image.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Image) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this image based on the context it is used
func (m *Image) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Image) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	if m.Images != nil {
		if err := m.Images.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
