// Code generated by go-swagger; DO NOT EDIT.

package integration

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

// NewUpdateIntegrationAsyncParams creates a new UpdateIntegrationAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIntegrationAsyncParams() *UpdateIntegrationAsyncParams {
	return &UpdateIntegrationAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIntegrationAsyncParamsWithTimeout creates a new UpdateIntegrationAsyncParams object
// with the ability to set a timeout on a request.
func NewUpdateIntegrationAsyncParamsWithTimeout(timeout time.Duration) *UpdateIntegrationAsyncParams {
	return &UpdateIntegrationAsyncParams{
		timeout: timeout,
	}
}

// NewUpdateIntegrationAsyncParamsWithContext creates a new UpdateIntegrationAsyncParams object
// with the ability to set a context for a request.
func NewUpdateIntegrationAsyncParamsWithContext(ctx context.Context) *UpdateIntegrationAsyncParams {
	return &UpdateIntegrationAsyncParams{
		Context: ctx,
	}
}

// NewUpdateIntegrationAsyncParamsWithHTTPClient creates a new UpdateIntegrationAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIntegrationAsyncParamsWithHTTPClient(client *http.Client) *UpdateIntegrationAsyncParams {
	return &UpdateIntegrationAsyncParams{
		HTTPClient: client,
	}
}

/*
UpdateIntegrationAsyncParams contains all the parameters to send to the API endpoint

	for the update integration async operation.

	Typically these are written to a http.Request.
*/
type UpdateIntegrationAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Integration details to be updated
	*/
	Body *models.UpdateIntegrationSpecification

	/* ID.

	   The ID of the integration
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update integration async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntegrationAsyncParams) WithDefaults() *UpdateIntegrationAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update integration async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntegrationAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update integration async params
func (o *UpdateIntegrationAsyncParams) WithTimeout(timeout time.Duration) *UpdateIntegrationAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update integration async params
func (o *UpdateIntegrationAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update integration async params
func (o *UpdateIntegrationAsyncParams) WithContext(ctx context.Context) *UpdateIntegrationAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update integration async params
func (o *UpdateIntegrationAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update integration async params
func (o *UpdateIntegrationAsyncParams) WithHTTPClient(client *http.Client) *UpdateIntegrationAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update integration async params
func (o *UpdateIntegrationAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update integration async params
func (o *UpdateIntegrationAsyncParams) WithAPIVersion(aPIVersion *string) *UpdateIntegrationAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update integration async params
func (o *UpdateIntegrationAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update integration async params
func (o *UpdateIntegrationAsyncParams) WithBody(body *models.UpdateIntegrationSpecification) *UpdateIntegrationAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update integration async params
func (o *UpdateIntegrationAsyncParams) SetBody(body *models.UpdateIntegrationSpecification) {
	o.Body = body
}

// WithID adds the id to the update integration async params
func (o *UpdateIntegrationAsyncParams) WithID(id string) *UpdateIntegrationAsyncParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update integration async params
func (o *UpdateIntegrationAsyncParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIntegrationAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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