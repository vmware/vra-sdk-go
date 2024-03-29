// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewGetRequestEventsUsingGET2Params creates a new GetRequestEventsUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRequestEventsUsingGET2Params() *GetRequestEventsUsingGET2Params {
	return &GetRequestEventsUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRequestEventsUsingGET2ParamsWithTimeout creates a new GetRequestEventsUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetRequestEventsUsingGET2ParamsWithTimeout(timeout time.Duration) *GetRequestEventsUsingGET2Params {
	return &GetRequestEventsUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetRequestEventsUsingGET2ParamsWithContext creates a new GetRequestEventsUsingGET2Params object
// with the ability to set a context for a request.
func NewGetRequestEventsUsingGET2ParamsWithContext(ctx context.Context) *GetRequestEventsUsingGET2Params {
	return &GetRequestEventsUsingGET2Params{
		Context: ctx,
	}
}

// NewGetRequestEventsUsingGET2ParamsWithHTTPClient creates a new GetRequestEventsUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRequestEventsUsingGET2ParamsWithHTTPClient(client *http.Client) *GetRequestEventsUsingGET2Params {
	return &GetRequestEventsUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetRequestEventsUsingGET2Params contains all the parameters to send to the API endpoint

	for the get request events using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetRequestEventsUsingGET2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* RequestID.

	   Request ID

	   Format: uuid
	*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get request events using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestEventsUsingGET2Params) WithDefaults() *GetRequestEventsUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get request events using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestEventsUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) WithTimeout(timeout time.Duration) *GetRequestEventsUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) WithContext(ctx context.Context) *GetRequestEventsUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) WithHTTPClient(client *http.Client) *GetRequestEventsUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) WithAPIVersion(aPIVersion *string) *GetRequestEventsUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithRequestID adds the requestID to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) WithRequestID(requestID strfmt.UUID) *GetRequestEventsUsingGET2Params {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get request events using g e t 2 params
func (o *GetRequestEventsUsingGET2Params) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRequestEventsUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
