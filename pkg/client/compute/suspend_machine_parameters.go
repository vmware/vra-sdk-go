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
)

// NewSuspendMachineParams creates a new SuspendMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuspendMachineParams() *SuspendMachineParams {
	return &SuspendMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuspendMachineParamsWithTimeout creates a new SuspendMachineParams object
// with the ability to set a timeout on a request.
func NewSuspendMachineParamsWithTimeout(timeout time.Duration) *SuspendMachineParams {
	return &SuspendMachineParams{
		timeout: timeout,
	}
}

// NewSuspendMachineParamsWithContext creates a new SuspendMachineParams object
// with the ability to set a context for a request.
func NewSuspendMachineParamsWithContext(ctx context.Context) *SuspendMachineParams {
	return &SuspendMachineParams{
		Context: ctx,
	}
}

// NewSuspendMachineParamsWithHTTPClient creates a new SuspendMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuspendMachineParamsWithHTTPClient(client *http.Client) *SuspendMachineParams {
	return &SuspendMachineParams{
		HTTPClient: client,
	}
}

/*
SuspendMachineParams contains all the parameters to send to the API endpoint

	for the suspend machine operation.

	Typically these are written to a http.Request.
*/
type SuspendMachineParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The id of the Machine.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suspend machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendMachineParams) WithDefaults() *SuspendMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suspend machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the suspend machine params
func (o *SuspendMachineParams) WithTimeout(timeout time.Duration) *SuspendMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suspend machine params
func (o *SuspendMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suspend machine params
func (o *SuspendMachineParams) WithContext(ctx context.Context) *SuspendMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suspend machine params
func (o *SuspendMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suspend machine params
func (o *SuspendMachineParams) WithHTTPClient(client *http.Client) *SuspendMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suspend machine params
func (o *SuspendMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the suspend machine params
func (o *SuspendMachineParams) WithAPIVersion(aPIVersion *string) *SuspendMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the suspend machine params
func (o *SuspendMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the suspend machine params
func (o *SuspendMachineParams) WithID(id string) *SuspendMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the suspend machine params
func (o *SuspendMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SuspendMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
