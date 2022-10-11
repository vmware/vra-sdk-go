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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewSubmitResourceActionRequestUsingPOST5Params creates a new SubmitResourceActionRequestUsingPOST5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitResourceActionRequestUsingPOST5Params() *SubmitResourceActionRequestUsingPOST5Params {
	return &SubmitResourceActionRequestUsingPOST5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitResourceActionRequestUsingPOST5ParamsWithTimeout creates a new SubmitResourceActionRequestUsingPOST5Params object
// with the ability to set a timeout on a request.
func NewSubmitResourceActionRequestUsingPOST5ParamsWithTimeout(timeout time.Duration) *SubmitResourceActionRequestUsingPOST5Params {
	return &SubmitResourceActionRequestUsingPOST5Params{
		timeout: timeout,
	}
}

// NewSubmitResourceActionRequestUsingPOST5ParamsWithContext creates a new SubmitResourceActionRequestUsingPOST5Params object
// with the ability to set a context for a request.
func NewSubmitResourceActionRequestUsingPOST5ParamsWithContext(ctx context.Context) *SubmitResourceActionRequestUsingPOST5Params {
	return &SubmitResourceActionRequestUsingPOST5Params{
		Context: ctx,
	}
}

// NewSubmitResourceActionRequestUsingPOST5ParamsWithHTTPClient creates a new SubmitResourceActionRequestUsingPOST5Params object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitResourceActionRequestUsingPOST5ParamsWithHTTPClient(client *http.Client) *SubmitResourceActionRequestUsingPOST5Params {
	return &SubmitResourceActionRequestUsingPOST5Params{
		HTTPClient: client,
	}
}

/*
SubmitResourceActionRequestUsingPOST5Params contains all the parameters to send to the API endpoint

	for the submit resource action request using p o s t 5 operation.

	Typically these are written to a http.Request.
*/
type SubmitResourceActionRequestUsingPOST5Params struct {

	/* ActionRequest.

	   actionRequest
	*/
	ActionRequest *models.ResourceActionRequest

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

// WithDefaults hydrates default values in the submit resource action request using p o s t 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitResourceActionRequestUsingPOST5Params) WithDefaults() *SubmitResourceActionRequestUsingPOST5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit resource action request using p o s t 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitResourceActionRequestUsingPOST5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) WithTimeout(timeout time.Duration) *SubmitResourceActionRequestUsingPOST5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) WithContext(ctx context.Context) *SubmitResourceActionRequestUsingPOST5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) WithHTTPClient(client *http.Client) *SubmitResourceActionRequestUsingPOST5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionRequest adds the actionRequest to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) WithActionRequest(actionRequest *models.ResourceActionRequest) *SubmitResourceActionRequestUsingPOST5Params {
	o.SetActionRequest(actionRequest)
	return o
}

// SetActionRequest adds the actionRequest to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) SetActionRequest(actionRequest *models.ResourceActionRequest) {
	o.ActionRequest = actionRequest
}

// WithAPIVersion adds the aPIVersion to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) WithAPIVersion(aPIVersion *string) *SubmitResourceActionRequestUsingPOST5Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithResourceID adds the resourceID to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) WithResourceID(resourceID strfmt.UUID) *SubmitResourceActionRequestUsingPOST5Params {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the submit resource action request using p o s t 5 params
func (o *SubmitResourceActionRequestUsingPOST5Params) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitResourceActionRequestUsingPOST5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ActionRequest != nil {
		if err := r.SetBodyParam(o.ActionRequest); err != nil {
			return err
		}
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
