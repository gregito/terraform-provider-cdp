// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceGroupRequest Configurations for instance group
//
// swagger:model InstanceGroupRequest
type InstanceGroupRequest struct {

	// The attached volume configuration. This does not include root volume.
	// Required: true
	AttachedVolumeConfiguration []*AttachedVolumeRequest `json:"attachedVolumeConfiguration"`

	// The instance group name.
	// Required: true
	InstanceGroupName *string `json:"instanceGroupName"`

	// The instance group type.
	// Required: true
	InstanceGroupType *string `json:"instanceGroupType"`

	// The cloud provider specific instance type to be used.
	// Required: true
	InstanceType *string `json:"instanceType"`

	// Number of instances in the instance group
	// Required: true
	NodeCount *int32 `json:"nodeCount"`

	// The names or CRNs of the recipes that would be applied to the instance group.
	RecipeNames []string `json:"recipeNames"`

	// Recovery mode for the instance group.
	RecoveryMode string `json:"recoveryMode,omitempty"`

	// The root volume size.
	RootVolumeSize int32 `json:"rootVolumeSize,omitempty"`

	// Max hourly price of spot instances.
	// Maximum: 255
	// Minimum: 0.001
	SpotMaxPrice float64 `json:"spotMaxPrice,omitempty"`

	// Percentage of spot instances in instance group.
	// Maximum: 100
	// Minimum: 0
	SpotPercentage *int32 `json:"spotPercentage,omitempty"`

	// The list of subnet IDs in case of multi-availability zone setup. Specifying this field overrides the datahub level subnet ID setup for the multi-availability zone configuration.
	SubnetIds []string `json:"subnetIds"`

	// Controls how to handle local storage in the instance group. Supported values: ATTACHED_VOLUMES - default behaviour. EPHEMERAL_VOLUMES - use local storage as cache. EPHEMERAL_VOLUMES_ONLY - use local storage for everything.
	// Enum: [ATTACHED_VOLUMES EPHEMERAL_VOLUMES EPHEMERAL_VOLUMES_ONLY]
	TemporaryStorage string `json:"temporaryStorage,omitempty"`

	// The volume encryption settings. This setting does not apply to Azure which always encrypts volumes.
	VolumeEncryption *VolumeEncryptionRequest `json:"volumeEncryption,omitempty"`
}

// Validate validates this instance group request
func (m *InstanceGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedVolumeConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpotMaxPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpotPercentage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemporaryStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeEncryption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupRequest) validateAttachedVolumeConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("attachedVolumeConfiguration", "body", m.AttachedVolumeConfiguration); err != nil {
		return err
	}

	for i := 0; i < len(m.AttachedVolumeConfiguration); i++ {
		if swag.IsZero(m.AttachedVolumeConfiguration[i]) { // not required
			continue
		}

		if m.AttachedVolumeConfiguration[i] != nil {
			if err := m.AttachedVolumeConfiguration[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachedVolumeConfiguration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachedVolumeConfiguration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroupRequest) validateInstanceGroupName(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupName", "body", m.InstanceGroupName); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupRequest) validateInstanceGroupType(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupType", "body", m.InstanceGroupType); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupRequest) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupRequest) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupRequest) validateSpotMaxPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.SpotMaxPrice) { // not required
		return nil
	}

	if err := validate.Minimum("spotMaxPrice", "body", m.SpotMaxPrice, 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("spotMaxPrice", "body", m.SpotMaxPrice, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupRequest) validateSpotPercentage(formats strfmt.Registry) error {
	if swag.IsZero(m.SpotPercentage) { // not required
		return nil
	}

	if err := validate.MinimumInt("spotPercentage", "body", int64(*m.SpotPercentage), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("spotPercentage", "body", int64(*m.SpotPercentage), 100, false); err != nil {
		return err
	}

	return nil
}

var instanceGroupRequestTypeTemporaryStoragePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ATTACHED_VOLUMES","EPHEMERAL_VOLUMES","EPHEMERAL_VOLUMES_ONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupRequestTypeTemporaryStoragePropEnum = append(instanceGroupRequestTypeTemporaryStoragePropEnum, v)
	}
}

const (

	// InstanceGroupRequestTemporaryStorageATTACHEDVOLUMES captures enum value "ATTACHED_VOLUMES"
	InstanceGroupRequestTemporaryStorageATTACHEDVOLUMES string = "ATTACHED_VOLUMES"

	// InstanceGroupRequestTemporaryStorageEPHEMERALVOLUMES captures enum value "EPHEMERAL_VOLUMES"
	InstanceGroupRequestTemporaryStorageEPHEMERALVOLUMES string = "EPHEMERAL_VOLUMES"

	// InstanceGroupRequestTemporaryStorageEPHEMERALVOLUMESONLY captures enum value "EPHEMERAL_VOLUMES_ONLY"
	InstanceGroupRequestTemporaryStorageEPHEMERALVOLUMESONLY string = "EPHEMERAL_VOLUMES_ONLY"
)

// prop value enum
func (m *InstanceGroupRequest) validateTemporaryStorageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, instanceGroupRequestTypeTemporaryStoragePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupRequest) validateTemporaryStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.TemporaryStorage) { // not required
		return nil
	}

	// value enum
	if err := m.validateTemporaryStorageEnum("temporaryStorage", "body", m.TemporaryStorage); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupRequest) validateVolumeEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeEncryption) { // not required
		return nil
	}

	if m.VolumeEncryption != nil {
		if err := m.VolumeEncryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeEncryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeEncryption")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this instance group request based on the context it is used
func (m *InstanceGroupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachedVolumeConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupRequest) contextValidateAttachedVolumeConfiguration(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttachedVolumeConfiguration); i++ {

		if m.AttachedVolumeConfiguration[i] != nil {
			if err := m.AttachedVolumeConfiguration[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachedVolumeConfiguration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachedVolumeConfiguration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroupRequest) contextValidateVolumeEncryption(ctx context.Context, formats strfmt.Registry) error {

	if m.VolumeEncryption != nil {
		if err := m.VolumeEncryption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeEncryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeEncryption")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupRequest) UnmarshalBinary(b []byte) error {
	var res InstanceGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
