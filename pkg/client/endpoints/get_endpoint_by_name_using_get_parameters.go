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

// NewGetEndpointByNameUsingGETParams creates a new GetEndpointByNameUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEndpointByNameUsingGETParams() *GetEndpointByNameUsingGETParams {
	return &GetEndpointByNameUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointByNameUsingGETParamsWithTimeout creates a new GetEndpointByNameUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetEndpointByNameUsingGETParamsWithTimeout(timeout time.Duration) *GetEndpointByNameUsingGETParams {
	return &GetEndpointByNameUsingGETParams{
		timeout: timeout,
	}
}

// NewGetEndpointByNameUsingGETParamsWithContext creates a new GetEndpointByNameUsingGETParams object
// with the ability to set a context for a request.
func NewGetEndpointByNameUsingGETParamsWithContext(ctx context.Context) *GetEndpointByNameUsingGETParams {
	return &GetEndpointByNameUsingGETParams{
		Context: ctx,
	}
}

// NewGetEndpointByNameUsingGETParamsWithHTTPClient creates a new GetEndpointByNameUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEndpointByNameUsingGETParamsWithHTTPClient(client *http.Client) *GetEndpointByNameUsingGETParams {
	return &GetEndpointByNameUsingGETParams{
		HTTPClient: client,
	}
}

/* GetEndpointByNameUsingGETParams contains all the parameters to send to the API endpoint
   for the get endpoint by name using g e t operation.

   Typically these are written to a http.Request.
*/
type GetEndpointByNameUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Name.

	   The name of the Endpoint
	*/
	Name string

	/* Project.

	   The project the Endpoint belongs to
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get endpoint by name using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointByNameUsingGETParams) WithDefaults() *GetEndpointByNameUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get endpoint by name using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointByNameUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithTimeout(timeout time.Duration) *GetEndpointByNameUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithContext(ctx context.Context) *GetEndpointByNameUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithHTTPClient(client *http.Client) *GetEndpointByNameUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithAuthorization(authorization string) *GetEndpointByNameUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithAPIVersion(aPIVersion string) *GetEndpointByNameUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithName(name string) *GetEndpointByNameUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) WithProject(project string) *GetEndpointByNameUsingGETParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get endpoint by name using get params
func (o *GetEndpointByNameUsingGETParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointByNameUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
