// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

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

// NewGetBlueprintRequestUsingGET1Params creates a new GetBlueprintRequestUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBlueprintRequestUsingGET1Params() *GetBlueprintRequestUsingGET1Params {
	return &GetBlueprintRequestUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlueprintRequestUsingGET1ParamsWithTimeout creates a new GetBlueprintRequestUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetBlueprintRequestUsingGET1ParamsWithTimeout(timeout time.Duration) *GetBlueprintRequestUsingGET1Params {
	return &GetBlueprintRequestUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetBlueprintRequestUsingGET1ParamsWithContext creates a new GetBlueprintRequestUsingGET1Params object
// with the ability to set a context for a request.
func NewGetBlueprintRequestUsingGET1ParamsWithContext(ctx context.Context) *GetBlueprintRequestUsingGET1Params {
	return &GetBlueprintRequestUsingGET1Params{
		Context: ctx,
	}
}

// NewGetBlueprintRequestUsingGET1ParamsWithHTTPClient creates a new GetBlueprintRequestUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetBlueprintRequestUsingGET1ParamsWithHTTPClient(client *http.Client) *GetBlueprintRequestUsingGET1Params {
	return &GetBlueprintRequestUsingGET1Params{
		HTTPClient: client,
	}
}

/*
GetBlueprintRequestUsingGET1Params contains all the parameters to send to the API endpoint

	for the get blueprint request using get1 operation.

	Typically these are written to a http.Request.
*/
type GetBlueprintRequestUsingGET1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* RequestID.

	   requestId

	   Format: uuid
	*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get blueprint request using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlueprintRequestUsingGET1Params) WithDefaults() *GetBlueprintRequestUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get blueprint request using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlueprintRequestUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) WithTimeout(timeout time.Duration) *GetBlueprintRequestUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) WithContext(ctx context.Context) *GetBlueprintRequestUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) WithHTTPClient(client *http.Client) *GetBlueprintRequestUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) WithAPIVersion(aPIVersion *string) *GetBlueprintRequestUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithRequestID adds the requestID to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) WithRequestID(requestID strfmt.UUID) *GetBlueprintRequestUsingGET1Params {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get blueprint request using get1 params
func (o *GetBlueprintRequestUsingGET1Params) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlueprintRequestUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
