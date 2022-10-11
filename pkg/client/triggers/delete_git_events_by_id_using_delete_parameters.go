// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewDeleteGitEventsByIDUsingDELETEParams creates a new DeleteGitEventsByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGitEventsByIDUsingDELETEParams() *DeleteGitEventsByIDUsingDELETEParams {
	return &DeleteGitEventsByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGitEventsByIDUsingDELETEParamsWithTimeout creates a new DeleteGitEventsByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteGitEventsByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteGitEventsByIDUsingDELETEParams {
	return &DeleteGitEventsByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteGitEventsByIDUsingDELETEParamsWithContext creates a new DeleteGitEventsByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteGitEventsByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteGitEventsByIDUsingDELETEParams {
	return &DeleteGitEventsByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteGitEventsByIDUsingDELETEParamsWithHTTPClient creates a new DeleteGitEventsByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGitEventsByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteGitEventsByIDUsingDELETEParams {
	return &DeleteGitEventsByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteGitEventsByIDUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete git events by Id using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteGitEventsByIDUsingDELETEParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete git events by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGitEventsByIDUsingDELETEParams) WithDefaults() *DeleteGitEventsByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete git events by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGitEventsByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteGitEventsByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteGitEventsByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteGitEventsByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteGitEventsByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteGitEventsByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) WithID(id string) *DeleteGitEventsByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete git events by Id using d e l e t e params
func (o *DeleteGitEventsByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGitEventsByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
