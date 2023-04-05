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

// Recipe Information about a recipe.
//
// swagger:model Recipe
type Recipe struct {

	// The CRN of the creator of the recipe.
	CreatorCrn string `json:"creatorCrn,omitempty"`

	// The CRN of the recipe.
	// Required: true
	Crn *string `json:"crn"`

	// The description of the recipe.
	Description string `json:"description,omitempty"`

	// The content of the recipe.
	RecipeContent string `json:"recipeContent,omitempty"`

	// The name of the recipe.
	// Required: true
	RecipeName *string `json:"recipeName"`

	// The type of recipe. Supported values are : PRE_CLOUDERA_MANAGER_START, PRE_TERMINATION, POST_CLOUDERA_MANAGER_START, POST_CLUSTER_INSTALL.
	Type string `json:"type,omitempty"`
}

// Validate validates this recipe
func (m *Recipe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipe) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *Recipe) validateRecipeName(formats strfmt.Registry) error {

	if err := validate.Required("recipeName", "body", m.RecipeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recipe based on context it is used
func (m *Recipe) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Recipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipe) UnmarshalBinary(b []byte) error {
	var res Recipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
