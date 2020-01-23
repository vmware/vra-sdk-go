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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetGerritListenerByNameUsingGETParams creates a new GetGerritListenerByNameUsingGETParams object
// with the default values initialized.
func NewGetGerritListenerByNameUsingGETParams() *GetGerritListenerByNameUsingGETParams {
	var ()
	return &GetGerritListenerByNameUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGerritListenerByNameUsingGETParamsWithTimeout creates a new GetGerritListenerByNameUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGerritListenerByNameUsingGETParamsWithTimeout(timeout time.Duration) *GetGerritListenerByNameUsingGETParams {
	var ()
	return &GetGerritListenerByNameUsingGETParams{

		timeout: timeout,
	}
}

// NewGetGerritListenerByNameUsingGETParamsWithContext creates a new GetGerritListenerByNameUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGerritListenerByNameUsingGETParamsWithContext(ctx context.Context) *GetGerritListenerByNameUsingGETParams {
	var ()
	return &GetGerritListenerByNameUsingGETParams{

		Context: ctx,
	}
}

// NewGetGerritListenerByNameUsingGETParamsWithHTTPClient creates a new GetGerritListenerByNameUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGerritListenerByNameUsingGETParamsWithHTTPClient(client *http.Client) *GetGerritListenerByNameUsingGETParams {
	var ()
	return &GetGerritListenerByNameUsingGETParams{
		HTTPClient: client,
	}
}

/*GetGerritListenerByNameUsingGETParams contains all the parameters to send to the API endpoint
for the get gerrit listener by name using g e t operation typically these are written to a http.Request
*/
type GetGerritListenerByNameUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Name
	  name

	*/
	Name string
	/*Project
	  project

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) WithTimeout(timeout time.Duration) *GetGerritListenerByNameUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) WithContext(ctx context.Context) *GetGerritListenerByNameUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) WithHTTPClient(client *http.Client) *GetGerritListenerByNameUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) WithAPIVersion(aPIVersion string) *GetGerritListenerByNameUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) WithName(name string) *GetGerritListenerByNameUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) WithProject(project string) *GetGerritListenerByNameUsingGETParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get gerrit listener by name using get params
func (o *GetGerritListenerByNameUsingGETParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetGerritListenerByNameUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
