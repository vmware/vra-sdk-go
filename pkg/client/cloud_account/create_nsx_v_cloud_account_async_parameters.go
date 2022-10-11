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

// NewCreateNsxVCloudAccountAsyncParams creates a new CreateNsxVCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNsxVCloudAccountAsyncParams() *CreateNsxVCloudAccountAsyncParams {
	return &CreateNsxVCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNsxVCloudAccountAsyncParamsWithTimeout creates a new CreateNsxVCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewCreateNsxVCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *CreateNsxVCloudAccountAsyncParams {
	return &CreateNsxVCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewCreateNsxVCloudAccountAsyncParamsWithContext creates a new CreateNsxVCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewCreateNsxVCloudAccountAsyncParamsWithContext(ctx context.Context) *CreateNsxVCloudAccountAsyncParams {
	return &CreateNsxVCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewCreateNsxVCloudAccountAsyncParamsWithHTTPClient creates a new CreateNsxVCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNsxVCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *CreateNsxVCloudAccountAsyncParams {
	return &CreateNsxVCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
CreateNsxVCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the create nsx v cloud account async operation.

	Typically these are written to a http.Request.
*/
type CreateNsxVCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Cloud Account NsxV Specification
	*/
	Body *models.CloudAccountNsxVSpecification

	/* ValidateOnly.

	   If provided, it only validates the credentials in the Cloud Account Specification, and cloud account will not be created.
	*/
	ValidateOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create nsx v cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNsxVCloudAccountAsyncParams) WithDefaults() *CreateNsxVCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nsx v cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNsxVCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *CreateNsxVCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) WithContext(ctx context.Context) *CreateNsxVCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *CreateNsxVCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) WithAPIVersion(aPIVersion *string) *CreateNsxVCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) WithBody(body *models.CloudAccountNsxVSpecification) *CreateNsxVCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) SetBody(body *models.CloudAccountNsxVSpecification) {
	o.Body = body
}

// WithValidateOnly adds the validateOnly to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) WithValidateOnly(validateOnly *string) *CreateNsxVCloudAccountAsyncParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create nsx v cloud account async params
func (o *CreateNsxVCloudAccountAsyncParams) SetValidateOnly(validateOnly *string) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNsxVCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
