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

// NewDeleteGerritListenerByNameUsingDELETEParams creates a new DeleteGerritListenerByNameUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGerritListenerByNameUsingDELETEParams() *DeleteGerritListenerByNameUsingDELETEParams {
	return &DeleteGerritListenerByNameUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGerritListenerByNameUsingDELETEParamsWithTimeout creates a new DeleteGerritListenerByNameUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteGerritListenerByNameUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteGerritListenerByNameUsingDELETEParams {
	return &DeleteGerritListenerByNameUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteGerritListenerByNameUsingDELETEParamsWithContext creates a new DeleteGerritListenerByNameUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteGerritListenerByNameUsingDELETEParamsWithContext(ctx context.Context) *DeleteGerritListenerByNameUsingDELETEParams {
	return &DeleteGerritListenerByNameUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteGerritListenerByNameUsingDELETEParamsWithHTTPClient creates a new DeleteGerritListenerByNameUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGerritListenerByNameUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteGerritListenerByNameUsingDELETEParams {
	return &DeleteGerritListenerByNameUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteGerritListenerByNameUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete gerrit listener by name using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteGerritListenerByNameUsingDELETEParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Name.

	   name
	*/
	Name string

	/* Project.

	   project
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete gerrit listener by name using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithDefaults() *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete gerrit listener by name using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithContext(ctx context.Context) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithAuthorization(authorization string) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithName(name string) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) WithProject(project string) *DeleteGerritListenerByNameUsingDELETEParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete gerrit listener by name using d e l e t e params
func (o *DeleteGerritListenerByNameUsingDELETEParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGerritListenerByNameUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
