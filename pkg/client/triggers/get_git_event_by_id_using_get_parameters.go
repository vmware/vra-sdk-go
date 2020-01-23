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

// NewGetGitEventByIDUsingGETParams creates a new GetGitEventByIDUsingGETParams object
// with the default values initialized.
func NewGetGitEventByIDUsingGETParams() *GetGitEventByIDUsingGETParams {
	var ()
	return &GetGitEventByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitEventByIDUsingGETParamsWithTimeout creates a new GetGitEventByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGitEventByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetGitEventByIDUsingGETParams {
	var ()
	return &GetGitEventByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetGitEventByIDUsingGETParamsWithContext creates a new GetGitEventByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGitEventByIDUsingGETParamsWithContext(ctx context.Context) *GetGitEventByIDUsingGETParams {
	var ()
	return &GetGitEventByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetGitEventByIDUsingGETParamsWithHTTPClient creates a new GetGitEventByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGitEventByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetGitEventByIDUsingGETParams {
	var ()
	return &GetGitEventByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetGitEventByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get git event by ID using g e t operation typically these are written to a http.Request
*/
type GetGitEventByIDUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetGitEventByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) WithContext(ctx context.Context) *GetGitEventByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetGitEventByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) WithAPIVersion(aPIVersion string) *GetGitEventByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) WithID(id string) *GetGitEventByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get git event by ID using get params
func (o *GetGitEventByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitEventByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
