// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// NewDeleteEndpointByIDUsingDELETEParams creates a new DeleteEndpointByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEndpointByIDUsingDELETEParams() *DeleteEndpointByIDUsingDELETEParams {
	return &DeleteEndpointByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEndpointByIDUsingDELETEParamsWithTimeout creates a new DeleteEndpointByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteEndpointByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteEndpointByIDUsingDELETEParams {
	return &DeleteEndpointByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteEndpointByIDUsingDELETEParamsWithContext creates a new DeleteEndpointByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteEndpointByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteEndpointByIDUsingDELETEParams {
	return &DeleteEndpointByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteEndpointByIDUsingDELETEParamsWithHTTPClient creates a new DeleteEndpointByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEndpointByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteEndpointByIDUsingDELETEParams {
	return &DeleteEndpointByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteEndpointByIDUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete endpoint by Id using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteEndpointByIDUsingDELETEParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Endpoint
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete endpoint by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEndpointByIDUsingDELETEParams) WithDefaults() *DeleteEndpointByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete endpoint by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEndpointByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteEndpointByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteEndpointByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteEndpointByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteEndpointByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteEndpointByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) WithID(id string) *DeleteEndpointByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete endpoint by Id using d e l e t e params
func (o *DeleteEndpointByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEndpointByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
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
