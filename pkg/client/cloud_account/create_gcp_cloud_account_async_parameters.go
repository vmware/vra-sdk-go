// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewCreateGcpCloudAccountAsyncParams creates a new CreateGcpCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGcpCloudAccountAsyncParams() *CreateGcpCloudAccountAsyncParams {
	return &CreateGcpCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGcpCloudAccountAsyncParamsWithTimeout creates a new CreateGcpCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewCreateGcpCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *CreateGcpCloudAccountAsyncParams {
	return &CreateGcpCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewCreateGcpCloudAccountAsyncParamsWithContext creates a new CreateGcpCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewCreateGcpCloudAccountAsyncParamsWithContext(ctx context.Context) *CreateGcpCloudAccountAsyncParams {
	return &CreateGcpCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewCreateGcpCloudAccountAsyncParamsWithHTTPClient creates a new CreateGcpCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGcpCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *CreateGcpCloudAccountAsyncParams {
	return &CreateGcpCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
CreateGcpCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the create gcp cloud account async operation.

	Typically these are written to a http.Request.
*/
type CreateGcpCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Cloud account Gcp specification
	*/
	Body *models.CloudAccountGcpSpecification

	/* ValidateOnly.

	   If provided, it only validates the credentials in the Cloud Account Specification, and cloud account will not be created.
	*/
	ValidateOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create gcp cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGcpCloudAccountAsyncParams) WithDefaults() *CreateGcpCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create gcp cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGcpCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *CreateGcpCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) WithContext(ctx context.Context) *CreateGcpCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *CreateGcpCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) WithAPIVersion(aPIVersion *string) *CreateGcpCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) WithBody(body *models.CloudAccountGcpSpecification) *CreateGcpCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) SetBody(body *models.CloudAccountGcpSpecification) {
	o.Body = body
}

// WithValidateOnly adds the validateOnly to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) WithValidateOnly(validateOnly *string) *CreateGcpCloudAccountAsyncParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create gcp cloud account async params
func (o *CreateGcpCloudAccountAsyncParams) SetValidateOnly(validateOnly *string) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGcpCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ValidateOnly != nil {

		// query param validateOnly
		var qrValidateOnly string

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := qrValidateOnly
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validateOnly", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}