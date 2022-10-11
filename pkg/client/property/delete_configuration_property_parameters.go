// Code generated by go-swagger; DO NOT EDIT.

package property

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

// NewDeleteConfigurationPropertyParams creates a new DeleteConfigurationPropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConfigurationPropertyParams() *DeleteConfigurationPropertyParams {
	return &DeleteConfigurationPropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfigurationPropertyParamsWithTimeout creates a new DeleteConfigurationPropertyParams object
// with the ability to set a timeout on a request.
func NewDeleteConfigurationPropertyParamsWithTimeout(timeout time.Duration) *DeleteConfigurationPropertyParams {
	return &DeleteConfigurationPropertyParams{
		timeout: timeout,
	}
}

// NewDeleteConfigurationPropertyParamsWithContext creates a new DeleteConfigurationPropertyParams object
// with the ability to set a context for a request.
func NewDeleteConfigurationPropertyParamsWithContext(ctx context.Context) *DeleteConfigurationPropertyParams {
	return &DeleteConfigurationPropertyParams{
		Context: ctx,
	}
}

// NewDeleteConfigurationPropertyParamsWithHTTPClient creates a new DeleteConfigurationPropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConfigurationPropertyParamsWithHTTPClient(client *http.Client) *DeleteConfigurationPropertyParams {
	return &DeleteConfigurationPropertyParams{
		HTTPClient: client,
	}
}

/*
DeleteConfigurationPropertyParams contains all the parameters to send to the API endpoint

	for the delete configuration property operation.

	Typically these are written to a http.Request.
*/
type DeleteConfigurationPropertyParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete configuration property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfigurationPropertyParams) WithDefaults() *DeleteConfigurationPropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete configuration property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfigurationPropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) WithTimeout(timeout time.Duration) *DeleteConfigurationPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) WithContext(ctx context.Context) *DeleteConfigurationPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) WithHTTPClient(client *http.Client) *DeleteConfigurationPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) WithAPIVersion(aPIVersion *string) *DeleteConfigurationPropertyParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) WithID(id string) *DeleteConfigurationPropertyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete configuration property params
func (o *DeleteConfigurationPropertyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfigurationPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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