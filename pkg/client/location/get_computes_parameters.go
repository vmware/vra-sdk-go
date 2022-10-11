// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewGetComputesParams creates a new GetComputesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComputesParams() *GetComputesParams {
	return &GetComputesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComputesParamsWithTimeout creates a new GetComputesParams object
// with the ability to set a timeout on a request.
func NewGetComputesParamsWithTimeout(timeout time.Duration) *GetComputesParams {
	return &GetComputesParams{
		timeout: timeout,
	}
}

// NewGetComputesParamsWithContext creates a new GetComputesParams object
// with the ability to set a context for a request.
func NewGetComputesParamsWithContext(ctx context.Context) *GetComputesParams {
	return &GetComputesParams{
		Context: ctx,
	}
}

// NewGetComputesParamsWithHTTPClient creates a new GetComputesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComputesParamsWithHTTPClient(client *http.Client) *GetComputesParams {
	return &GetComputesParams{
		HTTPClient: client,
	}
}

/*
GetComputesParams contains all the parameters to send to the API endpoint

	for the get computes operation.

	Typically these are written to a http.Request.
*/
type GetComputesParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the zone.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get computes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComputesParams) WithDefaults() *GetComputesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get computes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComputesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get computes params
func (o *GetComputesParams) WithTimeout(timeout time.Duration) *GetComputesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get computes params
func (o *GetComputesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get computes params
func (o *GetComputesParams) WithContext(ctx context.Context) *GetComputesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get computes params
func (o *GetComputesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get computes params
func (o *GetComputesParams) WithHTTPClient(client *http.Client) *GetComputesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get computes params
func (o *GetComputesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get computes params
func (o *GetComputesParams) WithAPIVersion(aPIVersion *string) *GetComputesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get computes params
func (o *GetComputesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get computes params
func (o *GetComputesParams) WithID(id string) *GetComputesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get computes params
func (o *GetComputesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetComputesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
