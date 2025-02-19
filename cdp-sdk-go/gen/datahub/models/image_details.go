// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageDetails The details of the image used for cluster instances.
//
// swagger:model ImageDetails
type ImageDetails struct {

	// The image catalog name.
	CatalogName string `json:"catalogName,omitempty"`

	// The image catalog URL.
	CatalogURL string `json:"catalogUrl,omitempty"`

	// The ID of the image used for cluster instances. This is internally generated by the cloud provider to uniquely identify the image.
	ID string `json:"id,omitempty"`

	// The name of the image used for cluster instances.
	Name string `json:"name,omitempty"`
}

// Validate validates this image details
func (m *ImageDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image details based on context it is used
func (m *ImageDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageDetails) UnmarshalBinary(b []byte) error {
	var res ImageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
