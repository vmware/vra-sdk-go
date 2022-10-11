// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewDeleteBlueprintUsingDELETE1Params creates a new DeleteBlueprintUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBlueprintUsingDELETE1Params() *DeleteBlueprintUsingDELETE1Params {
	return &DeleteBlueprintUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlueprintUsingDELETE1ParamsWithTimeout creates a new DeleteBlueprintUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteBlueprintUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteBlueprintUsingDELETE1Params {
	return &DeleteBlueprintUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteBlueprintUsingDELETE1ParamsWithContext creates a new DeleteBlueprintUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteBlueprintUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteBlueprintUsingDELETE1Params {
	return &DeleteBlueprintUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteBlueprintUsingDELETE1ParamsWithHTTPClient creates a new DeleteBlueprintUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBlueprintUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteBlueprintUsingDELETE1Params {
	return &DeleteBlueprintUsingDELETE1Params{
		HTTPClient: client,
	}
}

/*
DeleteBlueprintUsingDELETE1Params contains all the parameters to send to the API endpoint

	for the delete blueprint using d e l e t e 1 operation.

	Typically these are written to a http.Request.
*/
type DeleteBlueprintUsingDELETE1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* BlueprintID.

	   blueprintId

	   Format: uuid
	*/
	BlueprintID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete blueprint using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlueprintUsingDELETE1Params) WithDefaults() *DeleteBlueprintUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete blueprint using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlueprintUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteBlueprintUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) WithContext(ctx context.Context) *DeleteBlueprintUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteBlueprintUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) WithAPIVersion(aPIVersion *string) *DeleteBlueprintUsingDELETE1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprintID adds the blueprintID to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) WithBlueprintID(blueprintID strfmt.UUID) *DeleteBlueprintUsingDELETE1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the delete blueprint using d e l e t e 1 params
func (o *DeleteBlueprintUsingDELETE1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlueprintUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
