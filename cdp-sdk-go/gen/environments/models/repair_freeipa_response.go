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

// RepairFreeipaResponse Response object for an FreeIPA repair request.
//
// swagger:model RepairFreeipaResponse
type RepairFreeipaResponse struct {

	// Date when the operation ended. Omitted if operation has not ended.
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// If there is any error associated. The error will be populated on any error and it may be populated when the operation failure details are empty. The error will typically contain the high level information such as the assocated repair failure phase.
	Error string `json:"error,omitempty"`

	// List of operation details for failures. If the repair is only partially successful both successful and failure operation details will be populated.
	FailureOperationDetails []*RepairOperationDetails `json:"failureOperationDetails"`

	// Operation ID of the request for this operation. This ID can be used for geting status on the operation.
	OperationID string `json:"operationId,omitempty"`

	// Date when the operation started.
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Status of this operation.
	Status RepairStatus `json:"status,omitempty"`

	// List of operation details for all successes. If the repair is only partially successful both successful and failure operation details will be populated.
	SuccessfulOperationDetails []*RepairOperationDetails `json:"successfulOperationDetails"`
}

// Validate validates this repair freeipa response
func (m *RepairFreeipaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailureOperationDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessfulOperationDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepairFreeipaResponse) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RepairFreeipaResponse) validateFailureOperationDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.FailureOperationDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.FailureOperationDetails); i++ {
		if swag.IsZero(m.FailureOperationDetails[i]) { // not required
			continue
		}

		if m.FailureOperationDetails[i] != nil {
			if err := m.FailureOperationDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failureOperationDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failureOperationDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RepairFreeipaResponse) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RepairFreeipaResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *RepairFreeipaResponse) validateSuccessfulOperationDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.SuccessfulOperationDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.SuccessfulOperationDetails); i++ {
		if swag.IsZero(m.SuccessfulOperationDetails[i]) { // not required
			continue
		}

		if m.SuccessfulOperationDetails[i] != nil {
			if err := m.SuccessfulOperationDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("successfulOperationDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("successfulOperationDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this repair freeipa response based on the context it is used
func (m *RepairFreeipaResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailureOperationDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccessfulOperationDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepairFreeipaResponse) contextValidateFailureOperationDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FailureOperationDetails); i++ {

		if m.FailureOperationDetails[i] != nil {

			if swag.IsZero(m.FailureOperationDetails[i]) { // not required
				return nil
			}

			if err := m.FailureOperationDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failureOperationDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failureOperationDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RepairFreeipaResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *RepairFreeipaResponse) contextValidateSuccessfulOperationDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SuccessfulOperationDetails); i++ {

		if m.SuccessfulOperationDetails[i] != nil {

			if swag.IsZero(m.SuccessfulOperationDetails[i]) { // not required
				return nil
			}

			if err := m.SuccessfulOperationDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("successfulOperationDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("successfulOperationDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepairFreeipaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepairFreeipaResponse) UnmarshalBinary(b []byte) error {
	var res RepairFreeipaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
