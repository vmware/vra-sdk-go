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

// NewGetByIDUsingGET4Params creates a new GetByIDUsingGET4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetByIDUsingGET4Params() *GetByIDUsingGET4Params {
	return &GetByIDUsingGET4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetByIDUsingGET4ParamsWithTimeout creates a new GetByIDUsingGET4Params object
// with the ability to set a timeout on a request.
func NewGetByIDUsingGET4ParamsWithTimeout(timeout time.Duration) *GetByIDUsingGET4Params {
	return &GetByIDUsingGET4Params{
		timeout: timeout,
	}
}

// NewGetByIDUsingGET4ParamsWithContext creates a new GetByIDUsingGET4Params object
// with the ability to set a context for a request.
func NewGetByIDUsingGET4ParamsWithContext(ctx context.Context) *GetByIDUsingGET4Params {
	return &GetByIDUsingGET4Params{
		Context: ctx,
	}
}

// NewGetByIDUsingGET4ParamsWithHTTPClient creates a new GetByIDUsingGET4Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetByIDUsingGET4ParamsWithHTTPClient(client *http.Client) *GetByIDUsingGET4Params {
	return &GetByIDUsingGET4Params{
		HTTPClient: client,
	}
}

/* GetByIDUsingGET4Params contains all the parameters to send to the API endpoint
   for the get by ID using g e t 4 operation.

   Typically these are written to a http.Request.
*/
type GetByIDUsingGET4Params struct {

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

// WithDefaults hydrates default values in the get by ID using g e t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetByIDUsingGET4Params) WithDefaults() *GetByIDUsingGET4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get by ID using g e t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetByIDUsingGET4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) WithTimeout(timeout time.Duration) *GetByIDUsingGET4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) WithContext(ctx context.Context) *GetByIDUsingGET4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) WithHTTPClient(client *http.Client) *GetByIDUsingGET4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) WithAuthorization(authorization string) *GetByIDUsingGET4Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) WithAPIVersion(aPIVersion string) *GetByIDUsingGET4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) WithID(id string) *GetByIDUsingGET4Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get by ID using g e t 4 params
func (o *GetByIDUsingGET4Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetByIDUsingGET4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
