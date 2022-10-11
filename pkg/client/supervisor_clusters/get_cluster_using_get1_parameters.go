// Code generated by go-swagger; DO NOT EDIT.

package supervisor_clusters

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

// NewGetClusterUsingGET1Params creates a new GetClusterUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterUsingGET1Params() *GetClusterUsingGET1Params {
	return &GetClusterUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterUsingGET1ParamsWithTimeout creates a new GetClusterUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetClusterUsingGET1ParamsWithTimeout(timeout time.Duration) *GetClusterUsingGET1Params {
	return &GetClusterUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetClusterUsingGET1ParamsWithContext creates a new GetClusterUsingGET1Params object
// with the ability to set a context for a request.
func NewGetClusterUsingGET1ParamsWithContext(ctx context.Context) *GetClusterUsingGET1Params {
	return &GetClusterUsingGET1Params{
		Context: ctx,
	}
}

// NewGetClusterUsingGET1ParamsWithHTTPClient creates a new GetClusterUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterUsingGET1ParamsWithHTTPClient(client *http.Client) *GetClusterUsingGET1Params {
	return &GetClusterUsingGET1Params{
		HTTPClient: client,
	}
}

/*
GetClusterUsingGET1Params contains all the parameters to send to the API endpoint

	for the get cluster using get1 operation.

	Typically these are written to a http.Request.
*/
type GetClusterUsingGET1Params struct {

	/* EndpointSelfLinkID.

	   endpointSelfLinkId
	*/
	EndpointSelfLinkID string

	/* Moref.

	   moref
	*/
	Moref string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterUsingGET1Params) WithDefaults() *GetClusterUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) WithTimeout(timeout time.Duration) *GetClusterUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) WithContext(ctx context.Context) *GetClusterUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) WithHTTPClient(client *http.Client) *GetClusterUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointSelfLinkID adds the endpointSelfLinkID to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) WithEndpointSelfLinkID(endpointSelfLinkID string) *GetClusterUsingGET1Params {
	o.SetEndpointSelfLinkID(endpointSelfLinkID)
	return o
}

// SetEndpointSelfLinkID adds the endpointSelfLinkId to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) SetEndpointSelfLinkID(endpointSelfLinkID string) {
	o.EndpointSelfLinkID = endpointSelfLinkID
}

// WithMoref adds the moref to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) WithMoref(moref string) *GetClusterUsingGET1Params {
	o.SetMoref(moref)
	return o
}

// SetMoref adds the moref to the get cluster using get1 params
func (o *GetClusterUsingGET1Params) SetMoref(moref string) {
	o.Moref = moref
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpointSelfLinkId
	if err := r.SetPathParam("endpointSelfLinkId", o.EndpointSelfLinkID); err != nil {
		return err
	}

	// path param moref
	if err := r.SetPathParam("moref", o.Moref); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}