// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

// NewDeleteGcpStorageProfileParams creates a new DeleteGcpStorageProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGcpStorageProfileParams() *DeleteGcpStorageProfileParams {
	return &DeleteGcpStorageProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGcpStorageProfileParamsWithTimeout creates a new DeleteGcpStorageProfileParams object
// with the ability to set a timeout on a request.
func NewDeleteGcpStorageProfileParamsWithTimeout(timeout time.Duration) *DeleteGcpStorageProfileParams {
	return &DeleteGcpStorageProfileParams{
		timeout: timeout,
	}
}

// NewDeleteGcpStorageProfileParamsWithContext creates a new DeleteGcpStorageProfileParams object
// with the ability to set a context for a request.
func NewDeleteGcpStorageProfileParamsWithContext(ctx context.Context) *DeleteGcpStorageProfileParams {
	return &DeleteGcpStorageProfileParams{
		Context: ctx,
	}
}

// NewDeleteGcpStorageProfileParamsWithHTTPClient creates a new DeleteGcpStorageProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGcpStorageProfileParamsWithHTTPClient(client *http.Client) *DeleteGcpStorageProfileParams {
	return &DeleteGcpStorageProfileParams{
		HTTPClient: client,
	}
}

/*
DeleteGcpStorageProfileParams contains all the parameters to send to the API endpoint

	for the delete gcp storage profile operation.

	Typically these are written to a http.Request.
*/
type DeleteGcpStorageProfileParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the storage profile.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete gcp storage profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGcpStorageProfileParams) WithDefaults() *DeleteGcpStorageProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete gcp storage profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGcpStorageProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) WithTimeout(timeout time.Duration) *DeleteGcpStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) WithContext(ctx context.Context) *DeleteGcpStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) WithHTTPClient(client *http.Client) *DeleteGcpStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) WithAPIVersion(aPIVersion *string) *DeleteGcpStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) WithID(id string) *DeleteGcpStorageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete gcp storage profile params
func (o *DeleteGcpStorageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGcpStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
