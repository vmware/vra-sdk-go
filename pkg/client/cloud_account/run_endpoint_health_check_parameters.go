// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewRunEndpointHealthCheckParams creates a new RunEndpointHealthCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunEndpointHealthCheckParams() *RunEndpointHealthCheckParams {
	return &RunEndpointHealthCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunEndpointHealthCheckParamsWithTimeout creates a new RunEndpointHealthCheckParams object
// with the ability to set a timeout on a request.
func NewRunEndpointHealthCheckParamsWithTimeout(timeout time.Duration) *RunEndpointHealthCheckParams {
	return &RunEndpointHealthCheckParams{
		timeout: timeout,
	}
}

// NewRunEndpointHealthCheckParamsWithContext creates a new RunEndpointHealthCheckParams object
// with the ability to set a context for a request.
func NewRunEndpointHealthCheckParamsWithContext(ctx context.Context) *RunEndpointHealthCheckParams {
	return &RunEndpointHealthCheckParams{
		Context: ctx,
	}
}

// NewRunEndpointHealthCheckParamsWithHTTPClient creates a new RunEndpointHealthCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewRunEndpointHealthCheckParamsWithHTTPClient(client *http.Client) *RunEndpointHealthCheckParams {
	return &RunEndpointHealthCheckParams{
		HTTPClient: client,
	}
}

/*
RunEndpointHealthCheckParams contains all the parameters to send to the API endpoint

	for the run endpoint health check operation.

	Typically these are written to a http.Request.
*/
type RunEndpointHealthCheckParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Cloud account
	*/
	ID string

	/* PeriodicHealthCheckID.

	   If query param is provided then the endpoint health check is not started manually from the UI, but after a scheduled process.
	*/
	PeriodicHealthCheckID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run endpoint health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunEndpointHealthCheckParams) WithDefaults() *RunEndpointHealthCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run endpoint health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunEndpointHealthCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) WithTimeout(timeout time.Duration) *RunEndpointHealthCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) WithContext(ctx context.Context) *RunEndpointHealthCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) WithHTTPClient(client *http.Client) *RunEndpointHealthCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) WithAPIVersion(aPIVersion *string) *RunEndpointHealthCheckParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) WithID(id string) *RunEndpointHealthCheckParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) SetID(id string) {
	o.ID = id
}

// WithPeriodicHealthCheckID adds the periodicHealthCheckID to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) WithPeriodicHealthCheckID(periodicHealthCheckID *string) *RunEndpointHealthCheckParams {
	o.SetPeriodicHealthCheckID(periodicHealthCheckID)
	return o
}

// SetPeriodicHealthCheckID adds the periodicHealthCheckId to the run endpoint health check params
func (o *RunEndpointHealthCheckParams) SetPeriodicHealthCheckID(periodicHealthCheckID *string) {
	o.PeriodicHealthCheckID = periodicHealthCheckID
}

// WriteToRequest writes these params to a swagger request
func (o *RunEndpointHealthCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PeriodicHealthCheckID != nil {

		// query param periodicHealthCheckId
		var qrPeriodicHealthCheckID string

		if o.PeriodicHealthCheckID != nil {
			qrPeriodicHealthCheckID = *o.PeriodicHealthCheckID
		}
		qPeriodicHealthCheckID := qrPeriodicHealthCheckID
		if qPeriodicHealthCheckID != "" {

			if err := r.SetQueryParam("periodicHealthCheckId", qPeriodicHealthCheckID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
