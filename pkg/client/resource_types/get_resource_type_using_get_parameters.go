// Code generated by go-swagger; DO NOT EDIT.

package resource_types

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

// NewGetResourceTypeUsingGETParams creates a new GetResourceTypeUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceTypeUsingGETParams() *GetResourceTypeUsingGETParams {
	return &GetResourceTypeUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceTypeUsingGETParamsWithTimeout creates a new GetResourceTypeUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetResourceTypeUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceTypeUsingGETParams {
	return &GetResourceTypeUsingGETParams{
		timeout: timeout,
	}
}

// NewGetResourceTypeUsingGETParamsWithContext creates a new GetResourceTypeUsingGETParams object
// with the ability to set a context for a request.
func NewGetResourceTypeUsingGETParamsWithContext(ctx context.Context) *GetResourceTypeUsingGETParams {
	return &GetResourceTypeUsingGETParams{
		Context: ctx,
	}
}

// NewGetResourceTypeUsingGETParamsWithHTTPClient creates a new GetResourceTypeUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceTypeUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceTypeUsingGETParams {
	return &GetResourceTypeUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetResourceTypeUsingGETParams contains all the parameters to send to the API endpoint

	for the get resource type using g e t operation.

	Typically these are written to a http.Request.
*/
type GetResourceTypeUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* ResourceTypeID.

	   resourceTypeId
	*/
	ResourceTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource type using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceTypeUsingGETParams) WithDefaults() *GetResourceTypeUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource type using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceTypeUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceTypeUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) WithContext(ctx context.Context) *GetResourceTypeUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceTypeUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) WithAPIVersion(aPIVersion *string) *GetResourceTypeUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithResourceTypeID adds the resourceTypeID to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) WithResourceTypeID(resourceTypeID string) *GetResourceTypeUsingGETParams {
	o.SetResourceTypeID(resourceTypeID)
	return o
}

// SetResourceTypeID adds the resourceTypeId to the get resource type using get params
func (o *GetResourceTypeUsingGETParams) SetResourceTypeID(resourceTypeID string) {
	o.ResourceTypeID = resourceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceTypeUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resourceTypeId
	if err := r.SetPathParam("resourceTypeId", o.ResourceTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
