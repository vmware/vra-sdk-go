// Code generated by go-swagger; DO NOT EDIT.

package p_k_s_endpoints

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

// NewGetClustersUsingGETParams creates a new GetClustersUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClustersUsingGETParams() *GetClustersUsingGETParams {
	return &GetClustersUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersUsingGETParamsWithTimeout creates a new GetClustersUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetClustersUsingGETParamsWithTimeout(timeout time.Duration) *GetClustersUsingGETParams {
	return &GetClustersUsingGETParams{
		timeout: timeout,
	}
}

// NewGetClustersUsingGETParamsWithContext creates a new GetClustersUsingGETParams object
// with the ability to set a context for a request.
func NewGetClustersUsingGETParamsWithContext(ctx context.Context) *GetClustersUsingGETParams {
	return &GetClustersUsingGETParams{
		Context: ctx,
	}
}

// NewGetClustersUsingGETParamsWithHTTPClient creates a new GetClustersUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClustersUsingGETParamsWithHTTPClient(client *http.Client) *GetClustersUsingGETParams {
	return &GetClustersUsingGETParams{
		HTTPClient: client,
	}
}

/* GetClustersUsingGETParams contains all the parameters to send to the API endpoint
   for the get clusters using g e t operation.

   Typically these are written to a http.Request.
*/
type GetClustersUsingGETParams struct {

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get clusters using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersUsingGETParams) WithDefaults() *GetClustersUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get clusters using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get clusters using get params
func (o *GetClustersUsingGETParams) WithTimeout(timeout time.Duration) *GetClustersUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters using get params
func (o *GetClustersUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters using get params
func (o *GetClustersUsingGETParams) WithContext(ctx context.Context) *GetClustersUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters using get params
func (o *GetClustersUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clusters using get params
func (o *GetClustersUsingGETParams) WithHTTPClient(client *http.Client) *GetClustersUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clusters using get params
func (o *GetClustersUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get clusters using get params
func (o *GetClustersUsingGETParams) WithID(id string) *GetClustersUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get clusters using get params
func (o *GetClustersUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
