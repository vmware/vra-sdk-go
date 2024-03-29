// Code generated by go-swagger; DO NOT EDIT.

package supervisor_namespaces

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
	"github.com/go-openapi/swag"
)

// NewDeleteUsingDELETE2Params creates a new DeleteUsingDELETE2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsingDELETE2Params() *DeleteUsingDELETE2Params {
	return &DeleteUsingDELETE2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsingDELETE2ParamsWithTimeout creates a new DeleteUsingDELETE2Params object
// with the ability to set a timeout on a request.
func NewDeleteUsingDELETE2ParamsWithTimeout(timeout time.Duration) *DeleteUsingDELETE2Params {
	return &DeleteUsingDELETE2Params{
		timeout: timeout,
	}
}

// NewDeleteUsingDELETE2ParamsWithContext creates a new DeleteUsingDELETE2Params object
// with the ability to set a context for a request.
func NewDeleteUsingDELETE2ParamsWithContext(ctx context.Context) *DeleteUsingDELETE2Params {
	return &DeleteUsingDELETE2Params{
		Context: ctx,
	}
}

// NewDeleteUsingDELETE2ParamsWithHTTPClient creates a new DeleteUsingDELETE2Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsingDELETE2ParamsWithHTTPClient(client *http.Client) *DeleteUsingDELETE2Params {
	return &DeleteUsingDELETE2Params{
		HTTPClient: client,
	}
}

/*
DeleteUsingDELETE2Params contains all the parameters to send to the API endpoint

	for the delete using d e l e t e 2 operation.

	Typically these are written to a http.Request.
*/
type DeleteUsingDELETE2Params struct {

	/* Destroy.

	   destroy
	*/
	Destroy *bool

	/* SelfLinkID.

	   selfLinkId
	*/
	SelfLinkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete using d e l e t e 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE2Params) WithDefaults() *DeleteUsingDELETE2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete using d e l e t e 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) WithTimeout(timeout time.Duration) *DeleteUsingDELETE2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) WithContext(ctx context.Context) *DeleteUsingDELETE2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) WithHTTPClient(client *http.Client) *DeleteUsingDELETE2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestroy adds the destroy to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) WithDestroy(destroy *bool) *DeleteUsingDELETE2Params {
	o.SetDestroy(destroy)
	return o
}

// SetDestroy adds the destroy to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) SetDestroy(destroy *bool) {
	o.Destroy = destroy
}

// WithSelfLinkID adds the selfLinkID to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) WithSelfLinkID(selfLinkID string) *DeleteUsingDELETE2Params {
	o.SetSelfLinkID(selfLinkID)
	return o
}

// SetSelfLinkID adds the selfLinkId to the delete using d e l e t e 2 params
func (o *DeleteUsingDELETE2Params) SetSelfLinkID(selfLinkID string) {
	o.SelfLinkID = selfLinkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsingDELETE2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Destroy != nil {

		// query param destroy
		var qrDestroy bool

		if o.Destroy != nil {
			qrDestroy = *o.Destroy
		}
		qDestroy := swag.FormatBool(qrDestroy)
		if qDestroy != "" {

			if err := r.SetQueryParam("destroy", qDestroy); err != nil {
				return err
			}
		}
	}

	// path param selfLinkId
	if err := r.SetPathParam("selfLinkId", o.SelfLinkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
