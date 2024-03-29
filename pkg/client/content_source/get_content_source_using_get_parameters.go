// Code generated by go-swagger; DO NOT EDIT.

package content_source

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

// NewGetContentSourceUsingGETParams creates a new GetContentSourceUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentSourceUsingGETParams() *GetContentSourceUsingGETParams {
	return &GetContentSourceUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentSourceUsingGETParamsWithTimeout creates a new GetContentSourceUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetContentSourceUsingGETParamsWithTimeout(timeout time.Duration) *GetContentSourceUsingGETParams {
	return &GetContentSourceUsingGETParams{
		timeout: timeout,
	}
}

// NewGetContentSourceUsingGETParamsWithContext creates a new GetContentSourceUsingGETParams object
// with the ability to set a context for a request.
func NewGetContentSourceUsingGETParamsWithContext(ctx context.Context) *GetContentSourceUsingGETParams {
	return &GetContentSourceUsingGETParams{
		Context: ctx,
	}
}

// NewGetContentSourceUsingGETParamsWithHTTPClient creates a new GetContentSourceUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentSourceUsingGETParamsWithHTTPClient(client *http.Client) *GetContentSourceUsingGETParams {
	return &GetContentSourceUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetContentSourceUsingGETParams contains all the parameters to send to the API endpoint

	for the get content source using g e t operation.

	Typically these are written to a http.Request.
*/
type GetContentSourceUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about
	*/
	APIVersion *string

	/* ID.

	   ID of the content source

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get content source using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentSourceUsingGETParams) WithDefaults() *GetContentSourceUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get content source using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentSourceUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get content source using get params
func (o *GetContentSourceUsingGETParams) WithTimeout(timeout time.Duration) *GetContentSourceUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get content source using get params
func (o *GetContentSourceUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get content source using get params
func (o *GetContentSourceUsingGETParams) WithContext(ctx context.Context) *GetContentSourceUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get content source using get params
func (o *GetContentSourceUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get content source using get params
func (o *GetContentSourceUsingGETParams) WithHTTPClient(client *http.Client) *GetContentSourceUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get content source using get params
func (o *GetContentSourceUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get content source using get params
func (o *GetContentSourceUsingGETParams) WithAPIVersion(aPIVersion *string) *GetContentSourceUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get content source using get params
func (o *GetContentSourceUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get content source using get params
func (o *GetContentSourceUsingGETParams) WithID(id strfmt.UUID) *GetContentSourceUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get content source using get params
func (o *GetContentSourceUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentSourceUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
