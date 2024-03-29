// Code generated by go-swagger; DO NOT EDIT.

package custom_naming

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

// NewUpdateCustomNameParams creates a new UpdateCustomNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCustomNameParams() *UpdateCustomNameParams {
	return &UpdateCustomNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomNameParamsWithTimeout creates a new UpdateCustomNameParams object
// with the ability to set a timeout on a request.
func NewUpdateCustomNameParamsWithTimeout(timeout time.Duration) *UpdateCustomNameParams {
	return &UpdateCustomNameParams{
		timeout: timeout,
	}
}

// NewUpdateCustomNameParamsWithContext creates a new UpdateCustomNameParams object
// with the ability to set a context for a request.
func NewUpdateCustomNameParamsWithContext(ctx context.Context) *UpdateCustomNameParams {
	return &UpdateCustomNameParams{
		Context: ctx,
	}
}

// NewUpdateCustomNameParamsWithHTTPClient creates a new UpdateCustomNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCustomNameParamsWithHTTPClient(client *http.Client) *UpdateCustomNameParams {
	return &UpdateCustomNameParams{
		HTTPClient: client,
	}
}

/*
UpdateCustomNameParams contains all the parameters to send to the API endpoint

	for the update custom name operation.

	Typically these are written to a http.Request.
*/
type UpdateCustomNameParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   Custom Name specification
	*/
	Body *models.CustomNamingModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update custom name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomNameParams) WithDefaults() *UpdateCustomNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update custom name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update custom name params
func (o *UpdateCustomNameParams) WithTimeout(timeout time.Duration) *UpdateCustomNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update custom name params
func (o *UpdateCustomNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update custom name params
func (o *UpdateCustomNameParams) WithContext(ctx context.Context) *UpdateCustomNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update custom name params
func (o *UpdateCustomNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update custom name params
func (o *UpdateCustomNameParams) WithHTTPClient(client *http.Client) *UpdateCustomNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update custom name params
func (o *UpdateCustomNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update custom name params
func (o *UpdateCustomNameParams) WithAPIVersion(aPIVersion string) *UpdateCustomNameParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update custom name params
func (o *UpdateCustomNameParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update custom name params
func (o *UpdateCustomNameParams) WithBody(body *models.CustomNamingModel) *UpdateCustomNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update custom name params
func (o *UpdateCustomNameParams) SetBody(body *models.CustomNamingModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {

		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
