// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetUsingGETMixin5Params creates a new GetUsingGETMixin5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsingGETMixin5Params() *GetUsingGETMixin5Params {
	return &GetUsingGETMixin5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGETMixin5ParamsWithTimeout creates a new GetUsingGETMixin5Params object
// with the ability to set a timeout on a request.
func NewGetUsingGETMixin5ParamsWithTimeout(timeout time.Duration) *GetUsingGETMixin5Params {
	return &GetUsingGETMixin5Params{
		timeout: timeout,
	}
}

// NewGetUsingGETMixin5ParamsWithContext creates a new GetUsingGETMixin5Params object
// with the ability to set a context for a request.
func NewGetUsingGETMixin5ParamsWithContext(ctx context.Context) *GetUsingGETMixin5Params {
	return &GetUsingGETMixin5Params{
		Context: ctx,
	}
}

// NewGetUsingGETMixin5ParamsWithHTTPClient creates a new GetUsingGETMixin5Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsingGETMixin5ParamsWithHTTPClient(client *http.Client) *GetUsingGETMixin5Params {
	return &GetUsingGETMixin5Params{
		HTTPClient: client,
	}
}

/*
GetUsingGETMixin5Params contains all the parameters to send to the API endpoint

	for the get using g e t mixin5 operation.

	Typically these are written to a http.Request.
*/
type GetUsingGETMixin5Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* ID.

	   Тhe id of the project.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get using g e t mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGETMixin5Params) WithDefaults() *GetUsingGETMixin5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get using g e t mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGETMixin5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) WithTimeout(timeout time.Duration) *GetUsingGETMixin5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) WithContext(ctx context.Context) *GetUsingGETMixin5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) WithHTTPClient(client *http.Client) *GetUsingGETMixin5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) WithAPIVersion(aPIVersion *string) *GetUsingGETMixin5Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) WithID(id string) *GetUsingGETMixin5Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get using g e t mixin5 params
func (o *GetUsingGETMixin5Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGETMixin5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
