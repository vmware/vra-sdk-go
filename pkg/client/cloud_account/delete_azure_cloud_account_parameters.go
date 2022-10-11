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

// NewDeleteAzureCloudAccountParams creates a new DeleteAzureCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAzureCloudAccountParams() *DeleteAzureCloudAccountParams {
	return &DeleteAzureCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAzureCloudAccountParamsWithTimeout creates a new DeleteAzureCloudAccountParams object
// with the ability to set a timeout on a request.
func NewDeleteAzureCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteAzureCloudAccountParams {
	return &DeleteAzureCloudAccountParams{
		timeout: timeout,
	}
}

// NewDeleteAzureCloudAccountParamsWithContext creates a new DeleteAzureCloudAccountParams object
// with the ability to set a context for a request.
func NewDeleteAzureCloudAccountParamsWithContext(ctx context.Context) *DeleteAzureCloudAccountParams {
	return &DeleteAzureCloudAccountParams{
		Context: ctx,
	}
}

// NewDeleteAzureCloudAccountParamsWithHTTPClient creates a new DeleteAzureCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAzureCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteAzureCloudAccountParams {
	return &DeleteAzureCloudAccountParams{
		HTTPClient: client,
	}
}

/*
DeleteAzureCloudAccountParams contains all the parameters to send to the API endpoint

	for the delete azure cloud account operation.

	Typically these are written to a http.Request.
*/
type DeleteAzureCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Cloud Account
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete azure cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAzureCloudAccountParams) WithDefaults() *DeleteAzureCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete azure cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAzureCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteAzureCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) WithContext(ctx context.Context) *DeleteAzureCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteAzureCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) WithAPIVersion(aPIVersion *string) *DeleteAzureCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) WithID(id string) *DeleteAzureCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete azure cloud account params
func (o *DeleteAzureCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAzureCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
