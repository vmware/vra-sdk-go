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

// NewGetNamespaceUsingGETParams creates a new GetNamespaceUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNamespaceUsingGETParams() *GetNamespaceUsingGETParams {
	return &GetNamespaceUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceUsingGETParamsWithTimeout creates a new GetNamespaceUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetNamespaceUsingGETParamsWithTimeout(timeout time.Duration) *GetNamespaceUsingGETParams {
	return &GetNamespaceUsingGETParams{
		timeout: timeout,
	}
}

// NewGetNamespaceUsingGETParamsWithContext creates a new GetNamespaceUsingGETParams object
// with the ability to set a context for a request.
func NewGetNamespaceUsingGETParamsWithContext(ctx context.Context) *GetNamespaceUsingGETParams {
	return &GetNamespaceUsingGETParams{
		Context: ctx,
	}
}

// NewGetNamespaceUsingGETParamsWithHTTPClient creates a new GetNamespaceUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNamespaceUsingGETParamsWithHTTPClient(client *http.Client) *GetNamespaceUsingGETParams {
	return &GetNamespaceUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetNamespaceUsingGETParams contains all the parameters to send to the API endpoint

	for the get namespace using g e t operation.

	Typically these are written to a http.Request.
*/
type GetNamespaceUsingGETParams struct {

	/* ExpandStorage.

	   expandStorage
	*/
	ExpandStorage *bool

	/* SelfLinkID.

	   selfLinkId
	*/
	SelfLinkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get namespace using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceUsingGETParams) WithDefaults() *GetNamespaceUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get namespace using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get namespace using get params
func (o *GetNamespaceUsingGETParams) WithTimeout(timeout time.Duration) *GetNamespaceUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace using get params
func (o *GetNamespaceUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace using get params
func (o *GetNamespaceUsingGETParams) WithContext(ctx context.Context) *GetNamespaceUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace using get params
func (o *GetNamespaceUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace using get params
func (o *GetNamespaceUsingGETParams) WithHTTPClient(client *http.Client) *GetNamespaceUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace using get params
func (o *GetNamespaceUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpandStorage adds the expandStorage to the get namespace using get params
func (o *GetNamespaceUsingGETParams) WithExpandStorage(expandStorage *bool) *GetNamespaceUsingGETParams {
	o.SetExpandStorage(expandStorage)
	return o
}

// SetExpandStorage adds the expandStorage to the get namespace using get params
func (o *GetNamespaceUsingGETParams) SetExpandStorage(expandStorage *bool) {
	o.ExpandStorage = expandStorage
}

// WithSelfLinkID adds the selfLinkID to the get namespace using get params
func (o *GetNamespaceUsingGETParams) WithSelfLinkID(selfLinkID string) *GetNamespaceUsingGETParams {
	o.SetSelfLinkID(selfLinkID)
	return o
}

// SetSelfLinkID adds the selfLinkId to the get namespace using get params
func (o *GetNamespaceUsingGETParams) SetSelfLinkID(selfLinkID string) {
	o.SelfLinkID = selfLinkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespaceUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExpandStorage != nil {

		// query param expandStorage
		var qrExpandStorage bool

		if o.ExpandStorage != nil {
			qrExpandStorage = *o.ExpandStorage
		}
		qExpandStorage := swag.FormatBool(qrExpandStorage)
		if qExpandStorage != "" {

			if err := r.SetQueryParam("expandStorage", qExpandStorage); err != nil {
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
