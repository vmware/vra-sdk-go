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

// NewSyncStatusUsingGETParams creates a new SyncStatusUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSyncStatusUsingGETParams() *SyncStatusUsingGETParams {
	return &SyncStatusUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSyncStatusUsingGETParamsWithTimeout creates a new SyncStatusUsingGETParams object
// with the ability to set a timeout on a request.
func NewSyncStatusUsingGETParamsWithTimeout(timeout time.Duration) *SyncStatusUsingGETParams {
	return &SyncStatusUsingGETParams{
		timeout: timeout,
	}
}

// NewSyncStatusUsingGETParamsWithContext creates a new SyncStatusUsingGETParams object
// with the ability to set a context for a request.
func NewSyncStatusUsingGETParamsWithContext(ctx context.Context) *SyncStatusUsingGETParams {
	return &SyncStatusUsingGETParams{
		Context: ctx,
	}
}

// NewSyncStatusUsingGETParamsWithHTTPClient creates a new SyncStatusUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewSyncStatusUsingGETParamsWithHTTPClient(client *http.Client) *SyncStatusUsingGETParams {
	return &SyncStatusUsingGETParams{
		HTTPClient: client,
	}
}

/*
SyncStatusUsingGETParams contains all the parameters to send to the API endpoint

	for the sync status using g e t operation.

	Typically these are written to a http.Request.
*/
type SyncStatusUsingGETParams struct {

	/* NamespaceSelfLinkID.

	   namespaceSelfLinkId
	*/
	NamespaceSelfLinkID string

	/* RequestID.

	   requestId
	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sync status using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SyncStatusUsingGETParams) WithDefaults() *SyncStatusUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sync status using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SyncStatusUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sync status using get params
func (o *SyncStatusUsingGETParams) WithTimeout(timeout time.Duration) *SyncStatusUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync status using get params
func (o *SyncStatusUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync status using get params
func (o *SyncStatusUsingGETParams) WithContext(ctx context.Context) *SyncStatusUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync status using get params
func (o *SyncStatusUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync status using get params
func (o *SyncStatusUsingGETParams) WithHTTPClient(client *http.Client) *SyncStatusUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync status using get params
func (o *SyncStatusUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceSelfLinkID adds the namespaceSelfLinkID to the sync status using get params
func (o *SyncStatusUsingGETParams) WithNamespaceSelfLinkID(namespaceSelfLinkID string) *SyncStatusUsingGETParams {
	o.SetNamespaceSelfLinkID(namespaceSelfLinkID)
	return o
}

// SetNamespaceSelfLinkID adds the namespaceSelfLinkId to the sync status using get params
func (o *SyncStatusUsingGETParams) SetNamespaceSelfLinkID(namespaceSelfLinkID string) {
	o.NamespaceSelfLinkID = namespaceSelfLinkID
}

// WithRequestID adds the requestID to the sync status using get params
func (o *SyncStatusUsingGETParams) WithRequestID(requestID string) *SyncStatusUsingGETParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the sync status using get params
func (o *SyncStatusUsingGETParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *SyncStatusUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespaceSelfLinkId
	if err := r.SetPathParam("namespaceSelfLinkId", o.NamespaceSelfLinkID); err != nil {
		return err
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
