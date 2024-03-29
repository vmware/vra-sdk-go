// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

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

// NewDeleteUsingDELETE4Params creates a new DeleteUsingDELETE4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsingDELETE4Params() *DeleteUsingDELETE4Params {
	return &DeleteUsingDELETE4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsingDELETE4ParamsWithTimeout creates a new DeleteUsingDELETE4Params object
// with the ability to set a timeout on a request.
func NewDeleteUsingDELETE4ParamsWithTimeout(timeout time.Duration) *DeleteUsingDELETE4Params {
	return &DeleteUsingDELETE4Params{
		timeout: timeout,
	}
}

// NewDeleteUsingDELETE4ParamsWithContext creates a new DeleteUsingDELETE4Params object
// with the ability to set a context for a request.
func NewDeleteUsingDELETE4ParamsWithContext(ctx context.Context) *DeleteUsingDELETE4Params {
	return &DeleteUsingDELETE4Params{
		Context: ctx,
	}
}

// NewDeleteUsingDELETE4ParamsWithHTTPClient creates a new DeleteUsingDELETE4Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsingDELETE4ParamsWithHTTPClient(client *http.Client) *DeleteUsingDELETE4Params {
	return &DeleteUsingDELETE4Params{
		HTTPClient: client,
	}
}

/*
DeleteUsingDELETE4Params contains all the parameters to send to the API endpoint

	for the delete using d e l e t e 4 operation.

	Typically these are written to a http.Request.
*/
type DeleteUsingDELETE4Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* SourceID.

	   Catalog source ID

	   Format: uuid
	*/
	SourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete using d e l e t e 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE4Params) WithDefaults() *DeleteUsingDELETE4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete using d e l e t e 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) WithTimeout(timeout time.Duration) *DeleteUsingDELETE4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) WithContext(ctx context.Context) *DeleteUsingDELETE4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) WithHTTPClient(client *http.Client) *DeleteUsingDELETE4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) WithAPIVersion(aPIVersion *string) *DeleteUsingDELETE4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithSourceID adds the sourceID to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) WithSourceID(sourceID strfmt.UUID) *DeleteUsingDELETE4Params {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the delete using d e l e t e 4 params
func (o *DeleteUsingDELETE4Params) SetSourceID(sourceID strfmt.UUID) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsingDELETE4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
