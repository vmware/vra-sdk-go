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
)

// NewGetNamespaceQuotasUsingGETParams creates a new GetNamespaceQuotasUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNamespaceQuotasUsingGETParams() *GetNamespaceQuotasUsingGETParams {
	return &GetNamespaceQuotasUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceQuotasUsingGETParamsWithTimeout creates a new GetNamespaceQuotasUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetNamespaceQuotasUsingGETParamsWithTimeout(timeout time.Duration) *GetNamespaceQuotasUsingGETParams {
	return &GetNamespaceQuotasUsingGETParams{
		timeout: timeout,
	}
}

// NewGetNamespaceQuotasUsingGETParamsWithContext creates a new GetNamespaceQuotasUsingGETParams object
// with the ability to set a context for a request.
func NewGetNamespaceQuotasUsingGETParamsWithContext(ctx context.Context) *GetNamespaceQuotasUsingGETParams {
	return &GetNamespaceQuotasUsingGETParams{
		Context: ctx,
	}
}

// NewGetNamespaceQuotasUsingGETParamsWithHTTPClient creates a new GetNamespaceQuotasUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNamespaceQuotasUsingGETParamsWithHTTPClient(client *http.Client) *GetNamespaceQuotasUsingGETParams {
	return &GetNamespaceQuotasUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetNamespaceQuotasUsingGETParams contains all the parameters to send to the API endpoint

	for the get namespace quotas using g e t operation.

	Typically these are written to a http.Request.
*/
type GetNamespaceQuotasUsingGETParams struct {

	/* SelfLinkID.

	   selfLinkId
	*/
	SelfLinkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get namespace quotas using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceQuotasUsingGETParams) WithDefaults() *GetNamespaceQuotasUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get namespace quotas using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceQuotasUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) WithTimeout(timeout time.Duration) *GetNamespaceQuotasUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) WithContext(ctx context.Context) *GetNamespaceQuotasUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) WithHTTPClient(client *http.Client) *GetNamespaceQuotasUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelfLinkID adds the selfLinkID to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) WithSelfLinkID(selfLinkID string) *GetNamespaceQuotasUsingGETParams {
	o.SetSelfLinkID(selfLinkID)
	return o
}

// SetSelfLinkID adds the selfLinkId to the get namespace quotas using get params
func (o *GetNamespaceQuotasUsingGETParams) SetSelfLinkID(selfLinkID string) {
	o.SelfLinkID = selfLinkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespaceQuotasUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param selfLinkId
	if err := r.SetPathParam("selfLinkId", o.SelfLinkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
