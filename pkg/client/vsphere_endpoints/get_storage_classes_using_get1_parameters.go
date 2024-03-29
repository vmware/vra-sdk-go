// Code generated by go-swagger; DO NOT EDIT.

package vsphere_endpoints

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

// NewGetStorageClassesUsingGET1Params creates a new GetStorageClassesUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStorageClassesUsingGET1Params() *GetStorageClassesUsingGET1Params {
	return &GetStorageClassesUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageClassesUsingGET1ParamsWithTimeout creates a new GetStorageClassesUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetStorageClassesUsingGET1ParamsWithTimeout(timeout time.Duration) *GetStorageClassesUsingGET1Params {
	return &GetStorageClassesUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetStorageClassesUsingGET1ParamsWithContext creates a new GetStorageClassesUsingGET1Params object
// with the ability to set a context for a request.
func NewGetStorageClassesUsingGET1ParamsWithContext(ctx context.Context) *GetStorageClassesUsingGET1Params {
	return &GetStorageClassesUsingGET1Params{
		Context: ctx,
	}
}

// NewGetStorageClassesUsingGET1ParamsWithHTTPClient creates a new GetStorageClassesUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetStorageClassesUsingGET1ParamsWithHTTPClient(client *http.Client) *GetStorageClassesUsingGET1Params {
	return &GetStorageClassesUsingGET1Params{
		HTTPClient: client,
	}
}

/*
GetStorageClassesUsingGET1Params contains all the parameters to send to the API endpoint

	for the get storage classes using get1 operation.

	Typically these are written to a http.Request.
*/
type GetStorageClassesUsingGET1Params struct {

	/* EndpointSelfLinkID.

	   endpointSelfLinkId
	*/
	EndpointSelfLinkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get storage classes using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageClassesUsingGET1Params) WithDefaults() *GetStorageClassesUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get storage classes using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageClassesUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) WithTimeout(timeout time.Duration) *GetStorageClassesUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) WithContext(ctx context.Context) *GetStorageClassesUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) WithHTTPClient(client *http.Client) *GetStorageClassesUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointSelfLinkID adds the endpointSelfLinkID to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) WithEndpointSelfLinkID(endpointSelfLinkID string) *GetStorageClassesUsingGET1Params {
	o.SetEndpointSelfLinkID(endpointSelfLinkID)
	return o
}

// SetEndpointSelfLinkID adds the endpointSelfLinkId to the get storage classes using get1 params
func (o *GetStorageClassesUsingGET1Params) SetEndpointSelfLinkID(endpointSelfLinkID string) {
	o.EndpointSelfLinkID = endpointSelfLinkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageClassesUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpointSelfLinkId
	if err := r.SetPathParam("endpointSelfLinkId", o.EndpointSelfLinkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
