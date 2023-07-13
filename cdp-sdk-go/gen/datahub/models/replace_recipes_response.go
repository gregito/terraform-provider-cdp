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

// ReplaceRecipesResponse The response for replacing recipes.
//
// swagger:model ReplaceRecipesResponse
type ReplaceRecipesResponse struct {

	// The list of recipes, which will be attached to the cluster.
	AttachedRecipes []*InstanceGroupRecipeResponse `json:"attachedRecipes"`

	// The list of recipes, which will be detached from the cluster.
	DetachedRecipes []*InstanceGroupRecipeResponse `json:"detachedRecipes"`
}

// Validate validates this replace recipes response
func (m *ReplaceRecipesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedRecipes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetachedRecipes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceRecipesResponse) validateAttachedRecipes(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachedRecipes) { // not required
		return nil
	}

	for i := 0; i < len(m.AttachedRecipes); i++ {
		if swag.IsZero(m.AttachedRecipes[i]) { // not required
			continue
		}

		if m.AttachedRecipes[i] != nil {
			if err := m.AttachedRecipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachedRecipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachedRecipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReplaceRecipesResponse) validateDetachedRecipes(formats strfmt.Registry) error {
	if swag.IsZero(m.DetachedRecipes) { // not required
		return nil
	}

	for i := 0; i < len(m.DetachedRecipes); i++ {
		if swag.IsZero(m.DetachedRecipes[i]) { // not required
			continue
		}

		if m.DetachedRecipes[i] != nil {
			if err := m.DetachedRecipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detachedRecipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detachedRecipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this replace recipes response based on the context it is used
func (m *ReplaceRecipesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachedRecipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDetachedRecipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceRecipesResponse) contextValidateAttachedRecipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttachedRecipes); i++ {

		if m.AttachedRecipes[i] != nil {

			if swag.IsZero(m.AttachedRecipes[i]) { // not required
				return nil
			}

			if err := m.AttachedRecipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachedRecipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachedRecipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReplaceRecipesResponse) contextValidateDetachedRecipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DetachedRecipes); i++ {

		if m.DetachedRecipes[i] != nil {

			if swag.IsZero(m.DetachedRecipes[i]) { // not required
				return nil
			}

			if err := m.DetachedRecipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detachedRecipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detachedRecipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceRecipesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceRecipesResponse) UnmarshalBinary(b []byte) error {
	var res ReplaceRecipesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
