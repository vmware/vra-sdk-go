// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewDeleteVariableByIDUsingDELETEParams creates a new DeleteVariableByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVariableByIDUsingDELETEParams() *DeleteVariableByIDUsingDELETEParams {
	return &DeleteVariableByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVariableByIDUsingDELETEParamsWithTimeout creates a new DeleteVariableByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteVariableByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteVariableByIDUsingDELETEParams {
	return &DeleteVariableByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteVariableByIDUsingDELETEParamsWithContext creates a new DeleteVariableByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteVariableByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteVariableByIDUsingDELETEParams {
	return &DeleteVariableByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteVariableByIDUsingDELETEParamsWithHTTPClient creates a new DeleteVariableByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVariableByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteVariableByIDUsingDELETEParams {
	return &DeleteVariableByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteVariableByIDUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete variable by Id using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteVariableByIDUsingDELETEParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Variable
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete variable by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVariableByIDUsingDELETEParams) WithDefaults() *DeleteVariableByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete variable by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVariableByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteVariableByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteVariableByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteVariableByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteVariableByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteVariableByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) WithID(id string) *DeleteVariableByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete variable by Id using d e l e t e params
func (o *DeleteVariableByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVariableByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
