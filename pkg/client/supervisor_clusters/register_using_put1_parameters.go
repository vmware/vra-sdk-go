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

// NewRegisterUsingPUT1Params creates a new RegisterUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterUsingPUT1Params() *RegisterUsingPUT1Params {
	return &RegisterUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterUsingPUT1ParamsWithTimeout creates a new RegisterUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewRegisterUsingPUT1ParamsWithTimeout(timeout time.Duration) *RegisterUsingPUT1Params {
	return &RegisterUsingPUT1Params{
		timeout: timeout,
	}
}

// NewRegisterUsingPUT1ParamsWithContext creates a new RegisterUsingPUT1Params object
// with the ability to set a context for a request.
func NewRegisterUsingPUT1ParamsWithContext(ctx context.Context) *RegisterUsingPUT1Params {
	return &RegisterUsingPUT1Params{
		Context: ctx,
	}
}

// NewRegisterUsingPUT1ParamsWithHTTPClient creates a new RegisterUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterUsingPUT1ParamsWithHTTPClient(client *http.Client) *RegisterUsingPUT1Params {
	return &RegisterUsingPUT1Params{
		HTTPClient: client,
	}
}

/*
RegisterUsingPUT1Params contains all the parameters to send to the API endpoint

	for the register using p u t 1 operation.

	Typically these are written to a http.Request.
*/
type RegisterUsingPUT1Params struct {

	/* ClusterSelfLinkID.

	   clusterSelfLinkId
	*/
	ClusterSelfLinkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterUsingPUT1Params) WithDefaults() *RegisterUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) WithTimeout(timeout time.Duration) *RegisterUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) WithContext(ctx context.Context) *RegisterUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) WithHTTPClient(client *http.Client) *RegisterUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterSelfLinkID adds the clusterSelfLinkID to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) WithClusterSelfLinkID(clusterSelfLinkID string) *RegisterUsingPUT1Params {
	o.SetClusterSelfLinkID(clusterSelfLinkID)
	return o
}

// SetClusterSelfLinkID adds the clusterSelfLinkId to the register using p u t 1 params
func (o *RegisterUsingPUT1Params) SetClusterSelfLinkID(clusterSelfLinkID string) {
	o.ClusterSelfLinkID = clusterSelfLinkID
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterSelfLinkId
	if err := r.SetPathParam("clusterSelfLinkId", o.ClusterSelfLinkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
