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

// MachineUser Information about a Cloudera CDP machine user.
//
// swagger:model MachineUser
type MachineUser struct {

	// The list of Azure cloud identities assigned to the machine user.
	AzureCloudIdentities []*AzureCloudIdentity `json:"azureCloudIdentities"`

	// The date when this machine user record was created.
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// The CRN of the user.
	// Required: true
	Crn *string `json:"crn"`

	// The machine user name.
	// Required: true
	MachineUserName *string `json:"machineUserName"`

	// Information about the workload password for the machine user.
	WorkloadPasswordDetails *WorkloadPasswordDetails `json:"workloadPasswordDetails,omitempty"`

	// The username used in all the workload clusters of the machine user.
	WorkloadUsername string `json:"workloadUsername,omitempty"`
}

// Validate validates this machine user
func (m *MachineUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureCloudIdentities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadPasswordDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineUser) validateAzureCloudIdentities(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureCloudIdentities) { // not required
		return nil
	}

	for i := 0; i < len(m.AzureCloudIdentities); i++ {
		if swag.IsZero(m.AzureCloudIdentities[i]) { // not required
			continue
		}

		if m.AzureCloudIdentities[i] != nil {
			if err := m.AzureCloudIdentities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineUser) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MachineUser) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *MachineUser) validateMachineUserName(formats strfmt.Registry) error {

	if err := validate.Required("machineUserName", "body", m.MachineUserName); err != nil {
		return err
	}

	return nil
}

func (m *MachineUser) validateWorkloadPasswordDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadPasswordDetails) { // not required
		return nil
	}

	if m.WorkloadPasswordDetails != nil {
		if err := m.WorkloadPasswordDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadPasswordDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadPasswordDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this machine user based on the context it is used
func (m *MachineUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureCloudIdentities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkloadPasswordDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineUser) contextValidateAzureCloudIdentities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AzureCloudIdentities); i++ {

		if m.AzureCloudIdentities[i] != nil {

			if swag.IsZero(m.AzureCloudIdentities[i]) { // not required
				return nil
			}

			if err := m.AzureCloudIdentities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineUser) contextValidateWorkloadPasswordDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkloadPasswordDetails != nil {

		if swag.IsZero(m.WorkloadPasswordDetails) { // not required
			return nil
		}

		if err := m.WorkloadPasswordDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadPasswordDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadPasswordDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineUser) UnmarshalBinary(b []byte) error {
	var res MachineUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
