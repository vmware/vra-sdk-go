// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewGetExecutionByNameAndIndexUsingGETParams creates a new GetExecutionByNameAndIndexUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExecutionByNameAndIndexUsingGETParams() *GetExecutionByNameAndIndexUsingGETParams {
	return &GetExecutionByNameAndIndexUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExecutionByNameAndIndexUsingGETParamsWithTimeout creates a new GetExecutionByNameAndIndexUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetExecutionByNameAndIndexUsingGETParamsWithTimeout(timeout time.Duration) *GetExecutionByNameAndIndexUsingGETParams {
	return &GetExecutionByNameAndIndexUsingGETParams{
		timeout: timeout,
	}
}

// NewGetExecutionByNameAndIndexUsingGETParamsWithContext creates a new GetExecutionByNameAndIndexUsingGETParams object
// with the ability to set a context for a request.
func NewGetExecutionByNameAndIndexUsingGETParamsWithContext(ctx context.Context) *GetExecutionByNameAndIndexUsingGETParams {
	return &GetExecutionByNameAndIndexUsingGETParams{
		Context: ctx,
	}
}

// NewGetExecutionByNameAndIndexUsingGETParamsWithHTTPClient creates a new GetExecutionByNameAndIndexUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExecutionByNameAndIndexUsingGETParamsWithHTTPClient(client *http.Client) *GetExecutionByNameAndIndexUsingGETParams {
	return &GetExecutionByNameAndIndexUsingGETParams{
		HTTPClient: client,
	}
}

/* GetExecutionByNameAndIndexUsingGETParams contains all the parameters to send to the API endpoint
   for the get execution by name and index using g e t operation.

   Typically these are written to a http.Request.
*/
type GetExecutionByNameAndIndexUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Index.

	   The index of the Execution
	*/
	Index string

	/* Name.

	   The name of the Pipeline
	*/
	Name string

	/* Project.

	   The project the Pipeline belongs to
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get execution by name and index using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExecutionByNameAndIndexUsingGETParams) WithDefaults() *GetExecutionByNameAndIndexUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get execution by name and index using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExecutionByNameAndIndexUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithTimeout(timeout time.Duration) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithContext(ctx context.Context) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithHTTPClient(client *http.Client) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithAuthorization(authorization string) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithAPIVersion(aPIVersion string) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithIndex adds the index to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithIndex(index string) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetIndex(index string) {
	o.Index = index
}

// WithName adds the name to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithName(name string) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) WithProject(project string) *GetExecutionByNameAndIndexUsingGETParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get execution by name and index using get params
func (o *GetExecutionByNameAndIndexUsingGETParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetExecutionByNameAndIndexUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
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
