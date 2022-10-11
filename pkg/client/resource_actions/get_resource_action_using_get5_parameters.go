// Code generated by go-swagger; DO NOT EDIT.

package resource_actions

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

// NewGetResourceActionUsingGET5Params creates a new GetResourceActionUsingGET5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceActionUsingGET5Params() *GetResourceActionUsingGET5Params {
	return &GetResourceActionUsingGET5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceActionUsingGET5ParamsWithTimeout creates a new GetResourceActionUsingGET5Params object
// with the ability to set a timeout on a request.
func NewGetResourceActionUsingGET5ParamsWithTimeout(timeout time.Duration) *GetResourceActionUsingGET5Params {
	return &GetResourceActionUsingGET5Params{
		timeout: timeout,
	}
}

// NewGetResourceActionUsingGET5ParamsWithContext creates a new GetResourceActionUsingGET5Params object
// with the ability to set a context for a request.
func NewGetResourceActionUsingGET5ParamsWithContext(ctx context.Context) *GetResourceActionUsingGET5Params {
	return &GetResourceActionUsingGET5Params{
		Context: ctx,
	}
}

// NewGetResourceActionUsingGET5ParamsWithHTTPClient creates a new GetResourceActionUsingGET5Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceActionUsingGET5ParamsWithHTTPClient(client *http.Client) *GetResourceActionUsingGET5Params {
	return &GetResourceActionUsingGET5Params{
		HTTPClient: client,
	}
}

/*
GetResourceActionUsingGET5Params contains all the parameters to send to the API endpoint

	for the get resource action using g e t 5 operation.

	Typically these are written to a http.Request.
*/
type GetResourceActionUsingGET5Params struct {

	/* ActionID.

	   Action ID
	*/
	ActionID string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ResourceID.

	   Resource ID

	   Format: uuid
	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource action using g e t 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceActionUsingGET5Params) WithDefaults() *GetResourceActionUsingGET5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource action using g e t 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceActionUsingGET5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) WithTimeout(timeout time.Duration) *GetResourceActionUsingGET5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) WithContext(ctx context.Context) *GetResourceActionUsingGET5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) WithHTTPClient(client *http.Client) *GetResourceActionUsingGET5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) WithActionID(actionID string) *GetResourceActionUsingGET5Params {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithAPIVersion adds the aPIVersion to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) WithAPIVersion(aPIVersion *string) *GetResourceActionUsingGET5Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithResourceID adds the resourceID to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) WithResourceID(resourceID strfmt.UUID) *GetResourceActionUsingGET5Params {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get resource action using g e t 5 params
func (o *GetResourceActionUsingGET5Params) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceActionUsingGET5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}

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

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}