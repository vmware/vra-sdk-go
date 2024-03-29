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

// NewCreateVcfCloudAccountAsyncParams creates a new CreateVcfCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVcfCloudAccountAsyncParams() *CreateVcfCloudAccountAsyncParams {
	return &CreateVcfCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVcfCloudAccountAsyncParamsWithTimeout creates a new CreateVcfCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewCreateVcfCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *CreateVcfCloudAccountAsyncParams {
	return &CreateVcfCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewCreateVcfCloudAccountAsyncParamsWithContext creates a new CreateVcfCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewCreateVcfCloudAccountAsyncParamsWithContext(ctx context.Context) *CreateVcfCloudAccountAsyncParams {
	return &CreateVcfCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewCreateVcfCloudAccountAsyncParamsWithHTTPClient creates a new CreateVcfCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVcfCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *CreateVcfCloudAccountAsyncParams {
	return &CreateVcfCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
CreateVcfCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the create vcf cloud account async operation.

	Typically these are written to a http.Request.
*/
type CreateVcfCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   Cloud account VCF specification
	*/
	Body *models.CloudAccountVcfSpecification

	/* ValidateOnly.

	   If provided, it only validates the credentials in the Cloud Account Specification, and cloud account will not be created.
	*/
	ValidateOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vcf cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVcfCloudAccountAsyncParams) WithDefaults() *CreateVcfCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vcf cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVcfCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *CreateVcfCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) WithContext(ctx context.Context) *CreateVcfCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *CreateVcfCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) WithAPIVersion(aPIVersion string) *CreateVcfCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) WithBody(body *models.CloudAccountVcfSpecification) *CreateVcfCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) SetBody(body *models.CloudAccountVcfSpecification) {
	o.Body = body
}

// WithValidateOnly adds the validateOnly to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) WithValidateOnly(validateOnly *string) *CreateVcfCloudAccountAsyncParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create vcf cloud account async params
func (o *CreateVcfCloudAccountAsyncParams) SetValidateOnly(validateOnly *string) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVcfCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {

		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
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
