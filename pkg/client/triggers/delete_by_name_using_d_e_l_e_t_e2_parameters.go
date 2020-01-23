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

// NewDeleteByNameUsingDELETE2Params creates a new DeleteByNameUsingDELETE2Params object
// with the default values initialized.
func NewDeleteByNameUsingDELETE2Params() *DeleteByNameUsingDELETE2Params {
	var ()
	return &DeleteByNameUsingDELETE2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteByNameUsingDELETE2ParamsWithTimeout creates a new DeleteByNameUsingDELETE2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteByNameUsingDELETE2ParamsWithTimeout(timeout time.Duration) *DeleteByNameUsingDELETE2Params {
	var ()
	return &DeleteByNameUsingDELETE2Params{

		timeout: timeout,
	}
}

// NewDeleteByNameUsingDELETE2ParamsWithContext creates a new DeleteByNameUsingDELETE2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteByNameUsingDELETE2ParamsWithContext(ctx context.Context) *DeleteByNameUsingDELETE2Params {
	var ()
	return &DeleteByNameUsingDELETE2Params{

		Context: ctx,
	}
}

// NewDeleteByNameUsingDELETE2ParamsWithHTTPClient creates a new DeleteByNameUsingDELETE2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteByNameUsingDELETE2ParamsWithHTTPClient(client *http.Client) *DeleteByNameUsingDELETE2Params {
	var ()
	return &DeleteByNameUsingDELETE2Params{
		HTTPClient: client,
	}
}

/*DeleteByNameUsingDELETE2Params contains all the parameters to send to the API endpoint
for the delete by name using d e l e t e 2 operation typically these are written to a http.Request
*/
type DeleteByNameUsingDELETE2Params struct {

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

// WithTimeout adds the timeout to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) WithTimeout(timeout time.Duration) *DeleteByNameUsingDELETE2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) WithContext(ctx context.Context) *DeleteByNameUsingDELETE2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) WithHTTPClient(client *http.Client) *DeleteByNameUsingDELETE2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) WithAPIVersion(aPIVersion string) *DeleteByNameUsingDELETE2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) WithName(name string) *DeleteByNameUsingDELETE2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) WithProject(project string) *DeleteByNameUsingDELETE2Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete by name using d e l e t e 2 params
func (o *DeleteByNameUsingDELETE2Params) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteByNameUsingDELETE2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
