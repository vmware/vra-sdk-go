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

// NewDeleteGerritEventByIDUsingDELETEParams creates a new DeleteGerritEventByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGerritEventByIDUsingDELETEParams() *DeleteGerritEventByIDUsingDELETEParams {
	return &DeleteGerritEventByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGerritEventByIDUsingDELETEParamsWithTimeout creates a new DeleteGerritEventByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteGerritEventByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteGerritEventByIDUsingDELETEParams {
	return &DeleteGerritEventByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteGerritEventByIDUsingDELETEParamsWithContext creates a new DeleteGerritEventByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteGerritEventByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteGerritEventByIDUsingDELETEParams {
	return &DeleteGerritEventByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteGerritEventByIDUsingDELETEParamsWithHTTPClient creates a new DeleteGerritEventByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGerritEventByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteGerritEventByIDUsingDELETEParams {
	return &DeleteGerritEventByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteGerritEventByIDUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete gerrit event by Id using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteGerritEventByIDUsingDELETEParams struct {

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

// WithDefaults hydrates default values in the delete gerrit event by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGerritEventByIDUsingDELETEParams) WithDefaults() *DeleteGerritEventByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete gerrit event by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGerritEventByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteGerritEventByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteGerritEventByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteGerritEventByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteGerritEventByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteGerritEventByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) WithID(id string) *DeleteGerritEventByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete gerrit event by Id using d e l e t e params
func (o *DeleteGerritEventByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGerritEventByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
