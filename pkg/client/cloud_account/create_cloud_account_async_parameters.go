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

// NewCreateCloudAccountAsyncParams creates a new CreateCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCloudAccountAsyncParams() *CreateCloudAccountAsyncParams {
	return &CreateCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudAccountAsyncParamsWithTimeout creates a new CreateCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewCreateCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *CreateCloudAccountAsyncParams {
	return &CreateCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewCreateCloudAccountAsyncParamsWithContext creates a new CreateCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewCreateCloudAccountAsyncParamsWithContext(ctx context.Context) *CreateCloudAccountAsyncParams {
	return &CreateCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewCreateCloudAccountAsyncParamsWithHTTPClient creates a new CreateCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *CreateCloudAccountAsyncParams {
	return &CreateCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
CreateCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the create cloud account async operation.

	Typically these are written to a http.Request.
*/
type CreateCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   CloudAccount instance
	*/
	Body *models.CloudAccountSpecification

	/* ValidateOnly.

	   If provided, it only validates the credentials in the Cloud Account Specification, and cloud account will not be created.
	*/
	ValidateOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudAccountAsyncParams) WithDefaults() *CreateCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *CreateCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) WithContext(ctx context.Context) *CreateCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *CreateCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) WithAPIVersion(aPIVersion *string) *CreateCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) WithBody(body *models.CloudAccountSpecification) *CreateCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) SetBody(body *models.CloudAccountSpecification) {
	o.Body = body
}

// WithValidateOnly adds the validateOnly to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) WithValidateOnly(validateOnly *string) *CreateCloudAccountAsyncParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create cloud account async params
func (o *CreateCloudAccountAsyncParams) SetValidateOnly(validateOnly *string) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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