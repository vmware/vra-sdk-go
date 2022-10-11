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

// NewShutdownMachineParams creates a new ShutdownMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShutdownMachineParams() *ShutdownMachineParams {
	return &ShutdownMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShutdownMachineParamsWithTimeout creates a new ShutdownMachineParams object
// with the ability to set a timeout on a request.
func NewShutdownMachineParamsWithTimeout(timeout time.Duration) *ShutdownMachineParams {
	return &ShutdownMachineParams{
		timeout: timeout,
	}
}

// NewShutdownMachineParamsWithContext creates a new ShutdownMachineParams object
// with the ability to set a context for a request.
func NewShutdownMachineParamsWithContext(ctx context.Context) *ShutdownMachineParams {
	return &ShutdownMachineParams{
		Context: ctx,
	}
}

// NewShutdownMachineParamsWithHTTPClient creates a new ShutdownMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewShutdownMachineParamsWithHTTPClient(client *http.Client) *ShutdownMachineParams {
	return &ShutdownMachineParams{
		HTTPClient: client,
	}
}

/*
ShutdownMachineParams contains all the parameters to send to the API endpoint

	for the shutdown machine operation.

	Typically these are written to a http.Request.
*/
type ShutdownMachineParams struct {

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

// WithDefaults hydrates default values in the shutdown machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShutdownMachineParams) WithDefaults() *ShutdownMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the shutdown machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShutdownMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the shutdown machine params
func (o *ShutdownMachineParams) WithTimeout(timeout time.Duration) *ShutdownMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shutdown machine params
func (o *ShutdownMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shutdown machine params
func (o *ShutdownMachineParams) WithContext(ctx context.Context) *ShutdownMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shutdown machine params
func (o *ShutdownMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shutdown machine params
func (o *ShutdownMachineParams) WithHTTPClient(client *http.Client) *ShutdownMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shutdown machine params
func (o *ShutdownMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the shutdown machine params
func (o *ShutdownMachineParams) WithAPIVersion(aPIVersion *string) *ShutdownMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the shutdown machine params
func (o *ShutdownMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the shutdown machine params
func (o *ShutdownMachineParams) WithID(id string) *ShutdownMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the shutdown machine params
func (o *ShutdownMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ShutdownMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
