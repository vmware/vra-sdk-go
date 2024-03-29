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

// NewCancelBlueprintRequestUsingPOSTParams creates a new CancelBlueprintRequestUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelBlueprintRequestUsingPOSTParams() *CancelBlueprintRequestUsingPOSTParams {
	return &CancelBlueprintRequestUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelBlueprintRequestUsingPOSTParamsWithTimeout creates a new CancelBlueprintRequestUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCancelBlueprintRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *CancelBlueprintRequestUsingPOSTParams {
	return &CancelBlueprintRequestUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCancelBlueprintRequestUsingPOSTParamsWithContext creates a new CancelBlueprintRequestUsingPOSTParams object
// with the ability to set a context for a request.
func NewCancelBlueprintRequestUsingPOSTParamsWithContext(ctx context.Context) *CancelBlueprintRequestUsingPOSTParams {
	return &CancelBlueprintRequestUsingPOSTParams{
		Context: ctx,
	}
}

// NewCancelBlueprintRequestUsingPOSTParamsWithHTTPClient creates a new CancelBlueprintRequestUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelBlueprintRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *CancelBlueprintRequestUsingPOSTParams {
	return &CancelBlueprintRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CancelBlueprintRequestUsingPOSTParams contains all the parameters to send to the API endpoint

	for the cancel blueprint request using p o s t operation.

	Typically these are written to a http.Request.
*/
type CancelBlueprintRequestUsingPOSTParams struct {

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

// WithDefaults hydrates default values in the cancel blueprint request using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelBlueprintRequestUsingPOSTParams) WithDefaults() *CancelBlueprintRequestUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel blueprint request using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelBlueprintRequestUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *CancelBlueprintRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) WithContext(ctx context.Context) *CancelBlueprintRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *CancelBlueprintRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CancelBlueprintRequestUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithRequestID adds the requestID to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) WithRequestID(requestID strfmt.UUID) *CancelBlueprintRequestUsingPOSTParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the cancel blueprint request using p o s t params
func (o *CancelBlueprintRequestUsingPOSTParams) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelBlueprintRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
