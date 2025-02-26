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

// Cluster Information about a cluster.
//
// swagger:model Cluster
type Cluster struct {

	// The cloud platform.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// The Cloudera Manager details.
	ClouderaManager *ClouderaManagerDetails `json:"clouderaManager,omitempty"`

	// The name of the cluster.
	// Required: true
	ClusterName *string `json:"clusterName"`

	// The status of the cluster.
	ClusterStatus string `json:"clusterStatus,omitempty"`

	// The CRN of the cluster template used for the cluster creation.
	ClusterTemplateCrn string `json:"clusterTemplateCrn,omitempty"`

	// The date when the cluster was created.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The CRN of the credential.
	CredentialCrn string `json:"credentialCrn,omitempty"`

	// The CRN of the cluster.
	// Required: true
	Crn *string `json:"crn"`

	// The CRN of the attached datalake.
	DatalakeCrn string `json:"datalakeCrn,omitempty"`

	// The exposed service api endpoints.
	Endpoints *Endpoints `json:"endpoints,omitempty"`

	// The CRN of the environment.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The name of the environment.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The image details.
	ImageDetails *ImageDetails `json:"imageDetails,omitempty"`

	// The instance details.
	InstanceGroups []*InstanceGroup `json:"instanceGroups"`

	// Flag that toggles the multi availability zone for the given datahub cluster when you are not sure what subnet IDs can be used. This way the subnet IDs will be used what the environment suggests.
	MultiAz *bool `json:"multiAz,omitempty"`

	// The cluster node count.
	NodeCount int32 `json:"nodeCount,omitempty"`

	// Security related configurations for Data Hub clusters.
	Security *SecurityResponse `json:"security,omitempty"`

	// The status of the stack.
	Status string `json:"status,omitempty"`

	// The status reason.
	StatusReason string `json:"statusReason,omitempty"`

	// The workload type for the cluster.
	WorkloadType string `json:"workloadType,omitempty"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClouderaManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateClouderaManager(formats strfmt.Registry) error {
	if swag.IsZero(m.ClouderaManager) { // not required
		return nil
	}

	if m.ClouderaManager != nil {
		if err := m.ClouderaManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clouderaManager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clouderaManager")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	if m.Endpoints != nil {
		if err := m.Endpoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoints")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateImageDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageDetails) { // not required
		return nil
	}

	if m.ImageDetails != nil {
		if err := m.ImageDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageDetails")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateInstanceGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cluster) validateSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.Security) { // not required
		return nil
	}

	if m.Security != nil {
		if err := m.Security.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster based on the context it is used
func (m *Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClouderaManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) contextValidateClouderaManager(ctx context.Context, formats strfmt.Registry) error {

	if m.ClouderaManager != nil {

		if swag.IsZero(m.ClouderaManager) { // not required
			return nil
		}

		if err := m.ClouderaManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clouderaManager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clouderaManager")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if m.Endpoints != nil {

		if swag.IsZero(m.Endpoints) { // not required
			return nil
		}

		if err := m.Endpoints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoints")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) contextValidateImageDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageDetails != nil {

		if swag.IsZero(m.ImageDetails) { // not required
			return nil
		}

		if err := m.ImageDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageDetails")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) contextValidateInstanceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceGroups); i++ {

		if m.InstanceGroups[i] != nil {

			if swag.IsZero(m.InstanceGroups[i]) { // not required
				return nil
			}

			if err := m.InstanceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cluster) contextValidateSecurity(ctx context.Context, formats strfmt.Registry) error {

	if m.Security != nil {

		if swag.IsZero(m.Security) { // not required
			return nil
		}

		if err := m.Security.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
