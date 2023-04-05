// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewListGroupsForMachineUserParams creates a new ListGroupsForMachineUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGroupsForMachineUserParams() *ListGroupsForMachineUserParams {
	return &ListGroupsForMachineUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGroupsForMachineUserParamsWithTimeout creates a new ListGroupsForMachineUserParams object
// with the ability to set a timeout on a request.
func NewListGroupsForMachineUserParamsWithTimeout(timeout time.Duration) *ListGroupsForMachineUserParams {
	return &ListGroupsForMachineUserParams{
		timeout: timeout,
	}
}

// NewListGroupsForMachineUserParamsWithContext creates a new ListGroupsForMachineUserParams object
// with the ability to set a context for a request.
func NewListGroupsForMachineUserParamsWithContext(ctx context.Context) *ListGroupsForMachineUserParams {
	return &ListGroupsForMachineUserParams{
		Context: ctx,
	}
}

// NewListGroupsForMachineUserParamsWithHTTPClient creates a new ListGroupsForMachineUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGroupsForMachineUserParamsWithHTTPClient(client *http.Client) *ListGroupsForMachineUserParams {
	return &ListGroupsForMachineUserParams{
		HTTPClient: client,
	}
}

/*
ListGroupsForMachineUserParams contains all the parameters to send to the API endpoint

	for the list groups for machine user operation.

	Typically these are written to a http.Request.
*/
type ListGroupsForMachineUserParams struct {

	// Input.
	Input *models.ListGroupsForMachineUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list groups for machine user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGroupsForMachineUserParams) WithDefaults() *ListGroupsForMachineUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list groups for machine user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGroupsForMachineUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) WithTimeout(timeout time.Duration) *ListGroupsForMachineUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) WithContext(ctx context.Context) *ListGroupsForMachineUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) WithHTTPClient(client *http.Client) *ListGroupsForMachineUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) WithInput(input *models.ListGroupsForMachineUserRequest) *ListGroupsForMachineUserParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list groups for machine user params
func (o *ListGroupsForMachineUserParams) SetInput(input *models.ListGroupsForMachineUserRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListGroupsForMachineUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
