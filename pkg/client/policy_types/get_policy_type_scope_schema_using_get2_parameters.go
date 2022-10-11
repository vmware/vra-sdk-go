// Code generated by go-swagger; DO NOT EDIT.

package policy_types

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

// NewGetPolicyTypeScopeSchemaUsingGET2Params creates a new GetPolicyTypeScopeSchemaUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPolicyTypeScopeSchemaUsingGET2Params() *GetPolicyTypeScopeSchemaUsingGET2Params {
	return &GetPolicyTypeScopeSchemaUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyTypeScopeSchemaUsingGET2ParamsWithTimeout creates a new GetPolicyTypeScopeSchemaUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetPolicyTypeScopeSchemaUsingGET2ParamsWithTimeout(timeout time.Duration) *GetPolicyTypeScopeSchemaUsingGET2Params {
	return &GetPolicyTypeScopeSchemaUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetPolicyTypeScopeSchemaUsingGET2ParamsWithContext creates a new GetPolicyTypeScopeSchemaUsingGET2Params object
// with the ability to set a context for a request.
func NewGetPolicyTypeScopeSchemaUsingGET2ParamsWithContext(ctx context.Context) *GetPolicyTypeScopeSchemaUsingGET2Params {
	return &GetPolicyTypeScopeSchemaUsingGET2Params{
		Context: ctx,
	}
}

// NewGetPolicyTypeScopeSchemaUsingGET2ParamsWithHTTPClient creates a new GetPolicyTypeScopeSchemaUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetPolicyTypeScopeSchemaUsingGET2ParamsWithHTTPClient(client *http.Client) *GetPolicyTypeScopeSchemaUsingGET2Params {
	return &GetPolicyTypeScopeSchemaUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetPolicyTypeScopeSchemaUsingGET2Params contains all the parameters to send to the API endpoint

	for the get policy type scope schema using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetPolicyTypeScopeSchemaUsingGET2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ID.

	   Policy type ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get policy type scope schema using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WithDefaults() *GetPolicyTypeScopeSchemaUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get policy type scope schema using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WithTimeout(timeout time.Duration) *GetPolicyTypeScopeSchemaUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WithContext(ctx context.Context) *GetPolicyTypeScopeSchemaUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WithHTTPClient(client *http.Client) *GetPolicyTypeScopeSchemaUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WithAPIVersion(aPIVersion *string) *GetPolicyTypeScopeSchemaUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WithID(id string) *GetPolicyTypeScopeSchemaUsingGET2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get policy type scope schema using g e t 2 params
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyTypeScopeSchemaUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
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
