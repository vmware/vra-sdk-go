// Code generated by go-swagger; DO NOT EDIT.

package compute

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateMachineParams creates a new UpdateMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMachineParams() *UpdateMachineParams {
	return &UpdateMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMachineParamsWithTimeout creates a new UpdateMachineParams object
// with the ability to set a timeout on a request.
func NewUpdateMachineParamsWithTimeout(timeout time.Duration) *UpdateMachineParams {
	return &UpdateMachineParams{
		timeout: timeout,
	}
}

// NewUpdateMachineParamsWithContext creates a new UpdateMachineParams object
// with the ability to set a context for a request.
func NewUpdateMachineParamsWithContext(ctx context.Context) *UpdateMachineParams {
	return &UpdateMachineParams{
		Context: ctx,
	}
}

// NewUpdateMachineParamsWithHTTPClient creates a new UpdateMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMachineParamsWithHTTPClient(client *http.Client) *UpdateMachineParams {
	return &UpdateMachineParams{
		HTTPClient: client,
	}
}

/*
UpdateMachineParams contains all the parameters to send to the API endpoint

	for the update machine operation.

	Typically these are written to a http.Request.
*/
type UpdateMachineParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Machine Specification
	*/
	Body *models.UpdateMachineSpecification

	/* ID.

	   The ID of the Machine.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMachineParams) WithDefaults() *UpdateMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update machine params
func (o *UpdateMachineParams) WithTimeout(timeout time.Duration) *UpdateMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update machine params
func (o *UpdateMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update machine params
func (o *UpdateMachineParams) WithContext(ctx context.Context) *UpdateMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update machine params
func (o *UpdateMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update machine params
func (o *UpdateMachineParams) WithHTTPClient(client *http.Client) *UpdateMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update machine params
func (o *UpdateMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update machine params
func (o *UpdateMachineParams) WithAPIVersion(aPIVersion *string) *UpdateMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update machine params
func (o *UpdateMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update machine params
func (o *UpdateMachineParams) WithBody(body *models.UpdateMachineSpecification) *UpdateMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update machine params
func (o *UpdateMachineParams) SetBody(body *models.UpdateMachineSpecification) {
	o.Body = body
}

// WithID adds the id to the update machine params
func (o *UpdateMachineParams) WithID(id string) *UpdateMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update machine params
func (o *UpdateMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
