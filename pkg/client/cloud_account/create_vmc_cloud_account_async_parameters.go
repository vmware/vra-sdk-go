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

// NewCreateVmcCloudAccountAsyncParams creates a new CreateVmcCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVmcCloudAccountAsyncParams() *CreateVmcCloudAccountAsyncParams {
	return &CreateVmcCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVmcCloudAccountAsyncParamsWithTimeout creates a new CreateVmcCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewCreateVmcCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *CreateVmcCloudAccountAsyncParams {
	return &CreateVmcCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewCreateVmcCloudAccountAsyncParamsWithContext creates a new CreateVmcCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewCreateVmcCloudAccountAsyncParamsWithContext(ctx context.Context) *CreateVmcCloudAccountAsyncParams {
	return &CreateVmcCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewCreateVmcCloudAccountAsyncParamsWithHTTPClient creates a new CreateVmcCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVmcCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *CreateVmcCloudAccountAsyncParams {
	return &CreateVmcCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
CreateVmcCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the create vmc cloud account async operation.

	Typically these are written to a http.Request.
*/
type CreateVmcCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   Cloud account VMC specification
	*/
	Body *models.CloudAccountVmcSpecification

	/* ValidateOnly.

	   If provided, it only validates the credentials in the Cloud Account Specification, and cloud account will not be created.
	*/
	ValidateOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vmc cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVmcCloudAccountAsyncParams) WithDefaults() *CreateVmcCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vmc cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVmcCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *CreateVmcCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) WithContext(ctx context.Context) *CreateVmcCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *CreateVmcCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) WithAPIVersion(aPIVersion string) *CreateVmcCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) WithBody(body *models.CloudAccountVmcSpecification) *CreateVmcCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) SetBody(body *models.CloudAccountVmcSpecification) {
	o.Body = body
}

// WithValidateOnly adds the validateOnly to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) WithValidateOnly(validateOnly *string) *CreateVmcCloudAccountAsyncParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create vmc cloud account async params
func (o *CreateVmcCloudAccountAsyncParams) SetValidateOnly(validateOnly *string) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVmcCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
