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

// NewDeleteGerritListenerByIDUsingDELETEParams creates a new DeleteGerritListenerByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGerritListenerByIDUsingDELETEParams() *DeleteGerritListenerByIDUsingDELETEParams {
	return &DeleteGerritListenerByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGerritListenerByIDUsingDELETEParamsWithTimeout creates a new DeleteGerritListenerByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteGerritListenerByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteGerritListenerByIDUsingDELETEParams {
	return &DeleteGerritListenerByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteGerritListenerByIDUsingDELETEParamsWithContext creates a new DeleteGerritListenerByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteGerritListenerByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteGerritListenerByIDUsingDELETEParams {
	return &DeleteGerritListenerByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteGerritListenerByIDUsingDELETEParamsWithHTTPClient creates a new DeleteGerritListenerByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGerritListenerByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteGerritListenerByIDUsingDELETEParams {
	return &DeleteGerritListenerByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteGerritListenerByIDUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete gerrit listener by Id using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteGerritListenerByIDUsingDELETEParams struct {

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

// WithDefaults hydrates default values in the delete gerrit listener by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithDefaults() *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete gerrit listener by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) WithID(id string) *DeleteGerritListenerByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete gerrit listener by Id using d e l e t e params
func (o *DeleteGerritListenerByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGerritListenerByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
