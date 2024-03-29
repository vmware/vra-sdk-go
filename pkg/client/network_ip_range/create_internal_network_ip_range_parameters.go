// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

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

// NewCreateInternalNetworkIPRangeParams creates a new CreateInternalNetworkIPRangeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInternalNetworkIPRangeParams() *CreateInternalNetworkIPRangeParams {
	return &CreateInternalNetworkIPRangeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInternalNetworkIPRangeParamsWithTimeout creates a new CreateInternalNetworkIPRangeParams object
// with the ability to set a timeout on a request.
func NewCreateInternalNetworkIPRangeParamsWithTimeout(timeout time.Duration) *CreateInternalNetworkIPRangeParams {
	return &CreateInternalNetworkIPRangeParams{
		timeout: timeout,
	}
}

// NewCreateInternalNetworkIPRangeParamsWithContext creates a new CreateInternalNetworkIPRangeParams object
// with the ability to set a context for a request.
func NewCreateInternalNetworkIPRangeParamsWithContext(ctx context.Context) *CreateInternalNetworkIPRangeParams {
	return &CreateInternalNetworkIPRangeParams{
		Context: ctx,
	}
}

// NewCreateInternalNetworkIPRangeParamsWithHTTPClient creates a new CreateInternalNetworkIPRangeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInternalNetworkIPRangeParamsWithHTTPClient(client *http.Client) *CreateInternalNetworkIPRangeParams {
	return &CreateInternalNetworkIPRangeParams{
		HTTPClient: client,
	}
}

/*
CreateInternalNetworkIPRangeParams contains all the parameters to send to the API endpoint

	for the create internal network IP range operation.

	Typically these are written to a http.Request.
*/
type CreateInternalNetworkIPRangeParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Network IP Range Specification instance
	*/
	Body *models.NetworkIPRangeSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create internal network IP range params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInternalNetworkIPRangeParams) WithDefaults() *CreateInternalNetworkIPRangeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create internal network IP range params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInternalNetworkIPRangeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) WithTimeout(timeout time.Duration) *CreateInternalNetworkIPRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) WithContext(ctx context.Context) *CreateInternalNetworkIPRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) WithHTTPClient(client *http.Client) *CreateInternalNetworkIPRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) WithAPIVersion(aPIVersion *string) *CreateInternalNetworkIPRangeParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) WithBody(body *models.NetworkIPRangeSpecification) *CreateInternalNetworkIPRangeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create internal network IP range params
func (o *CreateInternalNetworkIPRangeParams) SetBody(body *models.NetworkIPRangeSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInternalNetworkIPRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
