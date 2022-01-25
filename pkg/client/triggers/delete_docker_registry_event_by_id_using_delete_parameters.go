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

// NewDeleteDockerRegistryEventByIDUsingDELETEParams creates a new DeleteDockerRegistryEventByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDockerRegistryEventByIDUsingDELETEParams() *DeleteDockerRegistryEventByIDUsingDELETEParams {
	return &DeleteDockerRegistryEventByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDockerRegistryEventByIDUsingDELETEParamsWithTimeout creates a new DeleteDockerRegistryEventByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteDockerRegistryEventByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	return &DeleteDockerRegistryEventByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteDockerRegistryEventByIDUsingDELETEParamsWithContext creates a new DeleteDockerRegistryEventByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteDockerRegistryEventByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	return &DeleteDockerRegistryEventByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteDockerRegistryEventByIDUsingDELETEParamsWithHTTPClient creates a new DeleteDockerRegistryEventByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDockerRegistryEventByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	return &DeleteDockerRegistryEventByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteDockerRegistryEventByIDUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete docker registry event by Id using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteDockerRegistryEventByIDUsingDELETEParams struct {

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

// WithDefaults hydrates default values in the delete docker registry event by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithDefaults() *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete docker registry event by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WithID(id string) *DeleteDockerRegistryEventByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete docker registry event by Id using d e l e t e params
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDockerRegistryEventByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
